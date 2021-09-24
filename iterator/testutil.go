package iterator

import (
	"testing"

	"github.com/tdakkota/algo2/slices"
	"github.com/tdakkota/algo2/constraints"
	"github.com/tdakkota/algo2/testutil"
)

type iteratorTestCase[T comparable] struct {
	name 	 string
	iterator Iterator[T]
	expected []T
}

func (test iteratorTestCase[T]) runner() func(t *testing.T) {
	return func(t *testing.T) {
		t.Helper()
		t.Run(test.name, func(t *testing.T) {
				var r []T
				test.iterator.Iterate(func(e T) bool {
					r = append(r, e)
					return true
				})
	
				testutil.EqualFn(t, test.expected, r, slices.Equal[T])
		})
	}
}

type testCaseBuilder[T comparable] interface {
	Name() string
	Iterator() Iterator[T]
	Expected() []T
}

func runTests[T constraints.Real, B testCaseBuilder[T]](tests ...B) func(t *testing.T) {
	return func(t *testing.T) {
		t.Helper()

		for _, test := range tests {
			iteratorTestCase[T]{
				name: test.Name(),
				iterator: test.Iterator(),
				expected: test.Expected(),
			}.runner()(t)
		}
	}
}