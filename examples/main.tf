terraform {
  required_providers {
    policy-sentry = {
      source = "reetasingh/policy-sentry"
      version = "1.0.8"
    }
  }
}

provider "policy-sentry" {
  # Configuration options
}

data "policy-sentry_document" "example" {
    read = list("arn:aws:ssm:us-east-1:123456789012:parameter/myparameter")
    write = list("arn:aws:ssm:us-east-1:123456789012:parameter/myparameter")
}

# Returns policy sentry document in json
output "policy-sentry_document_json" {
  value = data.policy-sentry_document.example.json
}