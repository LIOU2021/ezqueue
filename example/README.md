```bash
$ go run main.go
queue pm create ! interval: 3s, consumerNumber: 2 
queue rd create ! interval: 3s, consumerNumber: 2 
2023/03/05 23:00:14 pm send task: 0
2023/03/05 23:00:14 pm send task: 1
2023/03/05 23:00:14 rd send task: 0
2023/03/05 23:00:14 pm receive task: 0
2023/03/05 23:00:14 pm send task: 2
2023/03/05 23:00:14 pm receive task: 1
2023/03/05 23:00:14 rd send task: 1
2023/03/05 23:00:14 rd send task: 2
2023/03/05 23:00:14 rd send task: 3
2023/03/05 23:00:14 rd receive task: 0
2023/03/05 23:00:14 pm send task: 3
2023/03/05 23:00:14 pm send task: 4
2023/03/05 23:00:14 rd receive task: 1
2023/03/05 23:00:14 rd send task: 4
2023/03/05 23:00:15 before runtime.NumGoroutine():  12
2023/03/05 23:00:16 before runtime.NumGoroutine():  12
2023/03/05 23:00:17 before runtime.NumGoroutine():  12
2023/03/05 23:00:18 before runtime.NumGoroutine():  12
2023/03/05 23:00:19 before runtime.NumGoroutine():  12
2023/03/05 23:00:19 rd receive task: 2
2023/03/05 23:00:19 pm receive task: 2
2023/03/05 23:00:19 pm receive task: 3
2023/03/05 23:00:19 rd receive task: 3
2023/03/05 23:00:20 before runtime.NumGoroutine():  8
2023/03/05 23:00:21 before runtime.NumGoroutine():  8
2023/03/05 23:00:22 before runtime.NumGoroutine():  8
2023/03/05 23:00:23 before runtime.NumGoroutine():  8
2023/03/05 23:00:24 before runtime.NumGoroutine():  8
2023/03/05 23:00:24 pm receive task: 4
2023/03/05 23:00:24 rd receive task: 4
2023/03/05 23:00:25 before runtime.NumGoroutine():  6
2023/03/05 23:00:26 before runtime.NumGoroutine():  6
2023/03/05 23:00:27 before runtime.NumGoroutine():  6
2023/03/05 23:00:28 before runtime.NumGoroutine():  6
2023/03/05 23:00:29 before runtime.NumGoroutine():  6
2023/03/05 23:00:30 before runtime.NumGoroutine():  6
2023/03/05 23:00:31 before runtime.NumGoroutine():  6
2023/03/05 23:00:32 before runtime.NumGoroutine():  6
2023/03/05 23:00:33 before runtime.NumGoroutine():  6
2023/03/05 23:00:34 After runtime.NumGoroutine():  6
queue pm close !
queue rd close !
2023/03/05 23:00:34 before runtime.NumGoroutine():  2
2023/03/05 23:00:35 timeout runtime.NumGoroutine():  2
2023/03/05 23:00:35 finish ...