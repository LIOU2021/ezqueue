package queue

func Consumer() {
	for i := range h {
		go func(queueName string) {
			for msg := range h[queueName] {
				msg()
			}
		}(i)
	}
}
