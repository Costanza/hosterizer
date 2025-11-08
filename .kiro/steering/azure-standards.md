# Azure Coding Standards

## General Principles

- Follow Azure Well-Architected Framework
- Use Infrastructure as Code (Terraform, ARM templates, Bicep)
- Apply consistent naming conventions across all resources
- Use Azure Policy for governance and compliance
- Enable Azure Monitor and Application Insights

## Naming Conventions

- Format: `{resource-type}-{project}-{environment}-{region}-{instance}`
- Example: `vm-myapp-prod-eastus-01`, `sql-myapp-dev-westus-01`
- Use lowercase with hyphens
- Resource type abbreviations: `rg` (resource group), `vm` (virtual machine), `sql` (SQL database)

## Resource Organization

- Use Resource Groups to organize related resources
- One resource group per application/environment combination
- Use Management Groups for multi-subscription governance
- Apply locks to prevent accidental deletion of critical resources

## Tagging Strategy

- Required tags:
  - `Environment`: dev, staging, prod
  - `Project`: project name
  - `Owner`: team or email
  - `CostCenter`: billing allocation
  - `ManagedBy`: terraform, bicep, manual

```hcl
tags = {
  Environment = "prod"
  Project     = "myapp"
  Owner       = "platform-team@example.com"
  CostCenter  = "engineering"
  ManagedBy   = "terraform"
}
```

## Identity and Access Management

- Use Azure AD for identity management
- Enable MFA for all users
- Use Managed Identities for Azure resources (not service principals when possible)
- Apply RBAC at appropriate scope (subscription, resource group, resource)
- Use built-in roles before creating custom roles
- Follow least privilege principle

```bash
# Assign role to managed identity
az role assignment create \
  --assignee <managed-identity-id> \
  --role "Storage Blob Data Contributor" \
  --scope /subscriptions/{sub-id}/resourceGroups/{rg-name}
```

## Compute Services

### Virtual Machines
- Use Availability Sets or Availability Zones for high availability
- Enable Azure Backup for critical VMs
- Use managed disks (not unmanaged)
- Apply OS and security updates via Update Management
- Use VM Scale Sets for auto-scaling scenarios

### App Service
- Use deployment slots for zero-downtime deployments
- Enable Application Insights for monitoring
- Use App Service Plans appropriate for workload
- Configure auto-scaling rules
- Enable diagnostic logging

### Azure Functions
- Use consumption plan for event-driven workloads
- Use premium plan for VNet integration or longer execution times
- Enable Application Insights
- Use durable functions for stateful workflows
- Store secrets in Key Vault, not app settings

## Storage Services

### Storage Accounts
- Use Standard_LRS for dev, Standard_GRS or Standard_ZRS for prod
- Enable soft delete for blob storage
- Use lifecycle management policies
- Enable encryption at rest (enabled by default)
- Use private endpoints for secure access
- Enable storage analytics and logging

### Azure SQL Database
- Use elastic pools for multiple databases with variable usage
- Enable automatic tuning
- Configure geo-replication for disaster recovery
- Use Azure AD authentication
- Enable auditing and threat detection
- Store connection strings in Key Vault

## Networking

### Virtual Networks
- Use separate VNets for different environments
- Use subnets to segment workloads
- Implement Network Security Groups (NSGs) for traffic control
- Use Azure Firewall or Network Virtual Appliances for advanced scenarios
- Enable VNet flow logs for monitoring

### Application Gateway / Load Balancer
- Use Application Gateway for HTTP/HTTPS traffic with WAF
- Use Load Balancer for non-HTTP traffic
- Configure health probes
- Use availability zones for high availability

## Security

- Enable Azure Security Center (Defender for Cloud)
- Use Azure Key Vault for secrets, keys, and certificates
- Enable Azure AD Conditional Access
- Implement Azure Private Link for PaaS services
- Use Azure Policy for compliance enforcement
- Enable diagnostic logs for all resources

## Monitoring and Logging

- Enable Azure Monitor for all resources
- Use Application Insights for application telemetry
- Create action groups for alert notifications
- Use Log Analytics workspace for centralized logging
- Set up dashboards in Azure Portal or Grafana
- Configure diagnostic settings to send logs to Log Analytics

## Cost Optimization

- Use Azure Cost Management for tracking and optimization
- Right-size VMs based on metrics
- Use reserved instances for predictable workloads
- Implement auto-shutdown for dev/test resources
- Use Azure Advisor recommendations
- Delete unused resources (disks, IPs, snapshots)

## DevOps Integration

- Use Azure DevOps or GitHub Actions for CI/CD
- Store infrastructure code in version control
- Use Azure Pipelines for automated deployments
- Implement approval gates for production deployments
- Use Azure Artifacts for package management
