package v1

import (
	"context"
	"sync"

	maths "github.com/greenvine/hey-grpc-starter-interface/gen/go/maths/v1"
)

type Counter struct {
	sync.Mutex
	value uint64
}

var _ maths.CounterAPIServer = (*Handlers)(nil)
var atomicCounter = Counter{value: 0}

func (s *Handlers) Increment(
	_ context.Context, req *maths.IncrementCounterRequest) (*maths.IncrementCounterResponse, error) {
	atomicCounter.Lock()
	defer atomicCounter.Unlock()

	atomicCounter.value += req.GetStep()
	return &maths.IncrementCounterResponse{Value: atomicCounter.value}, nil
}
