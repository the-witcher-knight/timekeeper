package bcwatch

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/the-witcher-knight/timekeeper/internal/blockchain/contracts"
	"github.com/the-witcher-knight/timekeeper/internal/model"
	pkgerrors "github.com/the-witcher-knight/timekeeper/internal/pkg/errors"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/tracing"
	"github.com/the-witcher-knight/timekeeper/internal/repository"
)

func (ctrl controller) WatchAttendanceUpdated() func(context.Context) error {
	return func(ctx context.Context) error {
		tracer := tracing.NewTracer(ctrl.logger, tracing.String("service", "watch_attendance_updated"))

		watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}
		channel := make(chan *contracts.AttendanceAttendanceUpdated)

		ctr := ctrl.bc.AttendanceContract()

		// Subscribe to the AttendanceUpdated event
		sub, err := ctr.WatchAttendanceUpdated(watchOpts, channel, nil, nil)
		if err != nil {
			return pkgerrors.WithStack(err)
		}

		// Ensure the subscription is unsubscribed when the function exits
		defer sub.Unsubscribe()

		for {
			select {
			case event := <-channel:
				// Extract and log event details
				tr := tracer.With(
					tracing.String("event.address", event.Raw.Address.Hex()),
					tracing.String("event.tx_hash", event.Raw.TxHash.Hex()),
					tracing.Bytes("event.data", event.Raw.Data),
				)

				ctx := tracing.SetInCtx(ctx, tr)

				// Process the event by creating attendance
				if err := updateAttendance(ctx, ctrl, *event); err != nil {
					tr.Error(context.Background(), err, "Store updated attendance error", true)
					continue
				}

				// Log successful storage of attendance
				tr.Info(context.Background(), "Attendance updated successfully")

			case err := <-sub.Err():
				// Log and return subscription errors.
				tracer.Error(context.Background(), err, "Subscription error", false)

				return err
			case <-ctx.Done():
				// Handle context cancellation
				return ctx.Err()
			}
		}
	}
}

func updateAttendance(ctx context.Context, ctrl controller, event contracts.AttendanceAttendanceUpdated) error {
	ctx, cancel := context.WithTimeout(ctx, handleEventTimeout)
	defer cancel()

	// Map the event data to an Attendance model
	att := model.Attendance{
		ID:          event.Id.Int64(),
		EmployerID:  event.EmployeeId.Int64(),
		CheckInTime: time.Unix(event.CheckInTime.Int64(), 0),
		Notes:       event.Notes,
	}

	txFunc := func(registry repository.Registry) error {
		// Get and lock the attendance record by ID.
		oldAtt, err := registry.Attendance().GetAttendanceByID(ctx, att.ID, true)
		if err != nil {
			return err
		}

		// Update the attendance record with new data.
		oldAtt.CheckInTime = att.CheckInTime
		oldAtt.Notes = att.Notes

		// Update attendance record in the repository.
		if err := registry.Attendance().UpdateAttendance(ctx, oldAtt); err != nil {
			return err
		}

		return nil
	}

	// Execute the transaction function within a database transaction.
	if err := ctrl.repo.DoInTx(ctx, txFunc); err != nil {
		return err
	}

	return nil
}
