output "vpc_id" {
  description = "VPC ID."
  value       = module.vpc.vpc_id
}

output "public_subnet_id" {
  description = "Public subnet ID."
  value       = module.vpc.public_subnet_id
}

output "ec2_instance_id" {
  description = "EC2 instance ID."
  value       = module.ec2.instance_id
}

output "ec2_public_ip" {
  description = "EC2 public IPv4 address."
  value       = module.ec2.public_ip
}
