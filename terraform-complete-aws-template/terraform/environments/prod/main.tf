module "stack" {
  source = "../../"

  aws_region         = "ap-south-1"
  environment        = "prod"
  project_name       = "terraform-demo"
  vpc_cidr           = "10.30.0.0/16"
  public_subnet_cidr = "10.30.1.0/24"
  availability_zone  = "ap-south-1a"
  instance_type      = "t3.small"
}
