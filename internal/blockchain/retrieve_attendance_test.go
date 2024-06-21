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
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"

	"github.com/the-witcher-knight/timekeeper/internal/blockchain/contracts"
	"github.com/the-witcher-knight/timekeeper/internal/model"
)

func TestBlockchain_RetrieveAttendance(t *testing.T) {
	tcs := map[string]struct {
		givenInput AttendanceFilter
		expResult  []model.Attendance
		expErr     error
	}{
		"success": {
			givenInput: AttendanceFilter{
				EmployerID: 1,
				FromTime:   time.Date(2024, time.June, 20, 0, 0, 0, 0, time.UTC),
				ToTime:     time.Date(2024, time.June, 23, 0, 0, 0, 0, time.UTC),
			},
			expResult: []model.Attendance{
				{
					ID:          1,
					EmployerID:  1,
					CheckInTime: time.Date(2024, time.June, 20, 1, 0, 0, 0, time.UTC),
					Notes:       "Check in",
				},
			},
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

			for _, att := range tc.expResult {
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
			atts, err := bc.RetrieveAttendance(ctx, tc.givenInput)

			sim.Commit()

			// Then
			if tc.expErr != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assertModelSlice(t, tc.expResult, atts)
			}
		})
	}
}

func assertModelSlice[T any](t *testing.T, expected, actual []T, ignoreFields ...string) {
	var o T
	ignoreFieldsCmpopts := cmpopts.IgnoreFields(o, ignoreFields...)

	if !cmp.Equal(expected, actual, ignoreFieldsCmpopts) {
		t.Errorf("\n result mismatched. Diff: %+v", cmp.Diff(expected, actual, ignoreFieldsCmpopts))
		t.FailNow()
	}
}
