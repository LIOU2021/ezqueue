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
    func queue.Add(queueName ...string)
    ```
3. run consumer
    ```go
    func queue.Consumer()
    ```
4. dispatch your job
    ```go
    err := queue.Task(queuePM, func() {
			log.Println("pm receive task: ", v)
			time.Sleep(2 * time.Second)
		})
    ```
## more example
- usage example see example/