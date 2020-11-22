package sets

type Set[T any] interface {
	Contains(T) bool
	Add(T)
	Delete(T) bool
	Len() int
	Iterate(func(T) bool)
}
