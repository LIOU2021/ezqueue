```bash
$ go run main.go
queue pm create ! interval: 3s, consumerNumber: 2 
queue rd create ! interval: 3s, consumerNumber: 2 
2023/03/05 23:58:33 rd send task: 0
2023/03/05 23:58:33 rd send task: 1
2023/03/05 23:58:33 pm send task: 0
2023/03/05 23:58:33 rd receive task: 0
2023/03/05 23:58:33 rd send task: 2
2023/03/05 23:58:33 rd send task: 3
2023/03/05 23:58:33 rd send task: 4
2023/03/05 23:58:33 rd receive task: 1
2023/03/05 23:58:33 pm send task: 1
2023/03/05 23:58:33 pm send task: 2
2023/03/05 23:58:33 pm receive task: 0
2023/03/05 23:58:33 pm receive task: 1
2023/03/05 23:58:33 pm send task: 3
2023/03/05 23:58:33 pm send task: 4
2023/03/05 23:58:34 before runtime.NumGoroutine():  12
2023/03/05 23:58:35 before runtime.NumGoroutine():  12
2023/03/05 23:58:36 before runtime.NumGoroutine():  12
2023/03/05 23:58:37 before runtime.NumGoroutine():  12
2023/03/05 23:58:38 before runtime.NumGoroutine():  12
2023/03/05 23:58:38 rd receive task: 2
2023/03/05 23:58:38 rd receive task: 3
2023/03/05 23:58:38 pm receive task: 2
2023/03/05 23:58:38 pm receive task: 3
2023/03/05 23:58:39 before runtime.NumGoroutine():  8
2023/03/05 23:58:40 before runtime.NumGoroutine():  8
2023/03/05 23:58:41 before runtime.NumGoroutine():  8
2023/03/05 23:58:42 before runtime.NumGoroutine():  8
2023/03/05 23:58:43 After runtime.NumGoroutine():  8
close after waiting for 1 tasks for queue "pm"
close after waiting for 1 tasks for queue "rd"
2023/03/05 23:58:43 pm receive task: 4
2023/03/05 23:58:43 rd receive task: 4
2023/03/05 23:58:43 before runtime.NumGoroutine():  8
close after waiting for 1 tasks for queue "rd"
close after waiting for 1 tasks for queue "pm"
2023/03/05 23:58:44 before runtime.NumGoroutine():  8
close after waiting for 1 tasks for queue "pm"
close after waiting for 1 tasks for queue "rd"
2023/03/05 23:58:45 before runtime.NumGoroutine():  8
queue rd close !
queue pm close !
2023/03/05 23:58:46 before runtime.NumGoroutine():  4
2023/03/05 23:58:47 rd send task: 0
2023/03/05 23:58:47 rd send task: 1
2023/03/05 23:58:47 rd send task: 2
2023/03/05 23:58:47 pm send task: 0
2023/03/05 23:58:47 pm send task: 1
2023/03/05 23:58:47 pm send task: 2
2023/03/05 23:58:47 pm send task: 3
2023/03/05 23:58:47 rd send task: 3
2023/03/05 23:58:47 rd send task: 4
2023/03/05 23:58:47 pm send task: 4
2023/03/05 23:58:47 before runtime.NumGoroutine():  4
2023/03/05 23:58:48 before runtime.NumGoroutine():  2
2023/03/05 23:58:49 before runtime.NumGoroutine():  2
2023/03/05 23:58:50 before runtime.NumGoroutine():  2
2023/03/05 23:58:51 before runtime.NumGoroutine():  2
2023/03/05 23:58:52 timeout runtime.NumGoroutine():  2
2023/03/05 23:58:52 finish ...