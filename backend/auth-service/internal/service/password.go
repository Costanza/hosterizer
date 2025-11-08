package service

import (
	"errors"
	"fmt"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

const (
	// MinPasswordLength is the minimum required password length
	MinPasswordLength = 8

	// MaxPasswordLength is the maximum allowed password length
	MaxPasswordLength = 72 // bcrypt limitation

	// BcryptCost is the cost factor for bcrypt hashing
	BcryptCost = 12
)

var (
	// ErrPasswordTooShort is returned when password is too short
	ErrPasswordTooShort = errors.New("password must be at least 8 characters long")

	// ErrPasswordTooLong is returned when password is too long
	ErrPasswordTooLong = errors.New("password must not exceed 72 characters")

	// ErrPasswordNoUppercase is returned when password has no uppercase letter
	ErrPasswordNoUppercase = errors.New("password must contain at least one uppercase letter")

	// ErrPasswordNoLowercase is returned when password has no lowercase letter
	ErrPasswordNoLowercase = errors.New("password must contain at least one lowercase letter")

	// ErrPasswordNoDigit is returned when password has no digit
	ErrPasswordNoDigit = errors.New("password must contain at least one digit")

	// ErrPasswordNoSpecial is returned when password has no special character
	ErrPasswordNoSpecial = errors.New("password must contain at least one special character")
)

// PasswordService handles password hashing and validation
type PasswordService struct{}

// NewPasswordService creates a new password service
func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

// HashPassword hashes a password using bcrypt
func (s *PasswordService) HashPassword(password string) (string, error) {
	if err := s.ValidatePasswordStrength(password); err != nil {
		return "", err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), BcryptCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hash), nil
}

// ComparePassword compares a password with a hash
func (s *PasswordService) ComparePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return errors.New("invalid password")
		}
		return fmt.Errorf("failed to compare password: %w", err)
	}
	return nil
}

// ValidatePasswordStrength validates password strength requirements
func (s *PasswordService) ValidatePasswordStrength(password string) error {
	// Check length
	if len(password) < MinPasswordLength {
		return ErrPasswordTooShort
	}
	if len(password) > MaxPasswordLength {
		return ErrPasswordTooLong
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasDigit   bool
		hasSpecial bool
	)

	// Check character requirements
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return ErrPasswordNoUppercase
	}
	if !hasLower {
		return ErrPasswordNoLowercase
	}
	if !hasDigit {
		return ErrPasswordNoDigit
	}
	if !hasSpecial {
		return ErrPasswordNoSpecial
	}

	return nil
}

// IsPasswordValid checks if a password meets strength requirements without returning specific errors
func (s *PasswordService) IsPasswordValid(password string) bool {
	return s.ValidatePasswordStrength(password) == nil
}
