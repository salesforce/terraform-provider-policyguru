terraform {
  required_providers {
    policy-sentry = {
      source = "reetasingh/policy-sentry"
      version = "1.1.6"
    }
  }
}

provider "policy-sentry" {
  # Configuration options
}

data "policy-sentry_document" "example" {
    write = list("arn:aws:kms:us-east-1:123456789012:key/aaaa-bbbb-cccc")
    exclude_actions = list("s3:GetAccelerateConfiguration", "s3:GetAnalyticsConfiguration")
    read = list("arn:aws:s3:::mybucket")
    tagging = list("arn:aws:s3:::mybucket")
    permissions_management = list("arn:aws:s3:::mybucket")
}

# Returns policy sentry document in json
output "policy-sentry_document_json" {
  value = data.policy-sentry_document.example.json
}