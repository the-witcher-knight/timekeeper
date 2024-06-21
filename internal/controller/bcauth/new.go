package bcauth

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	"github.com/the-witcher-knight/timekeeper/internal/blockchain"
)

type Controller interface {
	// AuthorizeAccount grants permission to the given blockchain account address
	AuthorizeAccount(ctx context.Context, account common.Address) error

	// DeauthorizeAccount revokes permission for the given blockchain account address
	DeauthorizeAccount(ctx context.Context, account common.Address) error
}

type controller struct {
	bc blockchain.Blockchain
}

func New(bc blockchain.Blockchain) Controller {
	return &controller{
		bc: bc,
	}
}
