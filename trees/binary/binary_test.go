package binary

import (
	"testing"

	"github.com/tdakkota/algo2/alg"
	"github.com/tdakkota/algo2/slices"
	"github.com/tdakkota/algo2/testutil"
)

func TestBinary(t *testing.T) {
	const sizeTest = 5
	m := NewBinary[int, int]()
	values := slices.RepeatFn(func(i int) int {
		return i
	}, sizeTest)
	r := make([]int, sizeTest)

	for i, value := range values {
		_, found := m.Get(i)
		testutil.False(t, found)
		found = m.Delete(i)
		testutil.False(t, found)

		m.Put(i, value)
		v, _ := m.Get(i)
		testutil.Equal(t, value, v)
	}

	if min, ok := m.Min(); !ok || min.R != values[0] {
		t.Errorf("expected minimal is %v, got %v", values[0], min.R)
	}

	if max, ok := m.Max(); !ok || max.R != values[len(values)-1] {
		t.Errorf("expected maximal is %v, got %v", values[len(values)-1], max.R)
	}

	r = make([]int, sizeTest)
	m.Iterate(func(p alg.Pair[int, int]) bool {
		r[p.L] = p.R
		return true
	})
	testutil.EqualFn(t, values, r, slices.Equal[int])

	for i, value := range values {
		v, found := m.Get(i)
		testutil.True(t, found)
		testutil.Equal(t, value, v)

		m.Put(i, value)
		v, found = m.Get(i)
		testutil.True(t, found)
		testutil.Equal(t, value, v)
	}

	r = make([]int, sizeTest)
	m.Iterate(func(p alg.Pair[int, int]) bool {
		r[p.L] = p.R
		return true
	})
	testutil.EqualFn(t, values, r, slices.Equal[int])

	r = make([]int, sizeTest)
	m.IteratePreOrder(func(p alg.Pair[int, int]) bool {
		r[p.L] = p.R
		return true
	})
	testutil.EqualFn(t, values, r, slices.Equal[int])

	r = make([]int, sizeTest)
	m.IteratePostOrder(func(p alg.Pair[int, int]) bool {
		r[p.L] = p.R
		return true
	})
	testutil.EqualFn(t, values, r, slices.Equal[int])

	for i := range values {
		if found := m.Delete(i); !found {
			t.Errorf("unexpectedly not found when delete")
		}

		_, found := m.Get(i)
		if found {
			t.Errorf("expected delete value is not found")
		}
	}

	if min, ok := m.Min(); ok {
		t.Errorf("expected minimal not found, got %v", min.R)
	}

	if max, ok := m.Max(); ok {
		t.Errorf("expected maximal not found, got %v", max.R)
	}

	var i int
	m.Iterate(func(p alg.Pair[int, int]) bool {
		i++
		return true
	})
	testutil.Equal(t, 0, i)
}
