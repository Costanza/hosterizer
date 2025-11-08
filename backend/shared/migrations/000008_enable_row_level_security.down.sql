-- Disable Row-Level Security and drop policies
-- Drop policies and disable RLS on customers table
DROP POLICY IF EXISTS customer_own_record_policy ON customers;
ALTER TABLE customers DISABLE ROW LEVEL SECURITY;
-- Drop policies and disable RLS on cost_records table
DROP POLICY IF EXISTS admin_cost_records_policy ON cost_records;
DROP POLICY IF EXISTS customer_cost_records_policy ON cost_records;
ALTER TABLE cost_records DISABLE ROW LEVEL SECURITY;
-- Drop policies and disable RLS on ecommerce_integrations table
DROP POLICY IF EXISTS admin_ecommerce_policy ON ecommerce_integrations;
DROP POLICY IF EXISTS customer_ecommerce_policy ON ecommerce_integrations;
ALTER TABLE ecommerce_integrations DISABLE ROW LEVEL SECURITY;
-- Drop policies and disable RLS on deployments table
DROP POLICY IF EXISTS admin_deployments_policy ON deployments;
DROP POLICY IF EXISTS customer_deployments_policy ON deployments;
ALTER TABLE deployments DISABLE ROW LEVEL SECURITY;
-- Drop policies and disable RLS on sites table
DROP POLICY IF EXISTS admin_sites_policy ON sites;
DROP POLICY IF EXISTS customer_sites_policy ON sites;
ALTER TABLE sites DISABLE ROW LEVEL SECURITY;