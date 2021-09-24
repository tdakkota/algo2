package cache

import (
	"github.com/tdakkota/algo2/alg"
	"github.com/tdakkota/algo2/list"
)

type LRUCache[K comparable, V any] struct {
	capacity int
	cache    map[K]*list.Element[alg.Pair[K, V]]
	lruList  *list.LinkedList[alg.Pair[K, V]]
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] { 
	l := &LRUCache[K, V]{
		capacity: capacity,
		lruList:  list.NewLinkedList[alg.Pair[K, V]](),
	}

	if capacity > 0 {
		l.cache = make(map[K]*list.Element[alg.Pair[K, V]], capacity)
	} else {
		l.cache = map[K]*list.Element[alg.Pair[K, V]]{}
	}

	return l
}

func (l *LRUCache[K, V]) Get(key K) (v V, ok bool) {
	if found, ok := l.cache[key]; ok {
		l.lruList.MoveToFront(found)
		return found.Value.R, true
	}
	return
}

func (l *LRUCache[K, V]) Put(key K, value V) {
	if l.capacity == 0 {
		return
	}

	if found, ok := l.cache[key]; ok {
		found.Value.R = value
		l.lruList.MoveToFront(found)
	} else {
		if len(l.cache) >= l.capacity {
			l.Delete(l.lruList.Back().Value.L)
		}

		l.cache[key] = l.lruList.PushFront(alg.Two(key, value))
	}
}

func (l *LRUCache[K, V]) Delete(key K) bool {
	found, ok := l.cache[key]
	if !ok {
		return false
	}

	l.lruList.Remove(found)
	delete(l.cache, key)
	return true
}

func (l *LRUCache[K, V]) Len() int {
	return len(l.cache)
}

func (l *LRUCache[K, V]) Cap() int {
	return l.capacity
}