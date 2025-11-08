package service

import (
	"context"
	"fmt"
	"time"

	"github.com/hosterizer/auth-service/internal/domain"
)

const (
	// MaxFailedAttempts is the maximum number of failed login attempts before lockout
	MaxFailedAttempts = 3

	// LockoutDuration is the duration for which an account is locked
	LockoutDuration = 15 * time.Minute
)

// LockoutService handles account lockout logic
type LockoutService struct {
	userRepo          domain.UserRepository
	maxFailedAttempts int
	lockoutDuration   time.Duration
}

// LockoutConfig holds lockout service configuration
type LockoutConfig struct {
	MaxFailedAttempts int
	LockoutDuration   time.Duration
}

// NewLockoutService creates a new lockout service
func NewLockoutService(userRepo domain.UserRepository, config LockoutConfig) *LockoutService {
	maxAttempts := config.MaxFailedAttempts
	if maxAttempts == 0 {
		maxAttempts = MaxFailedAttempts
	}

	duration := config.LockoutDuration
	if duration == 0 {
		duration = LockoutDuration
	}

	return &LockoutService{
		userRepo:          userRepo,
		maxFailedAttempts: maxAttempts,
		lockoutDuration:   duration,
	}
}

// RecordFailedAttempt records a failed login attempt and locks the account if necessary
func (s *LockoutService) RecordFailedAttempt(ctx context.Context, user *domain.User) error {
	user.IncrementFailedAttempts()

	// Check if we should lock the account
	if user.FailedLoginAttempts >= s.maxFailedAttempts {
		user.LockAccount(s.lockoutDuration)
	}

	// Update the database
	err := s.userRepo.UpdateFailedAttempts(ctx, user.ID, user.FailedLoginAttempts, user.LockedUntil)
	if err != nil {
		return fmt.Errorf("failed to update failed attempts: %w", err)
	}

	return nil
}

// ResetFailedAttempts resets the failed login attempts for a user
func (s *LockoutService) ResetFailedAttempts(ctx context.Context, user *domain.User) error {
	user.ResetFailedAttempts()

	err := s.userRepo.UpdateFailedAttempts(ctx, user.ID, 0, nil)
	if err != nil {
		return fmt.Errorf("failed to reset failed attempts: %w", err)
	}

	return nil
}

// IsAccountLocked checks if an account is currently locked
func (s *LockoutService) IsAccountLocked(user *domain.User) bool {
	return user.IsLocked()
}

// UnlockAccount manually unlocks an account
func (s *LockoutService) UnlockAccount(ctx context.Context, user *domain.User) error {
	user.ResetFailedAttempts()

	err := s.userRepo.UpdateFailedAttempts(ctx, user.ID, 0, nil)
	if err != nil {
		return fmt.Errorf("failed to unlock account: %w", err)
	}

	return nil
}

// GetRemainingLockoutTime returns the remaining lockout time for a user
func (s *LockoutService) GetRemainingLockoutTime(user *domain.User) time.Duration {
	if !user.IsLocked() {
		return 0
	}

	remaining := time.Until(*user.LockedUntil)
	if remaining < 0 {
		return 0
	}

	return remaining
}
