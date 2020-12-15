# terraform-provider-policyguru

This is the Terraform Provider for [Policy Sentry](https://github.com/salesforce/policy_sentry/) - the IAM Least Privilege Policy Generator.

We have Policy Sentry hosted as a REST API and this Terraform provider points to the REST API. Using this Terraform provider, you can write Least Privilege IAM Policies without ever leaving your code editor! 

## Example

Consider the following example Terraform code:

```hcl-terraform
terraform {
  required_providers {
    policyguru = {
      source  = "salesforce/policyguru"
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

# This creates the policy document.
data "policyguru_document" "example" {
  actions_for_resources_at_access_level {
    write = list("arn:aws:kms:us-east-1:123456789012:key/aaaa-bbbb-cccc")
    read = list("arn:aws:s3:::mybucket/*")
  }

  actions_for_service_without_resource_constraint_support {
    include_single_actions = ["ssm:GetParameter"]
  }
  exclude_actions = list("kms:Decrypt*", "kms:Delete*", "kms:Disable", "kms:Schedule*")
}
# This creates an IAM Policy leveraging policyguru
resource "aws_iam_policy" "policy" {
  name        = "policyguru-example"
  path        = "/"
  description = "this uses the policyguru document data source"
  policy      = data.policyguru_document.example.json
}

# This shows you the output value
output "policy_document_json" {
  value = data.policyguru_document.example.json
}
```

That will create an IAM Policy titled `policyguru-example` which will contain the following policy:

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "S3ReadObject",
            "Effect": "Allow",
            "Action": [
                "s3:GetObject",
                "s3:GetObjectAcl",
                "s3:GetObjectLegalHold",
                "s3:GetObjectRetention",
                "s3:GetObjectTagging",
                "s3:GetObjectTorrent",
                "s3:GetObjectVersion",
                "s3:GetObjectVersionAcl",
                "s3:GetObjectVersionForReplication",
                "s3:GetObjectVersionTagging",
                "s3:GetObjectVersionTorrent"
            ],
            "Resource": [
                "arn:aws:s3:::mybucket/*"
            ]
        },
        {
            "Sid": "KmsWriteKey",
            "Effect": "Allow",
            "Action": [
                "kms:CancelKeyDeletion",
                "kms:CreateAlias",
                "kms:DisableKey",
                "kms:DisableKeyRotation",
                "kms:EnableKey",
                "kms:EnableKeyRotation",
                "kms:Encrypt",
                "kms:GenerateDataKey",
                "kms:GenerateDataKeyPair",
                "kms:GenerateDataKeyPairWithoutPlaintext",
                "kms:GenerateDataKeyWithoutPlaintext",
                "kms:ImportKeyMaterial",
                "kms:ReEncryptFrom",
                "kms:ReEncryptTo",
                "kms:Sign",
                "kms:UpdateAlias",
                "kms:UpdateKeyDescription",
                "kms:Verify"
            ],
            "Resource": [
                "arn:aws:kms:us-east-1:123456789012:key/aaaa-bbbb-cccc"
            ]
        },
        {
            "Sid": "SkipResourceConstraints",
            "Effect": "Allow",
            "Action": [
                "ssm:GetParameter"
            ],
            "Resource": [
                "*"
            ]
        }
    ]
}
```


## Contributing

Command             |      Description
--------------------| --------------------
```make build ```         | Build
```make install  ```      | Build and install
```make fmt```            | Format code
```make lint```           | Run golint
```make terraform-demo``` | Install and run terraform apply for file under examples folder


For documentation on using provider, check https://registry.terraform.io/providers/salesforce/policyguru/latest/docs