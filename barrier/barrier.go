package main

import "sync"

type Barrier struct {
	total int
	count int
	mutex *sync.Mutex
	cond  *sync.Cond
}

func (b *Barrier) Wait() {
	b.mutex.Lock()
	b.count = b.count - 1
	if b.count > 0 {
		b.cond.Wait()
	} else {
		b.count = b.total
		b.cond.Broadcast()
	}
	b.mutex.Unlock()
}

func NewBarrier(total int) *Barrier {
	mutex := sync.Mutex{}
	return &Barrier{
		total: total,
		count: total,
		mutex: &mutex,
		cond:  sync.NewCond(&mutex),
	}
}
