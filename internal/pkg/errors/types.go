package errors

import (
	pkgerrors "github.com/pkg/errors"
)

type StackTracer interface {
	StackTrace() pkgerrors.StackTrace
}

type WrappedError interface {
	Unwrap() error
}
