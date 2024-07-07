package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	apiKey string
	apiURL string
}

func NewClient(apiKey, apiURL string) *Client {
	return &Client{apiKey: apiKey, apiURL: apiURL}
}

func (c *Client) GetChatResponse(message string) (string, error) {
	requestBody := ChatRequest{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{Role: "system", Content: message},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshalling request body: %v", err)
	}

	req, err := http.NewRequest("POST", c.apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error createing request: %v", err)
	}

	req.Header.Set("Context-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request failed with status %d", resp.StatusCode)
	}

	var chatResponse ChatResponse
	err = json.NewDecoder(resp.Body).Decode(&chatResponse)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response body: %v", err)
	}

	if len(chatResponse.Choices) > 0 {
		return chatResponse.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response from OpenAI")
}
