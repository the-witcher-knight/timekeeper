package attendance

import (
	"context"

	"github.com/the-witcher-knight/timekeeper/internal/model"
)

func (ctrl controller) CreateAttendance(ctx context.Context, att model.Attendance) error {
	if err := ctrl.repo.Attendance().CreateAttendance(ctx, att); err != nil {
		return err
	}

	return nil
}
