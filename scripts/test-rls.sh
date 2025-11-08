#!/bin/bash

# Script to test Row-Level Security (RLS) for multi-tenancy
# This script runs the RLS test SQL file

set -e

# Default values
DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-5432}"
DB_USER="${DB_USER:-postgres}"
DB_PASSWORD="${DB_PASSWORD:-postgres}"
DB_NAME="${DB_NAME:-hosterizer}"

echo "Testing Row-Level Security..."
echo "Host: $DB_HOST"
echo "Port: $DB_PORT"
echo "Database: $DB_NAME"
echo ""

# Set password environment variable for psql
export PGPASSWORD=$DB_PASSWORD

# Run the test SQL script
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f "$(dirname "$0")/test-rls.sql"

echo ""
echo "RLS testing complete!"
