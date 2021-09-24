package stream

import "github.com/tdakkota/algo2/iterator"

type Stream[T any] struct {
	it iterator.Iterator[T]
}

func FromIterator[T any](it iterator.Iterator[T]) Stream[T] {
	return Stream[T]{it: it}
}

func Of[T any](e ...T) Stream[T] {
	return FromIterator[T](iterator.Slice(e))
}

var _ iterator.Iterator[int] = Stream[int]{}

func (s Stream[T]) Iterate(fn func(T) bool) {
	if s.it == nil {
		return
	}
	s.it.Iterate(fn)
}
