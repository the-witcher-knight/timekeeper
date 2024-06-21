package blockchain

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	pkgerrors "github.com/pkg/errors"

	"github.com/the-witcher-knight/timekeeper/internal/model"
)

type AttendanceFilter struct {
	EmployerID int64
	FromTime   time.Time
	ToTime     time.Time
}

func (bc blockchain) RetrieveAttendance(ctx context.Context, filter AttendanceFilter) ([]model.Attendance, error) {
	var from, to big.Int
	if !filter.FromTime.IsZero() {
		from = *big.NewInt(filter.FromTime.Unix())
	}

	if !filter.ToTime.IsZero() {
		to = *big.NewInt(filter.ToTime.Unix())
	}

	callOpts := &bind.CallOpts{Context: ctx, Pending: false}
	atts, err := bc.attContract.GetAttendance(callOpts,
		big.NewInt(filter.EmployerID),
		&from,
		&to,
	)
	if err != nil {
		if strings.Contains(err.Error(), errNotAuthorized) {
			return nil, ErrCurrentAccountNotAuthorized
		}

		return nil, pkgerrors.WithStack(err)
	}

	return toAttendanceSlice(atts), nil
}
