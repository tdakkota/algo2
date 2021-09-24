package iterator

import (
	"github.com/tdakkota/algo2/alg"
)

func Slice[T any](s []T) Iterator[T] {
	return ValueIterator[int, T](SliceWithIndex[T](s))
}

func SliceWithIndex[T any](s []T) Iterator[alg.Pair[int, T]] {
	return IteratorFunc[alg.Pair[int, T]](func(fn func(alg.Pair[int, T]) bool) {
		for i, e := range s {
			if !fn(alg.Two(i, e)) {
				return
			}
		}
	})
}

func Map[K comparable, V any](m map[K]V) Iterator[alg.Pair[K, V]] {
	return IteratorFunc[alg.Pair[K, V]](func(fn func(alg.Pair[K, V]) bool) {
		for k, v := range m {
			if !fn(alg.Two(k, v)) {
				return
			}
		}
	})
}

func Chan[T any](ch chan T) Iterator[T] {
	return ValueIterator[int, T](ChanWithIndex[T](ch))
}

func ChanWithIndex[T any](ch chan T) Iterator[alg.Pair[int, T]] {
	return IteratorFunc[alg.Pair[int, T]](func(fn func(alg.Pair[int, T]) bool) {
		i := 0
		for {
			e, ok := <-ch
			if !ok {
				return
			}

			if !fn(alg.Two(i, e)) {
				return
			}
			i++
		}
	})
}