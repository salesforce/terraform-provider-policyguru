package main

import (
	"fmt"
	abc "terraform-provider-policy-sentry/policy_sentry_rest"
)

func main() {
	client := abc.NewClient("")
	policyDocumentInput := new(abc.PolicyDocumentInput)

	s := "arn:aws:s3:::mybucket"
	exclude_actions := "s3:PutBucketPublicAccessBlock"

	//policyDocumentInput.Read = []*string{&s}
	//policyDocumentInput.Write = []*string{&s}
	//policyDocumentInput.Tagging = []*string{&s}
	//policyDocumentInput.PermissionsManagement = []*string{&s}
	//policyDocumentInput.List = []*string{&s}
	policyDocumentInput.ExcludeActions = []*string{&exclude_actions}

	actionForServices := new(abc.ActionsForServicesWithoutResourceConstraints)
	actionForResources := new(abc.ActionsForResourcesAtAccessLevel)
	overrides := new(abc.Overrides)

	policyDocumentInput.ActionsForServices = actionForServices
	policyDocumentInput.ActionsForResources = actionForResources
	policyDocumentInput.Overrides = overrides

	policyDocumentInput.ActionsForServices.Read = []*string{&s}

	fmt.Print(policyDocumentInput.ActionsForServices)
	fmt.Print(policyDocumentInput.ActionsForResources)
	fmt.Print(policyDocumentInput.Overrides)
	fmt.Print(policyDocumentInput.ExcludeActions)

	policyDocumentJsonString, err := client.GetPolicyDocumentJsonString(policyDocumentInput)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(policyDocumentJsonString)
}
