package main

import (
	"log"
	"runtime"
	"time"

	queue "github.com/LIOU2021/ezqueue"
)

const (
	queuePM = "pm"
	queueRD = "rd"
	counter = 5
)

func main() {
	// create two queue
	q1 := queue.NewQueue(queuePM)
	q2 := queue.NewQueue(queueRD)

	q1.SetConsumerNumber(2)         // option. default 1
	q1.SetInterval(time.Second * 3) // option. default 0
	q2.SetConsumerNumber(2)         // option. default 1
	q2.SetInterval(time.Second * 3) // option. default 0

	// add queue
	queue.Add(q1, q2)

	// create consumer
	queue.Consumer()

	go dispatch(q1)
	go dispatch(q2)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			log.Println("before runtime.NumGoroutine(): ", runtime.NumGoroutine())
		}
	}()

	select {
	case <-time.After(10 * time.Second):
		log.Println("After runtime.NumGoroutine(): ", runtime.NumGoroutine())
		// q1.Close()
		// q2.Close()
		queue.CloseAll()
		time.Sleep(1 * time.Second)

		go dispatch(q1)
		go dispatch(q2)
		time.Sleep(5 * time.Second)
		log.Println("timeout runtime.NumGoroutine(): ", runtime.NumGoroutine())
		log.Println("finish ...")
		break
	}
}

func dispatch(queue *queue.BasicQueue) {
	for i := 0; i < counter; i++ {
		v := i
		log.Printf("%s send task: %d\n", queue.GetName(), v)
		queue.Task(func() {
			log.Printf("%s receive task: %d", queue.GetName(), v)
			time.Sleep(2 * time.Second)
		})
	}
}
