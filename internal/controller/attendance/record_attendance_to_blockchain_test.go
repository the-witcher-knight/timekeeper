package attendance

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/the-witcher-knight/timekeeper/internal/blockchain"
	"github.com/the-witcher-knight/timekeeper/internal/model"
)

func TestController_RecordAttendanceToBlockchain(t *testing.T) {
	defer func(originFn func() time.Time) { timeNowWrapper = originFn }(timeNowWrapper)
	timeNowWrapper = func() time.Time { return time.Date(2020, time.June, 20, 1, 0, 0, 0, time.UTC) }

	type mockBlockchain struct {
		inInput model.Attendance
		outErr  error
	}

	tcs := map[string]struct {
		givenInput     int64
		mockBlockchain mockBlockchain
		expErr         error
	}{
		"success": {
			givenInput: 1,
			mockBlockchain: mockBlockchain{
				inInput: model.Attendance{
					EmployerID:  1,
					CheckInTime: time.Date(2020, time.June, 20, 1, 0, 0, 0, time.UTC),
					Notes:       "Check in",
				},
			},
		},
		"error - unexpected": {
			givenInput: 1,
			mockBlockchain: mockBlockchain{
				inInput: model.Attendance{
					EmployerID:  1,
					CheckInTime: time.Date(2020, time.June, 20, 1, 0, 0, 0, time.UTC),
					Notes:       "Check in",
				},
				outErr: errors.New("simulated error"),
			},
			expErr: errors.New("simulated error"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given
			ctx := context.Background()

			bcMock := new(blockchain.MockBlockchain)
			bcMock.On("RecordAttendance", ctx, tc.mockBlockchain.inInput).
				Return(tc.mockBlockchain.outErr)

			// When
			ctrl := New(bcMock, nil)
			err := ctrl.RecordAttendanceToBlockchain(ctx, tc.givenInput)

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
