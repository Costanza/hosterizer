-- Drop ecommerce_integrations table and related objects
DROP TRIGGER IF EXISTS ecommerce_integrations_updated_at ON ecommerce_integrations;
DROP INDEX IF EXISTS idx_ecommerce_configuration;
DROP INDEX IF EXISTS idx_ecommerce_site_platform;
DROP INDEX IF EXISTS idx_ecommerce_status;
DROP INDEX IF EXISTS idx_ecommerce_uuid;
DROP INDEX IF EXISTS idx_ecommerce_platform;
DROP INDEX IF EXISTS idx_ecommerce_site;
DROP TABLE IF EXISTS ecommerce_integrations;