-- Enable Row-Level Security (RLS) for multi-tenant isolation
-- Enable RLS on sites table
ALTER TABLE sites ENABLE ROW LEVEL SECURITY;
-- Create policy for sites table
-- This policy ensures users can only access sites belonging to their customer
CREATE POLICY customer_sites_policy ON sites FOR ALL TO app_user USING (
    customer_id = current_setting('app.current_customer_id', true)::BIGINT
);
-- Create policy for administrators (bypass RLS)
CREATE POLICY admin_sites_policy ON sites FOR ALL TO app_user USING (
    current_setting('app.current_user_role', true) = 'administrator'
);
-- Enable RLS on deployments table
ALTER TABLE deployments ENABLE ROW LEVEL SECURITY;
-- Create policy for deployments table
-- Users can only access deployments for sites they own
CREATE POLICY customer_deployments_policy ON deployments FOR ALL TO app_user USING (
    site_id IN (
        SELECT id
        FROM sites
        WHERE customer_id = current_setting('app.current_customer_id', true)::BIGINT
    )
);
-- Create policy for administrators on deployments
CREATE POLICY admin_deployments_policy ON deployments FOR ALL TO app_user USING (
    current_setting('app.current_user_role', true) = 'administrator'
);
-- Enable RLS on ecommerce_integrations table
ALTER TABLE ecommerce_integrations ENABLE ROW LEVEL SECURITY;
-- Create policy for ecommerce_integrations table
CREATE POLICY customer_ecommerce_policy ON ecommerce_integrations FOR ALL TO app_user USING (
    site_id IN (
        SELECT id
        FROM sites
        WHERE customer_id = current_setting('app.current_customer_id', true)::BIGINT
    )
);
-- Create policy for administrators on ecommerce_integrations
CREATE POLICY admin_ecommerce_policy ON ecommerce_integrations FOR ALL TO app_user USING (
    current_setting('app.current_user_role', true) = 'administrator'
);
-- Enable RLS on cost_records table
ALTER TABLE cost_records ENABLE ROW LEVEL SECURITY;
-- Create policy for cost_records table
CREATE POLICY customer_cost_records_policy ON cost_records FOR ALL TO app_user USING (
    customer_id = current_setting('app.current_customer_id', true)::BIGINT
);
-- Create policy for administrators on cost_records
CREATE POLICY admin_cost_records_policy ON cost_records FOR ALL TO app_user USING (
    current_setting('app.current_user_role', true) = 'administrator'
);
-- Enable RLS on customers table
ALTER TABLE customers ENABLE ROW LEVEL SECURITY;
-- Create policy for customers table
-- Customers can only see their own record
CREATE POLICY customer_own_record_policy ON customers FOR ALL TO app_user USING (
    id = current_setting('app.current_customer_id', true)::BIGINT
    OR current_setting('app.current_user_role', true) = 'administrator'
);
-- Add comments
COMMENT ON POLICY customer_sites_policy ON sites IS 'Customers can only access their own sites';
COMMENT ON POLICY admin_sites_policy ON sites IS 'Administrators can access all sites';
COMMENT ON POLICY customer_deployments_policy ON deployments IS 'Customers can only access deployments for their sites';
COMMENT ON POLICY admin_deployments_policy ON deployments IS 'Administrators can access all deployments';
COMMENT ON POLICY customer_ecommerce_policy ON ecommerce_integrations IS 'Customers can only access integrations for their sites';
COMMENT ON POLICY admin_ecommerce_policy ON ecommerce_integrations IS 'Administrators can access all integrations';
COMMENT ON POLICY customer_cost_records_policy ON cost_records IS 'Customers can only access their own cost records';
COMMENT ON POLICY admin_cost_records_policy ON cost_records IS 'Administrators can access all cost records';
COMMENT ON POLICY customer_own_record_policy ON customers IS 'Customers can only see their own customer record';