# Auth Service Endpoints

This folder contains all authentication and authorization endpoints for the Hosterizer platform.

## Endpoints Overview

### Health & Status
- **Health Check** - Verify service is running

### Authentication
- **Login - Success** - Successful login with valid credentials
- **Login - Invalid Credentials** - Test error handling for wrong password
- **Login - Missing Fields** - Test validation for missing required fields
- **Login - With MFA** - Login for users with MFA enabled
- **Logout** - End user session

### Token Management
- **Refresh Token** - Refresh access and refresh tokens
- **Refresh Token - Invalid** - Test error handling for invalid tokens

### User Information
- **Get Current User** - Retrieve authenticated user information
- **Get Current User - Unauthorized** - Test protected endpoint without auth

### Multi-Factor Authentication
- **MFA Setup** - Initialize MFA setup (get QR code)
- **MFA Verify** - Verify and enable MFA with TOTP code

## Request Flow

### Standard Login Flow
```
1. Health Check
   ↓
2. Login - Success
   ↓ (saves access_token and refresh_token)
3. Get Current User
   ↓ (uses access_token)
4. Refresh Token
   ↓ (uses refresh_token, gets new tokens)
5. Logout
```

### MFA Setup Flow
```
1. Login - Success
   ↓ (saves access_token)
2. MFA Setup
   ↓ (returns QR code and secret)
3. [Scan QR code with authenticator app]
   ↓
4. MFA Verify
   ↓ (enable MFA with TOTP code)
5. Logout
   ↓
6. Login - With MFA
   ↓ (requires TOTP code)
```

### Error Testing Flow
```
1. Login - Invalid Credentials
   ↓ (test wrong password)
2. Login - Missing Fields
   ↓ (test validation)
3. Get Current User - Unauthorized
   ↓ (test without token)
4. Refresh Token - Invalid
   ↓ (test with bad token)
```

## Authentication

Most endpoints require authentication via Bearer token:

```
Authorization: Bearer {{access_token}}
```

The `access_token` is automatically set after successful login.

## Environment Variables Used

- `{{auth_base_url}}` - Base URL for auth endpoints
- `{{access_token}}` - JWT access token (auto-set)
- `{{refresh_token}}` - JWT refresh token (auto-set)
- `{{user_id}}` - Current user ID (auto-set)
- `{{user_uuid}}` - Current user UUID (auto-set)
- `{{mfa_secret}}` - MFA secret (auto-set after setup)
- `{{mfa_qr_code}}` - MFA QR code URL (auto-set after setup)

## Response Codes

| Code | Meaning | When |
|------|---------|------|
| 200 | OK | Successful request |
| 400 | Bad Request | Invalid input or validation error |
| 401 | Unauthorized | Missing or invalid authentication |
| 405 | Method Not Allowed | Wrong HTTP method |
| 500 | Internal Server Error | Server-side error |

## Common Patterns

### Success Response
```json
{
  "access_token": "eyJhbGc...",
  "refresh_token": "eyJhbGc...",
  "user": {
    "id": 1,
    "uuid": "550e8400-e29b-41d4-a716-446655440000",
    "email": "user@example.com",
    "first_name": "John",
    "last_name": "Doe",
    "role": "customer",
    "mfa_enabled": false
  }
}
```

### Error Response
```json
{
  "error": "Unauthorized",
  "message": "invalid credentials"
}
```

## Testing Tips

1. **Run in Order**: Execute requests in sequence number order for best results
2. **Check Tests**: Each request has automated tests - check the "Tests" tab
3. **View Docs**: Click "Docs" tab for detailed information about each request
4. **Token Expiry**: Access tokens expire after 15 minutes - use refresh or login again
5. **MFA Testing**: You'll need a real authenticator app (Google Authenticator, Authy, etc.)

## Troubleshooting

### 401 Unauthorized
- Run "Login - Success" to get a fresh token
- Check if token has expired (15-minute lifetime)
- Verify you're using the correct environment

### MFA Verification Fails
- Ensure you're using the current code (expires every 30 seconds)
- Check device time is synchronized
- Verify you scanned the correct QR code

### Connection Refused
- Ensure auth service is running: `go run cmd/server/main.go`
- Check service is on port 8001: `curl http://localhost:8001/health`
- Verify environment URL matches your setup

## Related Documentation

- [Auth Service README](../../../backend/auth-service/README.md)
- [API Testing Guide](../../../docs/api-testing.md)
- [Quick Start Guide](../../QUICKSTART.md)
