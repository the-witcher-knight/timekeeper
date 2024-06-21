package blockchain

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	pkgerrors "github.com/the-witcher-knight/timekeeper/internal/pkg/errors"
)

func (bc blockchain) AuthorizeAccount(ctx context.Context, account common.Address) error {
	_, err := bc.attContract.AuthorizeAccount(bc.transactor, account, true)
	if err != nil {
		if strings.Contains(err.Error(), errOnlyOwner) {
			return ErrCurrentAccountNotAuthorized
		}

		return pkgerrors.WithStack(err)
	}

	return nil
}

func (bc blockchain) DeauthorizeAccount(ctx context.Context, account common.Address) error {
	_, err := bc.attContract.AuthorizeAccount(bc.transactor, account, false)
	if err != nil {
		if strings.Contains(err.Error(), errOnlyOwner) {
			return ErrCurrentAccountNotAuthorized
		}

		return pkgerrors.WithStack(err)
	}

	return nil
}
