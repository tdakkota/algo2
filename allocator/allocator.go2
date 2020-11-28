package allocator

type Allocator[T any] interface {
	New() (*T, bool)
	Free(*T)
}

type GCAllocator[T any] struct{}

func (GCAllocator[T]) New() (*T, bool) {
	return new(T), true
}

func (GCAllocator[T]) NewRange(size int) ([]T, bool) {
	return make([]T, size), true
}

func (GCAllocator[T]) Free(p *T) {}

func (GCAllocator[T]) FreeRange(p []T) {}

func Default[T any]() Allocator[T] {
	return GCAllocator[T]{}
}
