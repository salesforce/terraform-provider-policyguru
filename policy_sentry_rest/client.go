package policy_sentry_rest


import (
    "net/http"
    "fmt"
    "io/ioutil"
    "bytes"
)

const DefaultRestUrl string = "https://zeok878mvj.execute-api.us-east-1.amazonaws.com/dev/"

type Client struct {
	HttpClient *http.Client
	Host       string
	Base       string
}

func NewClient() *Client {
	return &Client{
		HttpClient: http.DefaultClient,
	}
}

func (c *Client) newRequest(path string, requestBody []byte) (*http.Request, error) {

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", DefaultRestUrl, path), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent {
		return body, err
	} else {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
}