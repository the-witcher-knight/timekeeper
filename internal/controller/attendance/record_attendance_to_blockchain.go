package attendance

import (
	"context"
	"time"

	"github.com/the-witcher-knight/timekeeper/internal/model"
)

var (
	timeNowWrapper = func() time.Time {
		return time.Now().UTC()
	}
)

func (ctrl controller) RecordAttendanceToBlockchain(ctx context.Context, employerID int64) error {
	att := model.Attendance{
		EmployerID:  employerID,
		CheckInTime: timeNowWrapper(),
		Notes:       "Check in", // Improve later
	}

	if err := ctrl.bc.RecordAttendance(ctx, att); err != nil {
		return err
	}

	return nil
}
