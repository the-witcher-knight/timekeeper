package bcauth

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/common"

	"github.com/the-witcher-knight/timekeeper/internal/blockchain"
)

func (ctrl controller) AuthorizeAccount(ctx context.Context, addr common.Address) error {
	if err := ctrl.bc.AuthorizeAccount(ctx, addr); err != nil {
		if errors.Is(err, blockchain.ErrCurrentAccountNotAuthorized) {
			return ErrCurrentAccountNotAuthorized
		}

		return err
	}

	return nil
}
