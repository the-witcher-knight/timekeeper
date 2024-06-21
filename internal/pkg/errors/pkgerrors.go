package errors

import (
	pkgerrors "github.com/pkg/errors"
)

func Wrap(err error, msg string) error {
	return pkgerrors.Wrap(err, msg)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return pkgerrors.Wrapf(err, format, args...)
}

func WithStack(err error) error {
	return pkgerrors.WithStack(err)
}
