package openai

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClientWithAPIKey(t *testing.T) {
	apiKey := "example-key"
	apiURL := "https://api.openai.com/v1/chat/competions"
	client := NewClient(apiKey, apiURL)

	if client.apiKey != apiKey {
		t.Errorf("expected apiKey to be %s, but got %s", apiKey, client.apiKey)
	}
}

func TestGetChatResponse(t *testing.T) {
	mockResponse := ChatResponse{
		Choices: []Choice{
			{
				Message: Message{Role: "assistant", Content: "I'm good, thank you!"},
			},
			{
				Message: Message{Role: "assistant", Content: "Context second"},
			},
		},
	}

	mockResponseData, _ := json.Marshal(mockResponse)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer test-api-key" {
			t.Fatalf("Expected Authorization header to be 'Bearer test-api-key', got %v", r.Header.Get("Authorization"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(mockResponseData)
	}))
	defer server.Close()

	client := NewClient("test-api-key", server.URL)

	response, err := client.GetChatResponse("Hello, how are you?")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResponse := "I'm good, thank you!"
	if response != expectedResponse {
		t.Fatalf("Expected response to be %v, got %v", expectedResponse, response)
	}
}
