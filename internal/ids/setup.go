package ids

import (
	pkgerrors "github.com/pkg/errors"
)

var (
	BlockChainAttendanceRecord IDGenerator
	User                       IDGenerator
	Attendance                 IDGenerator
)

// Setup init IDGenerator for snowflake id generator
func Setup() error {
	var err error

	if BlockChainAttendanceRecord == nil {
		BlockChainAttendanceRecord, err = NewIDGenerator()
		if err != nil {
			return pkgerrors.WithStack(err)
		}
	}

	if User == nil {
		User, err = NewIDGenerator()
		if err != nil {
			return pkgerrors.WithStack(err)
		}
	}

	if Attendance == nil {
		Attendance, err = NewIDGenerator()
		if err != nil {
			return pkgerrors.WithStack(err)
		}
	}

	return nil
}
