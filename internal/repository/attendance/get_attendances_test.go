package attendance

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"

	"github.com/the-witcher-knight/timekeeper/internal/model"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/postgres"
)

func TestRepository_ListAttendanceByFilter(t *testing.T) {
	tcs := map[string]struct {
		givenInput     FilterInput
		failedByDBconn bool
		mockConn       func(mock sqlmock.Sqlmock)
		expResult      []model.Attendance
		expErr         error
	}{
		"success": {
			givenInput: FilterInput{
				EmployerID: 1,
				FromTime:   time.Date(2024, time.June, 19, 0, 0, 0, 0, time.UTC),
				ToTime:     time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC),
			},
			expResult: []model.Attendance{
				{
					ID:          1,
					EmployerID:  1,
					CheckInTime: time.Date(2024, time.June, 20, 1, 0, 0, 0, time.UTC),
					Notes:       "Check in",
					CreatedAt:   time.Date(2024, time.June, 20, 1, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2024, time.June, 20, 1, 0, 0, 0, time.UTC),
				},
				{
					ID:          2,
					EmployerID:  1,
					CheckInTime: time.Date(2024, time.June, 20, 1, 0, 0, 0, time.UTC),
					Notes:       "Check in",
					CreatedAt:   time.Date(2024, time.June, 20, 1, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2024, time.June, 20, 1, 0, 0, 0, time.UTC),
				},
			},
		},
		"error - unexpected": {
			givenInput: FilterInput{
				EmployerID: 1,
				FromTime:   time.Date(2024, time.June, 19, 0, 0, 0, 0, time.UTC),
				ToTime:     time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC),
			},
			failedByDBconn: true,
			mockConn: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT "attendances".* FROM "attendances" WHERE \("attendances"."deleted_at" is null\) AND \("attendances"."employer_id" = \$1\) AND \("attendances"."check_in_time" >= \$2\) AND \("attendances"."check_in_time" <= \$3\)`).
					WithArgs(1, time.Date(2024, time.June, 19, 0, 0, 0, 0, time.UTC), time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC)).
					WillReturnError(errors.New("simulated error when select attendance"))
			},
			expErr: errors.New("ormmodel: failed to assign all query results to Attendance slice: bind failed to execute query: simulated error when select attendance"),
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
				atts, err := inst.ListAttendanceByFilter(ctx, tc.givenInput)

				// Then
				if tc.expErr != nil {
					require.EqualError(t, err, tc.expErr.Error())
				} else {
					require.NoError(t, err)
					assertModelSlice(t, tc.expResult, atts)
				}
			})
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

func TestRepository_GetAttendanceByID(t *testing.T) {
	tcs := map[string]struct {
		givenInput     int64
		failedByDBconn bool
		mockConn       func(mock sqlmock.Sqlmock)
		expResult      model.Attendance
		expErr         error
	}{
		"success": {
			givenInput: 1,
			expResult: model.Attendance{
				ID:          1,
				EmployerID:  1,
				CheckInTime: time.Date(2024, time.June, 20, 1, 0, 0, 0, time.UTC),
				Notes:       "Check in",
				CreatedAt:   time.Date(2024, time.June, 20, 1, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2024, time.June, 20, 1, 0, 0, 0, time.UTC),
			},
		},
		"success - record not found": {
			givenInput: 10,
		},
		"error - unexpected": {
			givenInput:     1,
			failedByDBconn: true,
			mockConn: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT "attendances".* FROM "attendances" WHERE \("attendances"."deleted_at" is null\) AND \("attendances"."id" = \$1\) LIMIT 1`).
					WithArgs(1).
					WillReturnError(errors.New("simulated error when select attendance"))
			},
			expErr: errors.New("ormmodel: failed to execute a one query for attendances: bind failed to execute query: simulated error when select attendance"),
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
				att, err := inst.GetAttendanceByID(ctx, tc.givenInput, false)

				// Then
				if tc.expErr != nil {
					require.EqualError(t, err, tc.expErr.Error())
				} else {
					require.NoError(t, err)
					assertModel(t, tc.expResult, att)
				}
			})
		})
	}
}

func assertModel[T any](t *testing.T, expected, actual T, ignoreFields ...string) {
	var o T
	ignoreFieldsCmpopts := cmpopts.IgnoreFields(o, ignoreFields...)

	if !cmp.Equal(expected, actual, ignoreFieldsCmpopts) {
		t.Errorf("\n result mismatched. Diff: %+v", cmp.Diff(expected, actual, ignoreFieldsCmpopts))
		t.FailNow()
	}
}
