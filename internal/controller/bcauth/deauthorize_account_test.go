package bcauth

import (
	"context"
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/the-witcher-knight/timekeeper/internal/blockchain"
)

func TestController_DeauthorizeAccount(t *testing.T) {
	type mockBlockchain struct {
		inInput common.Address
		outErr  error
	}

	tcs := map[string]struct {
		givenInput     common.Address
		mockBlockchain mockBlockchain
		expErr         error
	}{
		"success": {
			givenInput: common.HexToAddress("1"),
			mockBlockchain: mockBlockchain{
				inInput: common.HexToAddress("1"),
			},
		},
		"error - account not authorized": {
			givenInput: common.HexToAddress("1"),
			mockBlockchain: mockBlockchain{
				inInput: common.HexToAddress("1"),
				outErr:  blockchain.ErrCurrentAccountNotAuthorized,
			},
			expErr: ErrCurrentAccountNotAuthorized,
		},
		"error - unexpected": {
			givenInput: common.HexToAddress("1"),
			mockBlockchain: mockBlockchain{
				inInput: common.HexToAddress("1"),
				outErr:  errors.New("simulated error"),
			},
			expErr: errors.New("simulated error"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given
			ctx := context.Background()

			bcMock := new(blockchain.MockBlockchain)
			bcMock.On("DeauthorizeAccount", ctx, tc.mockBlockchain.inInput).
				Return(tc.mockBlockchain.outErr)

			// When
			ctrl := New(bcMock)
			err := ctrl.DeauthorizeAccount(ctx, tc.givenInput)

			// Then
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
			}
			bcMock.AssertExpectations(t)
		})
	}
}
