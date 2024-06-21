package bcauth

import (
	"errors"
)

var (
	ErrCurrentAccountNotAuthorized = errors.New("account address not authorized")
)
