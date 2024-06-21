package config

import (
	"errors"
)

var (
	ErrInvalidProjectName = errors.New("invalid project name")

	ErrInvalidEnvironment = errors.New("environment is invalid")

	ErrInvalidPort = errors.New("port is invalid")

	ErrInvalidDBURL = errors.New("db url is required")

	ErrInvalidBlockChainNetwork = errors.New("blockchain network is required")

	ErrInvalidBlockChainAccount = errors.New("blockchain account is required")

	ErrInvalidBlockChainPrivateKey = errors.New("blockchain private key is required")

	ErrInvalidBlockChainContractAddress = errors.New("blockchain contract address is required")

	ErrInvalidJWTSecret = errors.New("jwt secret is required")
)
