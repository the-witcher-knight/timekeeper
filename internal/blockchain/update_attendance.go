package blockchain

import (
	"context"
	"math/big"
	"strings"

	pkgerrors "github.com/pkg/errors"

	"github.com/the-witcher-knight/timekeeper/internal/model"
)

func (bc blockchain) UpdateAttendance(ctx context.Context, att model.Attendance) error {
	_, err := bc.attContract.UpdateAttendance(bc.transactor,
		big.NewInt(att.ID),
		big.NewInt(att.EmployerID),
		big.NewInt(att.CheckInTime.Unix()),
		att.Notes,
	)
	if err != nil {
		if strings.Contains(err.Error(), errRecordNotFound) {
			return ErrRecordNotFound
		}

		return pkgerrors.WithStack(err)
	}

	return nil
}
