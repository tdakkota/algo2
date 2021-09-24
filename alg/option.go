package alg

// Option is a parameterized functional option type.
type Option[T any] func(*T)

type Options[T any] []Option[T]

func (o Options[T]) Perform(t *T) {
	for _, op := range o {
		op(t)
	}
}
