package openai

const apiURL = "https://api.openai.com/v1/chat/competions"

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}

