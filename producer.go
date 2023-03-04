package ezqueue

import (
	"strconv"
	"sync"
)

func Producer() {
	go func() {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			for i := 0; i < 10; i++ {
				q["pm"] <- "p1: " + strconv.Itoa(i)
			}
			wg.Done()
		}()

		go func() {
			for i := 0; i < 10; i++ {
				q["rd"] <- "r1: " + strconv.Itoa(i)
			}
			wg.Done()
		}()
		wg.Wait()
		close(q["pm"])
		close(q["rd"])
	}()
}
