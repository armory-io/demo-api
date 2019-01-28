terraform {
  backend "s3" {}
}

provider "aws" {
  region = "us-west-2"
  profile = "prod"
}

variable "cluster_name" {
  default = "demo-api-demo"
}

module "cache" {
  source = "../../modules/elasticache"
  cluster_id = "${var.cluster_name}"
}