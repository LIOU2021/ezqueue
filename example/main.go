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
	queue.Add(queuePM, queueRD)

	// create consumer
	queue.Consumer()

	// create task
	go task1()
	go task2()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			log.Println("before runtime.NumGoroutine(): ", runtime.NumGoroutine())
		}
	}()

	select {
	case <-time.After(10 * time.Second):
		// queue.Close(queuePM, queueRD)
		log.Println("After runtime.NumGoroutine(): ", runtime.NumGoroutine())
		queue.CloseAll()
		time.Sleep(1 * time.Second)
		log.Println("timeout runtime.NumGoroutine(): ", runtime.NumGoroutine())
		log.Println("finish ...")
		break
	}
}

func task1() {
	for i := 0; i < counter; i++ {
		v := i
		log.Println("pm send task: ", v)
		err := queue.Task(queuePM, func() {
			log.Println("pm receive task: ", v)
			time.Sleep(2 * time.Second)
		})

		if err != nil {
			log.Fatal(err)
		}
	}
}

func task2() {
	for i := 0; i < counter; i++ {
		v := i
		log.Println("rd send task: ", v)
		err := queue.Task(queueRD, func() {
			log.Println("rd receive task: ", v)
			time.Sleep(2 * time.Second)
		})

		if err != nil {
			log.Fatal(err)
		}
	}
}
