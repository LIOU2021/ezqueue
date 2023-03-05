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
	counter = 10
)

func main() {
	// create two queue
	q1 := queue.NewQueue(queuePM)
	q2 := queue.NewQueue(queueRD)

	// q1.SetConsumerNumber(2) // option. default 1
	q2.SetConsumerNumber(2) // option. default 1

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
	case <-time.After(5 * time.Second):
		log.Println("After runtime.NumGoroutine(): ", runtime.NumGoroutine())
		// q1.Close()
		// q2.Close()
		log.Println("execute CloseAll")
		queue.CloseAll()
		time.Sleep(1 * time.Second)
		log.Println("dispatch again")
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
			log.Printf("%s process task: %d", queue.GetName(), v)
			time.Sleep(2 * time.Second)
			log.Printf("%s finish task: %d", queue.GetName(), v)
		})
	}
}
