package atomic

import "sync/atomic"

type Value[T any] struct {
	inner atomic.Value
}

func (v *Value[T]) Load() T {
	return v.inner.Load().(T)
}

func (v *Value[T]) Store(x T)  {
	v.inner.Store(x)
}