package queue

func (qu *BasicQueue) Task(f func()) {
	go func() {
		if qu.msgIsClose.Load() {
			return
		}
		qu.msgLen.Add(1)
		qu.msg <- f
	}()
}
