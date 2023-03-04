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
				value := i
				h["pm"] <- func() { log.Println("pm receive: ", value) }
			}
			wg.Done()
		}()

		go func() {
			for i := 0; i < 10; i++ {
				log.Println("rd send : ", i)
				value := i
				h["rd"] <- func() { log.Println("rd receive: ", value) }
			}
			wg.Done()
		}()
		wg.Wait()
		close(h["pm"])
		close(h["rd"])
	}()
}
