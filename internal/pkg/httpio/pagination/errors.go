package pagination

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrPageMustBeGreaterThanZero = errors.New("page must be greater than 0")
	ErrSizeMustBeGreaterThanZero = errors.New("size must be greater than 0")
	ErrExceededMaxPageSizeLimit  = errors.New(fmt.Sprintf("maximum size for each page cannot exceed %v", paginationMaxSizePerPage))
)
