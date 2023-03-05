package queue

func (qu *BasicQueue) Task(f func()) {
	go func() {
		qu.msg <- f
	}()
}
