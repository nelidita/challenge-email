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
