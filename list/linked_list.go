// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

// Element is an element of a linked list.
type Element[TElem any] struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *Element[TElem]

	// The list to which this element belongs.
	list *LinkedList[TElem]

	// The value stored with this element.
	Value TElem
}

// Next returns the next list element or nil.
func (e *Element[TElem]) Next() *Element[TElem] {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Prev returns the previous list element or nil.
func (e *Element[TElem]) Prev() *Element[TElem] {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func (e *Element[TElem]) List() *LinkedList[TElem] {
	return e.list
}

// LinkedList represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type LinkedList[TElem any] struct {
	root Element[TElem] // sentinel list element, only &root, root.prev, and root.next are used
	len  int            // current list length excluding (this) sentinel element
}

// Init initializes or clears list l.
func (l *LinkedList[TElem]) Init() *LinkedList[TElem] {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// NewLinkedList returns an initialized list.
func NewLinkedList[TElem any]() *LinkedList[TElem] { return new(LinkedList[TElem]).Init() }

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *LinkedList[TElem]) Len() int { return l.len }

// Front returns the first element of list l or nil if the list is empty.
func (l *LinkedList[TElem]) Front() *Element[TElem] {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back returns the last element of list l or nil if the list is empty.
func (l *LinkedList[TElem]) Back() *Element[TElem] {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// lazyInit lazily initializes a zero List value.
func (l *LinkedList[TElem]) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// insert inserts e after at, increments l.len, and returns e.
func (l *LinkedList[TElem]) insert(e, at *Element[TElem]) *Element[TElem] {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element[TElem]{Value: v}, at).
func (l *LinkedList[TElem]) insertValue(v TElem, at *Element[TElem]) *Element[TElem] {
	e := &Element[TElem]{}
	e.Value = v
	return l.insert(e, at)
}

// remove removes e from its list, decrements l.len, and returns e.
func (l *LinkedList[TElem]) remove(e *Element[TElem]) TElem {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--

	value := e.Value
	return value
}

// move moves e to next to at and returns e.
func (l *LinkedList[TElem]) move(e, at *Element[TElem]) *Element[TElem] {
	if e == at {
		return e
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e

	return e
}

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil.
func (l *LinkedList[TElem]) Remove(e *Element[TElem]) TElem {
	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		return l.remove(e)
	}
	return e.Value
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *LinkedList[TElem]) PushFront(v TElem) *Element[TElem] {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *LinkedList[TElem]) PushBack(v TElem) *Element[TElem] {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *LinkedList[TElem]) InsertBefore(v TElem, mark *Element[TElem]) *Element[TElem] {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark.prev)
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *LinkedList[TElem]) InsertAfter(v TElem, mark *Element[TElem]) *Element[TElem] {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark)
}

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *LinkedList[TElem]) MoveToFront(e *Element[TElem]) {
	if e.list != l || l.root.next == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, &l.root)
}

// MoveToBack moves element e to the back of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *LinkedList[TElem]) MoveToBack(e *Element[TElem]) {
	if e.list != l || l.root.prev == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, l.root.prev)
}

// MoveBefore moves element e to its new position before mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *LinkedList[TElem]) MoveBefore(e, mark *Element[TElem]) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark.prev)
}

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *LinkedList[TElem]) MoveAfter(e, mark *Element[TElem]) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark)
}

func (l LinkedList[TElem]) Iterate(fn func(TElem) bool) {
	for e := l.Front(); e != nil; e = e.Next() {
		if !fn(e.Value) {
			return
		}
	}
}

func (l LinkedList[TElem]) IterateElements(fn func(*Element[TElem]) bool) {
	for e := l.Front(); e != nil; e = e.Next() {
		if !fn(e) {
			return
		}
	}
}

// PushBackList inserts a copy of an other list at the back of list l.
// The lists l and other may be the same. They must not be nil.
func (l *LinkedList[TElem]) PushBackList(other *LinkedList[TElem]) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

// PushFrontList inserts a copy of an other list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (l *LinkedList[TElem]) PushFrontList(other *LinkedList[TElem]) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}

// Same as PushFront.
func (l *LinkedList[TElem]) Prepend(e ...TElem) {
	for _, v := range e {
		l.PushFront(v)
	}
}

// Same as PushBack.
func (l *LinkedList[TElem]) Append(e ...TElem) {
	for _, v := range e {
		l.PushBack(v)
	}
}

// Deletes and returns first element.
func (l *LinkedList[TElem]) PopFront() (v TElem, ok bool) {
	if e := l.Front(); e != nil {
		return l.Remove(e), true
	}

	return
}

// Deletes and returns last element.
func (l *LinkedList[TElem]) Pop() (v TElem, ok bool) {
	if e := l.Back(); e != nil {
		return l.Remove(e), true
	}

	return
}
