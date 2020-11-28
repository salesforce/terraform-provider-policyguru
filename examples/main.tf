terraform {
  required_providers {
    policy-sentry = {
      source = "reetasingh/policy-sentry"
      version = "1.3.1"
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
    actions_for_service_without_resource_constraint_support {
        write = list("s3")
        list = list("s3")
        read = ["s3"]
        include_single_actions = ["ssm:GetParameter"]
  }
  exclude_actions = list("kms:Delete*", "kms:Disable*", "kms:Schedule*")

}
# Returns policy sentry document in json
output "policy-sentry_document_json" {
  value = data.policy-sentry_document.example.json
}
resource "aws_iam_policy" "policy" {
  name        = "sample2"
  path        = "/"
  description = "this uses policy sentry document"
  policy      = data.policy-sentry_document.example.json
}