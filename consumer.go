package queue

import (
	"log"
	"time"
)

const (
	interval = 200 //Millisecond
)

func Consumer() {
	go func() {

		for {
			select {
			case msg, ok := <-h["pm"]:
				if len(h["pm"]) == 0 && !ok {
					// h["pm"] = nil
					log.Println("pm done------------------")
					return
				} else {
					msg()
				}
			}

			// if h["pm"] == nil {
			// 	log.Println("pm done------------------")
			// 	break
			// }
			time.Sleep(interval * time.Millisecond)
		}

	}()

	go func() {
		for {
			select {
			case msg, ok := <-h["rd"]:
				if len(h["rd"]) == 0 && !ok {
					// h["rd"] = nil
					log.Println("rd done------------------")
					return
				} else {
					msg()
				}
			}

			// if h["rd"] == nil {
			// 	log.Println("rd done------------------")
			// 	break
			// }
			time.Sleep(interval * time.Millisecond)
		}
	}()
}
