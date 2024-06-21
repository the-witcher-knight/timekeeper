package main

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	pkgerrors "github.com/pkg/errors"

	"github.com/the-witcher-knight/timekeeper/cmd/asciiart"
	"github.com/the-witcher-knight/timekeeper/internal/blockchain/contracts"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/config"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/system"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/tracing"
)

func main() {
	asciiart.Show()
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Printf("server exitted abnormally: %s\n", err.Error())
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.ReadConfigFromEnv()
	if err != nil {
		return err
	}

	s, err := system.New(cfg)
	if err != nil {
		return err
	}

	// Setup tracer for logging and tracing when job running
	tracer := tracing.NewTracer(s.Logger(),
		tracing.String("service", "deploy-contract"),
		tracing.String("network", cfg.BlockChain.NetworkHTTP),
	)
	defer func() {
		if p := recover(); p != nil {
			pErr, ok := p.(error)
			if !ok {
				pErr = fmt.Errorf("panic: %v", p)
			}

			// Log error and stacktrace when got panic
			tracer.Error(context.Background(), pErr, "caught a panic", false,
				tracing.Bytes("stacktrace", debug.Stack()))
		}
	}()

	ctx = tracing.SetInCtx(ctx, tracer)
	addr, txhash, err := deployContract(ctx, *s)
	if err != nil {
		tracer.Error(context.Background(), err, "failed to deploy contract", true)
		return err
	}

	tracer.With(
		tracing.String("contract.address", addr),
		tracing.String("contract.tx_hash", txhash),
	).Info(context.Background(), "Contract Deployed")

	return nil
}

func deployContract(ctx context.Context, s system.System) (string, string, error) {
	// Get current chain ID
	chainID, err := s.EthClientHTTP().ChainID(ctx)
	if err != nil {
		return "", "", pkgerrors.WithStack(err)
	}

	// Format private key
	privateKey, err := crypto.HexToECDSA(s.Config().BlockChain.PrivateKey)
	if err != nil {
		return "", "", pkgerrors.WithStack(err)
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", "", pkgerrors.WithStack(err)
	}

	addr, tx, _, err := contracts.DeployAttendance(txOpts, s.EthClientHTTP())
	if err != nil {
		return "", "", pkgerrors.WithStack(err)
	}

	_, err = bind.WaitDeployed(ctx, s.EthClientHTTP(), tx)
	if err != nil {
		return "", "", pkgerrors.WithStack(err)
	}

	return addr.Hex(), tx.Hash().Hex(), nil
}
