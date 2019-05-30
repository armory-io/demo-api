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
  bucket = "${var.environment_name}-branch-bucket"
  acl    = "public-read"

  tags = {
    Name = "Bucket for ${var.environment_name}"
  }
}

output "s3_bucket_arn" {
  value = "${aws_s3_bucket.b.arn}"
}
