# Database Migrations

This directory contains all database migrations for the Hosterizer platform.

## Migration Naming Convention

Migrations follow the format: `{version}_{description}.{up|down}.sql`

Example:
- `000001_create_users_table.up.sql`
- `000001_create_users_table.down.sql`

## Creating a New Migration

1. Determine the next version number (increment from the last migration)
2. Create two files:
   - `{version}_{description}.up.sql` - Contains the migration
   - `{version}_{description}.down.sql` - Contains the rollback

## Running Migrations

Migrations are automatically run on application startup. They can also be run manually using the migration tool.

## Migration Guidelines

- Each migration should be atomic and reversible
- Always provide a down migration
- Test migrations in development before applying to production
- Never modify existing migrations that have been applied to production
- Use transactions where appropriate
- Add comments to explain complex migrations
