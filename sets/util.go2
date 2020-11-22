package sets

import "github.com/tdakkota/algo2/alg"

// Creates copy of set
func Copy[T comparable](src Set[T]) (dst Set[T]) {
	dst = HashSet[T](src.Len())
	Fill(src, dst)
	return
}

// Fill src from dst
func Fill[T any](src, dst Set[T]) {
	src.Iterate(func(t T) bool {
		dst.Add(t)
		return true
	})
	return
}

// Add adds all the elements of b to a.
func Add[T any](a, b Set[T]) {
	b.Iterate(func(v T) bool {
		a.Add(v)
		return true
	})
}

// Sub removes all elements in a from b.
func Sub[T any](a, b Set[T]) {
	b.Iterate(func(v T) bool {
		a.Delete(v)
		return true
	})
}

// Union creates HashSet and puts to it union of a and b
func Union[T comparable](a, b Set[T]) (dst Set[T]) {
	dst = HashSet[T](a.Len() + b.Len())
	UnionTo(a, b, dst)
	return
}

// UnionTo puts union of a and b to dst
func UnionTo[T any](a, b, dst Set[T]) {
	Add[T](dst, a)
	Add[T](dst, b)
}

// Complement creates HashSet and puts to it complement(difference) of a and b
func Complement[T comparable](a, b Set[T]) (dst Set[T]) {
	dst = HashSet[T](a.Len())
	ComplementTo(a, b, dst)
	return
}

// ComplementTo puts complement(difference) of a and b to dst
func ComplementTo[T any](a, b, dst Set[T]) {
	Fill(a, dst)
	Sub[T](dst, b)
}

// Diff creates HashSet and puts to it symmetric difference of a and b
func Diff[T comparable](a, b Set[T]) (dst Set[T]) {
	dst = HashSet[T](a.Len())
	DiffTo(a, b, dst)
	return
}

// DiffTo puts symmetric difference of a and b to dst
func DiffTo[T any](a, b, dst Set[T]) {
	a.Iterate(func(v T) bool {
		if !b.Contains(v) {
			dst.Add(v)
		}
		return true
	})

	b.Iterate(func(v T) bool {
		if !a.Contains(v) {
			dst.Add(v)
		}
		return true
	})
}

// Intersect creates HashSet and puts to it intersection of a and b
func Intersect[T comparable](a, b Set[T]) (dst Set[T]) {
	dst = HashSet[T](a.Len())
	IntersectTo(a, b, dst)
	return
}

// IntersectTo puts intersection of a and b to dst
func IntersectTo[T any](a, b, dst Set[T]) {
	a.Iterate(func(v T) bool {
		if b.Contains(v) {
			dst.Add(v)
		}
		return true
	})
}

// Subset returns true if a is subset of b.
func Subset[T any](a, b Set[T]) (r bool) {
	r = true

	a.Iterate(func(v T) bool {
		if !b.Contains(v) {
			r = false
			return false
		}
		return true
	})

	return
}

// Equal returns true if a is subset of b and b is subset of a.
func Equal[T any](a, b Set[T]) bool {
	return Subset[T](a, b) && Subset[T](b, a)
}

// Cartesian returns cartesian product of a and b.
func Cartesian[T comparable](a, b Set[T]) (dst Set[alg.Pair[T, T]]) {
	dst = HashSet[alg.Pair[T, T]](a.Len() + b.Len())
	CartesianTo[T](a, b, dst)
	return
}

// CartesianTo puts cartesian product of a and b to dst
func CartesianTo[T any](a, b Set[T], dst Set[alg.Pair[T, T]]) {
	a.Iterate(func(l T) bool {
		b.Iterate(func(r T) bool {
			dst.Add(alg.Two(l, r))
			return true
		})
		return true
	})
}
