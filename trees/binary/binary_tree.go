package binary

import (
	"constraints"

	"github.com/tdakkota/algo2/alg"
)

// Binary represents the binary search tree.
type Binary[K, V any] struct {
	root *btnode[K, V]
	ord  alg.Ord[K]
	size int
}

func NewBinary[K constraints.Ordered, V any]() *Binary[K, V] {
	return NewBinaryWithOrd[K, V](alg.BuiltinOrd[K])
}

func NewBinaryWithOrd[K, V any](ord alg.Ord[K]) *Binary[K, V] {
	b := new(Binary[K, V])
	b.ord = ord
	return b
}

// Min returns the minimal element.
func (m *Binary[K, V]) Min() (v alg.Pair[K, V], ok bool) {
	if min := m.root.min(); min != nil {
		return alg.Two(min.key, min.value), true
	}
	return
}

// Max returns the maximal element.
func (m *Binary[K, V]) Max() (v alg.Pair[K, V], ok bool) {
	if max := m.root.max(); max != nil {
		return alg.Two(max.key, max.value), true
	}
	return
}

// Len returns size of tree.
func (m *Binary[K, V]) Len() int {
	return m.size
}

func (m *Binary[K, V]) Clear() {
	m.root = nil
}

func (m *Binary[K, V]) find(key K) **btnode[K, V] {
	pn := &m.root
	for *pn != nil {
		o := m.ord((*pn).key, key)
		switch o {
		case -1:
			pn = &(*pn).left
		case 1:
			pn = &(*pn).right
		default:
			return pn
		}
	}
	return pn
}

// Put inserts a new key/value into the map.
// If the key is already present, the value is replaced.
func (m *Binary[K, V]) Put(key K, val V) {
	m.size++
	pn := m.find(key)
	if *pn != nil {
		(*pn).value = val
		return
	}
	*pn = &btnode[K, V]{key: key, value: val, tree: m}
}

// Get returns the value associated with a key, or the zero value
// if not present. The found result reports whether the key was found.
func (m *Binary[K, V]) Get(key K) (v V, ok bool) {
	pn := m.find(key)
	if *pn == nil {
		return alg.Zero[V](), false
	}
	return (*pn).value, true
}

func (m *Binary[K, V]) delete(key K) (_ bool) {
	if m.root == nil {
		return
	}

	if m.ord(key, m.root.key) == 0 {
		tempRoot := &btnode[K, V]{tree: m}
		tempRoot.left = m.root
		r := m.root.del(tempRoot, key)
		m.root = tempRoot.left
		return r
	}
	return m.root.left.del(m.root, key) || m.root.right.del(m.root, key)
}

// Delete removes node with given key from tree and returns true if
// deletion is successful.
func (m *Binary[K, V]) Delete(key K) (ok bool) {
	ok = m.delete(key)
	if ok {
		m.size--
	}
	return
}

// Iterate traverse tree in order using given callback.
func (m *Binary[K, V]) iterate(o order, cb func(alg.Pair[K, V]) bool) {
	if m.root != nil {
		m.root.traverse(o, cb)
	}
}

// Iterate traverse tree in order using given callback.
func (m *Binary[K, V]) Iterate(cb func(alg.Pair[K, V]) bool) {
	m.IterateInOrder(cb)
}

// IterateInOrder traverse tree in order using given callback.
func (m *Binary[K, V]) IterateInOrder(cb func(alg.Pair[K, V]) bool) {
	m.iterate(inOrder, cb)
}

// IteratePreOrder traverse tree in pre-order using given callback.
func (m *Binary[K, V]) IteratePreOrder(cb func(alg.Pair[K, V]) bool) {
	m.iterate(preOrder, cb)
}

// IteratePostOrder traverse tree in post-order using given callback.
func (m *Binary[K, V]) IteratePostOrder(cb func(alg.Pair[K, V]) bool) {
	m.iterate(postOrder, cb)
}
