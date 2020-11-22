terraform {
  required_providers {
    policy-sentry = {
      source = "reetasingh/policy-sentry"
      version = "1.2.4"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}
# Configure the AWS Provider
provider "aws" {
  region = "us-east-1"
}
provider "policy-sentry" {
  endpoint = "https://rycbfaz4wl.execute-api.us-east-1.amazonaws.com/dev/write"
}
data "policy-sentry_document" "example" {
    exclude_actions = list("s3:GetAccelerateConfiguration", "s3:GetAnalyticsConfiguration")
}
# Returns policy sentry document in json
output "policy-sentry_document_json" {
  value = data.policy-sentry_document.example.json
}
resource "aws_iam_policy" "policy" {
  name        = "sample"
  path        = "/"
  description = "this uses policy sentry document"
  policy      = data.policy-sentry_document.example.json
}