package group

import (
	"fmt"

	"github.com/neatflowcv/ask/internal/pkg/printer"
)

var _ printer.Printer = (*Printer)(nil)

type Printer struct {
	printers []printer.Printer
}

func NewPrinter(printers ...printer.Printer) *Printer {
	return &Printer{
		printers: printers,
	}
}

func (p *Printer) Print(item string) error {
	for _, printer := range p.printers {
		err := printer.Print(item)
		if err != nil {
			return fmt.Errorf("print: %w", err)
		}
	}

	return nil
}
