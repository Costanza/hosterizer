.PHONY: help install build test clean dev

help:
	@echo "Hosterizer Build Commands"
	@echo "========================="
	@echo "install          - Install all dependencies"
	@echo "build            - Build all services and applications"
	@echo "test             - Run all tests"
	@echo "clean            - Clean build artifacts"
	@echo "dev              - Start development environment"
	@echo ""
	@echo "Backend Services:"
	@echo "  make backend-install    - Install backend dependencies"
	@echo "  make backend-build      - Build backend services"
	@echo "  make backend-test       - Test backend services"
	@echo ""
	@echo "Frontend Applications:"
	@echo "  make frontend-install   - Install frontend dependencies"
	@echo "  make frontend-build     - Build frontend applications"
	@echo "  make frontend-test      - Test frontend applications"

install: backend-install frontend-install

build: backend-build frontend-build

test: backend-test frontend-test

clean:
	@echo "Cleaning build artifacts..."
	@find backend -name "*.exe" -type f -delete
	@find backend -name "*.test" -type f -delete
	@rm -rf frontend/admin-portal/dist
	@rm -rf frontend/customer-portal/dist
	@rm -rf backend/cost-service/dist
	@echo "Clean complete"

# Backend targets
backend-install:
	@echo "Installing backend dependencies..."
	@cd backend/auth-service && go mod download
	@cd backend/customer-service && go mod download
	@cd backend/site-service && go mod download
	@cd backend/infrastructure-service && go mod download
	@cd backend/policy-service && go mod download
	@cd backend/ecommerce-service && go mod download
	@cd backend/cost-service && pip install -e .
	@echo "Backend dependencies installed"

backend-build:
	@echo "Building backend services..."
	@cd backend/auth-service && go build -o ../../bin/auth-service ./cmd/server
	@cd backend/customer-service && go build -o ../../bin/customer-service ./cmd/server
	@cd backend/site-service && go build -o ../../bin/site-service ./cmd/server
	@cd backend/infrastructure-service && go build -o ../../bin/infrastructure-service ./cmd/server
	@cd backend/policy-service && go build -o ../../bin/policy-service ./cmd/server
	@cd backend/ecommerce-service && go build -o ../../bin/ecommerce-service ./cmd/server
	@echo "Backend services built"

backend-test:
	@echo "Testing backend services..."
	@cd backend/auth-service && go test ./...
	@cd backend/customer-service && go test ./...
	@cd backend/site-service && go test ./...
	@cd backend/infrastructure-service && go test ./...
	@cd backend/policy-service && go test ./...
	@cd backend/ecommerce-service && go test ./...
	@cd backend/cost-service && pytest
	@echo "Backend tests complete"

# Frontend targets
frontend-install:
	@echo "Installing frontend dependencies..."
	@cd frontend/admin-portal && npm install
	@cd frontend/customer-portal && npm install
	@echo "Frontend dependencies installed"

frontend-build:
	@echo "Building frontend applications..."
	@cd frontend/admin-portal && npm run build
	@cd frontend/customer-portal && npm run build
	@echo "Frontend applications built"

frontend-test:
	@echo "Testing frontend applications..."
	@cd frontend/admin-portal && npm test -- --run
	@cd frontend/customer-portal && npm test -- --run
	@echo "Frontend tests complete"

# Development targets
dev-admin:
	@cd frontend/admin-portal && npm run dev

dev-customer:
	@cd frontend/customer-portal && npm run dev

dev-auth:
	@cd backend/auth-service && go run cmd/server/main.go

dev-customer-service:
	@cd backend/customer-service && go run cmd/server/main.go

dev-site:
	@cd backend/site-service && go run cmd/server/main.go

dev-infrastructure:
	@cd backend/infrastructure-service && go run cmd/server/main.go

dev-policy:
	@cd backend/policy-service && go run cmd/server/main.go

dev-ecommerce:
	@cd backend/ecommerce-service && go run cmd/server/main.go

dev-cost:
	@cd backend/cost-service && python src/cost_service/main.py
