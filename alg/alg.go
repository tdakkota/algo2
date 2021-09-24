// Package alg provides some common basic function.
package alg

func Zero[T any]() (_ T) { return }

func IsZero[T comparable](t T) bool {
	// check is equal to zero and it is not a NaN
	return t == Zero[T]() && t == t
}

func IsNaN[T comparable](t T) bool {
	return t != t
}

type Pair[L, R any] struct {
	L L
	R R
}

func Two[L, R any](l L, r R) Pair[L, R] {
	return Pair[L, R]{L: l, R: r}
}
