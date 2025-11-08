# Hosterizer Shared Package

This package contains shared code used across all Hosterizer backend services.

## Contents

- **database/**: Database connectivity, migrations, and RLS support
- **migrations/**: SQL migration files for database schema

## Installation

```bash
cd backend/shared
go mod download
```

## Database Setup

### Quick Start

1. **Initialize the database:**

   ```bash
   # Linux/Mac
   ./scripts/init-db.sh
   
   # Windows
   .\scripts\init-db.ps1
   ```

2. **Run migrations:**

   ```bash
   cd backend/shared
   go run cmd/migrate/main.go -action=up
   ```

3. **Verify RLS (optional):**

   ```bash
   # Linux/Mac
   ./scripts/test-rls.sh
   
   # Windows
   .\scripts\test-rls.ps1
   ```

### Database Configuration

Set environment variables for database connection:

```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=hosterizer
export DB_SSLMODE=disable
```

### Using in Services

Import the shared package in your service:

```go
import (
    "github.com/hosterizer/shared/database"
)

func main() {
    // Create database connection
    cfg := database.DefaultConfig()
    db, err := database.NewPostgresDB(cfg)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // Run migrations
    if err := database.MigrateUp(db.DB, migrations, "migrations"); err != nil {
        log.Fatal(err)
    }
    
    // Use database...
}
```

## Database Schema

### Tables

1. **users** - System users with authentication
2. **customers** - Customer organizations
3. **sites** - Website deployments
4. **deployments** - Deployment execution records
5. **policies** - Cloud policies
6. **ecommerce_integrations** - Ecommerce platform connections
7. **cost_records** - Cloud cost tracking

### Row-Level Security

Multi-tenant isolation is enforced using PostgreSQL Row-Level Security (RLS):

- Customer users can only access their own data
- Administrator users can access all data
- Isolation is enforced at the database level

See [database/README.md](database/README.md) for detailed RLS usage.

## Migrations

### Migration Files

Migrations are located in `migrations/` directory:

```
migrations/
├── 000001_create_users_table.up.sql
├── 000001_create_users_table.down.sql
├── 000002_create_customers_table.up.sql
├── 000002_create_customers_table.down.sql
...
```

### Creating Migrations

1. Determine next version number
2. Create `.up.sql` and `.down.sql` files
3. Write migration and rollback SQL
4. Test locally before committing

### Migration Commands

```bash
# Run all pending migrations
go run cmd/migrate/main.go -action=up

# Rollback last migration
go run cmd/migrate/main.go -action=down

# Check current version
go run cmd/migrate/main.go -action=version
```

## Development

### Prerequisites

- Go 1.21+
- PostgreSQL 15+

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build ./...
```

## Documentation

- [Database Package](database/README.md)
- [Migrations](migrations/README.md)

## License

Proprietary
