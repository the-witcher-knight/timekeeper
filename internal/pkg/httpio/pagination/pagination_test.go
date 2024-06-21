package pagination

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_Validate(t *testing.T) {
	tcs := map[string]struct {
		givenInput Input
		expErr     error
	}{
		"page 1 size 0": {
			givenInput: Input{
				Size: 0,
				Page: 1,
			},
			expErr: ErrSizeMustBeGreaterThanZero,
		},
		"page 1, pagination exceeded max": {
			givenInput: Input{
				Size: paginationMaxSizePerPage + 1,
				Page: 1,
			},
			expErr: ErrExceededMaxPageSizeLimit,
		},
		"page 1, pagination near max": {
			givenInput: Input{
				Size: paginationMaxSizePerPage,
				Page: 1,
			},
		},
		"limit -1, page 0": {
			givenInput: Input{
				Size: -1,
			},
			expErr: ErrSizeMustBeGreaterThanZero,
		},
		"limit 1, page 0": {
			givenInput: Input{
				Size: 1,
				Page: 0,
			},
			expErr: ErrPageMustBeGreaterThanZero,
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			err := tc.givenInput.Valid()
			require.Equal(t, tc.expErr, err)
		})
	}
}

func TestToSQLOffsetLimit(t *testing.T) {
	tcs := map[string]struct {
		givenInput Input
		expOffset  int
		expLimit   int
	}{
		"page: 1, size: 10": {
			givenInput: Input{
				Size: 10,
				Page: 1,
			},
			expOffset: 0,
			expLimit:  10,
		},
		"page: 2, size: 5": {
			givenInput: Input{
				Size: 5,
				Page: 2,
			},
			expOffset: 5,
			expLimit:  5,
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			offset, limit := tc.givenInput.ToSQLOffsetLimit()
			require.Equal(t, tc.expOffset, offset)
			require.Equal(t, tc.expLimit, limit)
		})
	}
}
