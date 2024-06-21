package attendance

import (
	"context"

	"github.com/the-witcher-knight/timekeeper/internal/blockchain"
	"github.com/the-witcher-knight/timekeeper/internal/model"
)

func (ctrl controller) RetrieveAttendanceFromBlockchain(ctx context.Context, filter FilterInput) ([]model.Attendance, error) {
	rs, err := ctrl.bc.RetrieveAttendance(ctx, blockchain.AttendanceFilter{
		EmployerID: filter.EmployerID,
		FromTime:   filter.FromTime,
		ToTime:     filter.ToTime,
	})
	if err != nil {
		return nil, err
	}

	return rs, nil
}
