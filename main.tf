terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.49.0"
    }
  }
  required_version = ">= 1.2.0"
}

provider "aws" {
  region = "us-west-2"
}
data "aws_vpc" "default" {
  default = true
}

resource "aws_security_group" "task_manager_sg_tf" {
  name        = "task-manager-sg-tf"
  description = "Allow HTTP & HTTPS to web server"
  vpc_id      = data.aws_vpc.default.id
}

resource "aws_security_group_rule" "allow_http" {
  type              = "ingress"
  description       = "HTTP ingress"
  from_port         = 80
  to_port           = 80
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.task_manager_sg_tf.id
}
resource "aws_security_group_rule" "allow_https" {
  type              = "ingress"
  description       = "HTTPS ingress"
  from_port         = 443
  to_port           = 443
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.task_manager_sg_tf.id
}

resource "aws_security_group_rule" "allow_all" {
  type              = "ingress"
  description       = "allow all"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.task_manager_sg_tf.id
}

resource "aws_security_group_rule" "allow_all_outbound" {
  type              = "egress"
  description       = "allow all"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.task_manager_sg_tf.id
}

resource "aws_instance" "task-manager" {
  ami                         = "ami-0cf2b4e024cdb6960"
  instance_type               = "t3.medium"
  vpc_security_group_ids      = [aws_security_group.task_manager_sg_tf.id]
  associate_public_ip_address = true
  key_name                    = "keypair"
  user_data                   = file("./scripts/init.sh")

  instance_market_options {
    market_type = "spot"
    spot_options {
      instance_interruption_behavior = "hibernate"
      max_price                      = 0.040
      spot_instance_type             = "persistent"
      valid_until                    = "2024-06-20T23:59:00Z"
    }
  }

  tags = {
    Terraform   = "true"
    Name        = "task-manager-test"
    Environment = "dev"
  }
  root_block_device {
    volume_type           = "gp3"
    volume_size           = "30"
    delete_on_termination = true
  }
}
resource "aws_ebs_encryption_by_default" "task-manager-ebs" {
  enabled = true
}
