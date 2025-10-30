package console

import (
	"fmt"

	"github.com/neatflowcv/ask/internal/pkg/printer"
)

var _ printer.Printer = (*Printer)(nil)

type Printer struct {
}

func NewPrinter() *Printer {
	return &Printer{}
}

func (p *Printer) Print(item string) error {
	fmt.Println(item) //nolint:forbidigo

	return nil
}
