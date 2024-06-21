package bcauth

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/common"

	"github.com/the-witcher-knight/timekeeper/internal/blockchain"
)

func (ctrl controller) DeauthorizeAccount(ctx context.Context, addr common.Address) error {
	if err := ctrl.bc.DeauthorizeAccount(ctx, addr); err != nil {
		if errors.Is(err, blockchain.ErrCurrentAccountNotAuthorized) {
			return ErrCurrentAccountNotAuthorized
		}

		return err
	}

	return nil
}
