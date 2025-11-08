-- Drop cost_records table and related objects
DROP INDEX IF EXISTS idx_cost_records_breakdown;
DROP INDEX IF EXISTS idx_cost_records_customer_date;
DROP INDEX IF EXISTS idx_cost_records_cloud_provider;
DROP INDEX IF EXISTS idx_cost_records_date;
DROP INDEX IF EXISTS idx_cost_records_customer;
DROP INDEX IF EXISTS idx_cost_records_site;
DROP TABLE IF EXISTS cost_records;