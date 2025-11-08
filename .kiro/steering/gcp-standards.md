# GCP (Google Cloud Platform) Coding Standards

## General Principles

- Follow Google Cloud Architecture Framework
- Use Infrastructure as Code (Terraform, Deployment Manager)
- Apply consistent naming and labeling
- Enable Cloud Logging and Cloud Monitoring
- Use Organization Policies for governance

## Naming Conventions

- Format: `{project}-{environment}-{service}-{resource}`
- Example: `myapp-prod-api-vm`, `myapp-dev-db-sql`
- Use lowercase with hyphens
- Keep names under 63 characters (GCP limit)
- Be descriptive but concise

## Resource Organization

- Use Projects to isolate environments and applications
- Use Folders to organize projects hierarchically
- Use Organizations for enterprise-wide governance
- Apply resource hierarchy: Organization → Folder → Project → Resource

## Labeling Strategy

- Required labels for all resources:
  - `environment`: dev, staging, prod
  - `project`: project name
  - `owner`: team or email
  - `cost-center`: billing allocation
  - `managed-by`: terraform, deployment-manager, manual

```hcl
labels = {
  environment = "prod"
  project     = "myapp"
  owner       = "platform-team"
  cost-center = "engineering"
  managed-by  = "terraform"
}
```

## Identity and Access Management

- Use Google Cloud IAM for access control
- Enable 2FA for all users
- Use service accounts for applications (not user accounts)
- Apply IAM policies at appropriate level (organization, folder, project, resource)
- Use predefined roles before creating custom roles
- Follow least privilege principle
- Use Workload Identity for GKE pods

```bash
# Grant role to service account
gcloud projects add-iam-policy-binding PROJECT_ID \
  --member="serviceAccount:SA_NAME@PROJECT_ID.iam.gserviceaccount.com" \
  --role="roles/storage.objectViewer"
```

## Compute Services

### Compute Engine
- Use instance templates for consistent VM configuration
- Use managed instance groups for auto-scaling and high availability
- Enable OS patch management
- Use persistent disks for data storage
- Apply appropriate machine types based on workload
- Use preemptible VMs for fault-tolerant batch workloads

### Cloud Run
- Use for containerized stateless applications
- Configure concurrency and timeout appropriately
- Use Cloud Run revisions for traffic splitting
- Enable Cloud Run authentication
- Use Secret Manager for sensitive configuration

### Cloud Functions
- Use for event-driven serverless workloads
- Set appropriate memory and timeout limits
- Use environment variables for configuration
- Enable Cloud Logging
- Use Secret Manager for secrets

### GKE (Google Kubernetes Engine)
- Use regional clusters for high availability
- Enable Workload Identity for pod authentication
- Use node pools for different workload types
- Enable cluster autoscaling
- Use Binary Authorization for image security
- Enable GKE monitoring and logging

## Storage Services

### Cloud Storage
- Use appropriate storage class: Standard, Nearline, Coldline, Archive
- Enable versioning for critical buckets
- Use lifecycle policies for cost optimization
- Enable uniform bucket-level access
- Use signed URLs for temporary access
- Enable audit logging

### Cloud SQL
- Use high availability configuration for production
- Enable automated backups with point-in-time recovery
- Use private IP for secure access
- Enable Cloud SQL Proxy for secure connections
- Use Secret Manager for credentials
- Enable query insights for performance monitoring

### Firestore / Datastore
- Design data model for query patterns
- Use composite indexes for complex queries
- Implement security rules
- Enable backups
- Monitor usage and costs

## Networking

### VPC (Virtual Private Cloud)
- Use shared VPC for multi-project networking
- Use separate VPCs for different environments
- Implement subnet ranges carefully (avoid overlaps)
- Use Cloud NAT for outbound internet access from private instances
- Enable VPC Flow Logs for network monitoring

### Cloud Load Balancing
- Use global load balancing for multi-region applications
- Configure health checks appropriately
- Use Cloud CDN for static content
- Enable Cloud Armor for DDoS protection and WAF

### Firewall Rules
- Use descriptive names and descriptions
- Apply least privilege: only allow required traffic
- Use service accounts as targets when possible
- Use tags for grouping instances
- Document the purpose of each rule

## Security

- Enable Security Command Center
- Use Cloud Armor for web application protection
- Use Cloud KMS for encryption key management
- Enable VPC Service Controls for data exfiltration protection
- Use Binary Authorization for container image security
- Implement Cloud Identity-Aware Proxy for application access
- Enable audit logging for all services

## Monitoring and Logging

- Enable Cloud Logging for all services
- Use Cloud Monitoring for metrics and alerting
- Create dashboards for key metrics
- Set up notification channels for alerts
- Use Cloud Trace for distributed tracing
- Use Cloud Profiler for performance analysis
- Implement log sinks for long-term retention

## Cost Optimization

- Use committed use discounts for predictable workloads
- Right-size instances based on monitoring data
- Use preemptible VMs for fault-tolerant workloads
- Implement autoscaling to match demand
- Delete unused resources (disks, snapshots, IPs)
- Use Cloud Storage lifecycle policies
- Monitor costs with Cloud Billing reports and budgets
- Use sustained use discounts automatically applied

## DevOps Integration

- Use Cloud Build for CI/CD pipelines
- Store infrastructure code in Cloud Source Repositories or GitHub
- Use Artifact Registry for container images and packages
- Implement deployment approval processes
- Use Cloud Deploy for continuous delivery to GKE
