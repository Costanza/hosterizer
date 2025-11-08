-- Drop customers table and related objects
DROP TRIGGER IF EXISTS customers_updated_at ON customers;
DROP INDEX IF EXISTS idx_customers_metadata;
DROP INDEX IF EXISTS idx_customers_white_label_config;
DROP INDEX IF EXISTS idx_customers_tier;
DROP INDEX IF EXISTS idx_customers_uuid;
DROP INDEX IF EXISTS idx_customers_owner;
DROP INDEX IF EXISTS idx_customers_status;
DROP TABLE IF EXISTS customers;