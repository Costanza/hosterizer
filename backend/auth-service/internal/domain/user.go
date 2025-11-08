package domain

import (
	"time"
)

// UserRole represents the role of a user in the system
type UserRole string

const (
	RoleAdministrator UserRole = "administrator"
	RoleCustomer      UserRole = "customer"
)

// User represents a system user with authentication credentials
type User struct {
	ID                  int64
	UUID                string
	Email               string
	PasswordHash        string
	FirstName           string
	LastName            string
	Role                UserRole
	MFAEnabled          bool
	MFASecret           string
	FailedLoginAttempts int
	LockedUntil         *time.Time
	LastLoginAt         *time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

// IsLocked checks if the user account is currently locked
func (u *User) IsLocked() bool {
	if u.LockedUntil == nil {
		return false
	}
	return time.Now().Before(*u.LockedUntil)
}

// CanLogin checks if the user can attempt to login
func (u *User) CanLogin() bool {
	return !u.IsLocked()
}

// IncrementFailedAttempts increments the failed login attempts counter
func (u *User) IncrementFailedAttempts() {
	u.FailedLoginAttempts++
}

// ResetFailedAttempts resets the failed login attempts counter
func (u *User) ResetFailedAttempts() {
	u.FailedLoginAttempts = 0
	u.LockedUntil = nil
}

// LockAccount locks the user account for the specified duration
func (u *User) LockAccount(duration time.Duration) {
	lockUntil := time.Now().Add(duration)
	u.LockedUntil = &lockUntil
}

// UpdateLastLogin updates the last login timestamp
func (u *User) UpdateLastLogin() {
	now := time.Now()
	u.LastLoginAt = &now
}
