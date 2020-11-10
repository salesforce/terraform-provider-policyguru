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
	if len(input.Write) > 0 {
	    inputBody["tagging"] = input.Tagging
	}
	if len(input.PermissionsManagement) > 0 {
	    inputBody["permissions-management"] = input.PermissionsManagement
	}
	if len(input.List) > 0 {
	    inputBody["list"] = input.List
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
