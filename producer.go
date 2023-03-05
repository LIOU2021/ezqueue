package queue

func (qu *BasicQueue) Task(f func()) {
	go func() {
		qu.msgLen.Add(1)
		qu.msg <- f
	}()
}
