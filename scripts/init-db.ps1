# Database initialization script for Hosterizer (PowerShell)
# This script creates the database and runs initialization SQL

param(
    [string]$DbHost = $env:DB_HOST ?? "localhost",
    [string]$DbPort = $env:DB_PORT ?? "5432",
    [string]$DbUser = $env:DB_USER ?? "postgres",
    [string]$DbPassword = $env:DB_PASSWORD ?? "postgres",
    [string]$DbName = $env:DB_NAME ?? "hosterizer"
)

Write-Host "Initializing Hosterizer database..."
Write-Host "Host: $DbHost"
Write-Host "Port: $DbPort"
Write-Host "Database: $DbName"

# Set password environment variable for psql
$env:PGPASSWORD = $DbPassword

# Check if PostgreSQL is accessible
try {
    $result = & psql -h $DbHost -p $DbPort -U $DbUser -lqt 2>&1
    if ($LASTEXITCODE -ne 0) {
        Write-Error "Cannot connect to PostgreSQL server"
        exit 1
    }
} catch {
    Write-Error "Cannot connect to PostgreSQL server: $_"
    exit 1
}

# Check if database exists
$dbExists = & psql -h $DbHost -p $DbPort -U $DbUser -tc "SELECT 1 FROM pg_database WHERE datname = '$DbName'" 2>&1
if ($dbExists -notmatch "1") {
    Write-Host "Creating database '$DbName'..."
    & psql -h $DbHost -p $DbPort -U $DbUser -c "CREATE DATABASE $DbName"
    if ($LASTEXITCODE -ne 0) {
        Write-Error "Failed to create database"
        exit 1
    }
}

Write-Host "Database '$DbName' is ready"

# Run initialization SQL
Write-Host "Running initialization SQL..."
$scriptPath = Join-Path $PSScriptRoot "init-db.sql"
& psql -h $DbHost -p $DbPort -U $DbUser -d $DbName -f $scriptPath

if ($LASTEXITCODE -eq 0) {
    Write-Host "Database initialization complete!" -ForegroundColor Green
} else {
    Write-Error "Database initialization failed"
    exit 1
}
