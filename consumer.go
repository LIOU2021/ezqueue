package queue

const (
	interval = 200 //Millisecond
)

func Consumer() {
	for i := range h {
		go func(queueName string) {
			for msg := range h[queueName] {
				msg()
			}
		}(i)
	}
}
