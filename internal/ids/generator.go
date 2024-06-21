package ids

import (
	"time"

	"github.com/sony/sonyflake"

	pkgerrors "github.com/pkg/errors"
)

type IDGenerator interface {
	NextID() (int64, error)
}

type idGenerator struct {
	sf *sonyflake.Sonyflake
}

func (generator idGenerator) NextID() (int64, error) {
	id, err := generator.sf.NextID()
	if err != nil {
		return 0, pkgerrors.WithStack(err)
	}

	return int64(id), nil
}

func NewIDGenerator() (IDGenerator, error) {
	st := sonyflake.Settings{
		StartTime: time.Now(),
	}

	sf, err := sonyflake.New(st)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	return idGenerator{
		sf: sf,
	}, nil
}
