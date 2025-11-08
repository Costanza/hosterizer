-- Database initialization script for Hosterizer
-- This script creates the database and necessary extensions
-- Create database (run this as postgres superuser)
-- CREATE DATABASE hosterizer;
-- Connect to hosterizer database
\ c hosterizer;
-- Enable required extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
-- Create application user role (if not exists)
DO $$ BEGIN IF NOT EXISTS (
    SELECT
    FROM pg_roles
    WHERE rolname = 'app_user'
) THEN CREATE ROLE app_user;
END IF;
END $$;
-- Grant necessary permissions
GRANT CONNECT ON DATABASE hosterizer TO app_user;
GRANT USAGE ON SCHEMA public TO app_user;
GRANT CREATE ON SCHEMA public TO app_user;
-- Set default privileges for future tables
ALTER DEFAULT PRIVILEGES IN SCHEMA public
GRANT SELECT,
    INSERT,
    UPDATE,
    DELETE ON TABLES TO app_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public
GRANT USAGE,
    SELECT ON SEQUENCES TO app_user;
-- Create a function to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column() RETURNS TRIGGER AS $$ BEGIN NEW.updated_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- Log initialization
DO $$ BEGIN RAISE NOTICE 'Database initialized successfully';
END $$;