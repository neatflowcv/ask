package inquirer

import (
	"context"
	"io"
)

type Inquirer interface {
	Ask(ctx context.Context, prompt string, writer io.Writer) error
}
