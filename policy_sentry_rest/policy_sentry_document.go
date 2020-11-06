package policy_sentry_rest

import "encoding/json"
import "strconv"
import "net/http"
import "fmt"

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

func (c *Client) GetPolicyDocument() (*PolicyDocument, error) {
	var policyDocument policyDocument

	requestBody, err := json.Marshal(map[string] interface{} {
    "mode": "crud",
    "read": []string{"arn:aws:s3:::example-org-s3-access-logs"} })

    if err != nil {
    return diag.FromErr(err)
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

func (c *Client) GetPolicyDocumentJsonString() (string, error) {
	var policyDocument policyDocument
	policyDocument, err := GetPolicyDocument()

	if err != nil {
		return nil, err
	}

	jsonDoc, err := json.Marshal(map[string] interface{} {
	    policyDocument
	}
    }
    if err != nil {
		return nil, err
	}

	jsonString := string(jsonDoc)

	return jsonString, err
}