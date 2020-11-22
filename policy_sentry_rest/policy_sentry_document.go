package policy_sentry_rest

import (
	"encoding/json"
)

const PolicyDocumentPath string = "write"
const DefaultRestUrl string = "https://tpzr780iv0.execute-api.us-east-1.amazonaws.com/dev/"

func (c *Client) GetPolicyDocument(input *PolicyDocumentInput) (*PolicyDocument, error) {
	var policyDocument PolicyDocument

	var inputBody map[string]interface{} = make(map[string]interface{})

	if len(input.ActionsForResources.Read) > 0 {
		inputBody["read"] = input.ActionsForResources.Read
	}
	if len(input.ActionsForResources.Write) > 0 {
		inputBody["write"] = input.ActionsForResources.Write
	}
	if len(input.ActionsForResources.Tagging) > 0 {
		inputBody["tagging"] = input.ActionsForResources.Tagging
	}
	if len(input.ActionsForResources.PermissionsManagement) > 0 {
		inputBody["permissions-management"] = input.ActionsForResources.PermissionsManagement
	}
	if len(input.ActionsForResources.List) > 0 {
		inputBody["list"] = input.ActionsForResources.List
	}

	if len(input.ActionsForServices.Read) > 0 {
		inputBody["service-read"] = input.ActionsForServices.Read
	}
	if len(input.ActionsForServices.Write) > 0 {
		inputBody["service-write"] = input.ActionsForServices.Write
	}
	if len(input.ActionsForServices.Tagging) > 0 {
		inputBody["service-tagging"] = input.ActionsForServices.Tagging
	}
	if len(input.ActionsForServices.PermissionsManagement) > 0 {
		inputBody["service-permissions-management"] = input.ActionsForServices.PermissionsManagement
	}
	if len(input.ActionsForServices.List) > 0 {
		inputBody["service-list"] = input.ActionsForServices.List
	}
	if len(input.ActionsForServices.SingleActions) > 0 {
		inputBody["single-actions"] = input.ActionsForServices.SingleActions
	}

	if len(input.Overrides.SkipResourceConstraints) > 0 {
		inputBody["skip-resource-constraints"] = input.Overrides.SkipResourceConstraints
	}

	if len(input.ExcludeActions) > 0 {
		inputBody["exclude-actions"] = input.ExcludeActions
	}

	requestBody, err := json.Marshal(inputBody)

	if err != nil {
		return nil, err
	}

	if len(c.Endpoint) == 0 {
		c.Endpoint = DefaultRestUrl + PolicyDocumentPath
	}

	req, err := c.newRequest(requestBody)
	if err != nil {
		return nil, err
	}

	resBody, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resBody, &policyDocument)
	if err != nil {
		return nil, err
	}

	return &policyDocument, nil
}

func (c *Client) GetPolicyDocumentJsonString(input *PolicyDocumentInput) (string, error) {
	policyDocument, err := c.GetPolicyDocument(input)

	if err != nil {
		return "", err
	}

	jsonDoc, err := json.Marshal(policyDocument)
	if err != nil {
		return "", err
	}

	jsonString := string(jsonDoc)

	return jsonString, err
}
