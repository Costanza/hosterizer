-- Create deployments table
CREATE TABLE deployments (
    id BIGSERIAL PRIMARY KEY,
    uuid UUID DEFAULT gen_random_uuid() UNIQUE NOT NULL,
    site_id BIGINT NOT NULL REFERENCES sites(id) ON DELETE CASCADE,
    deployment_type TEXT NOT NULL CHECK (
        deployment_type IN ('create', 'update', 'delete')
    ),
    status TEXT NOT NULL DEFAULT 'queued' CHECK (
        status IN (
            'queued',
            'running',
            'completed',
            'failed',
            'cancelled'
        )
    ),
    terraform_plan TEXT,
    terraform_output TEXT,
    error_message TEXT,
    started_at TIMESTAMPTZ,
    completed_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- Create indexes for deployments table
CREATE INDEX idx_deployments_site ON deployments(site_id);
CREATE INDEX idx_deployments_status ON deployments(status);
CREATE INDEX idx_deployments_uuid ON deployments(uuid);
CREATE INDEX idx_deployments_site_created ON deployments(site_id, created_at DESC);
CREATE INDEX idx_deployments_created ON deployments(created_at DESC);
-- Add comments to table
COMMENT ON TABLE deployments IS 'Deployment requests and their execution status';
COMMENT ON COLUMN deployments.deployment_type IS 'Type of deployment: create, update, or delete';
COMMENT ON COLUMN deployments.status IS 'Deployment status: queued, running, completed, failed, or cancelled';
COMMENT ON COLUMN deployments.terraform_plan IS 'Terraform plan output';
COMMENT ON COLUMN deployments.terraform_output IS 'Terraform apply output';
COMMENT ON COLUMN deployments.error_message IS 'Error message if deployment failed';