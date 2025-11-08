-- Create sites table
CREATE TABLE sites (
    id BIGSERIAL PRIMARY KEY,
    uuid UUID DEFAULT gen_random_uuid() UNIQUE NOT NULL,
    customer_id BIGINT NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    domain TEXT,
    cloud_provider TEXT NOT NULL CHECK (
        cloud_provider IN ('aws', 'azure', 'gcp', 'digitalocean', 'akamai')
    ),
    region TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'pending' CHECK (
        status IN (
            'pending',
            'provisioning',
            'active',
            'updating',
            'failed',
            'deleting',
            'deleted'
        )
    ),
    configuration JSONB NOT NULL DEFAULT '{}'::jsonb,
    infrastructure_metadata JSONB DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
-- Create indexes for sites table
CREATE INDEX idx_sites_customer ON sites(customer_id);
CREATE INDEX idx_sites_status ON sites(status);
CREATE INDEX idx_sites_cloud_provider ON sites(cloud_provider);
CREATE INDEX idx_sites_deleted ON sites(deleted_at)
WHERE deleted_at IS NOT NULL;
CREATE INDEX idx_sites_uuid ON sites(uuid);
CREATE INDEX idx_sites_customer_name ON sites(customer_id, name);
-- Create GIN indexes for JSONB columns
CREATE INDEX idx_sites_configuration ON sites USING GIN (configuration);
CREATE INDEX idx_sites_infrastructure_metadata ON sites USING GIN (infrastructure_metadata);
-- Create trigger to automatically update updated_at
CREATE TRIGGER sites_updated_at BEFORE
UPDATE ON sites FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
-- Add comments to table
COMMENT ON TABLE sites IS 'Website deployments managed by Hosterizer';
COMMENT ON COLUMN sites.cloud_provider IS 'Cloud provider: aws, azure, gcp, digitalocean, or akamai';
COMMENT ON COLUMN sites.status IS 'Site status: pending, provisioning, active, updating, failed, deleting, or deleted';
COMMENT ON COLUMN sites.configuration IS 'Site configuration parameters';
COMMENT ON COLUMN sites.infrastructure_metadata IS 'Metadata about provisioned infrastructure';
COMMENT ON COLUMN sites.deleted_at IS 'Soft delete timestamp';