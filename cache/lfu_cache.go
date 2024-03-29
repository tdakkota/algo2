package cache

import (
	"github.com/tdakkota/algo2/list"
)

type lfuCacheNode[K, V any] struct {
	key    K
	value  V
	parent *list.Element[freqNode[K, V]]
}

type freqNode[K, V any] struct {
	freq int
	list *list.LinkedList[lfuCacheNode[K, V]]
}

func (l *LFUCache[K, V]) newFreqNode(freq int) freqNode[K, V] {
	return freqNode[K, V]{
		freq: freq,
		list: list.NewLinkedList[lfuCacheNode[K, V]](),
	}
}

func (l *LFUCache[K, V]) evictNode(e *list.Element[lfuCacheNode[K, V]]) {
	// remove the cache node from the linkedList
	// remove the freqNode(parent) if it is empty
	// do nothing to the map
	node := e.Value
	parent := node.parent.Value
	parent.list.Remove(e)

	if parent.list.Len() == 0 {
		l.removeFreqNode(node.parent)
	}
}

func (l *LFUCache[K, V]) moveAddOneFreq(e *list.Element[lfuCacheNode[K, V]], key K, value V) *list.Element[lfuCacheNode[K, V]] {
	var freq *list.Element[freqNode[K, V]]
	n := e.Value
	parentElement := n.parent

	if next := parentElement.Next(); next != nil && next.Value.freq == (parentElement.Value.freq+1) {
		freq = next
	} else {
		freq = l.lfuHead.InsertAfter(l.newFreqNode(parentElement.Value.freq+1), parentElement)
	}

	cache := lfuCacheNode[K, V]{
		key:    key,
		value:  value,
		parent: freq,
	}
	listElement := freq.Value.list.PushFront(cache)
	l.evictNode(e)

	return listElement
}

func (l *LFUCache[K, V]) removeFreqNode(node *list.Element[freqNode[K, V]]) {
	if node.Value.freq == 0 {
		panic("should not remove the head")
	}
	l.lfuHead.Remove(node)
}

type LFUCache[K comparable, V any] struct {
	capacity int
	cache    map[K]*list.Element[lfuCacheNode[K, V]]
	lfuHead  *list.LinkedList[freqNode[K, V]]
}

func NewLFUCache[K comparable, V any](capacity int) *LFUCache[K, V] {
	l := &LFUCache[K, V]{
		capacity: capacity,
		lfuHead:  list.NewLinkedList[freqNode[K, V]](),
	}

	if capacity > 0 {
		l.cache = make(map[K]*list.Element[lfuCacheNode[K, V]], capacity)
	} else {
		l.cache = map[K]*list.Element[lfuCacheNode[K, V]]{}
	}

	return l
}

func (l *LFUCache[K, V]) Get(key K) (v V, ok bool) {
	if found, ok := l.cache[key]; ok {
		newCacheNode := l.moveAddOneFreq(found, key, found.Value.value)
		l.cache[key] = newCacheNode
		return found.Value.value, true
	}
	return
}

func (l *LFUCache[K, V]) Put(key K, value V) {
	if l.capacity == 0 {
		return
	}

	if found, ok := l.cache[key]; ok {
		l.cache[key] = l.moveAddOneFreq(found, key, value)
	} else {
		if len(l.cache) >= l.capacity {
			lfuNode := l.lfuHead.Front().Value.list.Back()
			l.Delete(lfuNode.Value.key)
		}

		var first *list.Element[freqNode[K, V]]
		if next := l.lfuHead.Front(); next != nil && next.Value.freq == 1 {
			first = next
		} else {
			first = l.lfuHead.PushFront(l.newFreqNode(1))
		}

		newCacheNode := lfuCacheNode[K, V]{
			key:    key,
			value:  value,
			parent: first,
		}
		l.cache[key] = first.Value.list.PushFront(newCacheNode)
	}
}

func (l *LFUCache[K, V]) Delete(key K) bool {
	found, ok := l.cache[key]
	if !ok {
		return false
	}

	l.evictNode(found)
	delete(l.cache, key)
	return true
}

func (l *LFUCache[K, V]) Len() int {
	return len(l.cache)
}

func (l *LFUCache[K, V]) Cap() int {
	return l.capacity
}
