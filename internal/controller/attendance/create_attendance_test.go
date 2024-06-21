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

func TestController_CreateAttendance(t *testing.T) {
	type mockRepo struct {
		inInput model.Attendance
		outErr  error
	}

	tcs := map[string]struct {
		givenInput model.Attendance
		mockRepo   mockRepo
		expErr     error
	}{
		"success": {
			givenInput: model.Attendance{
				ID:          1,
				EmployerID:  1,
				CheckInTime: time.Date(2024, time.June, 20, 0, 0, 0, 0, time.UTC),
				Notes:       "Check in",
			},
			mockRepo: mockRepo{
				inInput: model.Attendance{
					ID:          1,
					EmployerID:  1,
					CheckInTime: time.Date(2024, time.June, 20, 0, 0, 0, 0, time.UTC),
					Notes:       "Check in",
				},
			},
		},
		"error - unexpected": {
			givenInput: model.Attendance{
				ID:          1,
				EmployerID:  1,
				CheckInTime: time.Date(2024, time.June, 20, 0, 0, 0, 0, time.UTC),
				Notes:       "Check in",
			},
			mockRepo: mockRepo{
				inInput: model.Attendance{
					ID:          1,
					EmployerID:  1,
					CheckInTime: time.Date(2024, time.June, 20, 0, 0, 0, 0, time.UTC),
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

			attRepo := new(attendance.MockRepository)
			attRepo.On("CreateAttendance", ctx, tc.mockRepo.inInput).Return(tc.mockRepo.outErr)

			mockRegistry := new(repository.MockRegistry)
			mockRegistry.On("Attendance").Return(attRepo)

			// When
			ctrl := New(nil, mockRegistry)
			err := ctrl.CreateAttendance(ctx, tc.givenInput)

			// Then
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
			}
			attRepo.AssertExpectations(t)
		})
	}
}
