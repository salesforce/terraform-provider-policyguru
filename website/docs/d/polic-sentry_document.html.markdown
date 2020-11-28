---
subcategory: "IAM"
layout: "policy-sentry"
page_title: "Policy-Sentry: policy-sentry_document"
description: |-
  Generates an IAM policy document in JSON format
---

# Data Source: policy-sentry_document

Generates an IAM policy document in JSON format.

This is a data source which can be used to construct a JSON representation of
an IAM policy document, for use with resources which expect policy documents,
such as the `aws_iam_policy` resource.

-> For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://learn.hashicorp.com/terraform/aws/iam-policy).

```hcl
data "policy-sentry_document" "example" {
  actions_for_resources_at_access_level {
    write = list("arn:aws:kms:us-east-1:123456789012:key/aaaa-bbbb-cccc")
    read = list("arn:aws:s3:::mybucket")
    tagging = list("arn:aws:s3:::mybucket")
    permissions_management = list("arn:aws:s3:::mybucket")
  }

  actions_for_service_without_resource_constraint_support {
    write = list("s3")
    list = list("s3")
    read = ["s3"]
    include_single_actions = ["ssm:GetParameter"]
  }

  overrides {
    skip_resource_constraints_for_actions = []
  }

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
```

## Argument Reference

The following arguments are supported:

* `exclude_actions` (Optional) - A list of actions that should not be included in the resulting policy.
* `actions_for_resources_at_access_level` (Optional) - Provide Information about list of Amazon Resource Names (ARNs) that your role needs access to. actions_for_resources_at_access_level block is documented below.
* `actions_for_service_without_resource_constraint_support` (Optional) - Provide Information about AWS service actions that do not support resource ARN constraints. actions_for_service_without_resource_constraint_support block is documented below.
* `overrides` (Optional) - Provide Information about services to include actions that do not support resource ARN constraints. actions_for_service_without_resource_constraint_support block is documented below.

## Attributes Reference

The following attribute is exported:

* `json` - The above arguments serialized as a standard JSON policy document.