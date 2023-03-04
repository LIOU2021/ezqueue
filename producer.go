package queue

import (
	"log"
	"sync"
)

func Producer() {
	go func() {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			for i := 0; i < 10; i++ {
				log.Println("pm send : ", i)
				h["pm"] <- func() { log.Println("pm receive: ", i) }
			}
			wg.Done()
		}()

		go func() {
			for i := 0; i < 10; i++ {
				log.Println("rd send : ", i)
				h["rd"] <- func() { log.Println("rd receive: ", i) }
			}
			wg.Done()
		}()
		wg.Wait()
		close(h["pm"])
		close(h["rd"])
	}()
}
