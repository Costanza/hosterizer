# Hosterizer API - Bruno Collection

This Bruno collection provides comprehensive API testing for the Hosterizer platform.

## What is Bruno?

Bruno is an open-source API client similar to Postman, but with a focus on:
- Git-friendly (plain text files)
- Offline-first
- Privacy-focused (no cloud sync required)
- Fast and lightweight

## Installation

1. Download and install Bruno from: https://www.usebruno.com/downloads
2. Open Bruno
3. Click "Open Collection"
4. Navigate to this directory: `bruno/hosterizer-api`

## Collection Structure

```
hosterizer-api/
├── bruno.json                    # Collection configuration
├── environments/                 # Environment configurations
│   ├── local.bru                # Local development (default)
│   └── dev.bru                  # Development server
└── Auth Service/                # Auth service endpoints
    ├── Health Check.bru
    ├── Login - Success.bru
    ├── Login - Invalid Credentials.bru
    ├── Login - Missing Fields.bru
    ├── Login - With MFA.bru
    ├── Get Current User.bru
    ├── Get Current User - Unauthorized.bru
    ├── Refresh Token.bru
    ├── Refresh Token - Invalid.bru
    ├── MFA Setup.bru
    ├── MFA Verify.bru
    └── Logout.bru
```

## Environments

### Local (Default)
- Base URL: `http://localhost:8001`
- Use this for local development
- Requires auth service running locally

### Dev
- Base URL: `http://dev.hosterizer.com:8001`
- Use this for testing against development server

## Environment Variables

The collection uses the following environment variables:

| Variable | Description | Auto-populated |
|----------|-------------|----------------|
| `base_url` | Base URL for the API | No |
| `auth_base_url` | Auth service base URL | No |
| `access_token` | JWT access token | Yes (after login) |
| `refresh_token` | JWT refresh token | Yes (after login) |
| `user_id` | Current user ID | Yes (after login) |
| `user_uuid` | Current user UUID | Yes (after login) |
| `mfa_secret` | MFA secret key | Yes (after MFA setup) |
| `mfa_qr_code` | MFA QR code URL | Yes (after MFA setup) |

Variables marked "Yes" in Auto-populated are automatically set by post-response scripts.

## Prerequisites

Before testing, ensure:

1. **Database is running** with migrations applied:
   ```bash
   docker-compose up -d postgres
   cd backend/shared
   go run cmd/migrate/main.go up
   ```

2. **Redis is running**:
   ```bash
   docker-compose up -d redis
   ```

3. **Auth service is running**:
   ```bash
   cd backend/auth-service
   go run cmd/server/main.go
   ```

4. **Test user exists** in the database:
   ```sql
   INSERT INTO users (email, password_hash, first_name, last_name, role)
   VALUES (
     'admin@hosterizer.com',
     '$2a$12$LQv3c1yqBWVHxkd0LHAkCOYz6TtxMQJqhN8/LewY5GyYKKXqKKKKK', -- AdminPass123!
     'Admin',
     'User',
     'administrator'
   );
   ```

## Testing Workflow

### Basic Authentication Flow

1. **Health Check** - Verify service is running
2. **Login - Success** - Get access and refresh tokens
3. **Get Current User** - Verify authentication works
4. **Refresh Token** - Test token refresh
5. **Logout** - End session

### MFA Flow

1. **Login - Success** - Get authenticated
2. **MFA Setup** - Initialize MFA (get QR code)
3. Scan QR code with authenticator app (Google Authenticator, Authy, etc.)
4. **MFA Verify** - Enable MFA with code from app
5. **Logout**
6. **Login - With MFA** - Test MFA-protected login

### Error Testing

1. **Login - Invalid Credentials** - Test wrong password
2. **Login - Missing Fields** - Test validation
3. **Get Current User - Unauthorized** - Test without token
4. **Refresh Token - Invalid** - Test with invalid token

## Request Sequence

For best results, run requests in this order:

1. Health Check
2. Login - Success (saves tokens to environment)
3. Get Current User (uses saved token)
4. MFA Setup (uses saved token)
5. MFA Verify (uses saved token, requires manual code entry)
6. Refresh Token (uses saved refresh token)
7. Logout

## Tips

### Automatic Token Management
The collection automatically saves tokens after successful login:
- Access tokens are saved to `{{access_token}}`
- Refresh tokens are saved to `{{refresh_token}}`
- User info is saved to environment variables

### Testing MFA
1. Run "MFA Setup" request
2. Check the response for `qr_code_url`
3. Open the URL in a browser or scan with your phone
4. Add to your authenticator app
5. Get the 6-digit code from the app
6. Update "MFA Verify" request with the code
7. Run "MFA Verify" request

### Running All Tests
Bruno allows you to run all requests in a folder:
1. Right-click on "Auth Service" folder
2. Select "Run Folder"
3. View results in the runner

### Viewing Test Results
Each request includes automated tests that verify:
- Response status codes
- Response body structure
- Data types and values
- Business logic correctness

## Common Issues

### Connection Refused
- Ensure auth service is running on port 8001
- Check `base_url` in environment matches your setup

### 401 Unauthorized
- Run "Login - Success" first to get a valid token
- Check if token has expired (15-minute lifetime)
- Use "Refresh Token" to get a new token

### MFA Verification Fails
- Ensure you're using the current code (codes expire every 30 seconds)
- Check that you scanned the correct QR code
- Verify your device's time is synchronized

### Database Connection Errors
- Ensure PostgreSQL is running
- Verify database migrations are applied
- Check test user exists in database

## Adding New Requests

To add a new request:

1. Create a new `.bru` file in the appropriate folder
2. Use this template:

```
meta {
  name: Request Name
  type: http
  seq: 13
}

post {
  url: {{auth_base_url}}/endpoint
  body: json
  auth: bearer
}

auth:bearer {
  token: {{access_token}}
}

body:json {
  {
    "field": "value"
  }
}

docs {
  # Request Name
  
  Description of what this request does.
}

tests {
  test("should return 200 OK", function() {
    expect(res.status).to.equal(200);
  });
}
```

## Contributing

When adding new endpoints:
1. Create descriptive request names
2. Add comprehensive documentation in the `docs` section
3. Include automated tests in the `tests` section
4. Use post-response scripts to save relevant data to environment
5. Update this README with new endpoints

## Resources

- [Bruno Documentation](https://docs.usebruno.com/)
- [Bruno GitHub](https://github.com/usebruno/bruno)
- [Hosterizer API Documentation](../../docs/api/)
