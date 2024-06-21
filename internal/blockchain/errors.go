package blockchain

import (
	"errors"
)

const (
	errOnlyOwner      = "only owner can perform this action"
	errNotAuthorized  = "not authorized to perform this action"
	errRecordNotFound = "record not found"
)

var (
	ErrCurrentAccountNotAuthorized = errors.New("current account is not authorized to perform this action")
	ErrRecordNotFound              = errors.New("record not found")
)
