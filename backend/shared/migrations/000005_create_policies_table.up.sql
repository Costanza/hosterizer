-- Create policies table
CREATE TABLE policies (
    id BIGSERIAL PRIMARY KEY,
    uuid UUID DEFAULT gen_random_uuid() UNIQUE NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    policy_type TEXT NOT NULL CHECK (
        policy_type IN ('resource_limit', 'security', 'cost')
    ),
    rules JSONB NOT NULL,
    enabled BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- Create indexes for policies table
CREATE INDEX idx_policies_type ON policies(policy_type);
CREATE INDEX idx_policies_enabled ON policies(enabled);
CREATE INDEX idx_policies_uuid ON policies(uuid);
CREATE INDEX idx_policies_type_enabled ON policies(policy_type, enabled);
-- Create GIN index for JSONB rules column
CREATE INDEX idx_policies_rules ON policies USING GIN (rules);
-- Create trigger to automatically update updated_at
CREATE TRIGGER policies_updated_at BEFORE
UPDATE ON policies FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
-- Add comments to table
COMMENT ON TABLE policies IS 'Cloud policies for resource limits, security, and cost controls';
COMMENT ON COLUMN policies.policy_type IS 'Type of policy: resource_limit, security, or cost';
COMMENT ON COLUMN policies.rules IS 'Policy rules in JSON format';
COMMENT ON COLUMN policies.enabled IS 'Whether the policy is currently active';