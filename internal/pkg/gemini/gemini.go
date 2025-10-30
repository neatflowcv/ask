package gemini

import (
	"context"
	"fmt"

	"github.com/neatflowcv/ask/internal/pkg/inquirer"
	"google.golang.org/genai"
)

var _ inquirer.Inquirer = (*Client)(nil)

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}

func (c *Client) Ask(ctx context.Context, prompt string, channel chan<- string) error {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{ //nolint:exhaustruct
		APIKey:  c.apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return fmt.Errorf("new client: %w", err)
	}

	stream := client.Models.GenerateContentStream(
		ctx,
		"gemini-2.5-flash",
		genai.Text(prompt),
		nil,
	)

	if err != nil {
		return fmt.Errorf("generate content: %w", err)
	}

	for chunk, err := range stream {
		if err != nil {
			return fmt.Errorf("generate content: %w", err)
		}

		channel <- chunk.Text()
	}

	return nil
}
