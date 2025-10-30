package gemini

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/genai"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Ask(ctx context.Context, prompt string) (string, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{ //nolint:exhaustruct
		APIKey:  os.Getenv("KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return "", fmt.Errorf("new client: %w", err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("generate content: %w", err)
	}

	return result.Text(), nil
}
