package main

import (
	"time"

	"github.com/LIOU2021/ezqueue"
)

func main() {
	ezqueue.Producer()
	ezqueue.Consumer()

	for {
		time.Sleep(1 * time.Second)
	}
}
