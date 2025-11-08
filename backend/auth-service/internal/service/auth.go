package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/hosterizer/auth-service/internal/domain"
)

// AuthService orchestrates authentication operations
type AuthService struct {
	userRepo    domain.UserRepository
	passwordSvc *PasswordService
	jwtSvc      *JWTService
	mfaSvc      *MFAService
	lockoutSvc  *LockoutService
	sessionSvc  *SessionService
}

// AuthServiceConfig holds auth service configuration
type AuthServiceConfig struct {
	UserRepo    domain.UserRepository
	PasswordSvc *PasswordService
	JWTSvc      *JWTService
	MFASvc      *MFAService
	LockoutSvc  *LockoutService
	SessionSvc  *SessionService
}

// NewAuthService creates a new auth service
func NewAuthService(config AuthServiceConfig) *AuthService {
	return &AuthService{
		userRepo:    config.UserRepo,
		passwordSvc: config.PasswordSvc,
		jwtSvc:      config.JWTSvc,
		mfaSvc:      config.MFASvc,
		lockoutSvc:  config.LockoutSvc,
		sessionSvc:  config.SessionSvc,
	}
}

// LoginRequest represents a login request
type LoginRequest struct {
	Email    string
	Password string
	MFACode  string
}

// LoginResponse represents a login response
type LoginResponse struct {
	AccessToken  string
	RefreshToken string
	User         *domain.User
	RequiresMFA  bool
}

// Login authenticates a user and returns tokens
func (s *AuthService) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	// Get user by email
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, domain.ErrInvalidCredentials
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Check if account is locked
	if s.lockoutSvc.IsAccountLocked(user) {
		return nil, fmt.Errorf("account is locked until %v", user.LockedUntil)
	}

	// Verify password
	if err := s.passwordSvc.ComparePassword(user.PasswordHash, req.Password); err != nil {
		// Record failed attempt
		if err := s.lockoutSvc.RecordFailedAttempt(ctx, user); err != nil {
			return nil, fmt.Errorf("failed to record failed attempt: %w", err)
		}
		return nil, domain.ErrInvalidCredentials
	}

	// Check if MFA is enabled
	if user.MFAEnabled {
		if req.MFACode == "" {
			return &LoginResponse{
				RequiresMFA: true,
			}, nil
		}

		// Validate MFA code
		valid, err := s.mfaSvc.ValidateCodeWithWindow(user.MFASecret, req.MFACode, 1)
		if err != nil {
			return nil, fmt.Errorf("failed to validate MFA code: %w", err)
		}
		if !valid {
			if err := s.lockoutSvc.RecordFailedAttempt(ctx, user); err != nil {
				return nil, fmt.Errorf("failed to record failed attempt: %w", err)
			}
			return nil, errors.New("invalid MFA code")
		}
	}

	// Reset failed attempts on successful login
	if err := s.lockoutSvc.ResetFailedAttempts(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to reset failed attempts: %w", err)
	}

	// Update last login
	if err := s.userRepo.UpdateLastLogin(ctx, user.ID); err != nil {
		return nil, fmt.Errorf("failed to update last login: %w", err)
	}

	// Determine customer ID for token
	var customerID *int64
	// This would be populated from a customer relationship if the user is a customer

	// Generate tokens
	accessToken, err := s.jwtSvc.GenerateAccessToken(user, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err := s.jwtSvc.GenerateRefreshToken(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
		RequiresMFA:  false,
	}, nil
}

// RefreshTokens refreshes access and refresh tokens
func (s *AuthService) RefreshTokens(ctx context.Context, refreshToken string) (*LoginResponse, error) {
	// Validate refresh token
	claims, err := s.jwtSvc.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %w", err)
	}

	// Get user
	user, err := s.userRepo.GetByID(ctx, claims.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Check if account is locked
	if s.lockoutSvc.IsAccountLocked(user) {
		return nil, fmt.Errorf("account is locked")
	}

	// Generate new tokens
	accessToken, err := s.jwtSvc.GenerateAccessToken(user, claims.CustomerID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	newRefreshToken, err := s.jwtSvc.GenerateRefreshToken(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		User:         user,
	}, nil
}

// SetupMFA sets up MFA for a user
func (s *AuthService) SetupMFA(ctx context.Context, userID int64) (*MFASetupResult, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Generate MFA secret
	result, err := s.mfaSvc.GenerateSecret(user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to generate MFA secret: %w", err)
	}

	// Store the secret (but don't enable MFA yet)
	if err := s.userRepo.UpdateMFASecret(ctx, userID, result.Secret, false); err != nil {
		return nil, fmt.Errorf("failed to store MFA secret: %w", err)
	}

	return result, nil
}

// VerifyAndEnableMFA verifies MFA setup and enables it
func (s *AuthService) VerifyAndEnableMFA(ctx context.Context, userID int64, code string) error {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	if user.MFASecret == "" {
		return errors.New("MFA not set up")
	}

	// Validate the code
	valid, err := s.mfaSvc.ValidateCodeWithWindow(user.MFASecret, code, 1)
	if err != nil {
		return fmt.Errorf("failed to validate MFA code: %w", err)
	}
	if !valid {
		return errors.New("invalid MFA code")
	}

	// Enable MFA
	if err := s.userRepo.UpdateMFASecret(ctx, userID, user.MFASecret, true); err != nil {
		return fmt.Errorf("failed to enable MFA: %w", err)
	}

	return nil
}

// DisableMFA disables MFA for a user
func (s *AuthService) DisableMFA(ctx context.Context, userID int64) error {
	if err := s.userRepo.UpdateMFASecret(ctx, userID, "", false); err != nil {
		return fmt.Errorf("failed to disable MFA: %w", err)
	}
	return nil
}

// GetCurrentUser retrieves the current user information
func (s *AuthService) GetCurrentUser(ctx context.Context, userID int64) (*domain.User, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}
