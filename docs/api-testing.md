# API Testing with Bruno

This document describes how to test the Hosterizer API using Bruno, an open-source API client.

## Overview

The Hosterizer project includes a comprehensive Bruno collection for testing all API endpoints. Bruno was chosen for its:

- **Git-friendly format** - All requests stored as plain text files
- **Offline-first** - No cloud sync required
- **Privacy-focused** - Your data stays local
- **Fast and lightweight** - Native application
- **Open source** - Free and community-driven

## Quick Start

See [bruno/QUICKSTART.md](../bruno/QUICKSTART.md) for a 5-minute getting started guide.

## Collection Structure

```
bruno/hosterizer-api/
├── bruno.json                    # Collection metadata
├── collection.bru                # Collection documentation
├── README.md                     # Detailed documentation
├── environments/                 # Environment configurations
│   ├── local.bru                # Local development
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

## Features

### Automated Token Management

The collection automatically manages authentication tokens:

```javascript
// After successful login, tokens are saved to environment
script:post-response {
  if (res.status === 200 && res.body.access_token) {
    bru.setEnvVar("access_token", res.body.access_token);
    bru.setEnvVar("refresh_token", res.body.refresh_token);
  }
}
```

### Automated Testing

Each request includes automated tests:

```javascript
tests {
  test("should return 200 OK", function() {
    expect(res.status).to.equal(200);
  });
  
  test("should return access token", function() {
    expect(res.body.access_token).to.be.a('string');
  });
}
```

### Comprehensive Documentation

Every request includes:
- Description of what it does
- Prerequisites
- Expected responses
- Usage examples
- Related requests

## Testing Workflows

### Basic Authentication

1. **Health Check** - Verify service is running
2. **Login - Success** - Authenticate and get tokens
3. **Get Current User** - Verify token works
4. **Refresh Token** - Test token refresh
5. **Logout** - End session

### MFA Setup and Testing

1. **Login - Success** - Get authenticated
2. **MFA Setup** - Get QR code and secret
3. Scan QR code with authenticator app
4. **MFA Verify** - Enable MFA with TOTP code
5. **Logout**
6. **Login - With MFA** - Test MFA-protected login

### Error Handling

1. **Login - Invalid Credentials** - Wrong password
2. **Login - Missing Fields** - Validation errors
3. **Get Current User - Unauthorized** - No token
4. **Refresh Token - Invalid** - Invalid token

## Environment Variables

| Variable | Description | Auto-Set |
|----------|-------------|----------|
| `base_url` | API base URL | No |
| `auth_base_url` | Auth service URL | No |
| `access_token` | JWT access token | Yes |
| `refresh_token` | JWT refresh token | Yes |
| `user_id` | Current user ID | Yes |
| `user_uuid` | Current user UUID | Yes |
| `mfa_secret` | MFA secret | Yes |
| `mfa_qr_code` | MFA QR code URL | Yes |

## Test Users

Use the provided scripts to create test users:

**Linux/Mac:**
```bash
./scripts/create-test-user.sh
```

**Windows:**
```powershell
.\scripts\create-test-user.ps1
```

**Test Credentials:**
- Administrator: `admin@hosterizer.com` / `AdminPass123!`
- Customer: `customer@hosterizer.com` / `AdminPass123!`

## Running Tests

### Single Request
1. Open request in Bruno
2. Click "Send" or press `Ctrl+Enter`
3. View response and test results

### Folder (All Requests)
1. Right-click on folder (e.g., "Auth Service")
2. Select "Run Folder"
3. View results in runner

### Collection (All Folders)
1. Right-click on collection root
2. Select "Run Collection"
3. View comprehensive results

## CI/CD Integration

Bruno can be run from the command line for CI/CD:

```bash
# Install Bruno CLI
npm install -g @usebruno/cli

# Run collection
bru run bruno/hosterizer-api --env local

# Run specific folder
bru run bruno/hosterizer-api/Auth\ Service --env local

# Output formats
bru run bruno/hosterizer-api --env local --format json
bru run bruno/hosterizer-api --env local --format junit
```

## Best Practices

### Request Organization
- Group related requests in folders
- Use descriptive names
- Number requests for logical ordering
- Include both success and error cases

### Documentation
- Add comprehensive docs to each request
- Explain prerequisites
- Document expected responses
- Link to related requests

### Testing
- Test happy path and error cases
- Verify response structure
- Check data types
- Validate business logic

### Environment Variables
- Use variables for all URLs
- Auto-save tokens and IDs
- Keep sensitive data in environment
- Document all variables

## Troubleshooting

### Connection Issues
```
Error: connect ECONNREFUSED 127.0.0.1:8001
```
**Solution:** Ensure auth service is running on port 8001

### Authentication Errors
```
401 Unauthorized
```
**Solution:** Run "Login - Success" to get a fresh token

### MFA Issues
```
400 Bad Request - Invalid MFA code
```
**Solution:** 
- Use current code (expires every 30 seconds)
- Ensure device time is synchronized
- Verify correct QR code was scanned

### Database Errors
```
Error: User not found
```
**Solution:** Run `scripts/create-test-user.sh` to create test users

## Resources

- [Bruno Official Website](https://www.usebruno.com/)
- [Bruno Documentation](https://docs.usebruno.com/)
- [Bruno GitHub](https://github.com/usebruno/bruno)
- [Collection README](../bruno/hosterizer-api/README.md)
- [Quick Start Guide](../bruno/QUICKSTART.md)

## Contributing

When adding new API endpoints:

1. Create `.bru` file in appropriate folder
2. Add comprehensive documentation
3. Include automated tests
4. Use post-response scripts for variable management
5. Update README with new endpoints
6. Test both success and error cases

## Future Additions

As new services are added to Hosterizer, create new folders:

- `Customer Service/` - Customer management endpoints
- `Site Service/` - Site deployment endpoints
- `Infrastructure Service/` - Infrastructure management
- `Policy Service/` - Policy management
- `Cost Service/` - Cost tracking and reporting
- `Ecommerce Service/` - Ecommerce integration

Each service folder should follow the same patterns established in the Auth Service folder.
