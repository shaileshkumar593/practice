module "vpc" {
  source = "./modules/vpc"

  name                = local.name_prefix
  vpc_cidr            = var.vpc_cidr
  public_subnet_cidr  = var.public_subnet_cidr
  availability_zone   = var.availability_zone
}

module "ec2" {
  source = "./modules/ec2"

  name                 = local.name_prefix
  ami_id               = data.aws_ami.amazon_linux.id
  instance_type        = var.instance_type
  subnet_id            = module.vpc.public_subnet_id
  vpc_id               = module.vpc.vpc_id
  key_name             = var.key_name
}
