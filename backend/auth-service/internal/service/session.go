package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	// DefaultSessionTimeout is the default session timeout duration
	DefaultSessionTimeout = 30 * time.Minute

	// SessionKeyPrefix is the prefix for session keys in Redis
	SessionKeyPrefix = "session:"
)

// SessionData represents the data stored in a session
type SessionData struct {
	UserID     int64     `json:"user_id"`
	UUID       string    `json:"uuid"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	CustomerID *int64    `json:"customer_id,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	LastAccess time.Time `json:"last_access"`
}

// SessionService handles session management with Redis
type SessionService struct {
	client         *redis.Client
	sessionTimeout time.Duration
}

// SessionConfig holds session service configuration
type SessionConfig struct {
	RedisAddr      string
	RedisPassword  string
	RedisDB        int
	SessionTimeout time.Duration
}

// NewSessionService creates a new session service
func NewSessionService(config SessionConfig) (*SessionService, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	timeout := config.SessionTimeout
	if timeout == 0 {
		timeout = DefaultSessionTimeout
	}

	return &SessionService{
		client:         client,
		sessionTimeout: timeout,
	}, nil
}

// CreateSession creates a new session for a user
func (s *SessionService) CreateSession(ctx context.Context, sessionID string, data *SessionData) error {
	now := time.Now()
	data.CreatedAt = now
	data.LastAccess = now

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal session data: %w", err)
	}

	key := SessionKeyPrefix + sessionID
	err = s.client.Set(ctx, key, jsonData, s.sessionTimeout).Err()
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}

	return nil
}

// GetSession retrieves a session by ID
func (s *SessionService) GetSession(ctx context.Context, sessionID string) (*SessionData, error) {
	key := SessionKeyPrefix + sessionID
	jsonData, err := s.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("session not found")
		}
		return nil, fmt.Errorf("failed to get session: %w", err)
	}

	var data SessionData
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session data: %w", err)
	}

	return &data, nil
}

// UpdateSession updates an existing session
func (s *SessionService) UpdateSession(ctx context.Context, sessionID string, data *SessionData) error {
	data.LastAccess = time.Now()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal session data: %w", err)
	}

	key := SessionKeyPrefix + sessionID
	err = s.client.Set(ctx, key, jsonData, s.sessionTimeout).Err()
	if err != nil {
		return fmt.Errorf("failed to update session: %w", err)
	}

	return nil
}

// RefreshSession extends the session timeout
func (s *SessionService) RefreshSession(ctx context.Context, sessionID string) error {
	key := SessionKeyPrefix + sessionID

	// Check if session exists
	exists, err := s.client.Exists(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("failed to check session existence: %w", err)
	}
	if exists == 0 {
		return fmt.Errorf("session not found")
	}

	// Extend the TTL
	err = s.client.Expire(ctx, key, s.sessionTimeout).Err()
	if err != nil {
		return fmt.Errorf("failed to refresh session: %w", err)
	}

	return nil
}

// DeleteSession deletes a session
func (s *SessionService) DeleteSession(ctx context.Context, sessionID string) error {
	key := SessionKeyPrefix + sessionID
	err := s.client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}

	return nil
}

// DeleteUserSessions deletes all sessions for a user
func (s *SessionService) DeleteUserSessions(ctx context.Context, userID int64) error {
	// Scan for all session keys
	pattern := SessionKeyPrefix + "*"
	iter := s.client.Scan(ctx, 0, pattern, 0).Iterator()

	for iter.Next(ctx) {
		key := iter.Val()

		// Get session data
		jsonData, err := s.client.Get(ctx, key).Result()
		if err != nil {
			continue
		}

		var data SessionData
		if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
			continue
		}

		// Delete if it belongs to the user
		if data.UserID == userID {
			s.client.Del(ctx, key)
		}
	}

	if err := iter.Err(); err != nil {
		return fmt.Errorf("failed to scan sessions: %w", err)
	}

	return nil
}

// Close closes the Redis connection
func (s *SessionService) Close() error {
	return s.client.Close()
}
