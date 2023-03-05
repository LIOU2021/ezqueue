package queue

import (
	"time"
)

func Consumer() {
	for q := range h {
		go func(queue string) {
			for i := 0; i < h[queue].consumerNumber; i++ {
				go func() {
					for job := range h[queue].msg {
						job()
						time.Sleep(h[queue].interval)
					}
				}()
			}
		}(q)

	}
}
