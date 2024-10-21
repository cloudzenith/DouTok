package gofer

import (
	"github.com/TremblingV5/box/rearer"
	"github.com/panjf2000/ants/v2"
)

var (
	pool         *Pool
	panicHandler = ants.WithPanicHandler(func(err any) {
		rearer.LogRecoverStack(err)
	})
)

var defaultPoolSize = 1000

func SetPoolSize(size int) {
	defaultPoolSize = size
}

type Pool struct {
	poolWithFunc *ants.PoolWithFunc
	pool         *ants.Pool
}

func InitGlobalPool() {
	p, _ := ants.NewPool(defaultPoolSize, panicHandler)
	pool = &Pool{
		pool: p,
	}
}

func NewWithPoolFunc(poolSize int, f func(a any), options ...ants.Option) (*Pool, error) {
	options = append(options, panicHandler)

	pf, err := ants.NewPoolWithFunc(poolSize, f, options...)
	if err != nil {
		return nil, err
	}

	return &Pool{
		poolWithFunc: pf,
	}, nil
}

func (p *Pool) Release() {
	if p.pool != nil {
		p.Release()
	}

	if p.poolWithFunc != nil {
		p.poolWithFunc.Release()
	}
}

func (p *Pool) Invoke(args interface{}) error {
	return p.poolWithFunc.Invoke(args)
}

func (p *Pool) Running() int {
	return p.pool.Running()
}

func (p *Pool) RunningPF() int {
	return p.poolWithFunc.Running()
}

func (p *Pool) Submit(task func()) error {
	if p.pool == nil {
		pool, _ := ants.NewPool(defaultPoolSize, panicHandler)
		p.pool = pool
	}

	if err := p.pool.Submit(task); err != nil {
		// TODO: log errors here
		return err
	}

	return nil
}
