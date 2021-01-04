terraform {
  required_providers {
    policyguru = {
      source  = "salesforce/policyguru"
      version = "2.0.4"
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

data "policyguru_document" "example" {
  actions_for_resources_at_access_level {
    write = list("arn:aws:kms:us-east-1:123456789012:key/aaaa-bbbb-cccc")
    read = list("arn:aws:s3:::mybucket/*")
  }

  actions_for_service_without_resource_constraint_support {
    write                  = list("s3")
    list                   = list("s3")
    read                   = ["s3"]
    include_single_actions = ["ssm:GetParameter"]
  }
  exclude_actions = list("kms:Decrypt*", "kms:Delete*", "kms:Disable", "kms:Schedule*")
}

# Returns policyguru document in json
output "policyguru_document_json" {
  value = data.policyguru_document.example.json
}

resource "aws_iam_policy" "policy" {
  name        = "sample2"
  path        = "/"
  description = "this uses policyguru document"
  policy      = data.policyguru_document.example.json
}