# Hosterizer Terraform Infrastructure

Infrastructure as Code for deploying sites across multiple cloud providers.

## Structure

```
terraform/
├── modules/              # Reusable Terraform modules
│   ├── aws/             # AWS-specific modules
│   ├── azure/           # Azure-specific modules
│   ├── gcp/             # GCP-specific modules
│   ├── digitalocean/    # Digital Ocean modules
│   └── akamai/          # Akamai Cloud modules
└── templates/           # Site deployment templates
    └── site-deployment/ # Complete site deployment
```

## Supported Cloud Providers

1. **AWS** - Amazon Web Services
2. **Azure** - Microsoft Azure
3. **GCP** - Google Cloud Platform
4. **Digital Ocean** - Digital Ocean Cloud
5. **Akamai** - Akamai Cloud Computing (Linode)

## Module Organization

Each cloud provider has the following module structure:

- **networking**: VPC, subnets, security groups, firewall rules
- **compute**: Virtual machines, containers, load balancers
- **database**: Managed database services
- **storage**: Object storage, file storage

## Usage

### Prerequisites

- Terraform 1.5+
- Cloud provider credentials configured
- Remote state backend configured

### Deploying a Site

```bash
# Navigate to templates directory
cd templates/site-deployment

# Initialize Terraform
terraform init

# Create terraform.tfvars with site configuration
cat > terraform.tfvars <<EOF
site_name        = "example-site"
customer_id      = "customer-123"
cloud_provider   = "aws"
region           = "us-east-1"
instance_type    = "t3.medium"
EOF

# Plan deployment
terraform plan

# Apply deployment
terraform apply
```

### Module Variables

Each module accepts standard variables:

- `site_name`: Name of the site
- `customer_id`: Customer identifier
- `region`: Deployment region
- `environment`: Environment (dev, staging, prod)
- `tags`: Resource tags

### State Management

- Remote state stored in S3 (or equivalent)
- State locking enabled
- Separate state files per site deployment
- State encryption enabled

## Best Practices

- Use consistent naming conventions
- Apply appropriate tags to all resources
- Enable encryption at rest and in transit
- Use managed services when available
- Implement proper network segmentation
- Enable monitoring and logging
- Use least privilege IAM policies

## Development

### Adding a New Cloud Provider

1. Create provider directory under `modules/`
2. Implement networking, compute, database, storage modules
3. Create site deployment template
4. Document provider-specific requirements
5. Add provider to supported list

### Testing Modules

```bash
# Validate Terraform configuration
terraform validate

# Format Terraform files
terraform fmt -recursive

# Run security scan
tfsec .
```

## Security

- Never commit credentials or secrets
- Use secret management services
- Enable audit logging
- Implement network security controls
- Regular security scanning
- Follow cloud provider security best practices
