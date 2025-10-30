package inquirer

import "context"

type Inquirer interface {
	Ask(ctx context.Context, prompt string) (string, error)
}
