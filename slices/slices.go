package slices

import (
	"constraints"

	"github.com/tdakkota/algo2/alg"
)

func Repeat[T any](value T, count int) []T {
	switch {
	case count == 0:
		return []T{}
	case count == 1:
		return []T{value}
	case count < 0:
		panic("negative Repeat count")
	}

	nb := make([]T, count)
	bp := copy(nb, []T{value, value})
	for bp < len(nb) {
		copy(nb[bp:], nb[:bp])
		bp *= 2
	}
	return nb
}

func RepeatFn[T any](fn func(int) T, count int) []T {
	switch {
	case count == 0:
		return []T{}
	case count == 1:
		return []T{fn(0)}
	case count < 0:
		panic("negative Repeat count")
	}

	nb := make([]T, count)
	for i := range nb {
		nb[i] = fn(i)
	}
	return nb
}

// Reverse reverses slice.
func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// ReverseCopy is like Reverse, instead of mutating given slice,
// it returns a copy
func ReverseCopy[T any](i []T) []T {
	s := make([]T, len(i))
	copy(s, i)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// Equal reports whether two slices are equal: the same length and all
// elements equal. All floating point NaNs are considered equal.
func Equal[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v1 := range s1 {
		v2 := s2[i]
		if v1 != v2 {
			isNaN := func(f T) bool { return f != f }
			if !isNaN(v1) || !isNaN(v2) {
				return false
			}
		}
	}
	return true
}

// Map turns a []Elem1 to a []Elem2 using a mapping function.
func Map[Elem1, Elem2 any](s []Elem1, f func(Elem1) Elem2) []Elem2 {
	r := make([]Elem2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Reduce reduces a []Elem1 to a single value of type Elem2 using
// a reduction function.
func Reduce[Elem1, Elem2 any](s []Elem1, initializer Elem2, f func(Elem2, Elem1) Elem2) Elem2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// Filter filters values from a slice using a filter function.
func Filter[Elem any](s []Elem, f func(Elem) bool) []Elem {
	var r []Elem
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// Max returns the maximum element in a slice of some ordered type.
// If the slice is empty it returns the zero value of the element type.
func Max[Elem constraints.Ordered](s []Elem) Elem {
	if len(s) == 0 {
		var zero Elem
		return zero
	}
	return Reduce(s[1:], s[0], alg.Max[Elem])
}

// Min returns the minimum element in a slice of some ordered type.
// If the slice is empty it returns the zero value of the element type.
func Min[Elem constraints.Ordered](s []Elem) Elem {
	if len(s) == 0 {
		var zero Elem
		return zero
	}
	return Reduce(s[1:], s[0], alg.Min[Elem])
}
