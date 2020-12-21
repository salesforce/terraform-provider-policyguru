package policygururest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	HttpClient *http.Client
	Endpoint   string
}

func NewClient(endpoint string) *Client {
	return &Client{
		HttpClient: http.DefaultClient,
		Endpoint:   endpoint,
	}
}

func (c *Client) newRequest(requestBody []byte) (*http.Request, error) {

	req, err := http.NewRequest("POST", c.Endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

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
	}
	return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
}
