package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hosterizer/auth-service/internal/domain"
)

const (
	// DefaultAccessTokenDuration is the default duration for access tokens
	DefaultAccessTokenDuration = 15 * time.Minute

	// DefaultRefreshTokenDuration is the default duration for refresh tokens
	DefaultRefreshTokenDuration = 7 * 24 * time.Hour
)

var (
	// ErrInvalidToken is returned when a token is invalid
	ErrInvalidToken = errors.New("invalid token")

	// ErrExpiredToken is returned when a token has expired
	ErrExpiredToken = errors.New("token has expired")
)

// TokenClaims represents the JWT claims
type TokenClaims struct {
	UserID     int64           `json:"user_id"`
	UUID       string          `json:"uuid"`
	Email      string          `json:"email"`
	Role       domain.UserRole `json:"role"`
	CustomerID *int64          `json:"customer_id,omitempty"`
	TokenType  string          `json:"token_type"` // "access" or "refresh"
	jwt.RegisteredClaims
}

// JWTService handles JWT token generation and validation
type JWTService struct {
	secretKey            []byte
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

// JWTConfig holds JWT service configuration
type JWTConfig struct {
	SecretKey            string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

// NewJWTService creates a new JWT service
func NewJWTService(config JWTConfig) *JWTService {
	accessDuration := config.AccessTokenDuration
	if accessDuration == 0 {
		accessDuration = DefaultAccessTokenDuration
	}

	refreshDuration := config.RefreshTokenDuration
	if refreshDuration == 0 {
		refreshDuration = DefaultRefreshTokenDuration
	}

	return &JWTService{
		secretKey:            []byte(config.SecretKey),
		accessTokenDuration:  accessDuration,
		refreshTokenDuration: refreshDuration,
	}
}

// GenerateAccessToken generates an access token for a user
func (s *JWTService) GenerateAccessToken(user *domain.User, customerID *int64) (string, error) {
	now := time.Now()
	claims := TokenClaims{
		UserID:     user.ID,
		UUID:       user.UUID,
		Email:      user.Email,
		Role:       user.Role,
		CustomerID: customerID,
		TokenType:  "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(s.accessTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "hosterizer-auth",
			Subject:   user.UUID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign access token: %w", err)
	}

	return tokenString, nil
}

// GenerateRefreshToken generates a refresh token for a user
func (s *JWTService) GenerateRefreshToken(user *domain.User) (string, error) {
	now := time.Now()
	claims := TokenClaims{
		UserID:    user.ID,
		UUID:      user.UUID,
		Email:     user.Email,
		Role:      user.Role,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(s.refreshTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "hosterizer-auth",
			Subject:   user.UUID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return tokenString, nil
}

// ValidateToken validates a JWT token and returns the claims
func (s *JWTService) ValidateToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.secretKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, fmt.Errorf("%w: %v", ErrInvalidToken, err)
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

// ValidateAccessToken validates an access token
func (s *JWTService) ValidateAccessToken(tokenString string) (*TokenClaims, error) {
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.TokenType != "access" {
		return nil, errors.New("invalid token type: expected access token")
	}

	return claims, nil
}

// ValidateRefreshToken validates a refresh token
func (s *JWTService) ValidateRefreshToken(tokenString string) (*TokenClaims, error) {
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.TokenType != "refresh" {
		return nil, errors.New("invalid token type: expected refresh token")
	}

	return claims, nil
}

// GetAccessTokenDuration returns the access token duration
func (s *JWTService) GetAccessTokenDuration() time.Duration {
	return s.accessTokenDuration
}

// GetRefreshTokenDuration returns the refresh token duration
func (s *JWTService) GetRefreshTokenDuration() time.Duration {
	return s.refreshTokenDuration
}
