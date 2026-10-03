# AWS Terraform

Learning-oriented AWS foundation: VPC + public/private subnets. For a complete production platform, add EKS, ECR, RDS, IAM/Pod Identity, Secrets Manager, KMS, ALB Controller, CloudWatch and WAF/CloudFront as required.

Typical flow:
```text
Route53 -> CloudFront/WAF -> ALB -> EKS -> FastAPI -> RDS
```

```bash
terraform init
terraform plan
terraform apply
```
