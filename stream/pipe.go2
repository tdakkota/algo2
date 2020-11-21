package stream

import "github.com/tdakkota/algo2/iterator"

func Map[T, R any](s Stream[T], op func(T) R) Stream[R] {
	return Pipe[T, R](s, func(t T) (R, bool) {
		return op(t), true
	})
}

func (s Stream[T]) Map(fn func(T) T) Stream[T] {
	return Map[T, T](s, fn)
}

func Pipe[T, R any](s Stream[T], op func(d T) (R, bool)) Stream[R] {
	it := func(fn func(R) bool) {
		s.it.Iterate(func(t T) bool {
			r, ok := op(t)
			if !ok {
				return true
			}

			return fn(r)
		})
	}

	return FromIterator[R](iterator.IteratorFunc[R](it))
}

func (s Stream[T]) Pipe(op func(d T) (T, bool)) Stream[T] {
	return Pipe[T, T](s, op)
}
