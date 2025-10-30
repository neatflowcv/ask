package flow

import (
	"bytes"
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
	var buf bytes.Buffer

	err := s.inquirer.Ask(ctx, prompt, &buf)
	if err != nil {
		return "", fmt.Errorf("ask: %w", err)
	}

	return buf.String(), nil
}
