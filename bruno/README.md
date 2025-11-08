# Bruno API Collections

This directory contains Bruno API collections for testing the Hosterizer platform.

## What's Inside

```
bruno/
├── hosterizer-api/              # Main API collection
│   ├── bruno.json              # Collection configuration
│   ├── collection.bru          # Collection metadata
│   ├── README.md               # Detailed documentation
│   ├── environments/           # Environment configurations
│   │   ├── local.bru          # Local development
│   │   └── dev.bru            # Development server
│   └── Auth Service/          # Auth service endpoints (12 requests)
│       ├── README.md          # Service-specific docs
│       └── *.bru              # Individual requests
├── QUICKSTART.md               # 5-minute getting started guide
└── README.md                   # This file
```

## Quick Start

1. **Install Bruno**
   - Download from: https://www.usebruno.com/downloads
   - Available for Windows, macOS, and Linux

2. **Open Collection**
   - Launch Bruno
   - Click "Open Collection"
   - Navigate to `bruno/hosterizer-api`

3. **Setup Environment**
   - Select "local" environment (top-right dropdown)
   - Verify `base_url` is `http://localhost:8001`

4. **Create Test Users**
   ```bash
   # Linux/Mac
   ./scripts/create-test-user.sh
   
   # Windows
   .\scripts\create-test-user.ps1
   ```

5. **Start Testing**
   - Run "Health Check" to verify connectivity
   - Run "Login - Success" to authenticate
   - Explore other endpoints!

## Collections

### Hosterizer API
The main collection containing all API endpoints organized by service.

**Current Services:**
- ✅ Auth Service (12 endpoints)
- 🚧 Customer Service (coming soon)
- 🚧 Site Service (coming soon)
- 🚧 Infrastructure Service (coming soon)
- 🚧 Policy Service (coming soon)
- 🚧 Cost Service (coming soon)
- 🚧 Ecommerce Service (coming soon)

## Features

### 🔐 Automatic Token Management
Tokens are automatically saved after login and used in subsequent requests.

### ✅ Automated Testing
Every request includes automated tests to verify responses.

### 📚 Comprehensive Documentation
Each request has detailed documentation, prerequisites, and examples.

### 🌍 Multiple Environments
Switch between local, dev, staging, and production environments.

### 🔄 Request Chaining
Requests automatically pass data to subsequent requests via environment variables.

## Documentation

- **[Quick Start Guide](QUICKSTART.md)** - Get started in 5 minutes
- **[API Testing Guide](../docs/api-testing.md)** - Comprehensive testing documentation
- **[Collection README](hosterizer-api/README.md)** - Detailed collection documentation
- **[Auth Service README](hosterizer-api/Auth%20Service/README.md)** - Auth endpoints documentation

## Test Credentials

After running the test user creation script:

**Administrator:**
- Email: `admin@hosterizer.com`
- Password: `AdminPass123!`
- Role: `administrator`

**Customer:**
- Email: `customer@hosterizer.com`
- Password: `AdminPass123!`
- Role: `customer`

## Common Tasks

### Run All Tests in a Service
1. Right-click on service folder (e.g., "Auth Service")
2. Select "Run Folder"
3. View results in runner

### Export Request as cURL
1. Right-click on request
2. Select "Copy as cURL"
3. Paste in terminal

### View Environment Variables
1. Click environment dropdown (top-right)
2. Select current environment
3. View/edit variables

### Add New Request
1. Right-click on folder
2. Select "New Request"
3. Configure request details
4. Add documentation and tests

## CI/CD Integration

Bruno can be run from the command line:

```bash
# Install Bruno CLI
npm install -g @usebruno/cli

# Run entire collection
bru run bruno/hosterizer-api --env local

# Run specific folder
bru run bruno/hosterizer-api/Auth\ Service --env local

# Output as JSON
bru run bruno/hosterizer-api --env local --format json

# Output as JUnit XML
bru run bruno/hosterizer-api --env local --format junit
```

## Why Bruno?

We chose Bruno over alternatives like Postman because:

1. **Git-Friendly** - Plain text files that work great with version control
2. **Offline-First** - No cloud sync required, works completely offline
3. **Privacy-Focused** - Your API data stays on your machine
4. **Open Source** - Free and community-driven
5. **Fast** - Native application, not Electron-based
6. **Simple** - Clean UI without unnecessary complexity

## Contributing

When adding new endpoints:

1. Create `.bru` file in appropriate service folder
2. Follow naming convention: `Action - Variant.bru`
3. Add comprehensive documentation in `docs` section
4. Include automated tests in `tests` section
5. Use post-response scripts to save relevant data
6. Update service README with new endpoint
7. Test both success and error cases

## Support

### Issues
- Check service logs for errors
- Verify services are running
- Ensure database migrations are applied
- Confirm test users exist

### Resources
- [Bruno Documentation](https://docs.usebruno.com/)
- [Bruno GitHub](https://github.com/usebruno/bruno)
- [Hosterizer Documentation](../docs/)

## Roadmap

- [x] Auth Service endpoints
- [ ] Customer Service endpoints
- [ ] Site Service endpoints
- [ ] Infrastructure Service endpoints
- [ ] Policy Service endpoints
- [ ] Cost Service endpoints
- [ ] Ecommerce Service endpoints
- [ ] WebSocket testing support
- [ ] GraphQL endpoint support
- [ ] Performance testing scenarios
- [ ] Load testing configurations

## License

Same as main project - Proprietary
