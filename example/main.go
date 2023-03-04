package main

import (
	"time"

	queue "github.com/LIOU2021/ezqueue"
)

func main() {
	queue.Producer()
	queue.Consumer()

	for {
		time.Sleep(1 * time.Second)
	}
}
