package iterator

import (
	"constraints"
	"fmt"
	"math"
	"testing"
)

type rangeTestCase[T constraints.Integer] struct {
	from, to, step T
	expected       []T
}

func (t rangeTestCase[T]) Iterator() Iterator[T] {
	return Range[T](t.from, t.to, t.step)
}

func (t rangeTestCase[T]) Name() string {
	return fmt.Sprintf("from %v to %v by %v", t.from, t.to, t.step)
}

func (t rangeTestCase[T]) Expected() []T {
	return t.expected
}

func TestRange(t *testing.T) {
	t.Run(`int`, runTests[int, rangeTestCase[int]]([]rangeTestCase[int]{
		{1, 5, 1, []int{1, 2, 3, 4}},
		{1, 5, 5, []int{1}},
		{1, 5, 2, []int{1, 3}},
		{1, 5, -1, nil},
		{1, 0, -1, []int{1}},
		{5, 1, -1, []int{5, 4, 3, 2}},
		{5, 1, -2, []int{5, 3}},
	}...))

	t.Run(`float`, runTests[float64, rangeTestCase[float64]]([]rangeTestCase[float64]{
		{1, 5, 1, []float64{1, 2, 3, 4}},
		{1, 5, 5, []float64{1}},
		{1, 5, 2, []float64{1, 3}},
		{1, 5, -1, nil},
		{1, 0, -1, []float64{1}},
		{5, 1, -1, []float64{5, 4, 3, 2}},
		{5, 1, -2, []float64{5, 3}},

		{math.NaN(), 5, 1, nil},
		{1, math.NaN(), 1, nil},
		{1, 5, math.NaN(), nil},
	}...))
}

type repeatTestCase[T constraints.Real] struct {
	value    T
	times    uint
	expected []T
}

func (t repeatTestCase[T]) Iterator() Iterator[T] {
	return Repeat[T](t.value, t.times)
}

func (t repeatTestCase[T]) Name() string {
	return fmt.Sprintf("repeat %v %v times", t.value, t.times)
}

func (t repeatTestCase[T]) Expected() []T {
	return t.expected
}

func TestRepeat(t *testing.T) {
	runTests[int, repeatTestCase[int]]([]repeatTestCase[int]{
		{1, 4, []int{1, 1, 1, 1}},
		{1, 0, nil},
	}...)(t)
}
