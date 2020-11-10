terraform {
  required_providers {
    policy-sentry = {
      source = "reetasingh/policy-sentry"
      version = "1.1.5"
    }
  }
}

provider "policy-sentry" {
  # Configuration options
}

data "policy-sentry_document" "example" {
    write = list("arn:aws:kms:us-east-1:123456789012:key/aaaa-bbbb-cccc")
    exclude_actions = list("exclude-actions" , "kms:Delete*" , "kms:Disable*", "kms:Schedule*")
}

# Returns policy sentry document in json
output "policy-sentry_document_json" {
  value = data.policy-sentry_document.example.json
}