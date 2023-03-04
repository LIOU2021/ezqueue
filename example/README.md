```bash
$ go run main.go
queue pm create ! 
queue rd create ! 
2023/03/05 01:19:18 rd send task:  0
2023/03/05 01:19:18 rd send task:  1
2023/03/05 01:19:18 pm send task:  0
2023/03/05 01:19:18 rd receive task:  0
2023/03/05 01:19:18 rd receive task:  1
2023/03/05 01:19:18 rd send task:  2
2023/03/05 01:19:18 rd send task:  3
2023/03/05 01:19:18 pm send task:  1
2023/03/05 01:19:18 pm receive task:  0
2023/03/05 01:19:18 rd receive task:  2
2023/03/05 01:19:18 rd receive task:  3
2023/03/05 01:19:18 rd send task:  4
2023/03/05 01:19:18 pm send task:  2
2023/03/05 01:19:18 pm receive task:  1
2023/03/05 01:19:18 pm receive task:  2
2023/03/05 01:19:18 rd receive task:  4
2023/03/05 01:19:18 pm send task:  3
2023/03/05 01:19:18 pm send task:  4
2023/03/05 01:19:18 pm receive task:  3
2023/03/05 01:19:18 pm receive task:  4
queue pm close !
queue rd close !
2023/03/05 01:19:28 finish ...
```