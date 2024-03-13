terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"
}

provider "aws" {
  profile = "default"
  region = "us-west-2"
}

resource "aws_instance" "app_server" {
  ami           = "ami-830c94e3"
  instance_type = "t2.micro"
  key_name      = "key"
  security_groups = [aws_security_group.app_server.name]

  user_data = <<-EOF
  #!/bin/bash
  sudo apt-get update
  sudo apt-get install nginx
  echo t/var/www/html/index.nginx-debian.html
  EOF

  tags = {
    Name = "challenge-email"
  }
}

resource "aws_security_group" "app_server" {
  name        = "app_server_challenge_email"
  description = "Security group for the application server"

    ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["190.24.77.141/32"]  # Permitir solo tu dirección IP para SSH
  }

  ingress {
    from_port   = 4080
    to_port     = 4080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 5173
    to_port     = 5173
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "app_server_challenge_email"
  }
}
