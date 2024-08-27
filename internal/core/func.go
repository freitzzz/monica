package core

func Collect[T any](
	cbs ...func() (T, error),
) ([]T, error) {
	var err error
	var r []T

	for _, cb := range cbs {
		cbr, cberr := cb()

		if cberr != nil {
			err = cberr
			break
		}

		r = append(r, cbr)
	}

	return r, err
}
