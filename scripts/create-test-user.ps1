# PowerShell script to create test users for API testing
# Password for all test users: AdminPass123!
# Bcrypt hash: $2a$12$LQv3c1yqBWVHxkd0LHAkCOYz6TtxMQJqhN8/LewY5GyYKKXqKKKKK

param(
    [string]$DbHost = "localhost",
    [string]$DbPort = "5432",
    [string]$DbUser = "postgres",
    [string]$DbPassword = "postgres",
    [string]$DbName = "hosterizer"
)

Write-Host "Creating test users in database..." -ForegroundColor Cyan
Write-Host "Database: $DbName at ${DbHost}:${DbPort}" -ForegroundColor Gray
Write-Host ""

$env:PGPASSWORD = $DbPassword

$sql = @"
-- Create administrator user
INSERT INTO users (email, password_hash, first_name, last_name, role)
VALUES (
  'admin@hosterizer.com',
  '`$2a`$12`$LQv3c1yqBWVHxkd0LHAkCOYz6TtxMQJqhN8/LewY5GyYKKXqKKKKK',
  'Admin',
  'User',
  'administrator'
)
ON CONFLICT (email) DO UPDATE SET
  password_hash = EXCLUDED.password_hash,
  first_name = EXCLUDED.first_name,
  last_name = EXCLUDED.last_name,
  role = EXCLUDED.role;

-- Create customer user
INSERT INTO users (email, password_hash, first_name, last_name, role)
VALUES (
  'customer@hosterizer.com',
  '`$2a`$12`$LQv3c1yqBWVHxkd0LHAkCOYz6TtxMQJqhN8/LewY5GyYKKXqKKKKK',
  'Customer',
  'User',
  'customer'
)
ON CONFLICT (email) DO UPDATE SET
  password_hash = EXCLUDED.password_hash,
  first_name = EXCLUDED.first_name,
  last_name = EXCLUDED.last_name,
  role = EXCLUDED.role;

-- Display created users
SELECT 
  id, 
  email, 
  first_name, 
  last_name, 
  role,
  mfa_enabled,
  created_at
FROM users
WHERE email IN ('admin@hosterizer.com', 'customer@hosterizer.com')
ORDER BY email;
"@

try {
    $sql | psql -h $DbHost -p $DbPort -U $DbUser -d $DbName
    
    Write-Host ""
    Write-Host "✅ Test users created successfully!" -ForegroundColor Green
    Write-Host ""
    Write-Host "Test Users:" -ForegroundColor Cyan
    Write-Host "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" -ForegroundColor Gray
    Write-Host "Administrator:" -ForegroundColor Yellow
    Write-Host "  Email:    admin@hosterizer.com"
    Write-Host "  Password: AdminPass123!"
    Write-Host "  Role:     administrator"
    Write-Host ""
    Write-Host "Customer:" -ForegroundColor Yellow
    Write-Host "  Email:    customer@hosterizer.com"
    Write-Host "  Password: AdminPass123!"
    Write-Host "  Role:     customer"
    Write-Host "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" -ForegroundColor Gray
    Write-Host ""
    Write-Host "You can now use these credentials to test the API!" -ForegroundColor Green
}
catch {
    Write-Host "❌ Error creating test users: $_" -ForegroundColor Red
    exit 1
}
finally {
    Remove-Item Env:\PGPASSWORD -ErrorAction SilentlyContinue
}
