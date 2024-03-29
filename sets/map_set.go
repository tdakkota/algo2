package sets

import (
	"constraints"

	"github.com/tdakkota/algo2/alg"
	"github.com/tdakkota/algo2/maps"
	"github.com/tdakkota/algo2/trees/binary"
)

type MapSet[T any] struct {
	m maps.Map[T, struct{}]
}

func HashSet[T comparable](capacity int) MapSet[T] {
	return FromMap[T](maps.NewHashMap[T, struct{}](capacity))
}

func TreeSet[T constraints.Ordered]() MapSet[T] {
	return FromMap[T](binary.NewRedBlack[T, struct{}]())
}

func Of[T comparable](e ...T) MapSet[T] {
	set := HashSet[T](len(e))
	for _, t := range e {
		set.Add(t)
	}
	return set
}

func FromMap[T any](m maps.Map[T, struct{}]) MapSet[T] {
	return MapSet[T]{
		m: m,
	}
}

func (s MapSet[T]) Contains(t T) bool {
	_, ok := s.m.Get(t)
	return ok
}

func (s MapSet[T]) Add(t T) {
	s.m.Put(t, struct{}{})
}

func (s MapSet[T]) Delete(t T) bool {
	return s.m.Delete(t)
}

func (s MapSet[T]) Len() int {
	return s.m.Len()
}

func (s MapSet[T]) Iterate(cb func(T) bool) {
	s.m.Iterate(func(p alg.Pair[T, struct{}]) bool {
		return cb(p.L)
	})
}
