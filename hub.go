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
	msg            chan func()
	msgLen         atomic.Int32
	msgIsClose     atomic.Bool
}

func (qu *BasicQueue) GetName() string {
	return qu.name
}

func (qu *BasicQueue) Close() {
	for qu.msgLen.Load() > 0 {
		fmt.Printf("close after waiting for %d tasks for queue \"%s\"\n", qu.msgLen.Load(), qu.name)
		time.Sleep(1 * time.Second)
	}
	close(qu.msg)
	qu.msgIsClose.Store(true)
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

func NewQueue(name string) *BasicQueue {
	return &BasicQueue{
		name:           name,
		consumerNumber: 1,
		msg:            make(chan func()),
	}
}

// add queue by name
func Add(queue ...*BasicQueue) {
	for _, v := range queue {
		h[v.name] = v
		fmt.Printf("queue %s create ! consumerNumber: %d \n", v.name, v.consumerNumber)
	}
}

func CloseAll() {
	var wg sync.WaitGroup
	wg.Add(len(h))
	for _, queue := range h {
		go func(queue *BasicQueue) {
			queue.Close()
			wg.Done()
		}(queue)
	}
	wg.Wait()
}
