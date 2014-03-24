package concurrently

import (
	"runtime"
)

type Task interface {
	Run()
}

func RunNCPU(tasks []Task) {
	nCPU := runtime.NumCPU()
	if len(tasks) < nCPU {
		nCPU = len(tasks)
	}
	runtime.GOMAXPROCS(nCPU)
	run(tasks, nCPU)
}

func RunN(tasks []Task, nWorkers int) {
	run(tasks, nWorkers)
}

func run(tasks []Task, nWorkers int) {
	runtime.GOMAXPROCS(nWorkers)
	queue := make(chan *Task)

	//lets create some workers
	for i := 0; i < nWorkers; i++ {
		go do(queue)
	}

	//let's put the tasks onto the queue
	for i, _ := range tasks {
		queue <- &tasks[i]
	}

	//All items in queue. Lets shutdown the queue and workers
	for i := 0; i < nWorkers; i++ {
		queue <- nil
	}
}

func do(ch <-chan *Task) {
	for {
		task := <-ch
		if task == nil {
			break
		}

		(*task).Run()
	}
}
