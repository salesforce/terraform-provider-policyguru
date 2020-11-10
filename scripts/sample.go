package main

import (
	abc "terraform-provider-policy-sentry/policy_sentry_rest"
	"fmt"
)

func main() {
	client := abc.NewClient()
	policyDocumentInput := new(abc.PolicyDocumentInput)

	s := "arn:aws:s3:::mybucket"
	exclude_actions := "s3:PutBucketPublicAccessBlock"

	policyDocumentInput.Read = []*string{&s}
	policyDocumentInput.Write = []*string{&s}
	policyDocumentInput.Tagging = []*string{&s}
	policyDocumentInput.PermissionsManagement = []*string{&s}
	policyDocumentInput.List = []*string{&s}
	policyDocumentInput.ExcludeActions = []*string{&exclude_actions}

	policyDocumentJsonString, err := client.GetPolicyDocumentJsonString(policyDocumentInput)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(policyDocumentJsonString)
}
