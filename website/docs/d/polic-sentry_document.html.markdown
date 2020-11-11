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
    write = list("arn:aws:kms:us-east-1:123456789012:key/aaaa-bbbb-cccc")
    exclude_actions = list("s3:GetAccelerateConfiguration", "s3:GetAnalyticsConfiguration")
    read = list("arn:aws:s3:::mybucket")
    tagging = list("arn:aws:s3:::mybucket")
    permissions_management = list("arn:aws:s3:::mybucket")
}

resource "aws_iam_policy" "example" {
  name   = "example_policy"
  path   = "/"
  policy = data.policy-sentry_document.example.json
}
```

## Argument Reference

The following arguments are supported:

* `policy_id` (Optional) - An ID for the policy document.
* `source_json` (Optional) - An IAM policy document to import as a base for the
  current policy document.  Statements with non-blank `sid`s in the current
  policy document will overwrite statements with the same `sid` in the source
  json.  Statements without an `sid` cannot be overwritten.
* `override_json` (Optional) - An IAM policy document to import and override the
  current policy document.  Statements with non-blank `sid`s in the override
  document will overwrite statements with the same `sid` in the current document.
  Statements without an `sid` cannot be overwritten.
* `statement` (Optional) - A nested configuration block (described below)
  configuring one *statement* to be included in the policy document.
* `version` (Optional) - IAM policy document version. Valid values: `2008-10-17`, `2012-10-17`. Defaults to `2012-10-17`. For more information, see the [AWS IAM User Guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_version.html).

Each document configuration may have one or more `statement` blocks, which
each accept the following arguments:


Each policy may have either zero or more `principals` blocks or zero or more
`not_principals` blocks, both of which each accept the following arguments:

* `type` (Required) The type of principal. For AWS ARNs this is "AWS".  For AWS services (e.g. Lambda), this is "Service". For Federated access the type is "Federated".
* `identifiers` (Required) List of identifiers for principals. When `type`
  is "AWS", these are IAM user or role ARNs.  When `type` is "Service", these are AWS Service roles e.g. `lambda.amazonaws.com`. When `type` is "Federated", these are web identity users or SAML provider ARNs.

## Attributes Reference

The following attribute is exported:

* `json` - The above arguments serialized as a standard JSON policy document.
