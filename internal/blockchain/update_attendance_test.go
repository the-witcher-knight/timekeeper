package blockchain

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"

	"github.com/the-witcher-knight/timekeeper/internal/blockchain/contracts"
	"github.com/the-witcher-knight/timekeeper/internal/model"
)

func TestBlockchain_UpdateAttendance(t *testing.T) {
	tcs := map[string]struct {
		givenInput model.Attendance
		mockData   []model.Attendance
		expErr     error
	}{
		"success": {
			givenInput: model.Attendance{
				ID:          1,
				EmployerID:  1,
				CheckInTime: time.Date(2023, time.June, 20, 0, 0, 0, 0, time.UTC),
				Notes:       "Updated",
			},
			mockData: []model.Attendance{
				{
					ID:          1,
					EmployerID:  1,
					CheckInTime: time.Date(2023, time.June, 20, 0, 0, 0, 0, time.UTC),
					Notes:       "Check In",
				},
			},
		},
		"err record not found": {
			givenInput: model.Attendance{
				ID:          1,
				EmployerID:  1,
				CheckInTime: time.Date(2023, time.June, 20, 0, 0, 0, 0, time.UTC),
				Notes:       "Updated",
			},
			expErr: ErrRecordNotFound,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			chainID := params.AllDevChainProtocolChanges.ChainID

			key, err := crypto.GenerateKey()
			require.NoError(t, err)

			auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
			require.NoError(t, err)

			sim := simulated.NewBackend(map[common.Address]types.Account{
				auth.From: {Balance: big.NewInt(9e18)},
			})

			// Given
			_, tx, inst, err := contracts.DeployAttendance(auth, sim.Client())
			require.NoError(t, err)
			sim.Commit()

			for _, att := range tc.mockData {
				_, err = inst.RecordAttendance(auth, big.NewInt(att.ID), big.NewInt(att.EmployerID), big.NewInt(att.CheckInTime.Unix()), att.Notes)
				require.NoError(t, err)

				sim.Commit()

				ctx := context.Background()
				_, err = bind.WaitMined(ctx, sim.Client(), tx)
				require.NoError(t, err)
			}

			bc := blockchain{
				client:      sim.Client(),
				transactor:  auth,
				attContract: inst,
			}

			// When
			ctx := context.Background()
			err = bc.UpdateAttendance(ctx, tc.givenInput)

			sim.Commit()

			// Then
			if tc.expErr != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
