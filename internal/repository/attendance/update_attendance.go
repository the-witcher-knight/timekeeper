package attendance

import (
	"context"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/the-witcher-knight/timekeeper/internal/model"
	pkgerrors "github.com/the-witcher-knight/timekeeper/internal/pkg/errors"
	"github.com/the-witcher-knight/timekeeper/internal/repository/ormmodel"
)

func (repo repository) UpdateAttendance(ctx context.Context, att model.Attendance) error {
	orm := ormmodel.Attendance{
		ID:          att.ID,
		EmployerID:  att.EmployerID,
		CheckInTime: att.CheckInTime,
		Notes:       att.Notes,
		DeletedAt:   null.TimeFromPtr(att.DeletedAt),
	}
	if _, err := orm.Update(ctx, repo.dbConn, boil.Infer()); err != nil {
		return pkgerrors.WithStack(err)
	}

	return nil
}
