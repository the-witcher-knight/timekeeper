package pagination

const (
	paginationDefaultSize    = 10
	paginationDefaultPage    = 1
	paginationMaxSizePerPage = 1000
)

type Input struct {
	Page, Size    int
	IncludeTotals bool
}

func NewDefaultInput() Input {
	return Input{
		Page: paginationDefaultPage,
		Size: paginationDefaultSize,
	}
}

func (in *Input) Valid() error {
	if in == nil {
		return nil
	}

	if in.Size > paginationMaxSizePerPage {
		return ErrExceededMaxPageSizeLimit
	}
	if in.Size <= 0 {
		return ErrSizeMustBeGreaterThanZero
	}

	if in.Page <= 0 {
		return ErrPageMustBeGreaterThanZero
	}

	return nil
}

func (in *Input) ToSQLOffsetLimit() (int, int) {
	var offset, limit int
	if in == nil {
		in.Page = paginationDefaultPage
		in.Size = paginationDefaultSize
	}

	limit = in.Size
	offset = limit * (in.Page - 1)

	return offset, limit
}
