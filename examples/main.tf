terraform {
  required_providers {
    policysentry = {
      source = "github.com/reetasingh/terraform-provider-policy-sentry/policysentry"
    }
  }
}

provider "policysentry" {}

module "psl" {
  source  = "./policysentry"
}

output "psl" {
  value = module.psl.policysentry_document_json
}