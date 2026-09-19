output "vpc_id" {
  value = module.stack.vpc_id
}

output "ec2_instance_id" {
  value = module.stack.ec2_instance_id
}

output "ec2_public_ip" {
  value = module.stack.ec2_public_ip
}
