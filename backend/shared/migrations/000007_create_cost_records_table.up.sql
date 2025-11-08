-- Create cost_records table
CREATE TABLE cost_records (
    id BIGSERIAL PRIMARY KEY,
    site_id BIGINT NOT NULL REFERENCES sites(id) ON DELETE CASCADE,
    customer_id BIGINT NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
    cloud_provider TEXT NOT NULL,
    cost_date DATE NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    currency TEXT NOT NULL DEFAULT 'USD',
    resource_breakdown JSONB DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(site_id, cost_date)
);
-- Create indexes for cost_records table
CREATE INDEX idx_cost_records_site ON cost_records(site_id, cost_date DESC);
CREATE INDEX idx_cost_records_customer ON cost_records(customer_id, cost_date DESC);
CREATE INDEX idx_cost_records_date ON cost_records(cost_date DESC);
CREATE INDEX idx_cost_records_cloud_provider ON cost_records(cloud_provider);
CREATE INDEX idx_cost_records_customer_date ON cost_records(customer_id, cost_date DESC, cloud_provider);
-- Create GIN index for JSONB resource_breakdown column
CREATE INDEX idx_cost_records_breakdown ON cost_records USING GIN (resource_breakdown);
-- Add check constraint for positive amounts
ALTER TABLE cost_records
ADD CONSTRAINT cost_records_amount_positive CHECK (amount >= 0);
-- Add comments to table
COMMENT ON TABLE cost_records IS 'Daily cost records for sites from cloud providers';
COMMENT ON COLUMN cost_records.cloud_provider IS 'Cloud provider that generated the cost';
COMMENT ON COLUMN cost_records.cost_date IS 'Date for which the cost is recorded';
COMMENT ON COLUMN cost_records.amount IS 'Cost amount in the specified currency';
COMMENT ON COLUMN cost_records.currency IS 'Currency code (default: USD)';
COMMENT ON COLUMN cost_records.resource_breakdown IS 'Breakdown of costs by resource type';