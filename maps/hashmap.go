package maps

import (
	"github.com/tdakkota/algo2/alg"
)

var _ Map[int, struct{}] = HashMap[int, struct{}]{}

type HashMap[K comparable, V any] struct {
	m map[K]V
}

func NewHashMap[K comparable, V any](capacity int) HashMap[K, V] {
	var m map[K]V
	if capacity >= 1 {
		m = make(map[K]V, capacity)
	} else {
		m = map[K]V{}
	}

	return HashMap[K, V]{m: m}
}


func (m HashMap[K, V]) Get(k K) (v V, ok bool) {
	v, ok = m.m[k]
	return 
}

func (m HashMap[K, V]) Put(k K, v V) {
	m.m[k] = v
}

func (m HashMap[K, V]) Delete(k K) (ok bool) {
	_, ok = m.m[k]
	if ok {
		delete(m.m, k)
	}
	return ok
}

func (m HashMap[K, V]) Len() int  {
	return len(m.m)
}

func (m HashMap[K, V]) Iterate(cb func(alg.Pair[K, V]) bool) {
	for k, v := range m.m {
		if !cb(alg.Two(k, v)) {
			return
		}
	}
}