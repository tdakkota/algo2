package list

import "github.com/tdakkota/algo2/alg"

type SliceList[T any] []T

func (l SliceList[T]) Len() int { return len(l) }

// Adds element to the start of slice.
func (l *SliceList[T]) Prepend(v ...T) {
	*l = append(v, (*l)...)
}

// Adds element to the end of slice.
func (l *SliceList[T]) Append(v ...T) {
	*l = append((*l), v...)
}

// Deletes and returns first element.
func (l *SliceList[T]) PopFront() (v T, ok bool) {
	if len(*l) > 0 {
		v, *l = (*l)[0], (*l)[1:]
		return
	}

	return
}

// Deletes and returns last element.
func (l *SliceList[T]) Pop() (v T, ok bool) {
	if len(*l) > 0 {
		v, *l = (*l)[len(*l)-1], (*l)[:len(*l)-1]
		return
	}

	return
}

func (l *SliceList[T]) Remove(i int) (v T) {
	length := len(*l)
	if i == length-1 {
		v, _ = l.Pop()
		return v
	}

	a := *l
	v = a[i]
	// Remove the element at index i from a.
	copy(a[i:], a[i+1:])        // Shift a[i+1:] left one index.
	a[length-1] = alg.Zero[T]() // Erase last element (write zero value).
	a = a[:length-1]            // Truncate slice.
	*l = a
	return v
}

func (l SliceList[T]) Iterate(fn func(alg.Pair[int, T]) bool) {
	for i, e := range l {
		if !fn(alg.Two(i, e)) {
			return
		}
	}
}
