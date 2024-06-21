package errors

const (
	maxErrorDepth uint8 = 10
)

func GetRootErr(err error) error {
	originErr := err
	for i := 0; i < int(maxErrorDepth) && originErr != nil; i++ {
		wrapped, ok := originErr.(WrappedError)
		if !ok {
			break
		}

		originErr = wrapped.Unwrap()
	}

	return originErr
}
