package queue

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var h = make(map[string]*BasicQueue)

type BasicQueue struct {
	name           string
	consumerNumber int
	interval       time.Duration
	msg            chan func()
	msgLen         atomic.Int32
}

func (qu *BasicQueue) GetName() string {
	return qu.name
}

func (qu *BasicQueue) Close() {
	close(qu.msg)
	fmt.Printf("queue %s close !\n", qu.name)
}

func (qu *BasicQueue) SetConsumerNumber(number int) {
	if number < 1 {
		number = 1
	}
	qu.consumerNumber = number
}

func (qu *BasicQueue) GetConsumerNumber() int {
	return qu.consumerNumber
}

func (qu *BasicQueue) SetInterval(time time.Duration) {
	if time < 0 {
		time = 0
	}
	qu.interval = time
}

func (qu *BasicQueue) GetInterval() time.Duration {
	return qu.interval
}

func NewQueue(name string) *BasicQueue {
	return &BasicQueue{
		name:           name,
		consumerNumber: 1,
		interval:       time.Second * 0,
		msg:            make(chan func()),
	}
}

// add queue by name
func Add(queue ...*BasicQueue) {
	for _, v := range queue {
		h[v.name] = v
		fmt.Printf("queue %s create ! interval: %v, consumerNumber: %d \n", v.name, v.interval, v.consumerNumber)
	}
}

func CloseAll() {
	var wg sync.WaitGroup
	wg.Add(len(h))
	for _, queue := range h {
		go func(queue *BasicQueue) {
			for {
				fmt.Printf("close after waiting for %d tasks for queue \"%s\"\n", queue.msgLen.Load(), queue.name)
				if queue.msgLen.Load() == int32(0) {
					break
				}
				time.Sleep(1 * time.Second)
			}
			queue.Close()
			wg.Done()
		}(queue)
	}
	wg.Wait()
}
