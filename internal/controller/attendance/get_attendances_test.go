package attendance

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/the-witcher-knight/timekeeper/internal/model"
	"github.com/the-witcher-knight/timekeeper/internal/repository"
	"github.com/the-witcher-knight/timekeeper/internal/repository/attendance"
)

func TestController_ListAttendances(t *testing.T) {
	type mockRepo struct {
		inInput   attendance.FilterInput
		outResult []model.Attendance
		outErr    error
	}

	tcs := map[string]struct {
		givenInput FilterInput
		mockRepo   mockRepo
		expResult  []model.Attendance
		expErr     error
	}{
		"success": {
			givenInput: FilterInput{
				EmployerID: 1,
				FromTime:   time.Date(2024, time.June, 19, 0, 0, 0, 0, time.UTC),
				ToTime:     time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC),
			},
			mockRepo: mockRepo{
				inInput: attendance.FilterInput{
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
			mockRepo: mockRepo{
				inInput: attendance.FilterInput{
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

			attRepo := new(attendance.MockRepository)
			attRepo.On("ListAttendanceByFilter", ctx, tc.mockRepo.inInput).
				Return(tc.mockRepo.outResult, tc.mockRepo.outErr)

			mockRegistry := new(repository.MockRegistry)
			mockRegistry.On("Attendance").Return(attRepo)

			// When
			ctrl := New(nil, mockRegistry)
			atts, err := ctrl.ListAttendances(ctx, tc.givenInput)

			// Then
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.ElementsMatch(t, tc.expResult, atts)
			}
			attRepo.AssertExpectations(t)
		})
	}
}
