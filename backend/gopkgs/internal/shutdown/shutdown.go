package shutdown

import (
	"github.com/cloudzenith/DouTok/backend/gopkgs/gofer"
	"github.com/samber/lo"
	"os"
	"os/signal"
	"sort"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

var globalShutdownManager = newManager()

func init() {
	go func() {
		signals := make(chan os.Signal, 1)

		signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

		sig := <-signals
		globalShutdownManager.Fire(sig)
		signal.Stop(signals)
	}()
}

type orderHandler struct {
	fns []func()
}

type manager struct {
	mu         sync.RWMutex
	fired      atomic.Bool
	firedCh    chan struct{}
	doneCh     chan struct{}
	lastSignal atomic.Value
	orderMap   map[int]*orderHandler
}

func newManager() *manager {
	return &manager{
		firedCh:  make(chan struct{}),
		doneCh:   make(chan struct{}),
		orderMap: make(map[int]*orderHandler),
	}
}

func (s *manager) Fire(sig os.Signal) {
	if !s.fired.CompareAndSwap(false, true) {
		return
	}

	close(s.firedCh)
	s.lastSignal.Store(sig)

	s.mu.RLock()
	defer s.mu.RUnlock()

	keys := lo.Keys(s.orderMap)
	sort.Ints(keys)

	for _, key := range keys {
		s.fireOrder(s.orderMap[key])
	}

	close(s.doneCh)
}

func (s *manager) fireOrder(order *orderHandler) {
	wg := sync.WaitGroup{}

	for _, fn := range order.fns {
		wg.Add(1)
		tmp := fn
		gofer.Go(func() {
			defer wg.Done()
			tmp()
		})
	}

	wg.Wait()
}

func (s *manager) AddOrderHandler(order int, fn func()) {
	if s.fired.Load() {
		gofer.Go(fn)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	orderHandle, ok := s.orderMap[order]
	if !ok {
		orderHandle = &orderHandler{}
		s.orderMap[order] = orderHandle
	}

	orderHandle.fns = append(orderHandle.fns, fn)
}

func (s *manager) FiredCh() <-chan struct{} {
	return s.firedCh
}

func (s *manager) LastSignal() os.Signal {
	r, ok := s.lastSignal.Load().(os.Signal)
	if !ok {
		return nil
	}

	return r
}

func (s *manager) Wait(duration time.Duration) {
	if duration <= 0 {
		<-s.doneCh
		return
	}

	_ = gofer.GoWithTimeout(func() {
		<-s.doneCh
	}, duration)
}

func FrameworkShutdownHandler(fn func()) {
	globalShutdownManager.AddOrderHandler(0, fn)
}

func SDKShutdownHandler(fn func()) {
	globalShutdownManager.AddOrderHandler(10, fn)
}

func BizShutdownHandler(fn func()) {
	globalShutdownManager.AddOrderHandler(20, fn)
}

func Fire(sig os.Signal) {
	globalShutdownManager.Fire(sig)
}

func Wait(max time.Duration) {
	globalShutdownManager.Wait(max)
}

func FiredCh() <-chan struct{} {
	return globalShutdownManager.FiredCh()
}

func LastSignal() os.Signal {
	return globalShutdownManager.LastSignal()
}
