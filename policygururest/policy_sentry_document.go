package policygururest

import (
	"encoding/json"
	"fmt"
)

const policyDocumentPath string = "write"
const defaultRestUrl string = "https://api.policyguru.io/"

// GetPolicyDocument returns policy document by calling the REST API
func (c *Client) GetPolicyDocument(input *PolicyDocumentInput) (*PolicyDocument, error) {
	var policyDocument PolicyDocument

	var inputBody map[string]interface{} = make(map[string]interface{})

	if input.ActionsForServices != nil {
		if input.ActionsForServices.Read != nil && len(input.ActionsForServices.Read) > 0 {
			inputBody["service-read"] = input.ActionsForServices.Read
		}
		if input.ActionsForServices.Write != nil && len(input.ActionsForServices.Write) > 0 {
			inputBody["service-write"] = input.ActionsForServices.Write
		}
		if input.ActionsForServices.Tagging != nil && len(input.ActionsForServices.Tagging) > 0 {
			inputBody["service-tagging"] = input.ActionsForServices.Tagging
		}
		if input.ActionsForServices.PermissionsManagement != nil && len(input.ActionsForServices.PermissionsManagement) > 0 {
			inputBody["service-permissions-management"] = input.ActionsForServices.PermissionsManagement
		}
		if input.ActionsForServices.List != nil && len(input.ActionsForServices.List) > 0 {
			inputBody["service-list"] = input.ActionsForServices.List
		}
		if input.ActionsForServices.SingleActions != nil && len(input.ActionsForServices.SingleActions) > 0 {
			inputBody["single-actions"] = input.ActionsForServices.SingleActions
		}
	}

	if input.ActionsForResources != nil {
		if input.ActionsForResources.Read != nil && len(input.ActionsForResources.Read) > 0 {
			inputBody["read"] = input.ActionsForResources.Read
		}
		if input.ActionsForResources.Write != nil && len(input.ActionsForResources.Write) > 0 {
			inputBody["write"] = input.ActionsForResources.Write
		}
		if input.ActionsForResources.Tagging != nil && len(input.ActionsForResources.Tagging) > 0 {
			inputBody["tagging"] = input.ActionsForResources.Tagging
		}
		if input.ActionsForResources.PermissionsManagement != nil && len(input.ActionsForResources.PermissionsManagement) > 0 {
			inputBody["permissions-management"] = input.ActionsForResources.PermissionsManagement
		}
		if input.ActionsForResources.List != nil && len(input.ActionsForResources.List) > 0 {
			inputBody["list"] = input.ActionsForResources.List
		}
	}

	if input.Overrides != nil {
		if input.Overrides.SkipResourceConstraints != nil && len(input.Overrides.SkipResourceConstraints) > 0 {
			inputBody["skip-resource-constraints"] = input.Overrides.SkipResourceConstraints
		}
	}

	if len(input.ExcludeActions) > 0 {
		inputBody["exclude-actions"] = input.ExcludeActions
	}

	fmt.Println(inputBody)

	requestBody, err := json.Marshal(inputBody)

	if err != nil {
		return nil, err
	}

	if len(c.Endpoint) == 0 {
		c.Endpoint = defaultRestUrl + policyDocumentPath
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

// GetPolicyDocumentJsonString returns json form of policy document
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
