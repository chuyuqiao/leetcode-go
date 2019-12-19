package sync

import (
	"sync"
	"sync/atomic"
)

type Counter interface {
	Inc()
	Load() int64
}

type AtomicCounter struct {
	counter int64
}

func (ac *AtomicCounter) Inc() {
	atomic.AddInt64(&ac.counter, 1)
}

func (ac *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&ac.counter)
}

type MutexCounter struct {
	counter int64
	m       sync.Mutex
}

func (mc *MutexCounter) Inc() {
	mc.m.Lock()
	defer mc.m.Unlock()
	mc.counter++
}

func (mc *MutexCounter) Load() int64 {
	mc.m.Lock()
	defer mc.m.Unlock()
	return mc.counter
}
