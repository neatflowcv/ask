package file

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/neatflowcv/ask/internal/pkg/printer"
)

var _ printer.Printer = (*Printer)(nil)

type Printer struct {
	file *os.File
}

func NewPrinter(filename string) (*Printer, error) {
	cleaned := filepath.Clean(filename)

	file, err := os.Create(cleaned)
	if err != nil {
		return nil, fmt.Errorf("create file: %w", err)
	}

	return &Printer{
		file: file,
	}, nil
}

func (p *Printer) Print(item string) error {
	_, err := p.file.WriteString(item)
	if err != nil {
		return fmt.Errorf("write string: %w", err)
	}

	return nil
}

func (p *Printer) Close() {
	err := p.file.Close()
	if err != nil {
		log.Printf("close file: %v", err)
	}
}
