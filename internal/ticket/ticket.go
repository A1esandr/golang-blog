package main

import (
	"fmt"
	"sync"
)

type (
	Worker struct {
		tickets chan struct{}
	}

	Count struct {
		m     sync.RWMutex
		count int
	}
)

// Comment
func New(poolSize int) *Worker {
	pool := make(chan struct{}, poolSize)
	for i := 0; i < poolSize; i++ {
		pool <- struct{}{}
	}
	return &Worker{tickets: pool}
}

func (w *Worker) Do(count *Count) {
	defer func() {
		w.tickets <- struct{}{}
		count.Dec()
	}()
	fmt.Printf("Count %v \n", count.Get())
	var a int
	for i := 0; i < 10000000; i++ {
		a = i*2 - a
	}
	fmt.Printf("Do something %v \n", a)
}

func (w *Worker) Ticket() chan struct{} {
	return w.tickets
}

func (c *Count) Dec() {
	c.m.Lock()
	c.count--
	c.m.Unlock()
}

func (c *Count) Inc() {
	c.m.Lock()
	c.count++
	c.m.Unlock()
}

func (c *Count) Get() int {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.count
}

func main() {
	count := &Count{}
	worker := New(10)
	for i := 0; i < 100; i++ {
		select {
		case <-worker.Ticket():
			go func() {
				count.Inc()
				worker.Do(count)
			}()
		}
	}

	countS := &Count{}
	for i := 0; i < 100; i++ {
		<-worker.Ticket()
		go func() {
			countS.Inc()
			worker.Do(countS)
		}()
	}
}
