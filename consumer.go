package queue

import (
	"log"
	"time"
)

func Consumer() {
	go func() {
		for {
			select {
			case msg, ok := <-h["pm"]:
				if !ok {
					h["pm"] = nil
				} else {
					msg()
				}

			case msg, ok := <-h["rd"]:
				if !ok {
					h["rd"] = nil
				} else {
					msg()
				}
			}

			if h["rd"] == nil && h["pm"] == nil {
				log.Println("done------------------")
				break
				// 這裡有問題，到底怎退出的?
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()
}
