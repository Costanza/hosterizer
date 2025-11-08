# Auth Service

The authentication and authorization service for the Hosterizer platform.

## Features

- User authentication with email/password
- JWT token generation and validation (access and refresh tokens)
- Multi-factor authentication (MFA) using TOTP
- Account lockout mechanism after failed login attempts
- Session management with Redis
- Password strength validation
- Secure password hashing with bcrypt

## Architecture

### Domain Layer
- `internal/domain/user.go` - User domain model with business logic
- `internal/domain/repository.go` - Repository interface definitions

### Repository Layer
- `internal/repository/user_postgres.go` - PostgreSQL implementation of UserRepository

### Service Layer
- `internal/service/auth.go` - Main authentication service orchestrating all operations
- `internal/service/password.go` - Password hashing and validation
- `internal/service/jwt.go` - JWT token generation and validation
- `internal/service/mfa.go` - Multi-factor authentication (TOTP)
- `internal/service/lockout.go` - Account lockout mechanism
- `internal/service/session.go` - Session management with Redis

### Handler Layer
- `internal/handler/auth.go` - HTTP handlers for authentication endpoints

## API Endpoints

### POST /api/v1/auth/login
Login with email and password. Returns access and refresh tokens.

**Request:**
```json
{
  "email": "user@example.com",
  "password": "SecurePass123!",
  "mfa_code": "123456"
}
```

**Response:**
```json
{
  "access_token": "eyJhbGc...",
  "refresh_token": "eyJhbGc...",
  "requires_mfa": false,
  "user": {
    "id": 1,
    "uuid": "550e8400-e29b-41d4-a716-446655440000",
    "email": "user@example.com",
    "first_name": "John",
    "last_name": "Doe",
    "role": "customer",
    "mfa_enabled": true
  }
}
```

### POST /api/v1/auth/logout
Logout the current user.

### POST /api/v1/auth/refresh
Refresh access and refresh tokens.

**Request:**
```json
{
  "refresh_token": "eyJhbGc..."
}
```

### POST /api/v1/auth/mfa/setup
Setup MFA for the authenticated user. Returns QR code URL and secret.

**Response:**
```json
{
  "secret": "JBSWY3DPEHPK3PXP",
  "qr_code_url": "otpauth://totp/..."
}
```

### POST /api/v1/auth/mfa/verify
Verify MFA setup with a TOTP code.

**Request:**
```json
{
  "code": "123456"
}
```

### GET /api/v1/auth/me
Get current user information. Requires authentication.

## Configuration

Environment variables:

- `PORT` - Server port (default: 8001)
- `DB_HOST` - PostgreSQL host (default: localhost)
- `DB_PORT` - PostgreSQL port (default: 5432)
- `DB_USER` - PostgreSQL user (default: postgres)
- `DB_PASSWORD` - PostgreSQL password (default: postgres)
- `DB_NAME` - PostgreSQL database name (default: hosterizer)
- `DB_SSLMODE` - PostgreSQL SSL mode (default: disable)
- `JWT_SECRET` - Secret key for JWT signing (required in production)
- `REDIS_ADDR` - Redis address (default: localhost:6379)
- `REDIS_PASSWORD` - Redis password (default: empty)

## Security Features

### Password Requirements
- Minimum 8 characters
- At least one uppercase letter
- At least one lowercase letter
- At least one digit
- At least one special character

### Account Lockout
- Maximum 3 failed login attempts
- Account locked for 15 minutes after exceeding limit
- Automatic unlock after timeout

### Token Expiration
- Access tokens: 15 minutes
- Refresh tokens: 7 days
- Session timeout: 30 minutes

## Running the Service

```bash
# Set environment variables
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=hosterizer
export JWT_SECRET=your-secret-key
export REDIS_ADDR=localhost:6379

# Run the service
go run cmd/server/main.go
```

## Dependencies

- `github.com/lib/pq` - PostgreSQL driver
- `github.com/golang-jwt/jwt/v5` - JWT implementation
- `github.com/redis/go-redis/v9` - Redis client
- `github.com/pquerna/otp` - TOTP implementation
- `golang.org/x/crypto` - bcrypt password hashing
- `github.com/hosterizer/shared` - Shared database utilities

## Testing

The service includes comprehensive unit tests for all components. Run tests with:

```bash
go test ./...
```
