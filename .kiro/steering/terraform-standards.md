# Terraform Coding Standards

## Project Structure

```
terraform/
├── environments/
│   ├── dev/
│   │   ├── main.tf
│   │   ├── variables.tf
│   │   ├── outputs.tf
│   │   └── terraform.tfvars
│   ├── staging/
│   └── prod/
├── modules/
│   ├── networking/
│   ├── compute/
│   └── database/
└── README.md
```

## File Organization

- `main.tf`: Primary resource definitions
- `variables.tf`: Input variable declarations
- `outputs.tf`: Output value declarations
- `versions.tf`: Provider and Terraform version constraints
- `terraform.tfvars`: Variable values (don't commit sensitive values)

## Naming Conventions

- Resources: `{resource_type}_{descriptive_name}`
- Variables: `snake_case`
- Modules: `kebab-case` for directory names
- Use descriptive names that indicate purpose

```hcl
resource "aws_instance" "web_server" {
  ami           = var.ami_id
  instance_type = var.instance_type
}

variable "instance_type" {
  description = "EC2 instance type"
  type        = string
  default     = "t3.micro"
}
```

## Variables

- Always include description
- Specify type constraints
- Provide defaults for optional variables
- Use validation rules when appropriate
- Group related variables together

```hcl
variable "environment" {
  description = "Environment name (dev, staging, prod)"
  type        = string
  
  validation {
    condition     = contains(["dev", "staging", "prod"], var.environment)
    error_message = "Environment must be dev, staging, or prod."
  }
}

variable "tags" {
  description = "Common tags to apply to all resources"
  type        = map(string)
  default     = {}
}
```

## Modules

- Create reusable modules for common patterns
- Keep modules focused and single-purpose
- Document module inputs and outputs
- Version modules using Git tags or Terraform Registry
- Use semantic versioning for module releases

```hcl
module "vpc" {
  source  = "./modules/networking"
  version = "1.2.0"
  
  vpc_cidr    = var.vpc_cidr
  environment = var.environment
  tags        = var.common_tags
}
```

## State Management

- Use remote state backend (S3, Azure Storage, Terraform Cloud)
- Enable state locking (DynamoDB for S3, built-in for others)
- Never commit state files to version control
- Use separate state files per environment
- Enable versioning on state storage

```hcl
terraform {
  backend "s3" {
    bucket         = "myapp-terraform-state"
    key            = "prod/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "terraform-state-lock"
  }
}
```

## Provider Configuration

- Pin provider versions to avoid breaking changes
- Use `required_providers` block
- Configure providers in root module, not in reusable modules
- Use provider aliases for multi-region or multi-account scenarios

```hcl
terraform {
  required_version = ">= 1.5.0"
  
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
  
  default_tags {
    tags = var.common_tags
  }
}
```

## Resource Configuration

- Use `for_each` over `count` for creating multiple similar resources
- Use `depends_on` sparingly (Terraform handles most dependencies)
- Use lifecycle blocks to prevent accidental deletion
- Enable `prevent_destroy` for critical resources

```hcl
resource "aws_s3_bucket" "data" {
  for_each = toset(var.bucket_names)
  
  bucket = each.value
  
  lifecycle {
    prevent_destroy = true
  }
}
```

## Data Sources

- Use data sources to reference existing resources
- Prefer data sources over hardcoded values
- Cache data source results when possible

```hcl
data "aws_ami" "ubuntu" {
  most_recent = true
  owners      = ["099720109477"] # Canonical
  
  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-*"]
  }
}
```

## Outputs

- Output values needed by other modules or for reference
- Include descriptions for all outputs
- Mark sensitive outputs appropriately

```hcl
output "vpc_id" {
  description = "ID of the VPC"
  value       = aws_vpc.main.id
}

output "database_password" {
  description = "Database master password"
  value       = aws_db_instance.main.password
  sensitive   = true
}
```

## Best Practices

- Run `terraform fmt` before committing
- Use `terraform validate` to check syntax
- Run `terraform plan` before apply
- Use workspaces for environment separation (or separate state files)
- Implement CI/CD for Terraform deployments
- Use `tflint` for additional linting
- Document complex logic with comments
- Use locals for computed values used multiple times
- Avoid hardcoding values; use variables
- Use `terraform-docs` to generate module documentation

## Security

- Never commit sensitive values (use environment variables or secret managers)
- Use encrypted state storage
- Implement least privilege for Terraform execution role
- Scan for security issues with `tfsec` or `checkov`
- Use `.gitignore` to exclude sensitive files

```gitignore
# .gitignore
*.tfstate
*.tfstate.*
.terraform/
*.tfvars
!terraform.tfvars.example
```

## Testing

- Use `terraform plan` for validation
- Implement automated testing with Terratest
- Test modules independently
- Use `terraform-compliance` for policy testing
