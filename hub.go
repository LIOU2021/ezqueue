package queue

import (
	"errors"
	"fmt"
)

var h = make(map[string]chan func())

// add queue by name
func Add(queueName ...string) {
	for _, v := range queueName {
		h[v] = make(chan func())
		fmt.Printf("queue %s create ! \n", v)
	}
}

func Close(queueName ...string) error {
	for _, v := range queueName {
		if h[v] == nil {
			return errors.New(fmt.Sprintf("queue name %s not exists", v))
		}
		close(h[v])
		fmt.Printf("queue %s close ! \n", v)
	}

	return nil
}

func CloseAll() {
	for i := range h {
		Close(i)
	}
}
