package gofer

import (
	"context"
	"errors"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/TremblingV5/box/rearer"
)

type Group struct {
	wg sync.WaitGroup

	isErrorGroup bool

	errOnce sync.Once
	err     error

	numG   int   // number of usable G
	gCount int64 // number of running G

	queueSize int               // size of wait queue
	queue     chan func() error // wait queue

	cancel    func()
	ctx       context.Context
	waitFired atomic.Bool
}

func NewGroup(ctx context.Context, options ...GroupOption) *Group {
	g := &Group{
		wg: sync.WaitGroup{},
	}

	for _, option := range options {
		option(g)
	}

	if g.isErrorGroup {
		if g.numG == 0 {
			g.numG = runtime.NumCPU()
		}

		if g.queueSize == 0 {
			g.queueSize = g.numG
		}

		ctx, cancel := context.WithCancel(ctx)
		g.ctx = ctx
		g.cancel = cancel
		g.queue = make(chan func() error, g.queueSize)
		return g
	}

	return g

}

func (g *Group) Run(f func() error) error {
	if f == nil {
		return errors.New("submitted task can't be nil")
	}

	if g.isErrorGroup && g.waitFired.Load() {
		return errors.New("can't submit task to a error group after Wait()")
	}

	if g.isErrorGroup {
		select {
		case <-g.ctx.Done():
			return nil
		case g.queue <- f:
			g.checkAndStartG()
		}

		return nil
	}

	g.wg.Add(1)

	Go(func() {
		defer g.wg.Done()
		f()
	})

	return nil
}

func (g *Group) checkAndStartG() {
	if atomic.LoadInt64(&g.gCount) >= int64(g.numG) {
		return
	}

	g.startG()
}

func (g *Group) startG() {
	g.wg.Add(1)
	atomic.AddInt64(&g.gCount, 1)
	go func() {
		defer func() {
			panicError := recover()
			if panicError != nil {
				rearer.LogRecoverStack(panicError, rearer.WithCtx(g.ctx))

				var err error
				switch v := panicError.(type) {
				case error:
					err = v
				case string:
					err = errors.New(v)
				default:
					err = errors.New("something panic")
				}

				g.setError(err)
			}
		}()
		defer g.wg.Done()
		defer atomic.AddInt64(&g.gCount, -1)

		for {
			select {
			case <-g.ctx.Done():
				g.setError(errors.New("context canceled"))
				return
			case f, ok := <-g.queue:
				if !ok {
					// channel is closed
					return
				}

				if err := f(); err != nil {
					g.setError(err)
				}
			}
		}
	}()
}

func (g *Group) setError(err error) {
	g.errOnce.Do(func() {
		g.err = err
		g.cancel()
	})
}

func (g *Group) Wait() error {
	if g.isErrorGroup {
		if g.waitFired.CompareAndSwap(false, true) {
			close(g.queue)
		}
	}

	g.wg.Wait()

	if g.isErrorGroup {
		g.cancel()
	}

	return g.err
}
