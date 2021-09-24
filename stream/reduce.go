package stream

func Reduce[T, U any](s Stream[T], identity U, acc func(U, T) U) (r U) {
	r = identity
	s.Iterate(func(t T) bool {
		r = acc(r, t)
		return true
	})

	return r
}

func (s Stream[T]) ToSlice() []T {
	return Reduce[T, []T](s, nil, func(acc []T, e T) []T {
		return append(acc, e)
	})
}
