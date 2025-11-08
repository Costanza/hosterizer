package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/hosterizer/auth-service/internal/domain"
	"github.com/lib/pq"
)

// PostgresUserRepository implements UserRepository using PostgreSQL
type PostgresUserRepository struct {
	db *sql.DB
}

// NewPostgresUserRepository creates a new PostgreSQL user repository
func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

// Create creates a new user
func (r *PostgresUserRepository) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (
			email, password_hash, first_name, last_name, role,
			mfa_enabled, mfa_secret, failed_login_attempts
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, uuid, created_at, updated_at
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		user.Email,
		user.PasswordHash,
		user.FirstName,
		user.LastName,
		user.Role,
		user.MFAEnabled,
		user.MFASecret,
		user.FailedLoginAttempts,
	).Scan(&user.ID, &user.UUID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" { // unique_violation
			return domain.ErrUserAlreadyExists
		}
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

// GetByID retrieves a user by ID
func (r *PostgresUserRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	query := `
		SELECT 
			id, uuid, email, password_hash, first_name, last_name, role,
			mfa_enabled, mfa_secret, failed_login_attempts, locked_until,
			last_login_at, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	user := &domain.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.Role,
		&user.MFAEnabled,
		&user.MFASecret,
		&user.FailedLoginAttempts,
		&user.LockedUntil,
		&user.LastLoginAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return user, nil
}

// GetByUUID retrieves a user by UUID
func (r *PostgresUserRepository) GetByUUID(ctx context.Context, uuid string) (*domain.User, error) {
	query := `
		SELECT 
			id, uuid, email, password_hash, first_name, last_name, role,
			mfa_enabled, mfa_secret, failed_login_attempts, locked_until,
			last_login_at, created_at, updated_at
		FROM users
		WHERE uuid = $1
	`

	user := &domain.User{}
	err := r.db.QueryRowContext(ctx, query, uuid).Scan(
		&user.ID,
		&user.UUID,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.Role,
		&user.MFAEnabled,
		&user.MFASecret,
		&user.FailedLoginAttempts,
		&user.LockedUntil,
		&user.LastLoginAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user by uuid: %w", err)
	}

	return user, nil
}

// GetByEmail retrieves a user by email
func (r *PostgresUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
		SELECT 
			id, uuid, email, password_hash, first_name, last_name, role,
			mfa_enabled, mfa_secret, failed_login_attempts, locked_until,
			last_login_at, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	user := &domain.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.Role,
		&user.MFAEnabled,
		&user.MFASecret,
		&user.FailedLoginAttempts,
		&user.LockedUntil,
		&user.LastLoginAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return user, nil
}

// Update updates an existing user
func (r *PostgresUserRepository) Update(ctx context.Context, user *domain.User) error {
	query := `
		UPDATE users
		SET 
			email = $1,
			password_hash = $2,
			first_name = $3,
			last_name = $4,
			role = $5,
			mfa_enabled = $6,
			mfa_secret = $7,
			failed_login_attempts = $8,
			locked_until = $9,
			last_login_at = $10,
			updated_at = NOW()
		WHERE id = $11
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		user.Email,
		user.PasswordHash,
		user.FirstName,
		user.LastName,
		user.Role,
		user.MFAEnabled,
		user.MFASecret,
		user.FailedLoginAttempts,
		user.LockedUntil,
		user.LastLoginAt,
		user.ID,
	)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" { // unique_violation
			return domain.ErrUserAlreadyExists
		}
		return fmt.Errorf("failed to update user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

// Delete deletes a user by ID
func (r *PostgresUserRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

// UpdateFailedAttempts updates the failed login attempts and lock status
func (r *PostgresUserRepository) UpdateFailedAttempts(ctx context.Context, id int64, attempts int, lockedUntil *time.Time) error {
	query := `
		UPDATE users
		SET 
			failed_login_attempts = $1,
			locked_until = $2,
			updated_at = NOW()
		WHERE id = $3
	`

	result, err := r.db.ExecContext(ctx, query, attempts, lockedUntil, id)
	if err != nil {
		return fmt.Errorf("failed to update failed attempts: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

// UpdateLastLogin updates the last login timestamp
func (r *PostgresUserRepository) UpdateLastLogin(ctx context.Context, id int64) error {
	query := `
		UPDATE users
		SET 
			last_login_at = NOW(),
			updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to update last login: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

// UpdateMFASecret updates the MFA secret for a user
func (r *PostgresUserRepository) UpdateMFASecret(ctx context.Context, id int64, secret string, enabled bool) error {
	query := `
		UPDATE users
		SET 
			mfa_secret = $1,
			mfa_enabled = $2,
			updated_at = NOW()
		WHERE id = $3
	`

	result, err := r.db.ExecContext(ctx, query, secret, enabled, id)
	if err != nil {
		return fmt.Errorf("failed to update MFA secret: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}
