package queue

import (
	"log"
	"time"
)

const (
	interval = 200 //Millisecond
)

func Consumer() {
	for i, _ := range h {
		go func(queueName string) {
			for {
				select {
				case msg, ok := <-h[queueName]:
					if len(h[queueName]) == 0 && !ok {
						log.Println(queueName, " done------------------")
						return
					} else {
						msg()
					}
				}

				time.Sleep(interval * time.Millisecond)
			}
		}(i)
	}
}
