package binary

import (
	"github.com/tdakkota/algo2/alg"
)

// btnode is a binary search tree node.
type btnode[K, V any] struct {
	key         K
	value       V
	tree        *Binary[K, V]
	left, right *btnode[K, V]
}

type order int

const (
	inOrder order = iota
	preOrder
	postOrder
)

func (m *btnode[K, V]) traverse(o order, cb func(alg.Pair[K, V]) bool) (_ bool) {
	switch o {
	case inOrder:
		if m.left != nil && !m.left.traverse(o, cb) {
			return
		}

		if !cb(alg.Two(m.key, m.value)) {
			return
		}

		if m.right != nil && !m.right.traverse(o, cb) {
			return
		}
	case preOrder:
		if !cb(alg.Two(m.key, m.value)) {
			return
		}

		if m.left != nil && !m.left.traverse(o, cb) {
			return
		}

		if m.right != nil && !m.right.traverse(o, cb) {
			return
		}
	case postOrder:
		if m.left != nil && !m.left.traverse(o, cb) {
			return
		}

		if m.right != nil && !m.right.traverse(o, cb) {
			return
		}

		if !cb(alg.Two(m.key, m.value)) {
			return
		}
	default:
		panic("binary tree: invalid traverse order")
	}

	return true
}

func (n *btnode[K, V]) min() *btnode[K, V] {
	if n == nil {
		return nil
	}
	tmp := n

	if min := n.left.min(); min != nil {
		if n.tree.ord(tmp.key, min.key) > 0 {
			tmp = min
		}
	}

	if min := n.right.min(); min != nil {
		if n.tree.ord(tmp.key, min.key) > 0 {
			tmp = min
		}
	}

	return tmp
}

func (n *btnode[K, V]) max() *btnode[K, V] {
	if n == nil {
		return nil
	}
	tmp := n

	if max := n.left.max(); max != nil {
		if n.tree.ord(tmp.key, max.key) < 0 {
			tmp = max
		}
	}

	if max := n.right.max(); max != nil {
		if n.tree.ord(tmp.key, max.key) < 0 {
			tmp = max
		}
	}

	return tmp
}

func (parent *btnode[K, V]) link(n *btnode[K, V]) {
	if parent.left == n {
		if n.left != nil {
			parent.left = n.left
		} else {
			parent.left = n.right
		}
	} else if parent.right == n {
		if n.left != nil {
			parent.right = n.left
		} else {
			parent.right = n.right
		}
	}
	n.tree = nil
}

func (n *btnode[K, V]) del(parent *btnode[K, V], k K) bool {
	if n == nil {
		return false
	}

	o := n.tree.ord(n.key, k)
	switch o {
	case 0:
		if m := n.right.min(); n.left != nil && m != nil {
			n.key = m.key
			return n.right.del(n, n.key)
		}
		parent.link(n)
		return true
	case 1:
		if n.left == nil {
			return false
		}
		return n.left.del(n, k)
	case -1:
		if n.right == nil {
			return false
		}
		return n.right.del(n, k)
	}
	return false
}
