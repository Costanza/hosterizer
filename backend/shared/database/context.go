package database

import (
	"context"
	"database/sql"
	"fmt"
)

// TenantContext holds tenant-specific information for RLS
type TenantContext struct {
	CustomerID int64
	UserRole   string
}

// SetTenantContext sets the PostgreSQL session variables for row-level security
func SetTenantContext(ctx context.Context, db *sql.DB, tc TenantContext) error {
	conn, err := db.Conn(ctx)
	if err != nil {
		return fmt.Errorf("failed to get connection: %w", err)
	}
	defer conn.Close()

	// Set customer ID
	if tc.CustomerID > 0 {
		_, err = conn.ExecContext(ctx, "SET LOCAL app.current_customer_id = $1", tc.CustomerID)
		if err != nil {
			return fmt.Errorf("failed to set customer_id: %w", err)
		}
	}

	// Set user role
	if tc.UserRole != "" {
		_, err = conn.ExecContext(ctx, "SET LOCAL app.current_user_role = $1", tc.UserRole)
		if err != nil {
			return fmt.Errorf("failed to set user_role: %w", err)
		}
	}

	return nil
}

// WithTenantContext executes a function within a transaction with tenant context set
func WithTenantContext(ctx context.Context, db *sql.DB, tc TenantContext, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Set customer ID
	if tc.CustomerID > 0 {
		_, err = tx.ExecContext(ctx, "SET LOCAL app.current_customer_id = $1", tc.CustomerID)
		if err != nil {
			return fmt.Errorf("failed to set customer_id: %w", err)
		}
	}

	// Set user role
	if tc.UserRole != "" {
		_, err = tx.ExecContext(ctx, "SET LOCAL app.current_user_role = $1", tc.UserRole)
		if err != nil {
			return fmt.Errorf("failed to set user_role: %w", err)
		}
	}

	// Execute the function
	if err := fn(tx); err != nil {
		return err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// ClearTenantContext clears the PostgreSQL session variables
func ClearTenantContext(ctx context.Context, db *sql.DB) error {
	conn, err := db.Conn(ctx)
	if err != nil {
		return fmt.Errorf("failed to get connection: %w", err)
	}
	defer conn.Close()

	_, err = conn.ExecContext(ctx, "RESET app.current_customer_id")
	if err != nil {
		return fmt.Errorf("failed to reset customer_id: %w", err)
	}

	_, err = conn.ExecContext(ctx, "RESET app.current_user_role")
	if err != nil {
		return fmt.Errorf("failed to reset user_role: %w", err)
	}

	return nil
}
