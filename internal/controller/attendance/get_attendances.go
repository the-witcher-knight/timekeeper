package attendance

import (
	"context"
	"time"

	"github.com/the-witcher-knight/timekeeper/internal/model"
	"github.com/the-witcher-knight/timekeeper/internal/repository/attendance"
)

type FilterInput struct {
	EmployerID int64
	FromTime   time.Time
	ToTime     time.Time
}

func (ctrl controller) ListAttendances(ctx context.Context, filter FilterInput) ([]model.Attendance, error) {
	rs, err := ctrl.repo.Attendance().ListAttendanceByFilter(ctx, attendance.FilterInput{
		EmployerID: filter.EmployerID,
		FromTime:   filter.FromTime,
		ToTime:     filter.ToTime,
	})
	if err != nil {
		return nil, err
	}

	return rs, nil
}
