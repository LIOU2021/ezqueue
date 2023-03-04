package ezqueue

import (
	"log"
	"time"
)

func Consumer() {
	go func() {
		for {
			select {
			case msg, ok := <-q["pm"]:
				if !ok {
					q["pm"] = nil
				} else {
					log.Printf("pm receive: %s\n", msg)
				}

			case msg, ok := <-q["rd"]:
				if !ok {
					q["rd"] = nil
				} else {
					log.Printf("rd receive: %s\n", msg)
				}
			}

			if q["rd"] == nil && q["pm"] == nil {
				log.Println("done------------------")
				break
				// 這裡有問題，到底怎退出的?
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()
}
