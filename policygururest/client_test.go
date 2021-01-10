package policygururest

import "testing"

type ClientInput struct {
	endpoint string
	expected string
}

func TestNewClient(t *testing.T) {
	testDataItems := []ClientInput{
		{endpoint: "http:/api.policyguru/write", expected: "http:/api.policyguru/write"},
		{endpoint: "", expected: "https://api.policyguru.io/write-iam-policy"},
	}

	for _, item := range testDataItems {
		client := NewClient(item.endpoint)
		if client.Endpoint != item.expected {
			t.Errorf("Client endpoint not set correctly. Expected %v but got %v", item.expected, client.Endpoint)
		}
	}
}
