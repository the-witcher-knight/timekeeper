package bcwatch

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/the-witcher-knight/timekeeper/internal/blockchain"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/logging"
	"github.com/the-witcher-knight/timekeeper/internal/repository"
)

type Controller interface {
	// WatchAttendanceRecorded watches for attendance recorded event
	WatchAttendanceRecorded() func(context.Context) error

	// WatchAttendanceUpdated watches for attendance updated event
	WatchAttendanceUpdated() func(context.Context) error
}

type controller struct {
	logger    logging.Logger
	ethClient *ethclient.Client
	bc        blockchain.Blockchain
	repo      repository.Registry
}

func New(
	logger logging.Logger,
	wsClient ethclient.Client,
	bc blockchain.Blockchain,
	repo repository.Registry,
) Controller {
	return &controller{
		logger:    logger,
		ethClient: &wsClient,
		bc:        bc,
		repo:      repo,
	}
}
