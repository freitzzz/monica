package core

func Collect[T any](
	cbs ...func() (T, error),
) ([]T, error) {
	var err error
	var r []T

	for _, cb := range cbs {
		cbr, err := cb()

		if err != nil {
			break
		}

		r = append(r, cbr)
	}

	return r, err
}
