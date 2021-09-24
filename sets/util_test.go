package sets

import (
	"sort"
	"testing"

	"github.com/tdakkota/algo2/alg"
	"github.com/tdakkota/algo2/iterator"
	"github.com/tdakkota/algo2/slices"
	"github.com/tdakkota/algo2/testutil"
)

func equal(t *testing.T, set Set[int], expected []int) {
	t.Helper()
	values := iterator.Values[int](set)
	sort.Ints(values)
	testutil.EqualFn(t, expected, values, slices.Equal[int])
}

func TestCopy(t *testing.T) {
	cp := Copy[int](Of(1, 2, 3))
	equal(t, cp, []int{1, 2, 3})
}

func TestUnion(t *testing.T) {
	a, b := Of(1, 2, 3), Of(3, 4, 5)
	equal(t, Union[int](a, b), []int{1, 2, 3, 4, 5})
	testutil.EqualFn(t, Set[int](a), Union[int](a, a), Equal[int])
}

func TestComplement(t *testing.T) {
	a, b := Of(1, 2, 3, 4), Of(3, 4, 5, 6, 7)
	c := Complement[int](a, b)
	equal(t, c, []int{1, 2})

	c = Complement[int](b, a)
	equal(t, c, []int{5, 6, 7})

	c = Complement[int](a, a)
	testutil.Equal(t, 0, c.Len())
}

func TestDiff(t *testing.T) {
	a, b := Of(1, 2, 3, 4, 5), Of(3, 4, 5, 6, 7)
	c := Diff[int](a, b)
	equal(t, c, []int{1, 2, 6, 7})

	c = Complement[int](a, a)
	testutil.Equal(t, 0, c.Len())
}

func TestIntersect(t *testing.T) {
	a, b := Of(1, 2, 3, 4), Of(3, 4, 5, 6, 7)
	c := Intersect[int](a, b)
	equal(t, c, []int{3, 4})
	testutil.EqualFn(t, Set[int](a), Intersect[int](a, a), Equal[int])
}

func TestCartesian(t *testing.T) {
	a, b := Of(1, 2), Of(3, 4)
	c := Cartesian[int](a, b)
	v := HashSet[alg.Pair[int, int]](0)
	v.Add(alg.Two(1, 3))
	v.Add(alg.Two(1, 4))
	v.Add(alg.Two(2, 3))
	v.Add(alg.Two(2, 4))

	c.Iterate(func(p alg.Pair[int, int]) bool {
		v.Delete(p)
		return true
	})

	if v.Len() != 0 {
		t.Fatal("expected all pairs were met")
	}
}
