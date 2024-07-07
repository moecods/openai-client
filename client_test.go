package openai

import (
	"testing"
)

func TestNewClientWithAPIKey(t *testing.T) {
	apiKey := "example-key"
	client := NewClient(apiKey)

	if client.apiKey != apiKey {
		t.Errorf("expected apiKey to be %s, but got %s", apiKey, client.apiKey)
	}
}