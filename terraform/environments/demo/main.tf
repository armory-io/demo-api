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

resource "aws_route53_record" "www" {
  zone_id = "Z641P3484FWS1"
  name    = "${var.cluster_name}-cache.tf-resources.com."
  type    = "CNAME"
  ttl     = "300"
  records = ["${module.cache.hostname}"]
}