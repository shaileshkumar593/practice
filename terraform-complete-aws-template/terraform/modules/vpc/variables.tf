variable "name" {
  description = "Name prefix."
  type        = string
}

variable "vpc_cidr" {
  description = "VPC CIDR."
  type        = string
}

variable "public_subnet_cidr" {
  description = "Public subnet CIDR."
  type        = string
}

variable "availability_zone" {
  description = "Availability zone."
  type        = string
}
