package policy_sentry_rest

import "encoding/json"

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

func (c *Client) GetPolicyDocument(mode string) (*PolicyDocument, error) {
	var policyDocument PolicyDocument

	inputBody := map[string]interface{}{
		"mode": mode,
		"read": []string{"arn:aws:s3:::example-org-s3-access-logs"}}

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

func (c *Client) GetPolicyDocumentJsonString(mode string) (string, error) {
	policyDocument, err := c.GetPolicyDocument(mode)

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
