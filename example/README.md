```bash
$ go run main.go 
queue pm create ! 
queue rd create ! 
2023/03/05 21:08:31 pm send task:  0
2023/03/05 21:08:31 pm send task:  1
2023/03/05 21:08:31 rd send task:  0
2023/03/05 21:08:31 pm receive task:  0
2023/03/05 21:08:31 pm send task:  2
2023/03/05 21:08:31 pm send task:  3
2023/03/05 21:08:31 pm send task:  4
2023/03/05 21:08:31 rd send task:  1
2023/03/05 21:08:31 rd send task:  2
2023/03/05 21:08:31 rd receive task:  0
2023/03/05 21:08:31 rd send task:  3
2023/03/05 21:08:31 rd send task:  4
2023/03/05 21:08:32 before runtime.NumGoroutine():  12
2023/03/05 21:08:33 before runtime.NumGoroutine():  12
2023/03/05 21:08:33 rd receive task:  1
2023/03/05 21:08:33 pm receive task:  1
2023/03/05 21:08:34 before runtime.NumGoroutine():  10
2023/03/05 21:08:35 before runtime.NumGoroutine():  10
2023/03/05 21:08:35 pm receive task:  2
2023/03/05 21:08:35 rd receive task:  2
2023/03/05 21:08:36 before runtime.NumGoroutine():  8
2023/03/05 21:08:37 before runtime.NumGoroutine():  8
2023/03/05 21:08:37 rd receive task:  3
2023/03/05 21:08:37 pm receive task:  3
2023/03/05 21:08:38 before runtime.NumGoroutine():  6
2023/03/05 21:08:39 pm receive task:  4
2023/03/05 21:08:39 rd receive task:  4
2023/03/05 21:08:39 before runtime.NumGoroutine():  6
2023/03/05 21:08:40 before runtime.NumGoroutine():  4
2023/03/05 21:08:41 After runtime.NumGoroutine():  4
queue pm close !
queue rd close !
2023/03/05 21:08:41 before runtime.NumGoroutine():  4
2023/03/05 21:08:42 timeout runtime.NumGoroutine():  2
2023/03/05 21:08:42 finish ...