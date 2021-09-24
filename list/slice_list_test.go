package list

import (
	"testing"

	"github.com/tdakkota/algo2/alg"
	"github.com/tdakkota/algo2/slices"
	"github.com/tdakkota/algo2/testutil"
)


func TestSliceList(t *testing.T) {
	l := SliceList[int]([]int{1, 2, 3, 4, 5})
	testutil.Equal(t, 5, l.Len())

	l.Append(6)
	testutil.Equal(t, 6, l.Len())
	v, _ := l.Pop()
	testutil.Equal(t, 5, l.Len())
	testutil.Equal(t, 6, v)

	v, _ = l.PopFront()
	testutil.Equal(t, 4, l.Len())
	testutil.Equal(t, 1, v)

	v = l.Remove(1)
	testutil.Equal(t, 3, l.Len())
	testutil.Equal(t, 3, v)

	l.Prepend(1)
	testutil.Equal(t, 4, l.Len())

	r := make([]int, l.Len())
	l.Iterate(func(p alg.Pair[int, int]) bool {
		r[p.L] = p.R
		return true
	})
	testutil.EqualFn(t, []int{1, 2, 4, 5}, r, slices.Equal[int])
}