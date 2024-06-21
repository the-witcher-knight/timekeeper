package httpio

import (
	"context"
)

// Validator is an object that can be validated.
type Validator interface {
	Valid(ctx context.Context) error
}
