// Package alg provides some common basic function.
package alg

import "constraints"

// Max returns the maximum of two values of some ordered type.
func Max[T constraints.Ordered](a, b T) T {
	if Lt(a, b) {
		return b
	}
	return a
}

// Min returns the minimum of two values of some ordered type.
func Min[T constraints.Ordered](a, b T) T {
	if Lt(a, b) {
		return a
	}
	return b
}

func Gt[T constraints.Ordered](a, b T) bool {
	return a > b
}

func Eq[T comparable](a, b T) bool {
	return a == b
}

func Lt[T constraints.Ordered](a, b T) bool {
	return a < b
}

func GtEq[T constraints.Ordered](a, b T) bool {
	return Eq[T](a, b) || Gt[T](a, b)
}

func LtEq[T constraints.Ordered](a, b T) bool {
	return Eq[T](a, b) || Lt[T](a, b)
}

type Order int

const (
	Lesser  Order = -1
	Equal         = 0
	Greater       = 1
)

type Ord[T any] func(a, b T) Order

func BuiltinOrd[T constraints.Ordered](a, b T) (_ Order) {
	switch {
	case a < b:
		return -1
	case a == b:
		return 0
	case a > b:
		return 1
	}

	return
}
