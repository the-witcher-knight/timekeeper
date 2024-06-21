package blockchain

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/the-witcher-knight/timekeeper/internal/blockchain/contracts"
	"github.com/the-witcher-knight/timekeeper/internal/model"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/config"
	pkgerrors "github.com/the-witcher-knight/timekeeper/internal/pkg/errors"
)

type Blockchain interface {
	// AttendanceContract returns attendance contract instance
	AttendanceContract() contracts.Attendance

	// RecordAttendance records an attendance to blockchain
	RecordAttendance(ctx context.Context, att model.Attendance) error

	// RetrieveAttendance returns a list of attendance by filter given
	RetrieveAttendance(ctx context.Context, filter AttendanceFilter) ([]model.Attendance, error)

	// UpdateAttendance update an attendance
	UpdateAttendance(ctx context.Context, att model.Attendance) error

	// AuthorizeAccount authorize account for a contract
	AuthorizeAccount(ctx context.Context, account common.Address) error

	// DeauthorizeAccount deauthorize account for a contract
	DeauthorizeAccount(ctx context.Context, account common.Address) error
}

type blockchain struct {
	client      bind.DeployBackend
	transactor  *bind.TransactOpts
	attContract *contracts.Attendance
}

func New(cfg config.AppConfig, ethclient ethclient.Client) (Blockchain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get current chain ID
	chainID, err := ethclient.ChainID(ctx)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	// Format private key
	privateKey, err := crypto.HexToECDSA(cfg.BlockChain.PrivateKey)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	contract := common.HexToAddress(cfg.BlockChain.ContractAddress)
	attContract, err := contracts.NewAttendance(contract, &ethclient)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	return &blockchain{
		client:      &ethclient,
		transactor:  transactor,
		attContract: attContract,
	}, nil
}

func (bc blockchain) AttendanceContract() contracts.Attendance {
	return *bc.attContract
}
