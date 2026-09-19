# Terraform Complete AWS Template

A beginner-to-advanced Terraform project structure using reusable VPC and EC2 modules, with separate dev/staging/prod environment roots.

## Prerequisites

- Terraform >= 1.6
- AWS CLI configured
- An AWS account with permissions to create VPC/EC2 resources

## Structure

```text
terraform/
├── main.tf
├── variables.tf
├── outputs.tf
├── providers.tf
├── versions.tf
├── locals.tf
├── data.tf
├── terraform.tfvars.example
├── backend.tf.example
├── modules/
│   ├── vpc/
│   │   ├── main.tf
│   │   ├── variables.tf
│   │   └── outputs.tf
│   └── ec2/
│       ├── main.tf
│       ├── variables.tf
│       └── outputs.tf
└── environments/
    ├── dev/
    ├── staging/
    └── prod/
```

## Quick start

The root configuration is a simple learning example. Copy the example variables file:

```bash
cp terraform.tfvars.example terraform.tfvars
```

Set a unique S3 bucket name if using the example.

Then:

```bash
terraform init
terraform fmt -recursive
terraform validate
terraform plan
terraform apply
terraform destroy
```

## Recommended environment workflow

Each environment is an independent root module and can have its own backend/state.

Example:

```bash
cd environments/dev
terraform init
terraform fmt -recursive
terraform validate
terraform plan
terraform apply
```

The dev/staging/prod examples use the same VPC and EC2 modules with different variables.

## Important

The example backend is intentionally disabled as a real backend because a remote state bucket must exist before Terraform can use it. Copy `backend.tf.example` to `backend.tf` only after creating/configuring your remote state backend.

Do not commit:

- `terraform.tfstate`
- `terraform.tfstate.*`
- `.terraform/`
- `terraform.tfvars` when it contains secrets

The example does not contain AWS credentials.
