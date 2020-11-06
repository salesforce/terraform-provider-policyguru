terraform {
  required_providers {
    policy-sentry = {
      source = "reetasingh/policy-sentry"
      version = "1.0.2"
    }
  }
}

provider "policy-sentry" {
  # Configuration options
}

data "policy-sentry_document" "example" {}

# Returns policy sentry document in json
output "policy-sentry_document_json" {
  value = data.policy-sentry_document.example.json
}