package processor

import (
	"sync"

	"github.com/georgezeimp/parallel-requests/concurrency"
	"github.com/georgezeimp/parallel-requests/hasher"
	"github.com/georgezeimp/parallel-requests/output"
)

type requestService interface {
	Get(address string) []byte
}

type Processor struct {
	h  *hasher.Hasher
	op *output.Presenter
	rs requestService
}

func NewProcessor(hasher *hasher.Hasher, outputPresenter *output.Presenter, requestService requestService) *Processor {
	return &Processor{
		h:  hasher,
		op: outputPresenter,
		rs: requestService,
	}
}

func (s *Processor) Process(addresses []string, npr int) []string {
	// Prepare concurrency related instances
	results := concurrency.ConcurrentSlice{}
	ch := make(chan int, npr)
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(addresses))

	// Spawn concurrent goroutines for each address
	for _, address := range addresses {
		go func(a string) {
			// Chek if there was an error or channel was closed
			_, ok := <-ch
			if !ok {
				waitGroup.Done()
				return
			}

			// Safely append result to the rest of the results
			results.Append(s.op.Prepare([]string{a, s.h.ToMD5(s.rs.Get(a))}))
			waitGroup.Done()
		}(address)
	}

	// Unblock channel to start processing the input
	for i := 0; i < len(addresses); i++ {
		ch <- i
	}

	// Finalise concurrency related code
	close(ch)
	waitGroup.Wait()

	// Return result items
	return results.GetItems()
}
