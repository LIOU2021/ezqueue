package queue

import (
	"errors"
	"fmt"
)

// add task like producer
func Task(queueName string, f func()) error {
	if h[queueName] == nil {
		return errors.New(fmt.Sprintf("queue name %s not exists", queueName))
	}

	go func() {
		h[queueName] <- f
	}()

	return nil
}
