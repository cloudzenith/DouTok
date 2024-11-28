package gofer

import (
	"context"
	"sync"
	"time"

	"github.com/TremblingV5/box/rearer"
)

var useGlobalPool bool
var setUseGlobalPoolOnce sync.Once

func SetUseGlobalPool(value bool) {
	setUseGlobalPoolOnce.Do(func() {
		useGlobalPool = value
	})
}

func Go(f func()) {
	if useGlobalPool {
		if pool == nil {
			InitGlobalPool()
		}

		_ = pool.Submit(f)
		return
	}

	go func() {
		defer rearer.Recover()
		f()
	}()
}

func GoWithCtx(ctx context.Context, f func(context.Context)) {
	if useGlobalPool {
		_ = pool.Submit(func() {
			f(ctx)
		})
		return
	}

	go func() {
		defer rearer.RecoverWithCtx(ctx)
		f(ctx)
	}()
}

func GoWithTimeout(f func(), d time.Duration) (isFinish bool) {
	ch := make(chan struct{})

	Go(func() {
		f()
		close(ch)
	})

	timer := time.NewTimer(d)
	select {
	case <-timer.C:
		return false
	case <-ch:
		timer.Stop()
		return true
	}
}
