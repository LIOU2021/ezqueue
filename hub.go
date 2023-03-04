package queue

import (
	"errors"
	"fmt"
)

var h = make(map[string]chan func())

// add queue by name
func Add(queueName string) {
	h[queueName] = make(chan func())
}

func Close(queueName string) error {
	if h[queueName] == nil {
		return errors.New(fmt.Sprintf("queue name %s not exists", queueName))
	}
	close(h[queueName])
	fmt.Printf("queue %s close ! \n", queueName)
	return nil
}
