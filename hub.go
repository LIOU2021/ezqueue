package ezqueue

var q = make(map[string]chan string)

func init() {
	q["pm"] = make(chan string)
	q["rd"] = make(chan string)
}
