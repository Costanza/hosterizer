# Hosterizer Database Package

This package provides database connectivity, connection pooling, and migration management for the Hosterizer platform.

## Features

- PostgreSQL connection with configurable connection pooling
- Database migration management using golang-migrate
- Row-level security (RLS) support for multi-tenant isolation
- Tenant context management for RLS policies

## Usage

### Creating a Database Connection

```go
import "github.com/hosterizer/shared/database"

// Create configuration
cfg := database.Config{
    Host:            "localhost",
    Port:            5432,
    User:            "postgres",
    Password:        "postgres",
    Database:        "hosterizer",
    SSLMode:         "disable",
    MaxOpenConns:    25,
    MaxIdleConns:    5,
    ConnMaxLifetime: 5 * time.Minute,
    ConnMaxIdleTime: 10 * time.Minute,
}

// Connect to database
db, err := database.NewPostgresDB(cfg)
if err != nil {
    log.Fatal(err)
}
defer db.Close()
```

### Running Migrations

Migrations are embedded in the application and run automatically on startup:

```go
import (
    "embed"
    "github.com/hosterizer/shared/database"
)

//go:embed migrations/*.sql
var migrations embed.FS

func main() {
    db, err := database.NewPostgresDB(cfg)
    if err != nil {
        log.Fatal(err)
    }
    
    // Run migrations
    if err := database.MigrateUp(db.DB, migrations, "migrations"); err != nil {
        log.Fatal(err)
    }
}
```

### Using Tenant Context for RLS

For customer users, set the tenant context to enforce row-level security:

```go
import "github.com/hosterizer/shared/database"

// Set tenant context for a customer user
tc := database.TenantContext{
    CustomerID: 123,
    UserRole:   "customer",
}

// Execute queries within tenant context
err := database.WithTenantContext(ctx, db.DB, tc, func(tx *sql.Tx) error {
    // All queries in this transaction will be filtered by customer_id
    rows, err := tx.QueryContext(ctx, "SELECT * FROM sites")
    // ... process rows
    return nil
})
```

For administrator users:

```go
tc := database.TenantContext{
    UserRole: "administrator",
}

// Administrators can access all data
err := database.WithTenantContext(ctx, db.DB, tc, func(tx *sql.Tx) error {
    // All queries in this transaction will bypass RLS
    rows, err := tx.QueryContext(ctx, "SELECT * FROM sites")
    // ... process rows
    return nil
})
```

## Database Schema

The database includes the following tables:

### Core Tables
- **users**: System users with authentication credentials
- **customers**: Organizations or individuals using Hosterizer
- **sites**: Website deployments managed by Hosterizer
- **deployments**: Deployment requests and execution status

### Supporting Tables
- **policies**: Cloud policies for resource limits, security, and cost controls
- **ecommerce_integrations**: Ecommerce platform integrations
- **cost_records**: Daily cost records from cloud providers

## Row-Level Security

RLS is enabled on the following tables to ensure tenant isolation:
- sites
- deployments
- ecommerce_integrations
- cost_records
- customers

### RLS Policies

- **Customer users**: Can only access data belonging to their customer_id
- **Administrator users**: Can access all data (bypass RLS)

### Session Variables

RLS policies use PostgreSQL session variables:
- `app.current_customer_id`: The customer ID for the current session
- `app.current_user_role`: The user role (administrator or customer)

## Migration Management

### Creating a New Migration

1. Determine the next version number (e.g., 000010)
2. Create two files in `backend/shared/migrations/`:
   - `{version}_{description}.up.sql` - Migration
   - `{version}_{description}.down.sql` - Rollback

Example:
```
000010_add_site_tags.up.sql
000010_add_site_tags.down.sql
```

### Running Migrations Manually

Use the migration CLI tool:

```bash
cd backend/shared
go run cmd/migrate/main.go -action=up
go run cmd/migrate/main.go -action=version
go run cmd/migrate/main.go -action=down
```

## Database Initialization

### Using Scripts

Initialize the database using the provided scripts:

**Linux/Mac:**
```bash
./scripts/init-db.sh
```

**Windows:**
```powershell
.\scripts\init-db.ps1
```

### Environment Variables

Configure database connection using environment variables:
- `DB_HOST`: Database host (default: localhost)
- `DB_PORT`: Database port (default: 5432)
- `DB_USER`: Database user (default: postgres)
- `DB_PASSWORD`: Database password (default: postgres)
- `DB_NAME`: Database name (default: hosterizer)
- `DB_SSLMODE`: SSL mode (default: disable)

## Testing RLS

Test row-level security using the provided test scripts:

**Linux/Mac:**
```bash
./scripts/test-rls.sh
```

**Windows:**
```powershell
.\scripts\test-rls.ps1
```

This will run a series of queries to verify that:
- Customers can only see their own data
- Administrators can see all data
- Tenant isolation is properly enforced

## Connection Pooling

The database package configures connection pooling with the following defaults:
- **MaxOpenConns**: 25 (maximum number of open connections)
- **MaxIdleConns**: 5 (maximum number of idle connections)
- **ConnMaxLifetime**: 5 minutes (maximum lifetime of a connection)
- **ConnMaxIdleTime**: 10 minutes (maximum idle time before closing)

These values can be adjusted based on your workload requirements.

## Health Checks

Check database health:

```go
ctx := context.Background()
if err := db.HealthCheck(ctx); err != nil {
    log.Printf("Database health check failed: %v", err)
}
```
