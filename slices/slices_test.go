package slices

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/tdakkota/algo2/testutil"
)

func TestRepeat(t *testing.T) {
	tests := []struct {
		count int
		value int
	}{
		{0, 1},
		{1, 1},
		{2, 1},
		{1023, 0},
		{1023, 1},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Test Repeat %d %d times", test.value, test.count), func(t *testing.T) {
			a := Repeat[int](test.value, test.count)
			testutil.Equal(t, len(a), test.count)
			testutil.EqualFn(t, a, Repeat[int](test.value, test.count), Equal[int])
		})
	}
}

var indexFn = func(i int) int { return i }

func TestRepeatFn(t *testing.T) {
	tests := []struct {
		count int
		value func(int) int
	}{
		{0, indexFn},
		{1, indexFn},
		{2, indexFn},
		{2, func(i int) int { return 1 }},
		{1023, indexFn},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Test RepeatFn function %d times", test.count), func(t *testing.T) {
			a := RepeatFn[int](test.value, test.count)
			testutil.Equal(t, len(a), test.count)
			testutil.EqualFn(t, a, RepeatFn[int](test.value, test.count), Equal[int])
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Test Reverse function %v", test.input), func(t *testing.T) {
			Reverse[int](test.input)
			testutil.EqualFn(t, test.expected, test.input, Equal[int])
		})
	}
}

func TestReverseCopy(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Test ReverseCopy function %v", test.input), func(t *testing.T) {
			testutil.EqualFn(t, test.expected, ReverseCopy[int](test.input), Equal[int])
		})
	}
}

func TestEqual(t *testing.T) {
	s1 := []int{1, 2, 3}
	if !Equal(s1, s1) {
		t.Errorf("Equal(%v, %v) = false, want true", s1, s1)
	}
	s2 := []int{1, 2, 3}
	if !Equal(s1, s2) {
		t.Errorf("Equal(%v, %v) = false, want true", s1, s2)
	}
	s2 = append(s2, 4)
	if Equal(s1, s2) {
		t.Errorf("Equal(%v, %v) = true, want false", s1, s2)
	}

	s3 := []float64{1, 2, math.NaN()}
	if !Equal(s3, s3) {
		t.Errorf("Equal(%v, %v) = false, want true", s3, s3)
	}

	if Equal(s1, nil) {
		t.Errorf("Equal(%v, nil) = true, want false", s1)
	}
	if Equal(nil, s1) {
		t.Errorf("Equal(nil, %v) = true, want false", s1)
	}
	if !Equal(s1[:0], nil) {
		t.Errorf("Equal(%v, nil = false, want true", s1[:0])
	}
}

func TestMap(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := Map(s1, func(i int) float64 { return float64(i) * 2.5 })
	if want := []float64{2.5, 5, 7.5}; !Equal(s2, want) {
		t.Errorf("Map(%v, ...) = %v, want %v", s1, s2, want)
	}

	s3 := []string{"Hello", "World"}
	s4 := Map(s3, strings.ToLower)
	if want := []string{"hello", "world"}; !Equal(s4, want) {
		t.Errorf("Map(%v, strings.ToLower) = %v, want %v", s3, s4, want)
	}

	s5 := Map(nil, func(i int) int { return i })
	if len(s5) != 0 {
		t.Errorf("Map(nil, identity) = %v, want empty slice", s5)
	}
}

func TestReduce(t *testing.T) {
	s1 := []int{1, 2, 3}
	r := Reduce(s1, 0, func(f float64, i int) float64 { return float64(i)*2.5 + f })
	if want := 15.0; r != want {
		t.Errorf("Reduce(%v, 0, ...) = %v, want %v", s1, r, want)
	}

	if got := Reduce(nil, 0, func(i, j int) int { return i + j }); got != 0 {
		t.Errorf("Reduce(nil, 0, add) = %v, want 0", got)
	}
}

func TestFilter(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := Filter(s1, func(i int) bool { return i%2 == 0 })
	if want := []int{2}; !Equal(s2, want) {
		t.Errorf("Filter(%v, even) = %v, want %v", s1, s2, want)
	}

	if s3 := Filter(s1[:0], func(i int) bool { return true }); len(s3) > 0 {
		t.Errorf("Filter(%v, identity) = %v, want empty slice", s1[:0], s3)
	}
}

func TestMax(t *testing.T) {
	s1 := []int{1, 2, 3, -5}
	if got, want := Max(s1), 3; got != want {
		t.Errorf("Max(%v) = %d, want %d", s1, got, want)
	}

	s2 := []string{"aaa", "a", "aa", "aaaa"}
	if got, want := Max(s2), "aaaa"; got != want {
		t.Errorf("Max(%v) = %q, want %q", s2, got, want)
	}

	if got, want := Max(s2[:0]), ""; got != want {
		t.Errorf("Max(%v) = %q, want %q", s2[:0], got, want)
	}
}

func TestMin(t *testing.T) {
	s1 := []int{1, 2, 3, -5}
	if got, want := Min(s1), -5; got != want {
		t.Errorf("Min(%v) = %d, want %d", s1, got, want)
	}

	s2 := []string{"aaa", "a", "aa", "aaaa"}
	if got, want := Min(s2), "a"; got != want {
		t.Errorf("Min(%v) = %q, want %q", s2, got, want)
	}

	if got, want := Min(s2[:0]), ""; got != want {
		t.Errorf("Min(%v) = %q, want %q", s2[:0], got, want)
	}
}
