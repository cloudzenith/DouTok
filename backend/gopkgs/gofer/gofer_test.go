package gofer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestSetUseGlobalPool(t *testing.T) {
	assert.Equal(t, false, useGlobalPool)

	SetUseGlobalPool(true)
	assert.Equal(t, true, useGlobalPool)
}

func submitGo() bool {
	var wg sync.WaitGroup
	wg.Add(1)
	var flag bool
	Go(func() {
		flag = true
		wg.Done()
	})

	wg.Wait()
	return flag
}

func TestGo(t *testing.T) {
	assert.Equal(t, true, submitGo())

	InitGlobalPool()
	SetUseGlobalPool(true)
	assert.Equal(t, true, submitGo())
}

func submitGoWithCtx() bool {
	var wg sync.WaitGroup
	wg.Add(1)

	var flag bool
	ctx := context.Background()
	GoWithCtx(ctx, func(ctx context.Context) {
		flag = true
		wg.Done()
	})

	wg.Wait()
	return flag
}

func TestGoWithCtx(t *testing.T) {
	assert.Equal(t, true, submitGoWithCtx())

	InitGlobalPool()
	SetUseGlobalPool(true)
	assert.Equal(t, true, submitGoWithCtx())
}
