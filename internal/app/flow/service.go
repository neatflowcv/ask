package flow

import (
	"context"
	"fmt"

	"github.com/neatflowcv/ask/internal/pkg/inquirer"
)

type Service struct {
	inquirer inquirer.Inquirer
}

func NewService(inquirer inquirer.Inquirer) *Service {
	return &Service{
		inquirer: inquirer,
	}
}

func (s *Service) Ask(ctx context.Context, prompt string) (string, error) {
	answer, err := s.inquirer.Ask(ctx, prompt)
	if err != nil {
		return "", fmt.Errorf("ask: %w", err)
	}

	return answer, nil
}
