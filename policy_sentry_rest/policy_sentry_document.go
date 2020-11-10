package policy_sentry_rest

import (
    "encoding/json"
)

type PolicyDocumentInput struct {
    Read   []*string
	Write   []*string
	Tagging []*string
	List    []*string
	PermissionsManagement []*string
	ExcludeActions []*string
	SkipResourceConstraints []*string
	ServiceRead []*string
	ServiceWrite []*string
	ServiceList []*string
	ServiceTagging []*string
	ServicePermissionsManagement []*string
	SingleActions []*string
}

type PolicyDocument struct {
	Statement []struct {
		Action   []string `json:"Action"`
		Effect   string   `json:"Effect"`
		Resource []string `json:"Resource"`
		Sid      string   `json:"Sid"`
	} `json:"Statement"`
	Version string `json:"Version"`
}

const PolicyDocumentPath string = "write"

func (c *Client) GetPolicyDocument(input *PolicyDocumentInput) (*PolicyDocument, error) {
	var policyDocument PolicyDocument

	var inputBody map[string]interface{} =  make(map[string]interface{})
	if len(input.Read) > 0 {
	    inputBody["read"] = input.Read
	}
	if len(input.Write) > 0 {
	    inputBody["write"] = input.Write
	}
	if len(input.Tagging) > 0 {
	    inputBody["tagging"] = input.Tagging
	}
	if len(input.PermissionsManagement) > 0 {
	    inputBody["permissions-management"] = input.PermissionsManagement
	}
	if len(input.List) > 0 {
	    inputBody["list"] = input.List
	}
	if len(input.SkipResourceConstraints) > 0 {
	    inputBody["skip-resource-constraints"] = input.SkipResourceConstraints
	}
	if len(input.ExcludeActions) > 0 {
	    inputBody["exclude-actions"] = input.ExcludeActions
	}
	if len(input.ServiceRead) > 0 {
	    inputBody["service-read"] = input.ServiceRead
	}
	if len(input.ServiceWrite) > 0 {
	    inputBody["service-write"] = input.ServiceWrite
	}
	if len(input.ServiceList) > 0 {
	    inputBody["service-list"] = input.ServiceList
	}
	if len(input.ServiceTagging) > 0 {
	    inputBody["service-tagging"] = input.ServiceTagging
	}
	if len(input.ServicePermissionsManagement) > 0 {
	    inputBody["service-permissions-management"] = input.ServicePermissionsManagement
	}
	if len(input.SingleActions) > 0 {
	    inputBody["single-actions"] = input.SingleActions
	}

	requestBody, err := json.Marshal(inputBody)

	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(PolicyDocumentPath, requestBody)
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
