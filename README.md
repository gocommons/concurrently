concurrently - A concurrency library for Go language
=============================================

Even thought the basic building blocks for concurrency are built into the Go language itself, this library aims to make it a bit easier for you to use these concepts. 

Sample Usage
-----

```go
package main

import (
	"fmt"
	"github.com/gocommons/concurrently"
	"time"
)

type MyTask struct {
	ipAddress string
	ch        chan string
}

func (t MyTask) Run() {
	fmt.Println("Processing ", t.ipAddress)
	time.Sleep(10 * time.Second)
	t.ch <- fmt.Sprintf("Processing %s ends ", t.ipAddress)
}

func main() {
	t := time.Now()
	taskList := make([]concurrently.Task, 10)
	ch := make(chan string)
	go func() {
		for {
			i, ok := <-ch
			if !ok {
				break
			}
			fmt.Println(i)
		}
	}()
	for i := 0; i < 10; i++ {
		taskList[i] = MyTask{fmt.Sprintf("%s-%d", "Test", i), ch}
	}
	concurrently.RunN(taskList, 10)
	close(ch)
	t2 := time.Now()
	fmt.Println(t2.Sub(t))
}

```
