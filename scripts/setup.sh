#!/bin/bash

set -e

echo "========================================="
echo "Hosterizer Project Setup"
echo "========================================="
echo ""

# Check prerequisites
echo "Checking prerequisites..."

# Check Go
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21 or higher."
    exit 1
fi
echo "✓ Go $(go version | awk '{print $3}')"

# Check Node.js
if ! command -v node &> /dev/null; then
    echo "❌ Node.js is not installed. Please install Node.js 18 or higher."
    exit 1
fi
echo "✓ Node.js $(node --version)"

# Check Python
if ! command -v python3 &> /dev/null; then
    echo "❌ Python is not installed. Please install Python 3.11 or higher."
    exit 1
fi
echo "✓ Python $(python3 --version | awk '{print $2}')"

# Check Terraform
if ! command -v terraform &> /dev/null; then
    echo "⚠️  Terraform is not installed. Install Terraform 1.5+ for infrastructure automation."
else
    echo "✓ Terraform $(terraform version | head -n1 | awk '{print $2}')"
fi

# Check Docker
if ! command -v docker &> /dev/null; then
    echo "⚠️  Docker is not installed. Install Docker for local development environment."
else
    echo "✓ Docker $(docker --version | awk '{print $3}' | sed 's/,//')"
fi

echo ""
echo "Installing dependencies..."
echo ""

# Install backend dependencies
echo "📦 Installing backend service dependencies..."
cd backend/auth-service && go mod download && cd ../..
cd backend/customer-service && go mod download && cd ../..
cd backend/site-service && go mod download && cd ../..
cd backend/infrastructure-service && go mod download && cd ../..
cd backend/policy-service && go mod download && cd ../..
cd backend/ecommerce-service && go mod download && cd ../..

echo "📦 Installing Python cost service dependencies..."
cd backend/cost-service
python3 -m venv venv
source venv/bin/activate 2>/dev/null || . venv/Scripts/activate 2>/dev/null
pip install -e . > /dev/null
deactivate 2>/dev/null || true
cd ../..

# Install frontend dependencies
echo "📦 Installing frontend dependencies..."
cd frontend/admin-portal && npm install && cd ../..
cd frontend/customer-portal && npm install && cd ../..

echo ""
echo "✅ Setup complete!"
echo ""
echo "Next steps:"
echo "1. Copy .env.example to .env and configure your environment variables"
echo "2. Start the database and Redis: docker-compose up -d postgres redis"
echo "3. Run database migrations (will be implemented in subsequent tasks)"
echo "4. Start services using the Makefile commands"
echo ""
echo "Development commands:"
echo "  make help              - Show all available commands"
echo "  make dev-admin         - Start admin portal (http://localhost:3000)"
echo "  make dev-customer      - Start customer portal (http://localhost:3001)"
echo "  make dev-auth          - Start auth service"
echo ""
echo "For full documentation, see README.md"
