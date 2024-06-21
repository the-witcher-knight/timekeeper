package attendance

import (
	"context"
	"errors"

	"github.com/the-witcher-knight/timekeeper/internal/blockchain"
	"github.com/the-witcher-knight/timekeeper/internal/model"
)

func (ctrl controller) UpdateAttendanceToBlockchain(ctx context.Context, att model.Attendance) error {
	if err := ctrl.bc.UpdateAttendance(ctx, att); err != nil {
		if errors.Is(err, blockchain.ErrRecordNotFound) {
			return ErrRecordNotFound
		}

		return err
	}

	return nil
}
