package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) (err error) {
	var (
		i        int
		errCount uint32
		wg       = &sync.WaitGroup{}
		tasksCh  = make(chan Task)
	)

	for i = 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range tasksCh {
				err := task()
				if m > 0 && err != nil {
					atomic.AddUint32(&errCount, 1)
				}
			}
		}()
	}

	for i = range tasks {
		if m > 0 && atomic.LoadUint32(&errCount) >= uint32(m) {
			err = ErrErrorsLimitExceeded
			break
		}
		tasksCh <- tasks[i]
	}
	close(tasksCh)
	wg.Wait()
	return
}
