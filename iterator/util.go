package iterator

import (
	"constraints"

	"github.com/tdakkota/algo2/alg"
)

func Range[T constraints.Integer](from, to, step T) Iterator[T] {
	if alg.IsZero(step) {
		return Empty[T]()
	}

	if step < T(0) {
		return IteratorFunc[T](func(cb func(i T) bool) {
			n := from
			for n > to {
				if !cb(n) {
					return
				}
				n += step
			}
		})
	} else {
		return IteratorFunc[T](func(cb func(i T) bool) {
			n := from
			for n < to {
				if !cb(n) {
					return
				}
				n += step
			}
		})
	}
}

func FromTo[T constraints.Integer](from, to T) Iterator[T] {
	return Range[T](from, to, T(1))
}

func Repeat[T any](value T, count uint) Iterator[T] {
	if count <= 0 {
		return Empty[T]()
	}

	return IteratorFunc[T](func(cb func(i T) bool) {
		for i := uint(0); i < count; i++ {
			if !cb(value) {
				return
			}
		}
	})
}

func Const[T any](value T) Iterator[T] {
	return IteratorFunc[T](func(cb func(i T) bool) {
		for {
			if !cb(value) {
				return
			}
		}
	})
}

type sequential[T any] struct {
	iterators []Iterator[T]
}

func (s sequential[T]) Iterate(fn func(T) bool) {
	var stop bool
	for _, iterator := range s.iterators {
		iterator.Iterate(func(e T) bool {
			stop = fn(e)
			return stop
		})

		if stop {
			return
		}
	}
}

// Sequential function joins multiple iterators
func Sequential[T any](iterators ...Iterator[T]) Iterator[T] {
	return sequential[T]{
		iterators: iterators,
	}
}

func Values[T any](iterators ...Iterator[T]) (r []T) {
	Sequential[T](iterators...).Iterate(func(e T) bool {
		r = append(r, e)
		return true
	})
	return
}
