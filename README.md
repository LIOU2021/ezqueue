# golang easy queue
EzQueue is a lightweight Go package that provides a simple implementation of a queue and channels. You can create multiple independent queues, each with its own goroutine as a consumer, so the consumption behavior of different queues will not affect each other. Each time a producer distributes a task, it is dispatched through a separate goroutine, ensuring that the main thread is not blocked by task distribution. For example, if you need to send an email for a user registration API, you can call a third-party API to send an email and put it into the task for dispatching. Your API will not be waiting or blocked by the program in the task. This package is small and suitable for scenarios where development requirements are relatively simple. It can help you manage your tasks efficiently.

## why use?
EzQueue is an excellent choice when you need a lightweight and easy-to-use package for managing your queues. With support for creating multiple independent queues, each with its own goroutine as a consumer, EzQueue is an efficient way to manage your tasks. Unlike other queue packages, EzQueue does not block when dispatching tasks, making it ideal for scenarios where you need to process large volumes of tasks quickly. EzQueue is especially useful for managing simple cases such as sending emails or managing user registration, where you need a simple yet reliable solution.
## how to use?
1. import
    ```bash
    go get github.com/LIOU2021/ezqueue
    ```
2. create your queue
    ```go
    q1:=queue.NewQueue("myQueue01")
    q2:=queue.NewQueue("myQueue02")
    ```
3. custom config
    ```go
    // q1.SetConsumerNumber(2) // option. default 1
	q2.SetConsumerNumber(2) // option. default 1
    ```
4. add queue
    ```go
    queue.Add(q1, q2)
    ```
5. run consumer
    ```go
    func queue.Consumer()
    ```
6. dispatch your job
    ```go
    q1.Task(func() {
        log.Println("do something start ... ")
        time.Sleep(2 * time.Second)
        log.Println("do something end ... ")
    })

    q2.Task(func() {
        log.Println("do something start ... ")
        time.Sleep(2 * time.Second)
        log.Println("do something end ... ")
    })
    ```
## full example
- your main.go
    ```go
    package main

    import (
        "log"
        "time"

        queue "github.com/LIOU2021/ezqueue"
    )

    const (
        queuePM = "pm"
        queueRD = "rd"
    )

    func main() {
        // create two queue
        q1 := queue.NewQueue(queuePM)
        q2 := queue.NewQueue(queueRD)

        // q1.SetConsumerNumber(2) // option. default 1
        q2.SetConsumerNumber(2) // option. default 1

        // add queue
        queue.Add(q1, q2)

        // create consumer
        queue.Consumer()

        q1.Task(func() {
            log.Println("q1-1 do something start ... ")
            time.Sleep(2 * time.Second)
            log.Println("q1-1 do something end ... ")
        })

        q1.Task(func() {
            log.Println("q1-2 do something start ... ")
            time.Sleep(2 * time.Second)
            log.Println("q1-2 do something end ... ")
        })

        q2.Task(func() {
            log.Println("q2-1 do something start ... ")
            time.Sleep(2 * time.Second)
            log.Println("q2-1 do something end ... ")
        })

        q2.Task(func() {
            log.Println("q2-2 do something start ... ")
            time.Sleep(2 * time.Second)
            log.Println("q2-2 do something end ... ")
        })

        for {
            time.Sleep(1 * time.Second)
        }
    }
    ```
- output
    ```bash
    $ go run main.go
    queue pm create ! consumerNumber: 1 
    queue rd create ! consumerNumber: 2 
    2023/03/26 20:29:10 q1-1 do something start ... 
    2023/03/26 20:29:10 q2-2 do something start ... 
    2023/03/26 20:29:10 q2-1 do something start ...
    2023/03/26 20:29:13 q2-1 do something end ... 
    2023/03/26 20:29:13 q2-2 do something end ... 
    2023/03/26 20:29:13 q1-1 do something end ...
    2023/03/26 20:29:13 q1-2 do something start ...
    2023/03/26 20:29:15 q1-2 do something end ...
    ```    
## more example
- usage example see example/