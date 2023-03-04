package queue

import (
	"errors"
	"fmt"
)

// func Producer() {
// 	go func() {
// 		var wg sync.WaitGroup
// 		wg.Add(2)
// 		go func() {
// 			for i := 0; i < 10; i++ {
// 				go func(index int) {
// 					log.Println("pm send : ", index)
// 					value := index
// 					h["pm"] <- func() { log.Println("pm receive: ", value) }
// 				}(i)
// 			}
// 			wg.Done()
// 		}()

// 		go func() {
// 			for i := 0; i < 10; i++ {
// 				go func(index int) {
// 					log.Println("rd send : ", index)
// 					value := index
// 					h["rd"] <- func() { log.Println("rd receive: ", value) }
// 				}(i)
// 			}
// 			wg.Done()
// 		}()
// 		wg.Wait()
// 		close(h["pm"])
// 		close(h["rd"])
// 	}()
// }

func Task(queueName string, f func()) error {
	if h[queueName] == nil {
		return errors.New(fmt.Sprintf("queue name %s not exists", queueName))
	}

	go func() {
		h[queueName] <- f
	}()

	return nil
}
