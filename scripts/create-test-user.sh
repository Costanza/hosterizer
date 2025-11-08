#!/bin/bash

# Script to create test users for API testing
# Password for all test users: AdminPass123!
# Bcrypt hash: $2a$12$LQv3c1yqBWVHxkd0LHAkCOYz6TtxMQJqhN8/LewY5GyYKKXqKKKKK

set -e

# Database connection settings
DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-5432}"
DB_USER="${DB_USER:-postgres}"
DB_PASSWORD="${DB_PASSWORD:-postgres}"
DB_NAME="${DB_NAME:-hosterizer}"

echo "Creating test users in database..."
echo "Database: $DB_NAME at $DB_HOST:$DB_PORT"
echo ""

# Create test users
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME <<EOF

-- Create administrator user
INSERT INTO users (email, password_hash, first_name, last_name, role)
VALUES (
  'admin@hosterizer.com',
  '\$2a\$12\$LQv3c1yqBWVHxkd0LHAkCOYz6TtxMQJqhN8/LewY5GyYKKXqKKKKK',
  'Admin',
  'User',
  'administrator'
)
ON CONFLICT (email) DO UPDATE SET
  password_hash = EXCLUDED.password_hash,
  first_name = EXCLUDED.first_name,
  last_name = EXCLUDED.last_name,
  role = EXCLUDED.role;

-- Create customer user
INSERT INTO users (email, password_hash, first_name, last_name, role)
VALUES (
  'customer@hosterizer.com',
  '\$2a\$12\$LQv3c1yqBWVHxkd0LHAkCOYz6TtxMQJqhN8/LewY5GyYKKXqKKKKK',
  'Customer',
  'User',
  'customer'
)
ON CONFLICT (email) DO UPDATE SET
  password_hash = EXCLUDED.password_hash,
  first_name = EXCLUDED.first_name,
  last_name = EXCLUDED.last_name,
  role = EXCLUDED.role;

-- Display created users
SELECT 
  id, 
  email, 
  first_name, 
  last_name, 
  role,
  mfa_enabled,
  created_at
FROM users
WHERE email IN ('admin@hosterizer.com', 'customer@hosterizer.com')
ORDER BY email;

EOF

echo ""
echo "✅ Test users created successfully!"
echo ""
echo "Test Users:"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "Administrator:"
echo "  Email:    admin@hosterizer.com"
echo "  Password: AdminPass123!"
echo "  Role:     administrator"
echo ""
echo "Customer:"
echo "  Email:    customer@hosterizer.com"
echo "  Password: AdminPass123!"
echo "  Role:     customer"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "You can now use these credentials to test the API!"
