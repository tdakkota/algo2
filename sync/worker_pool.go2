package sync

import (
	"context"
)

type GoroutinePool interface {
	Do(func(context.Context))
}

type LimitedGoroutinePool struct {
	parent context.Context
	queue  chan struct{}
}

func NewLimitedPool(c context.Context, limit int) LimitedGoroutinePool {
	return LimitedGoroutinePool{
		queue:  make(chan struct{}, limit),
		parent: c,
	}
}

func (p LimitedGoroutinePool) Do(cb func(context.Context)) {
	p.queue <- struct{}{}
	go func() {
		select {
		case <-p.parent.Done():
			return
		default:
			cb(p.parent)
			<-p.queue
		}
	}()
}
