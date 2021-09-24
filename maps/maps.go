package maps

import (
	"github.com/tdakkota/algo2/alg"
)

type Map[K, V any] interface {
	Get(K) (V, bool)
	Put(K, V)
	Delete(K) bool
	Len() int
	Iterate(func(alg.Pair[K, V]) bool)
}
