package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := wrapDone(in, done)

	for _, stage := range stages {
		out = wrapDone(stage(out), done)
	}
	return out
}

func wrapDone(in In, done In) Out {
	out := make(Bi)

	go func() {
		defer func() {
			close(out)
			// Drain required if send data after done
			// for range in {}
		}()

		for {
			select {
			case <-done:
				return
			default:
			}
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				out <- v
			}
		}
	}()
	return out
}
