package attendance

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	"github.com/the-witcher-knight/timekeeper/internal/model"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/postgres"
)

func TestRepository_UpdateAttendance(t *testing.T) {
	tcs := map[string]struct {
		givenInput     model.Attendance
		failedByDBconn bool
		mockConn       func(mock sqlmock.Sqlmock)
		expErr         error
	}{
		"success": {
			givenInput: model.Attendance{
				ID:          2,
				EmployerID:  1,
				CheckInTime: time.Date(2024, time.June, 20, 0, 0, 0, 0, time.UTC),
				Notes:       "Check in",
			},
		},
		"error - unexpected": {
			givenInput: model.Attendance{
				ID:          10,
				EmployerID:  1,
				CheckInTime: time.Date(2024, time.June, 20, 0, 0, 0, 0, time.UTC),
				Notes:       "Updated",
			},
			failedByDBconn: true,
			mockConn: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`UPDATE "attendances" SET "employer_id"=\$1,"check_in_time"=\$2,"notes"=\$3,"updated_at"=\$4,"deleted_at"=\$5 WHERE "id"=\$6`).
					WithArgs(1, time.Date(2024, time.June, 20, 0, 0, 0, 0, time.UTC), "Updated", sqlmock.AnyArg(), nil, 10).
					WillReturnError(errors.New("simulated error when update attendance"))
			},
			expErr: errors.New("ormmodel: unable to update attendances row: simulated error when update attendance"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			postgres.RunTestInTx(t, func(t *testing.T, tx postgres.ContextExecutor) {
				ctx := context.Background()
				postgres.LoadSqlTestFile(t, tx, "testdata/attendance.sql")

				// Given
				inst := repository{dbConn: tx}
				if tc.failedByDBconn {
					dbmock, mock, err := sqlmock.New()
					require.NoError(t, err)
					defer dbmock.Close()
					tc.mockConn(mock)
					inst = repository{dbConn: dbmock}
				}

				// When
				err := inst.UpdateAttendance(ctx, tc.givenInput)

				// Then
				if tc.expErr != nil {
					require.EqualError(t, err, tc.expErr.Error())
				} else {
					require.NoError(t, err)
				}
			})
		})
	}
}
