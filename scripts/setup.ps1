# Hosterizer Project Setup Script for Windows

Write-Host "=========================================" -ForegroundColor Cyan
Write-Host "Hosterizer Project Setup" -ForegroundColor Cyan
Write-Host "=========================================" -ForegroundColor Cyan
Write-Host ""

# Check prerequisites
Write-Host "Checking prerequisites..." -ForegroundColor Yellow

# Check Go
try {
    $goVersion = go version
    Write-Host "✓ $goVersion" -ForegroundColor Green
} catch {
    Write-Host "❌ Go is not installed. Please install Go 1.21 or higher." -ForegroundColor Red
    exit 1
}

# Check Node.js
try {
    $nodeVersion = node --version
    Write-Host "✓ Node.js $nodeVersion" -ForegroundColor Green
} catch {
    Write-Host "❌ Node.js is not installed. Please install Node.js 18 or higher." -ForegroundColor Red
    exit 1
}

# Check Python
try {
    $pythonVersion = python --version
    Write-Host "✓ $pythonVersion" -ForegroundColor Green
} catch {
    Write-Host "❌ Python is not installed. Please install Python 3.11 or higher." -ForegroundColor Red
    exit 1
}

# Check Terraform
try {
    $tfVersion = terraform version
    Write-Host "✓ Terraform $(($tfVersion -split '\n')[0])" -ForegroundColor Green
} catch {
    Write-Host "⚠️  Terraform is not installed. Install Terraform 1.5+ for infrastructure automation." -ForegroundColor Yellow
}

# Check Docker
try {
    $dockerVersion = docker --version
    Write-Host "✓ $dockerVersion" -ForegroundColor Green
} catch {
    Write-Host "⚠️  Docker is not installed. Install Docker for local development environment." -ForegroundColor Yellow
}

Write-Host ""
Write-Host "Installing dependencies..." -ForegroundColor Yellow
Write-Host ""

# Install backend dependencies
Write-Host "📦 Installing backend service dependencies..." -ForegroundColor Cyan
Set-Location backend\auth-service
go mod download
Set-Location ..\..

Set-Location backend\customer-service
go mod download
Set-Location ..\..

Set-Location backend\site-service
go mod download
Set-Location ..\..

Set-Location backend\infrastructure-service
go mod download
Set-Location ..\..

Set-Location backend\policy-service
go mod download
Set-Location ..\..

Set-Location backend\ecommerce-service
go mod download
Set-Location ..\..

Write-Host "📦 Installing Python cost service dependencies..." -ForegroundColor Cyan
Set-Location backend\cost-service
python -m venv venv
.\venv\Scripts\Activate.ps1
pip install -e . | Out-Null
deactivate
Set-Location ..\..

# Install frontend dependencies
Write-Host "📦 Installing frontend dependencies..." -ForegroundColor Cyan
Set-Location frontend\admin-portal
npm install
Set-Location ..\..

Set-Location frontend\customer-portal
npm install
Set-Location ..\..

Write-Host ""
Write-Host "✅ Setup complete!" -ForegroundColor Green
Write-Host ""
Write-Host "Next steps:"
Write-Host "1. Copy .env.example to .env and configure your environment variables"
Write-Host "2. Start the database and Redis: docker-compose up -d postgres redis"
Write-Host "3. Run database migrations (will be implemented in subsequent tasks)"
Write-Host "4. Start services using the Makefile commands"
Write-Host ""
Write-Host "Development commands:"
Write-Host "  make help              - Show all available commands"
Write-Host "  make dev-admin         - Start admin portal (http://localhost:3000)"
Write-Host "  make dev-customer      - Start customer portal (http://localhost:3001)"
Write-Host "  make dev-auth          - Start auth service"
Write-Host ""
Write-Host "For full documentation, see README.md"
