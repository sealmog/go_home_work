package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func worker(jobs <-chan Task, errCounter *int32) {
	for job := range jobs {
		err := job()
		if err != nil {
			atomic.AddInt32(errCounter, 1)
		}
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	jobs := make(chan Task)
	var errCounter int32
	var wg sync.WaitGroup

	for w := 1; w <= n; w++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker(jobs, &errCounter)
		}()
	}

	wg.Add(1)

	go func() {
		defer wg.Done()
		for _, task := range tasks {
			jobs <- task
			if atomic.LoadInt32(&errCounter) >= int32(m) {
				break
			}
		}
		close(jobs)
	}()

	wg.Wait()
	if errCounter >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
