package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/hosterizer/auth-service/internal/service"
)

// AuthHandler handles authentication HTTP requests
type AuthHandler struct {
	authSvc *service.AuthService
	jwtSvc  *service.JWTService
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authSvc *service.AuthService, jwtSvc *service.JWTService) *AuthHandler {
	return &AuthHandler{
		authSvc: authSvc,
		jwtSvc:  jwtSvc,
	}
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// LoginRequest represents a login request
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	MFACode  string `json:"mfa_code,omitempty"`
}

// LoginResponse represents a login response
type LoginResponse struct {
	AccessToken  string    `json:"access_token,omitempty"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	RequiresMFA  bool      `json:"requires_mfa"`
	User         *UserInfo `json:"user,omitempty"`
}

// UserInfo represents user information in responses
type UserInfo struct {
	ID         int64  `json:"id"`
	UUID       string `json:"uuid"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Role       string `json:"role"`
	MFAEnabled bool   `json:"mfa_enabled"`
}

// Login handles login requests
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Validate input
	if req.Email == "" || req.Password == "" {
		h.sendError(w, http.StatusBadRequest, "email and password are required")
		return
	}

	// Perform login
	resp, err := h.authSvc.Login(r.Context(), service.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
		MFACode:  req.MFACode,
	})
	if err != nil {
		h.sendError(w, http.StatusUnauthorized, err.Error())
		return
	}

	// Build response
	loginResp := LoginResponse{
		RequiresMFA: resp.RequiresMFA,
	}

	if !resp.RequiresMFA {
		loginResp.AccessToken = resp.AccessToken
		loginResp.RefreshToken = resp.RefreshToken
		loginResp.User = &UserInfo{
			ID:         resp.User.ID,
			UUID:       resp.User.UUID,
			Email:      resp.User.Email,
			FirstName:  resp.User.FirstName,
			LastName:   resp.User.LastName,
			Role:       string(resp.User.Role),
			MFAEnabled: resp.User.MFAEnabled,
		}
	}

	h.sendJSON(w, http.StatusOK, loginResp)
}

// Logout handles logout requests
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// In a stateless JWT system, logout is handled client-side by removing the token
	// If using sessions, we would invalidate the session here

	h.sendJSON(w, http.StatusOK, map[string]string{
		"message": "logged out successfully",
	})
}

// RefreshRequest represents a refresh token request
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// Refresh handles token refresh requests
func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.RefreshToken == "" {
		h.sendError(w, http.StatusBadRequest, "refresh token is required")
		return
	}

	// Refresh tokens
	resp, err := h.authSvc.RefreshTokens(r.Context(), req.RefreshToken)
	if err != nil {
		h.sendError(w, http.StatusUnauthorized, err.Error())
		return
	}

	h.sendJSON(w, http.StatusOK, LoginResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		User: &UserInfo{
			ID:         resp.User.ID,
			UUID:       resp.User.UUID,
			Email:      resp.User.Email,
			FirstName:  resp.User.FirstName,
			LastName:   resp.User.LastName,
			Role:       string(resp.User.Role),
			MFAEnabled: resp.User.MFAEnabled,
		},
	})
}

// MFASetupResponse represents MFA setup response
type MFASetupResponse struct {
	Secret    string `json:"secret"`
	QRCodeURL string `json:"qr_code_url"`
}

// SetupMFA handles MFA setup requests
func (h *AuthHandler) SetupMFA(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// Extract user ID from token
	userID, err := h.getUserIDFromToken(r)
	if err != nil {
		h.sendError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	// Setup MFA
	result, err := h.authSvc.SetupMFA(r.Context(), userID)
	if err != nil {
		h.sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.sendJSON(w, http.StatusOK, MFASetupResponse{
		Secret:    result.Secret,
		QRCodeURL: result.QRCodeURL,
	})
}

// VerifyMFARequest represents MFA verification request
type VerifyMFARequest struct {
	Code string `json:"code"`
}

// VerifyMFA handles MFA verification requests
func (h *AuthHandler) VerifyMFA(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// Extract user ID from token
	userID, err := h.getUserIDFromToken(r)
	if err != nil {
		h.sendError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	var req VerifyMFARequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Code == "" {
		h.sendError(w, http.StatusBadRequest, "code is required")
		return
	}

	// Verify and enable MFA
	if err := h.authSvc.VerifyAndEnableMFA(r.Context(), userID, req.Code); err != nil {
		h.sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	h.sendJSON(w, http.StatusOK, map[string]string{
		"message": "MFA enabled successfully",
	})
}

// GetMe handles current user info requests
func (h *AuthHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.sendError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// Extract user ID from token
	userID, err := h.getUserIDFromToken(r)
	if err != nil {
		h.sendError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	// Get user
	user, err := h.authSvc.GetCurrentUser(r.Context(), userID)
	if err != nil {
		h.sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.sendJSON(w, http.StatusOK, UserInfo{
		ID:         user.ID,
		UUID:       user.UUID,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Role:       string(user.Role),
		MFAEnabled: user.MFAEnabled,
	})
}

// Helper methods

func (h *AuthHandler) getUserIDFromToken(r *http.Request) (int64, error) {
	// Extract token from Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 0, http.ErrNoCookie
	}

	// Remove "Bearer " prefix
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return 0, http.ErrNoCookie
	}

	// Validate token
	claims, err := h.jwtSvc.ValidateAccessToken(tokenString)
	if err != nil {
		return 0, err
	}

	return claims.UserID, nil
}

func (h *AuthHandler) sendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *AuthHandler) sendError(w http.ResponseWriter, status int, message string) {
	h.sendJSON(w, status, ErrorResponse{
		Error:   http.StatusText(status),
		Message: message,
	})
}

// RegisterRoutes registers all auth routes
func (h *AuthHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/auth/login", h.Login)
	mux.HandleFunc("/api/v1/auth/logout", h.Logout)
	mux.HandleFunc("/api/v1/auth/refresh", h.Refresh)
	mux.HandleFunc("/api/v1/auth/mfa/setup", h.SetupMFA)
	mux.HandleFunc("/api/v1/auth/mfa/verify", h.VerifyMFA)
	mux.HandleFunc("/api/v1/auth/me", h.GetMe)
}
