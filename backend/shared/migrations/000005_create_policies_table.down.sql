-- Drop policies table and related objects
DROP TRIGGER IF EXISTS policies_updated_at ON policies;
DROP INDEX IF EXISTS idx_policies_rules;
DROP INDEX IF EXISTS idx_policies_type_enabled;
DROP INDEX IF EXISTS idx_policies_uuid;
DROP INDEX IF EXISTS idx_policies_enabled;
DROP INDEX IF EXISTS idx_policies_type;
DROP TABLE IF EXISTS policies;