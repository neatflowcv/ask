package flow

import (
	"context"
	"fmt"
	"log"

	"github.com/neatflowcv/ask/internal/pkg/inquirer"
	"github.com/neatflowcv/ask/internal/pkg/printer"
)

type Service struct {
	inquirer inquirer.Inquirer
	printer  printer.Printer
}

func NewService(inquirer inquirer.Inquirer, printer printer.Printer) *Service {
	return &Service{
		inquirer: inquirer,
		printer:  printer,
	}
}

func (s *Service) Ask(ctx context.Context, prompt string) error {
	channel := make(chan string)

	go func() {
		for content := range channel {
			err := s.printer.Print(content)
			if err != nil {
				log.Printf("print: %v", err)

				return
			}
		}
	}()

	err := s.inquirer.Ask(ctx, prompt, channel)
	if err != nil {
		return fmt.Errorf("ask: %w", err)
	}

	close(channel)

	return nil
}
