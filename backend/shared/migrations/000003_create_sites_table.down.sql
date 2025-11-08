-- Drop sites table and related objects
DROP TRIGGER IF EXISTS sites_updated_at ON sites;
DROP INDEX IF EXISTS idx_sites_infrastructure_metadata;
DROP INDEX IF EXISTS idx_sites_configuration;
DROP INDEX IF EXISTS idx_sites_customer_name;
DROP INDEX IF EXISTS idx_sites_uuid;
DROP INDEX IF EXISTS idx_sites_deleted;
DROP INDEX IF EXISTS idx_sites_cloud_provider;
DROP INDEX IF EXISTS idx_sites_status;
DROP INDEX IF EXISTS idx_sites_customer;
DROP TABLE IF EXISTS sites;