data "aws_vpc" "selected" {
  id = var.vpc_id
}

resource "aws_security_group" "ec2" {
  name        = "${var.name}-instance-sg"
  description = "Security group for EC2 instance."
  vpc_id      = data.aws_vpc.selected.id

  ingress {
    description = "SSH - restrict this CIDR in production."
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    description = "Allow outbound traffic."
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${var.name}-instance-sg"
  }
}

resource "aws_instance" "this" {
  ami                         = var.ami_id
  instance_type               = var.instance_type
  subnet_id                   = var.subnet_id
  key_name                    = var.key_name
  vpc_security_group_ids      = [aws_security_group.ec2.id]
  associate_public_ip_address = true

  metadata_options {
    http_endpoint = "enabled"
    http_tokens   = "required"
  }

  tags = {
    Name = "${var.name}-ec2"
  }
}
