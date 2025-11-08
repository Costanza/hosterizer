package domain

import (
	"context"
	"errors"
	"time"
)

var (
	// ErrUserNotFound is returned when a user is not found
	ErrUserNotFound = errors.New("user not found")

	// ErrUserAlreadyExists is returned when attempting to create a user with an existing email
	ErrUserAlreadyExists = errors.New("user already exists")

	// ErrInvalidCredentials is returned when login credentials are invalid
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	// Create creates a new user
	Create(ctx context.Context, user *User) error

	// GetByID retrieves a user by ID
	GetByID(ctx context.Context, id int64) (*User, error)

	// GetByUUID retrieves a user by UUID
	GetByUUID(ctx context.Context, uuid string) (*User, error)

	// GetByEmail retrieves a user by email
	GetByEmail(ctx context.Context, email string) (*User, error)

	// Update updates an existing user
	Update(ctx context.Context, user *User) error

	// Delete deletes a user by ID
	Delete(ctx context.Context, id int64) error

	// UpdateFailedAttempts updates the failed login attempts and lock status
	UpdateFailedAttempts(ctx context.Context, id int64, attempts int, lockedUntil *time.Time) error

	// UpdateLastLogin updates the last login timestamp
	UpdateLastLogin(ctx context.Context, id int64) error

	// UpdateMFASecret updates the MFA secret for a user
	UpdateMFASecret(ctx context.Context, id int64, secret string, enabled bool) error
}
