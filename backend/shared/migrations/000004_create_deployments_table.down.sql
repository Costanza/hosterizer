-- Drop deployments table and related objects
DROP INDEX IF EXISTS idx_deployments_created;
DROP INDEX IF EXISTS idx_deployments_site_created;
DROP INDEX IF EXISTS idx_deployments_uuid;
DROP INDEX IF EXISTS idx_deployments_status;
DROP INDEX IF EXISTS idx_deployments_site;
DROP TABLE IF EXISTS deployments;