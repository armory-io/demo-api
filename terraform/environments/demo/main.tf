terraform {
  backend "s3" {}
}

provider "aws" {
  region  = "us-west-2"
  profile = "prod"
}

variable "environment_name" {
  default = "demo-api-demo"
}

resource "aws_s3_bucket" "b" {
  bucket = "${var.environment_name}-bucket"
  acl    = "public-read"

  tags = {
    Name = "Bucket for ${var.environment_name}"
  }
}
