package binary

import (
	"constraints"

	"github.com/tdakkota/algo2/alg"
)

// RedBlack represents the red black search tree.
type RedBlack[K, V any] struct {
	root *rbnode[K, V]
	ord  alg.Ord[K]
	size int
}

func NewRedBlack[K constraints.Ordered, V any]() *RedBlack[K, V] {
	return NewRedBlackWithOrd[K, V](alg.BuiltinOrd[K])
}

func NewRedBlackWithOrd[K, V any](ord alg.Ord[K]) *RedBlack[K, V] {
	b := new(RedBlack[K, V])
	b.ord = ord
	return b
}

// Min returns the minimal element.
func (tree *RedBlack[K, V]) Min() (v alg.Pair[K, V], ok bool) {
	if min := tree.root.min(); min != nil {
		return alg.Two(min.key, min.value), true
	}
	return
}

// Max returns the maximal element.
func (tree *RedBlack[K, V]) Max() (v alg.Pair[K, V], ok bool) {
	if max := tree.root.max(); max != nil {
		return alg.Two(max.key, max.value), true
	}
	return
}

// Len returns size of tree.
func (tree *RedBlack[K, V]) Len() int {
	return tree.size
}

func (tree *RedBlack[K, V]) Clear() {
	tree.root = nil
}

func (tree *RedBlack[K, V]) Put(key K, value V) {
	var insertedNode *rbnode[K, V]
	if tree.root == nil {
		tree.root = &rbnode[K, V]{key: key, value: value, color: red, tree: tree}
		insertedNode = tree.root
	} else {
		node := tree.root
		loop := true
		for loop {
			switch tree.ord(key, node.key) {
			case 0:
				node.key = key
				node.value = value
				return
			case -1:
				if node.left == nil {
					node.left = &rbnode[K, V]{key: key, value: value, color: red, tree: tree}
					insertedNode = node.left
					loop = false
				} else {
					node = node.left
				}
			case 1:
				if node.right == nil {
					node.right = &rbnode[K, V]{key: key, value: value, color: red, tree: tree}
					insertedNode = node.right
					loop = false
				} else {
					node = node.right
				}
			}
		}
		insertedNode.parent = node
	}
	tree.insertCase1(insertedNode)
	tree.size++
}

func (tree *RedBlack[K, V]) Get(key K) (value V, found bool) {
	if node := tree.lookup(key); node != nil {
		return node.value, true
	}
	return
}

func (tree *RedBlack[K, V]) Delete(key K) (ok bool) {
	var child *rbnode[K, V]
	node := tree.lookup(key)
	if node == nil {
		return
	}
	if node.left != nil && node.right != nil {
		pred := node.left.maximumNode()
		node.key = pred.key
		node.value = pred.value
		node = pred
	}
	if node.left == nil || node.right == nil {
		if node.right == nil {
			child = node.left
		} else {
			child = node.right
		}
		if node.color == black {
			node.color = nodeColor(child)
			tree.deleteCase1(node)
		}
		tree.replaceNode(node, child)
		if node.parent == nil && child != nil {
			child.color = black
		}
	}
	node.tree = nil
	tree.size--
	return true
}

func (tree *RedBlack[K, V]) lookup(key K) *rbnode[K, V] {
	node := tree.root
	for node != nil {
		o := tree.ord(key, node.key)
		switch o {
		case 0:
			return node
		case -1:
			node = node.left
		case 1:
			node = node.right
		}
	}
	return nil
}

func (tree *RedBlack[K, V]) rotateLeft(node *rbnode[K, V]) {
	right := node.right
	tree.replaceNode(node, right)
	node.right = right.left
	if right.left != nil {
		right.left.parent = node
	}
	right.left = node
	node.parent = right
}

func (tree *RedBlack[K, V]) rotateRight(node *rbnode[K, V]) {
	left := node.left
	tree.replaceNode(node, left)
	node.left = left.right
	if left.right != nil {
		left.right.parent = node
	}
	left.right = node
	node.parent = left
}

func (tree *RedBlack[K, V]) replaceNode(old, new *rbnode[K, V]) {
	if old.parent == nil {
		tree.root = new
	} else {
		if old == old.parent.left {
			old.parent.left = new
		} else {
			old.parent.right = new
		}
	}
	if new != nil {
		new.parent = old.parent
	}
}

func (tree *RedBlack[K, V]) insertCase1(node *rbnode[K, V]) {
	if node.parent == nil {
		node.color = black
	} else {
		tree.insertCase2(node)
	}
}

func (tree *RedBlack[K, V]) insertCase2(node *rbnode[K, V]) {
	if nodeColor(node.parent) == black {
		return
	}
	tree.insertCase3(node)
}

func (tree *RedBlack[K, V]) insertCase3(node *rbnode[K, V]) {
	uncle := node.uncle()
	if nodeColor(uncle) == red {
		node.parent.color = black
		uncle.color = black
		node.grandparent().color = red
		tree.insertCase1(node.grandparent())
	} else {
		tree.insertCase4(node)
	}
}

func (tree *RedBlack[K, V]) insertCase4(node *rbnode[K, V]) {
	grandparent := node.grandparent()
	if node == node.parent.right && node.parent == grandparent.left {
		tree.rotateLeft(node.parent)
		node = node.left
	} else if node == node.parent.left && node.parent == grandparent.right {
		tree.rotateRight(node.parent)
		node = node.right
	}
	tree.insertCase5(node)
}

func (tree *RedBlack[K, V]) insertCase5(node *rbnode[K, V]) {
	node.parent.color = black
	grandparent := node.grandparent()
	grandparent.color = red
	if node == node.parent.left && node.parent == grandparent.left {
		tree.rotateRight(grandparent)
	} else if node == node.parent.right && node.parent == grandparent.right {
		tree.rotateLeft(grandparent)
	}
}

func (tree *RedBlack[K, V]) deleteCase1(node *rbnode[K, V]) {
	if node.parent == nil {
		return
	}
	tree.deleteCase2(node)
}

func (tree *RedBlack[K, V]) deleteCase2(node *rbnode[K, V]) {
	sibling := node.sibling()
	if nodeColor(sibling) == red {
		node.parent.color = red
		sibling.color = black
		if node == node.parent.left {
			tree.rotateLeft(node.parent)
		} else {
			tree.rotateRight(node.parent)
		}
	}
	tree.deleteCase3(node)
}

func (tree *RedBlack[K, V]) deleteCase3(node *rbnode[K, V]) {
	sibling := node.sibling()
	if nodeColor(node.parent) == black &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.left) == black &&
		nodeColor(sibling.right) == black {
		sibling.color = red
		tree.deleteCase1(node.parent)
	} else {
		tree.deleteCase4(node)
	}
}

func (tree *RedBlack[K, V]) deleteCase4(node *rbnode[K, V]) {
	sibling := node.sibling()
	if nodeColor(node.parent) == red &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.left) == black &&
		nodeColor(sibling.right) == black {
		sibling.color = red
		node.parent.color = black
	} else {
		tree.deleteCase5(node)
	}
}

func (tree *RedBlack[K, V]) deleteCase5(node *rbnode[K, V]) {
	sibling := node.sibling()
	if node == node.parent.left &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.left) == red &&
		nodeColor(sibling.right) == black {
		sibling.color = red
		sibling.left.color = black
		tree.rotateRight(sibling)
	} else if node == node.parent.right &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.right) == red &&
		nodeColor(sibling.left) == black {
		sibling.color = red
		sibling.right.color = black
		tree.rotateLeft(sibling)
	}
	tree.deleteCase6(node)
}

func (tree *RedBlack[K, V]) deleteCase6(node *rbnode[K, V]) {
	sibling := node.sibling()
	sibling.color = nodeColor(node.parent)
	node.parent.color = black
	if node == node.parent.left && nodeColor(sibling.right) == red {
		sibling.right.color = black
		tree.rotateLeft(node.parent)
	} else if nodeColor(sibling.left) == red {
		sibling.left.color = black
		tree.rotateRight(node.parent)
	}
}

// Iterate traverse tree in order using given callback.
func (tree *RedBlack[K, V]) iterate(o order, cb func(alg.Pair[K, V]) bool) {
	if tree.root != nil {
		tree.root.traverse(o, cb)
	}
}

// Iterate traverse tree in order using given callback.
func (tree *RedBlack[K, V]) Iterate(cb func(alg.Pair[K, V]) bool) {
	tree.IterateInOrder(cb)
}

// IterateInOrder traverse tree in order using given callback.
func (tree *RedBlack[K, V]) IterateInOrder(cb func(alg.Pair[K, V]) bool) {
	tree.iterate(inOrder, cb)
}

// IteratePreOrder traverse tree in pre-order using given callback.
func (tree *RedBlack[K, V]) IteratePreOrder(cb func(alg.Pair[K, V]) bool) {
	tree.iterate(preOrder, cb)
}

// IteratePostOrder traverse tree in post-order using given callback.
func (tree *RedBlack[K, V]) IteratePostOrder(cb func(alg.Pair[K, V]) bool) {
	tree.iterate(postOrder, cb)
}
