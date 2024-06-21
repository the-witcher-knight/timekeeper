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

func TestController_RetrieveAttendanceFromBlockchain(t *testing.T) {
	type mockBlockchain struct {
		inInput   blockchain.AttendanceFilter
		outResult []model.Attendance
		outErr    error
	}

	tcs := map[string]struct {
		givenInput     FilterInput
		mockBlockchain mockBlockchain
		expResult      []model.Attendance
		expErr         error
	}{
		"success": {
			givenInput: FilterInput{
				EmployerID: 1,
				FromTime:   time.Date(2024, time.June, 19, 0, 0, 0, 0, time.UTC),
				ToTime:     time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC),
			},
			mockBlockchain: mockBlockchain{
				inInput: blockchain.AttendanceFilter{
					EmployerID: 1,
					FromTime:   time.Date(2024, time.June, 19, 0, 0, 0, 0, time.UTC),
					ToTime:     time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC),
				},
				outResult: []model.Attendance{
					{
						ID:          1,
						EmployerID:  1,
						CheckInTime: time.Date(2024, time.June, 19, 1, 0, 0, 0, time.UTC),
						Notes:       "Check in",
					},
					{
						ID:          2,
						EmployerID:  1,
						CheckInTime: time.Date(2024, time.June, 20, 1, 0, 0, 0, time.UTC),
						Notes:       "Check in",
					},
				},
			},
			expResult: []model.Attendance{
				{
					ID:          1,
					EmployerID:  1,
					CheckInTime: time.Date(2024, time.June, 19, 1, 0, 0, 0, time.UTC),
					Notes:       "Check in",
				},
				{
					ID:          2,
					EmployerID:  1,
					CheckInTime: time.Date(2024, time.June, 20, 1, 0, 0, 0, time.UTC),
					Notes:       "Check in",
				},
			},
		},
		"error - unexpected": {
			givenInput: FilterInput{
				EmployerID: 1,
				FromTime:   time.Date(2024, time.June, 19, 0, 0, 0, 0, time.UTC),
				ToTime:     time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC),
			},
			mockBlockchain: mockBlockchain{
				inInput: blockchain.AttendanceFilter{
					EmployerID: 1,
					FromTime:   time.Date(2024, time.June, 19, 0, 0, 0, 0, time.UTC),
					ToTime:     time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC),
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
			bcMock.On("RetrieveAttendance", ctx, tc.mockBlockchain.inInput).
				Return(tc.mockBlockchain.outResult, tc.mockBlockchain.outErr)

			// When
			ctrl := New(bcMock, nil)
			atts, err := ctrl.RetrieveAttendanceFromBlockchain(ctx, tc.givenInput)

			// Then
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.ElementsMatch(t, tc.expResult, atts)
			}
			bcMock.AssertExpectations(t)
		})
	}
}
