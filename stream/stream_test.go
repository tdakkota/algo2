package stream

import (
	"testing"

	"github.com/tdakkota/algo2/slices"
	"github.com/tdakkota/algo2/testutil"
)

func TestStream(t *testing.T) {
	s := Of[int](1, 2, 3, 4, 5, 6, 7, 8, 9)
	s = s.Skip(3).
		Filter(func(i int) bool {
			return i%2 == 0
		}).
		Limit(2)
	s = Map[int, int](s, func(e int) int {
		return e * 2
	}).
		Skip(0).
		Filter(func(i int) bool {
			return i%2 == 0
		}).
		Limit(2)
	expected := []int{8, 12}
	testutil.EqualFn(t, expected, s.ToSlice(), slices.Equal[int])
}
