package queue

var h = make(map[string]chan func())

func init() {
	h["pm"] = make(chan func())
	h["rd"] = make(chan func())
}

func Add(queueName string) {
	h[queueName] = make(chan func())
}
