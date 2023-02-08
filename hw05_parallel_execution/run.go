package hw05parallelexecution

import (
	"errors"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func worker(jobs <-chan Task, errs chan<- error, quit <-chan struct{}) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}

			err := job()
			if err != nil {
				errs <- err
			} else {
				errs <- nil
			}
		case <-quit:
			return
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
	errs := make(chan error)
	quit := make(chan struct{})
	// var wg sync.WaitGroup

	for w := 1; w <= n; w++ {
		// wg.Add(1)

		go func() {
			// defer wg.Done()
			worker(jobs, errs, quit)
		}()
	}

	go func() {
		for _, task := range tasks {
			select {
			case jobs <- task:
			case <-quit:
				return
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

			for b := 1; b <= n; b++ {
				<-errs
			}
			// wg.Wait()
			return ErrErrorsLimitExceeded
		}
	}
	// wg.Wait()
	close(quit)
	return nil
}
