# Script to test Row-Level Security (RLS) for multi-tenancy (PowerShell)
# This script runs the RLS test SQL file

param(
    [string]$DbHost = $env:DB_HOST ?? "localhost",
    [string]$DbPort = $env:DB_PORT ?? "5432",
    [string]$DbUser = $env:DB_USER ?? "postgres",
    [string]$DbPassword = $env:DB_PASSWORD ?? "postgres",
    [string]$DbName = $env:DB_NAME ?? "hosterizer"
)

Write-Host "Testing Row-Level Security..."
Write-Host "Host: $DbHost"
Write-Host "Port: $DbPort"
Write-Host "Database: $DbName"
Write-Host ""

# Set password environment variable for psql
$env:PGPASSWORD = $DbPassword

# Run the test SQL script
$scriptPath = Join-Path $PSScriptRoot "test-rls.sql"
& psql -h $DbHost -p $DbPort -U $DbUser -d $DbName -f $scriptPath

if ($LASTEXITCODE -eq 0) {
    Write-Host ""
    Write-Host "RLS testing complete!" -ForegroundColor Green
} else {
    Write-Error "RLS testing failed"
    exit 1
}
