# Hosterizer Backend Quick Start

Get the Hosterizer backend up and running in minutes.

## Prerequisites

- PostgreSQL 15+
- Go 1.21+
- Git

## Step 1: Clone and Setup

```bash
# Navigate to backend directory
cd backend

# Install Go dependencies for shared package
cd shared
go mod download
cd ..
```

## Step 2: Start PostgreSQL

Ensure PostgreSQL is running:

```bash
# Check if PostgreSQL is running
pg_isready

# If not running, start it (varies by OS)
# macOS (Homebrew):
brew services start postgresql@15

# Linux (systemd):
sudo systemctl start postgresql

# Windows:
# Start PostgreSQL service from Services app
```

## Step 3: Initialize Database

Run the initialization script:

**Linux/Mac:**
```bash
./scripts/init-db.sh
```

**Windows PowerShell:**
```powershell
.\scripts\init-db.ps1
```

**Custom Configuration:**
```bash
# Set environment variables for custom configuration
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=your_password
export DB_NAME=hosterizer

# Then run the script
./scripts/init-db.sh
```

## Step 4: Run Migrations

```bash
cd shared
go run cmd/migrate/main.go -action=up
```

Expected output:
```
Connected to database: postgres@localhost:5432/hosterizer
Running migrations up...
Migrations completed successfully
```

## Step 5: Verify Setup

Check migration version:
```bash
go run cmd/migrate/main.go -action=version
```

Test row-level security (optional):
```bash
cd ../..
./scripts/test-rls.sh  # or .ps1 on Windows
```

## Step 6: Start a Service

Example with auth-service:

```bash
cd auth-service
go run cmd/server/main.go
```

## Environment Variables

Create a `.env` file in the backend directory:

```bash
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=hosterizer
DB_SSLMODE=disable

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379

# Service Ports
AUTH_SERVICE_PORT=8001
CUSTOMER_SERVICE_PORT=8002
SITE_SERVICE_PORT=8003
INFRASTRUCTURE_SERVICE_PORT=8004
POLICY_SERVICE_PORT=8005
ECOMMERCE_SERVICE_PORT=8006
```

## Common Commands

### Database Operations

```bash
# Initialize database
./scripts/init-db.sh

# Run migrations
cd shared && go run cmd/migrate/main.go -action=up

# Check migration version
cd shared && go run cmd/migrate/main.go -action=version

# Rollback last migration
cd shared && go run cmd/migrate/main.go -action=down

# Test RLS
./scripts/test-rls.sh
```

### Development

```bash
# Run a service
cd <service-name>
go run cmd/server/main.go

# Run tests
go test ./...

# Format code
go fmt ./...

# Lint code
golangci-lint run
```

### Database Access

```bash
# Connect to database
psql -h localhost -U postgres -d hosterizer

# Useful queries
# List all tables
\dt

# Describe a table
\d users

# Check RLS policies
SELECT * FROM pg_policies;

# View current migration version
SELECT * FROM schema_migrations;
```

## Troubleshooting

### Cannot connect to database

**Error:** `failed to ping database: connection refused`

**Solution:**
1. Verify PostgreSQL is running: `pg_isready`
2. Check connection parameters in environment variables
3. Verify PostgreSQL is listening on the correct port: `netstat -an | grep 5432`

### Migration fails

**Error:** `migration failed: relation already exists`

**Solution:**
1. Check current version: `go run cmd/migrate/main.go -action=version`
2. If dirty, manually fix the issue in the database
3. Update schema_migrations table if needed

### RLS test fails

**Error:** `permission denied for table sites`

**Solution:**
1. Verify RLS policies are created: `SELECT * FROM pg_policies;`
2. Ensure session variables are set correctly
3. Re-run migration 000008

### Port already in use

**Error:** `bind: address already in use`

**Solution:**
1. Check what's using the port: `lsof -i :8001` (Mac/Linux) or `netstat -ano | findstr :8001` (Windows)
2. Kill the process or use a different port
3. Update the PORT environment variable

## Next Steps

1. **Implement Services**: Start with auth-service (Task 3)
2. **Add Domain Models**: Create domain models in each service
3. **Implement Repositories**: Create repository layer for database access
4. **Build APIs**: Implement REST API endpoints
5. **Add Tests**: Write unit and integration tests

## Resources

- [Database Setup Guide](DATABASE_SETUP.md)
- [Shared Package README](shared/README.md)
- [Database Package README](shared/database/README.md)
- [Requirements Document](../.kiro/specs/hosterizer/requirements.md)
- [Design Document](../.kiro/specs/hosterizer/design.md)

## Getting Help

- Check the documentation in `backend/shared/database/README.md`
- Review migration files in `backend/shared/migrations/`
- Consult the design document for architecture details
- Review the implementation summary in `backend/shared/IMPLEMENTATION_SUMMARY.md`

## Development Workflow

1. Pull latest changes
2. Run migrations: `cd shared && go run cmd/migrate/main.go -action=up`
3. Make your changes
4. Write tests
5. Run tests: `go test ./...`
6. Format code: `go fmt ./...`
7. Commit and push

Happy coding! 🚀
