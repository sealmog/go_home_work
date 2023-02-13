package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func worker(jobs <-chan Task, errs chan<- error, quit <-chan struct{}) {
	for {
		select {
		case <-quit:
			return
		default:
		}
		select {
		case <-quit:
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			errs <- job()
		}
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	numJobs := len(tasks)
	jobs := make(chan Task)
	errs := make(chan error, numJobs)
	quit := make(chan struct{})
	var wg sync.WaitGroup

	for w := 1; w <= n; w++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker(jobs, errs, quit)
		}()
	}

	wg.Add(1)

	go func() {
		defer wg.Done()
		defer close(jobs)

		for _, task := range tasks {
			select {
			case <-quit:
				return
			default:
			}
			select {
			case <-quit:
				return
			case jobs <- task:
			}
		}
	}()

	errCounter := 0

	for a := 1; a <= numJobs; a++ {
		err := <-errs
		if err != nil {
			errCounter++
		}

		if errCounter >= m {
			close(quit)
			wg.Wait()
			close(errs)
			return ErrErrorsLimitExceeded
		}
	}

	close(quit)
	wg.Wait()
	close(errs)
	return nil
}
