package attendance

import (
	"context"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/the-witcher-knight/timekeeper/internal/ids"
	"github.com/the-witcher-knight/timekeeper/internal/model"
	pkgerrors "github.com/the-witcher-knight/timekeeper/internal/pkg/errors"
	"github.com/the-witcher-knight/timekeeper/internal/repository/ormmodel"
)

func (repo repository) CreateAttendance(ctx context.Context, att model.Attendance) error {
	if att.ID == 0 {
		userID, err := ids.Attendance.NextID()
		if err != nil {
			return pkgerrors.WithStack(err)
		}

		att.ID = userID
	}

	orm := ormmodel.Attendance{
		ID:          att.ID,
		EmployerID:  att.EmployerID,
		CheckInTime: att.CheckInTime,
		Notes:       att.Notes,
		CreatedAt:   att.CreatedAt,
		UpdatedAt:   att.UpdatedAt,
		DeletedAt:   null.TimeFromPtr(att.DeletedAt),
	}
	if err := orm.Insert(ctx, repo.dbConn, boil.Infer()); err != nil {
		return pkgerrors.WithStack(err)
	}

	return nil
}
