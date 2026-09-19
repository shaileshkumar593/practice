module "stack" {
  source = "../../"

  aws_region         = "ap-south-1"
  environment        = "dev"
  project_name       = "terraform-demo"
  vpc_cidr           = "10.10.0.0/16"
  public_subnet_cidr = "10.10.1.0/24"
  availability_zone  = "ap-south-1a"
  instance_type      = "t3.micro"
}
