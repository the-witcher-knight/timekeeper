package attendance

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/the-witcher-knight/timekeeper/internal/model"
	pkgerrors "github.com/the-witcher-knight/timekeeper/internal/pkg/errors"
	"github.com/the-witcher-knight/timekeeper/internal/repository/ormmodel"
)

type FilterInput struct {
	EmployerID int64
	FromTime   time.Time
	ToTime     time.Time
}

func (repo repository) ListAttendanceByFilter(ctx context.Context, filters FilterInput) ([]model.Attendance, error) {
	qms := []qm.QueryMod{
		ormmodel.AttendanceWhere.DeletedAt.IsNull(),
	}

	if filters.EmployerID > 0 {
		qms = append(qms, ormmodel.AttendanceWhere.EmployerID.EQ(filters.EmployerID))
	}

	if !filters.FromTime.IsZero() {
		qms = append(qms, ormmodel.AttendanceWhere.CheckInTime.GTE(filters.FromTime))
	}

	if !filters.ToTime.IsZero() {
		qms = append(qms, ormmodel.AttendanceWhere.CheckInTime.LTE(filters.ToTime))
	}

	orms, err := ormmodel.Attendances(qms...).All(ctx, repo.dbConn)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	return toAttendanceSlice(orms), nil
}

func (repo repository) GetAttendanceByID(ctx context.Context, id int64, lock bool) (model.Attendance, error) {
	qms := []qm.QueryMod{
		ormmodel.AttendanceWhere.DeletedAt.IsNull(),
	}

	if id > 0 {
		qms = append(qms, ormmodel.AttendanceWhere.ID.EQ(id))
	}

	if lock {
		qms = append(qms, qm.For("UPDATE"))
	}

	orm, err := ormmodel.Attendances(qms...).One(ctx, repo.dbConn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Attendance{}, nil
		}

		return model.Attendance{}, pkgerrors.WithStack(err)
	}

	return toAttendanceModel(*orm), nil

}
