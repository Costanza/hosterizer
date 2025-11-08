# AWS Coding Standards

## General Principles

- Follow AWS Well-Architected Framework pillars
- Use Infrastructure as Code (Terraform, CloudFormation, CDK)
- Tag all resources consistently for cost tracking and organization
- Enable CloudTrail and CloudWatch logging for all accounts
- Use least privilege principle for IAM policies

## Naming Conventions

- Use consistent naming: `{project}-{environment}-{service}-{resource}`
- Example: `myapp-prod-api-lambda`, `myapp-dev-db-rds`
- Use lowercase with hyphens for resource names
- Include environment in all resource names

## Tagging Strategy

- Required tags for all resources:
  - `Environment`: dev, staging, prod
  - `Project`: project name
  - `Owner`: team or individual email
  - `CostCenter`: for billing allocation
  - `ManagedBy`: terraform, cloudformation, manual

```hcl
tags = {
  Environment = "prod"
  Project     = "myapp"
  Owner       = "platform-team@example.com"
  CostCenter  = "engineering"
  ManagedBy   = "terraform"
}
```

## IAM Best Practices

- Never use root account for daily operations
- Enable MFA for all users, especially privileged accounts
- Use IAM roles for EC2, Lambda, and other services (not access keys)
- Create service-specific roles with minimal permissions
- Use AWS managed policies when appropriate
- Rotate credentials regularly
- Use AWS Organizations for multi-account management

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:GetObject",
        "s3:PutObject"
      ],
      "Resource": "arn:aws:s3:::my-bucket/*"
    }
  ]
}
```

## Compute Services

### Lambda
- Use environment variables for configuration
- Set appropriate timeout and memory limits
- Enable X-Ray tracing for debugging
- Use Lambda layers for shared dependencies
- Implement proper error handling and retries
- Use provisioned concurrency for latency-sensitive functions

### EC2
- Use Auto Scaling Groups for high availability
- Store AMI IDs in variables, not hardcoded
- Use Systems Manager for patching and configuration
- Enable detailed monitoring for production instances
- Use placement groups for low-latency requirements

## Storage Services

### S3
- Enable versioning for critical buckets
- Use lifecycle policies to transition to cheaper storage classes
- Enable server-side encryption (SSE-S3 or SSE-KMS)
- Block public access by default
- Use bucket policies and IAM policies together
- Enable access logging for audit trails

### RDS
- Use Multi-AZ for production databases
- Enable automated backups with appropriate retention
- Use parameter groups for configuration management
- Enable encryption at rest
- Use secrets manager for database credentials
- Monitor with CloudWatch and Performance Insights

## Networking

### VPC
- Use separate VPCs for different environments
- Use at least 2 availability zones for high availability
- Public subnets for load balancers, private subnets for applications
- Use NAT Gateways for outbound internet access from private subnets
- Enable VPC Flow Logs for network monitoring

### Security Groups
- Use descriptive names and descriptions
- Follow least privilege: only open required ports
- Reference other security groups instead of CIDR blocks when possible
- Document the purpose of each rule

## Monitoring and Logging

- Enable CloudWatch Logs for all services
- Set up CloudWatch Alarms for critical metrics
- Use CloudWatch Dashboards for visualization
- Enable AWS Config for compliance monitoring
- Use EventBridge for event-driven architectures
- Implement centralized logging with CloudWatch Logs Insights

## Cost Optimization

- Use Reserved Instances or Savings Plans for predictable workloads
- Right-size instances based on CloudWatch metrics
- Use Spot Instances for fault-tolerant workloads
- Implement auto-scaling to match demand
- Delete unused resources (EBS volumes, snapshots, elastic IPs)
- Use AWS Cost Explorer and Budgets for monitoring

## Security

- Enable GuardDuty for threat detection
- Use AWS WAF for web application protection
- Enable AWS Shield for DDoS protection
- Use KMS for encryption key management
- Implement VPC endpoints for private AWS service access
- Regular security audits with AWS Security Hub
