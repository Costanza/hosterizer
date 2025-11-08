# Bruno API Testing - Quick Start Guide

Get started testing the Hosterizer API in 5 minutes!

## Step 1: Install Bruno

Download and install Bruno from: https://www.usebruno.com/downloads

Available for:
- Windows
- macOS
- Linux

## Step 2: Start Required Services

```bash
# Start PostgreSQL and Redis
docker-compose up -d postgres redis

# Apply database migrations
cd backend/shared
go run cmd/migrate/main.go up

# Start the auth service
cd ../auth-service
go run cmd/server/main.go
```

## Step 3: Create Test User

Run this SQL to create a test user:

```sql
-- Connect to the database
psql -h localhost -U postgres -d hosterizer

-- Create test user (password: AdminPass123!)
INSERT INTO users (email, password_hash, first_name, last_name, role)
VALUES (
  'admin@hosterizer.com',
  '$2a$12$LQv3c1yqBWVHxkd0LHAkCOYz6TtxMQJqhN8/LewY5GyYKKXqKKKKK',
  'Admin',
  'User',
  'administrator'
);
```

Or use the provided script:

```bash
cd scripts
./create-test-user.sh
```

## Step 4: Open Collection in Bruno

1. Open Bruno
2. Click "Open Collection"
3. Navigate to: `bruno/hosterizer-api`
4. Click "Select Folder"

## Step 5: Run Your First Request

1. In Bruno, expand "Auth Service" folder
2. Click on "Health Check"
3. Click the "Send" button (or press Ctrl+Enter)
4. You should see: `200 OK` with body `"OK"`

## Step 6: Test Authentication

1. Click on "Login - Success"
2. Click "Send"
3. You should see:
   - Status: `200 OK`
   - `access_token` and `refresh_token` in response
   - User information

The tokens are automatically saved to environment variables!

## Step 7: Test Protected Endpoint

1. Click on "Get Current User"
2. Click "Send"
3. You should see your user information

Notice the request automatically uses the `{{access_token}}` from the previous login!

## What's Next?

### Test MFA Flow
1. Run "MFA Setup"
2. Scan the QR code with Google Authenticator or Authy
3. Run "MFA Verify" with the code from your app
4. Try "Login - With MFA"

### Test Error Cases
- "Login - Invalid Credentials"
- "Login - Missing Fields"
- "Get Current User - Unauthorized"
- "Refresh Token - Invalid"

### Run All Tests
Right-click on "Auth Service" folder and select "Run Folder" to run all requests sequentially.

## Environment Variables

The collection uses these variables (auto-populated after login):

- `access_token` - Your JWT access token
- `refresh_token` - Your JWT refresh token
- `user_id` - Your user ID
- `user_uuid` - Your user UUID

You can view/edit these in Bruno:
1. Click the environment dropdown (top-right)
2. Select "local"
3. View the variables

## Troubleshooting

### "Connection refused"
- Make sure auth service is running: `go run cmd/server/main.go`
- Check it's on port 8001: `curl http://localhost:8001/health`

### "Invalid credentials"
- Verify test user exists in database
- Check password is exactly: `AdminPass123!`
- Email is: `admin@hosterizer.com`

### "Unauthorized" errors
- Run "Login - Success" first to get a token
- Tokens expire after 15 minutes - login again if needed

## Tips

- **Keyboard Shortcuts**: 
  - `Ctrl+Enter` (or `Cmd+Enter` on Mac) to send request
  - `Ctrl+/` to toggle sidebar
  
- **View Tests**: Click the "Tests" tab to see automated test results

- **View Docs**: Click the "Docs" tab to see request documentation

- **Copy as cURL**: Right-click request → "Copy as cURL"

## Need Help?

- Check the full README: `bruno/hosterizer-api/README.md`
- View Bruno docs: https://docs.usebruno.com/
- Check auth service logs for errors

Happy testing! 🚀
