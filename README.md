# Hosterizer

Multi-tenant cloud hosting platform that enables users to deploy and manage multiple websites across various cloud providers.

## Project Structure

```
hosterizer/
├── backend/                    # Backend services (Golang)
│   ├── auth-service/          # Authentication and authorization
│   ├── customer-service/      # Customer management
│   ├── site-service/          # Site management
│   ├── infrastructure-service/ # Infrastructure automation
│   ├── cost-service/          # Cost management (Python)
│   ├── policy-service/        # Policy management
│   └── ecommerce-service/     # Ecommerce integration
├── frontend/                   # Frontend applications (React + TypeScript)
│   ├── admin-portal/          # Admin portal
│   └── customer-portal/       # Customer portal
├── terraform/                  # Infrastructure as Code
│   ├── modules/               # Terraform modules
│   │   ├── aws/
│   │   ├── azure/
│   │   ├── gcp/
│   │   ├── digitalocean/
│   │   └── akamai/
│   └── templates/             # Deployment templates
└── docs/                       # Documentation
```

## Technology Stack

### Backend
- **Primary Language**: Golang
- **Secondary Language**: Python (for cost service)
- **Database**: PostgreSQL with row-level security
- **Cache**: Redis
- **API Style**: RESTful

### Frontend
- **Framework**: React 18+ with TypeScript
- **Styling**: Tailwind CSS (configurable)
- **State Management**: React Query + Context API
- **Routing**: React Router

### Infrastructure
- **IaC**: Terraform
- **Cloud Providers**: AWS, Azure, GCP, Digital Ocean, Akamai Cloud
- **Observability**: LGTM Stack (Loki, Grafana, Tempo, Mimir)

## Getting Started

### Prerequisites
- Go 1.21+
- Node.js 18+
- Python 3.11+
- Terraform 1.5+
- PostgreSQL 15+
- Redis 7+

### Development Setup

1. Clone the repository
2. Set up backend services (see `backend/README.md`)
3. Set up frontend applications (see `frontend/README.md`)
4. Configure environment variables
5. Run database migrations
6. Start services

## Documentation

- [Requirements](/.kiro/specs/hosterizer/requirements.md)
- [Design](/.kiro/specs/hosterizer/design.md)
- [Implementation Tasks](/.kiro/specs/hosterizer/tasks.md)

## License

Proprietary
