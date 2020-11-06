
provider "policysentry" {}

module "psl" {
  source  = "./policysentry"
}

output "psl" {
  value = module.psl.policysentry_document_json
}