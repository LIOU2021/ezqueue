package main

import (
	"log"
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

	select {
	case <-time.After(10 * time.Second):
		// queue.Close(queuePM, queueRD)
		queue.CloseAll()
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
		})

		if err != nil {
			log.Fatal(err)
		}
	}
}
