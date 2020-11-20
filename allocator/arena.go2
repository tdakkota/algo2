package allocator

import "github.com/tdakkota/algo2/alg"

type ArenaAllocator[T any] struct {
	index int
	arena []T
}

func NewArena[T any](size int) *ArenaAllocator[T] {
	return &ArenaAllocator[T]{
		arena: make([]T, size),
	}
}

func (a ArenaAllocator[T]) Len() int {
	return len(a.arena)
}

func (a ArenaAllocator[T]) Used() int {
	return a.index
}

func (a *ArenaAllocator[T]) New() (*T, bool) {
	if a.index >= len(a.arena) {
		return nil, false
	}
	e := &a.arena[a.index]
	a.index++
	return e, true
}

func (a *ArenaAllocator[T]) Free(p *T) {
	if &a.arena[a.index] == p {
		a.arena[a.index] = alg.Zero[T]()
		a.index--
	}
}
