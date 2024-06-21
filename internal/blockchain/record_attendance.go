package blockchain

import (
	"context"
	"math/big"
	"strings"

	pkgerrors "github.com/pkg/errors"

	"github.com/the-witcher-knight/timekeeper/internal/ids"
	"github.com/the-witcher-knight/timekeeper/internal/model"
)

func (bc blockchain) RecordAttendance(ctx context.Context, att model.Attendance) error {
	if att.ID == 0 {
		newID, err := ids.BlockChainAttendanceRecord.NextID()
		if err != nil {
			return err
		}

		att.ID = newID
	}

	_, err := bc.attContract.RecordAttendance(bc.transactor,
		big.NewInt(att.ID), big.NewInt(att.EmployerID), big.NewInt(att.CheckInTime.Unix()), att.Notes)
	if err != nil {
		if strings.Contains(err.Error(), errNotAuthorized) {
			return ErrCurrentAccountNotAuthorized
		}

		return pkgerrors.WithStack(err)
	}

	return nil
}
