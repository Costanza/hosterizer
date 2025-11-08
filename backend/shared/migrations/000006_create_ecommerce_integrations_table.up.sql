-- Create ecommerce_integrations table
CREATE TABLE ecommerce_integrations (
    id BIGSERIAL PRIMARY KEY,
    uuid UUID DEFAULT gen_random_uuid() UNIQUE NOT NULL,
    site_id BIGINT NOT NULL REFERENCES sites(id) ON DELETE CASCADE,
    platform TEXT NOT NULL CHECK (
        platform IN ('shopify', 'woocommerce', 'bigcommerce')
    ),
    credentials_encrypted TEXT NOT NULL,
    configuration JSONB DEFAULT '{}'::jsonb,
    status TEXT NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'inactive', 'error')),
    last_sync_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(site_id, platform)
);
-- Create indexes for ecommerce_integrations table
CREATE INDEX idx_ecommerce_site ON ecommerce_integrations(site_id);
CREATE INDEX idx_ecommerce_platform ON ecommerce_integrations(platform);
CREATE INDEX idx_ecommerce_uuid ON ecommerce_integrations(uuid);
CREATE INDEX idx_ecommerce_status ON ecommerce_integrations(status);
CREATE INDEX idx_ecommerce_site_platform ON ecommerce_integrations(site_id, platform);
-- Create GIN index for JSONB configuration column
CREATE INDEX idx_ecommerce_configuration ON ecommerce_integrations USING GIN (configuration);
-- Create trigger to automatically update updated_at
CREATE TRIGGER ecommerce_integrations_updated_at BEFORE
UPDATE ON ecommerce_integrations FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
-- Add comments to table
COMMENT ON TABLE ecommerce_integrations IS 'Ecommerce platform integrations for sites';
COMMENT ON COLUMN ecommerce_integrations.platform IS 'Ecommerce platform: shopify, woocommerce, or bigcommerce';
COMMENT ON COLUMN ecommerce_integrations.credentials_encrypted IS 'Encrypted API credentials';
COMMENT ON COLUMN ecommerce_integrations.configuration IS 'Platform-specific configuration';
COMMENT ON COLUMN ecommerce_integrations.status IS 'Integration status: active, inactive, or error';
COMMENT ON COLUMN ecommerce_integrations.last_sync_at IS 'Last successful synchronization timestamp';