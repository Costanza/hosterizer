-- Create customers table
CREATE TABLE customers (
    id BIGSERIAL PRIMARY KEY,
    uuid UUID DEFAULT gen_random_uuid() UNIQUE NOT NULL,
    name TEXT NOT NULL,
    tier TEXT NOT NULL DEFAULT 'standard' CHECK (tier IN ('standard', 'premium')),
    status TEXT NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'inactive', 'suspended')),
    owner_user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    white_label_config JSONB DEFAULT '{}'::jsonb,
    metadata JSONB DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- Create indexes for customers table
CREATE INDEX idx_customers_status ON customers(status);
CREATE INDEX idx_customers_owner ON customers(owner_user_id);
CREATE INDEX idx_customers_uuid ON customers(uuid);
CREATE INDEX idx_customers_tier ON customers(tier);
-- Create GIN index for JSONB columns
CREATE INDEX idx_customers_white_label_config ON customers USING GIN (white_label_config);
CREATE INDEX idx_customers_metadata ON customers USING GIN (metadata);
-- Create trigger to automatically update updated_at
CREATE TRIGGER customers_updated_at BEFORE
UPDATE ON customers FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
-- Add comments to table
COMMENT ON TABLE customers IS 'Organizations or individuals who use Hosterizer to host websites';
COMMENT ON COLUMN customers.tier IS 'Customer tier: standard or premium';
COMMENT ON COLUMN customers.status IS 'Customer status: active, inactive, or suspended';
COMMENT ON COLUMN customers.white_label_config IS 'Custom branding configuration (logo, colors, domain)';
COMMENT ON COLUMN customers.metadata IS 'Additional customer metadata';