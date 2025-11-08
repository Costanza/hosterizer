# Hosterizer Backend Services

Microservices architecture for the Hosterizer platform.

## Services

### Golang Services

1. **Auth Service** (Port 8001)
   - User authentication and authorization
   - JWT token management
   - MFA support
   - Session management

2. **Customer Service** (Port 8002)
   - Customer CRUD operations
   - White-label configuration management

3. **Site Service** (Port 8003)
   - Site lifecycle management
   - Deployment tracking

4. **Infrastructure Service** (Port 8004)
   - Terraform workflow orchestration
   - Cloud provider integration

5. **Policy Service** (Port 8005)
   - Policy definition and validation
   - Compliance checking

6. **Ecommerce Service** (Port 8006)
   - Ecommerce platform integration
   - Credential management

### Python Services

7. **Cost Service** (Port 8007)
   - Cloud cost collection
   - Cost aggregation and reporting
   - Budget monitoring

## Development Setup

### Prerequisites
- Go 1.21+
- Python 3.11+
- PostgreSQL 15+
- Redis 7+

### Running Golang Services

```bash
# Navigate to service directory
cd auth-service

# Download dependencies
go mod download

# Run the service
go run cmd/server/main.go
```

### Running Python Service

```bash
# Navigate to cost service
cd cost-service

# Create virtual environment
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate

# Install dependencies
pip install -e .

# Run the service
python src/cost_service/main.py
```

## Environment Variables

Each service requires the following environment variables:

```
PORT=<service-port>
DATABASE_URL=postgresql://user:password@localhost:5432/hosterizer
REDIS_URL=redis://localhost:6379
LOG_LEVEL=info
```

## Project Structure

Each Golang service follows this structure:
```
service-name/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── domain/
│   ├── repository/
│   ├── service/
│   └── handlers/
├── go.mod
└── go.sum
```

Python service structure:
```
cost-service/
├── src/
│   └── cost_service/
│       ├── __init__.py
│       ├── main.py
│       ├── domain/
│       ├── repository/
│       └── services/
├── tests/
└── pyproject.toml
```
