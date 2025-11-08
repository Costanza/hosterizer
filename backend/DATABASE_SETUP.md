# Hosterizer Database Setup Guide

This guide provides instructions for setting up and managing the Hosterizer PostgreSQL database.

## Overview

The Hosterizer platform uses PostgreSQL 15+ with the following features:
- Connection pooling for optimal performance
- Automated migrations using golang-migrate
- Row-level security (RLS) for multi-tenant isolation
- JSONB columns for flexible data storage
- Comprehensive indexing for query performance

## Quick Start

### 1. Prerequisites

Ensure you have the following installed:
- PostgreSQL 15 or higher
- Go 1.21 or higher
- psql command-line tool

### 2. Initialize Database

Run the initialization script to create the database and required extensions:

**Linux/Mac:**
```bash
./scripts/init-db.sh
```

**Windows PowerShell:**
```powershell
.\scripts\init-db.ps1
```

This script will:
- Create the `hosterizer` database
- Enable required PostgreSQL extensions (uuid-ossp, pgcrypto)
- Create the `app_user` role
- Set up necessary permissions
- Create helper functions (e.g., update_updated_at_column)

### 3. Run Migrations

Execute all database migrations:

```bash
cd backend/shared
go run cmd/migrate/main.go -action=up
```

This will create all tables, indexes, constraints, and RLS policies.

### 4. Verify Setup (Optional)

Test row-level security to ensure tenant isolation:

**Linux/Mac:**
```bash
./scripts/test-rls.sh
```

**Windows PowerShell:**
```powershell
.\scripts\test-rls.ps1
```

## Database Schema

### Core Tables

#### users
Stores system users with authentication credentials.
- Supports both administrator and customer roles
- Includes MFA support
- Tracks failed login attempts and account lockouts

#### customers
Stores customer organizations.
- Links to owner user
- Supports white-label configuration (JSONB)
- Tracks customer tier (standard/premium) and status

#### sites
Stores website deployments.
- Links to customer
- Supports multiple cloud providers (AWS, Azure, GCP, Digital Ocean, Akamai)
- Tracks deployment status and configuration
- Implements soft delete with deleted_at timestamp

#### deployments
Tracks deployment execution history.
- Links to site
- Stores Terraform plan and output
- Tracks deployment status and timing

### Supporting Tables

#### policies
Cloud policies for governance.
- Supports resource limits, security, and cost policies
- Stores rules in JSONB format
- Can be enabled/disabled

#### ecommerce_integrations
Ecommerce platform connections.
- Supports Shopify, WooCommerce, BigCommerce
- Stores encrypted credentials
- Tracks integration status and last sync

#### cost_records
Daily cost tracking from cloud providers.
- Links to both site and customer
- Stores cost breakdown by resource type (JSONB)
- Unique constraint on site_id + cost_date

## Row-Level Security (RLS)

### Overview

RLS ensures complete data isolation between customers:
- Customer users can only access their own data
- Administrator users can access all data
- Isolation is enforced at the database level

### Affected Tables

RLS is enabled on:
- sites
- deployments
- ecommerce_integrations
- cost_records
- customers

### Session Variables

RLS policies use these PostgreSQL session variables:
- `app.current_customer_id`: Customer ID for the session
- `app.current_user_role`: User role (administrator or customer)

### Usage in Application Code

```go
import "github.com/hosterizer/shared/database"

// For customer users
tc := database.TenantContext{
    CustomerID: 123,
    UserRole:   "customer",
}

err := database.WithTenantContext(ctx, db.DB, tc, func(tx *sql.Tx) error {
    // Queries here are automatically filtered by customer_id
    rows, err := tx.QueryContext(ctx, "SELECT * FROM sites")
    // ...
    return nil
})

// For administrator users
tc := database.TenantContext{
    UserRole: "administrator",
}

err := database.WithTenantContext(ctx, db.DB, tc, func(tx *sql.Tx) error {
    // Queries here can access all data
    rows, err := tx.QueryContext(ctx, "SELECT * FROM sites")
    // ...
    return nil
})
```

## Connection Pooling

The database package configures connection pooling with these defaults:

| Setting | Default | Description |
|---------|---------|-------------|
| MaxOpenConns | 25 | Maximum number of open connections |
| MaxIdleConns | 5 | Maximum number of idle connections |
| ConnMaxLifetime | 5 minutes | Maximum lifetime of a connection |
| ConnMaxIdleTime | 10 minutes | Maximum idle time before closing |

Adjust these values based on your workload:

```go
cfg := database.Config{
    Host:            "localhost",
    Port:            5432,
    User:            "postgres",
    Password:        "postgres",
    Database:        "hosterizer",
    SSLMode:         "disable",
    MaxOpenConns:    50,  // Increase for high traffic
    MaxIdleConns:    10,
    ConnMaxLifetime: 5 * time.Minute,
    ConnMaxIdleTime: 10 * time.Minute,
}
```

## Migration Management

### Migration Files

Migrations are stored in `backend/shared/migrations/`:

```
000001_create_users_table.up.sql
000001_create_users_table.down.sql
000002_create_customers_table.up.sql
000002_create_customers_table.down.sql
...
```

### Creating New Migrations

1. Determine the next version number (e.g., 000010)
2. Create two files:
   - `{version}_{description}.up.sql` - Migration
   - `{version}_{description}.down.sql` - Rollback

Example:
```sql
-- 000010_add_site_tags.up.sql
ALTER TABLE sites ADD COLUMN tags TEXT[] DEFAULT ARRAY[]::TEXT[];
CREATE INDEX idx_sites_tags ON sites USING GIN (tags);

-- 000010_add_site_tags.down.sql
DROP INDEX IF EXISTS idx_sites_tags;
ALTER TABLE sites DROP COLUMN IF EXISTS tags;
```

### Migration Commands

```bash
cd backend/shared

# Run all pending migrations
go run cmd/migrate/main.go -action=up

# Rollback last migration
go run cmd/migrate/main.go -action=down

# Check current migration version
go run cmd/migrate/main.go -action=version
```

### Migration Best Practices

- Always provide both up and down migrations
- Test migrations in development before production
- Never modify existing migrations that have been applied to production
- Use transactions where appropriate
- Add comments to explain complex migrations
- Keep migrations atomic and focused

## Environment Variables

Configure database connection using environment variables:

```bash
# Database connection
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=hosterizer
export DB_SSLMODE=disable

# For production, use SSL
export DB_SSLMODE=require
```

## Troubleshooting

### Connection Issues

**Problem:** Cannot connect to database

**Solution:**
1. Verify PostgreSQL is running: `pg_isready`
2. Check connection parameters
3. Verify user has necessary permissions
4. Check firewall settings

### Migration Failures

**Problem:** Migration fails with "dirty" state

**Solution:**
1. Check migration version: `go run cmd/migrate/main.go -action=version`
2. Manually fix the database issue
3. Update schema_migrations table to mark as clean
4. Re-run migrations

### RLS Issues

**Problem:** Queries return no results for customer users

**Solution:**
1. Verify session variables are set correctly
2. Check that customer_id matches the data
3. Verify RLS policies are enabled: `SELECT * FROM pg_policies;`
4. Run RLS test script to validate setup

### Performance Issues

**Problem:** Slow query performance

**Solution:**
1. Check query execution plan: `EXPLAIN ANALYZE SELECT ...`
2. Verify indexes are being used
3. Adjust connection pool settings
4. Consider adding additional indexes
5. Monitor with `pg_stat_statements`

## Monitoring

### Key Metrics to Monitor

- Connection pool utilization
- Query performance (slow queries)
- Database size and growth
- Index usage
- Replication lag (if using replicas)

### Useful Queries

```sql
-- Check connection pool usage
SELECT count(*) FROM pg_stat_activity WHERE datname = 'hosterizer';

-- Find slow queries
SELECT query, mean_exec_time, calls 
FROM pg_stat_statements 
ORDER BY mean_exec_time DESC 
LIMIT 10;

-- Check table sizes
SELECT 
    schemaname,
    tablename,
    pg_size_pretty(pg_total_relation_size(schemaname||'.'||tablename)) AS size
FROM pg_tables
WHERE schemaname = 'public'
ORDER BY pg_total_relation_size(schemaname||'.'||tablename) DESC;

-- Check index usage
SELECT 
    schemaname,
    tablename,
    indexname,
    idx_scan,
    idx_tup_read,
    idx_tup_fetch
FROM pg_stat_user_indexes
ORDER BY idx_scan DESC;
```

## Backup and Recovery

### Backup

```bash
# Full database backup
pg_dump -h localhost -U postgres -d hosterizer -F c -f hosterizer_backup.dump

# Schema only
pg_dump -h localhost -U postgres -d hosterizer -s -f hosterizer_schema.sql

# Data only
pg_dump -h localhost -U postgres -d hosterizer -a -f hosterizer_data.sql
```

### Restore

```bash
# Restore from custom format
pg_restore -h localhost -U postgres -d hosterizer -c hosterizer_backup.dump

# Restore from SQL file
psql -h localhost -U postgres -d hosterizer -f hosterizer_backup.sql
```

## Production Considerations

### Security

1. **Use SSL/TLS**: Set `DB_SSLMODE=require` or `verify-full`
2. **Strong passwords**: Use complex passwords for database users
3. **Least privilege**: Grant only necessary permissions
4. **Network security**: Restrict database access to application servers
5. **Audit logging**: Enable PostgreSQL audit logging

### Performance

1. **Connection pooling**: Use PgBouncer or similar for connection pooling
2. **Read replicas**: Set up read replicas for read-heavy workloads
3. **Monitoring**: Use tools like pg_stat_statements, pgBadger
4. **Regular maintenance**: Run VACUUM and ANALYZE regularly
5. **Index optimization**: Monitor and optimize indexes

### High Availability

1. **Replication**: Set up streaming replication
2. **Failover**: Implement automatic failover with tools like Patroni
3. **Backups**: Automated daily backups with point-in-time recovery
4. **Monitoring**: Set up alerts for replication lag and failures

## Additional Resources

- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [golang-migrate Documentation](https://github.com/golang-migrate/migrate)
- [PostgreSQL Row-Level Security](https://www.postgresql.org/docs/current/ddl-rowsecurity.html)
- [Shared Package README](backend/shared/README.md)
- [Database Package README](backend/shared/database/README.md)
