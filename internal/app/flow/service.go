package flow

import (
	"context"
	"fmt"

	"github.com/neatflowcv/ask/internal/pkg/gemini"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Ask(ctx context.Context, prompt string) (string, error) {
	client := gemini.NewClient()

	answer, err := client.Ask(ctx, prompt)
	if err != nil {
		return "", fmt.Errorf("ask: %w", err)
	}

	return answer, nil
}
