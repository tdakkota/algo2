package stream

func (s Stream[T]) Peek(fn func(T)) Stream[T] {
	return Pipe[T, T](s, func(t T) (T, bool) {
		fn(t)
		return t, true
	})
}

func (s Stream[T]) PeekPtr(fn func(*T)) Stream[T] {
	return Pipe[T, T](s, func(t T) (T, bool) {
		fn(&t)
		return t, true
	})
}

func (s Stream[T]) Filter(fn func(T) bool) Stream[T] {
	return Pipe[T, T](s, func(t T) (T, bool) {
		return t, fn(t)
	})
}

func (s Stream[T]) TakeWhile(fn func(T) bool) Stream[T] {
	r := true
	return Pipe[T, T](s, func(t T) (T, bool) {
		if r && !fn(t) {
			r = false
		}
		return t, r
	})
}

func (s Stream[T]) DropWhile(fn func(T) bool) Stream[T] {
	r := false
	return Pipe[T, T](s, func(t T) (T, bool) {
		if !r && !fn(t) {
			r = true
		}
		return t, r
	})
}

func (s Stream[T]) Skip(i int) Stream[T] {
	return s.DropWhile(func(T) bool {
		i--
		return i >= 0
	})
}

func (s Stream[T]) Limit(i int) Stream[T] {
	return s.TakeWhile(func(T) bool {
		i--
		return i >= 0
	})
}

func (s Stream[T]) Count() (i int) {
	s.Iterate(func(T) bool {
		i++
		return true
	})
	return
}

func (s Stream[T]) First() (r T, ok bool) {
	s.Iterate(func(t T) bool {
		r, ok = t, true
		return false
	})
	return
}

func (s Stream[T]) Any(fn func(T) bool) (r bool) {
	_, r = s.Filter(fn).First()
	return
}

func (s Stream[T]) All(fn func(T) bool) (r bool) {
	r = true
	s.Iterate(func(t T) bool {
		if !fn(t) {
			r = false
			return false
		}
		return true
	})
	return
}
