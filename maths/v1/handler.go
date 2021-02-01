package v1

import maths "github.com/greenvine/hey-grpc-starter-interface/gen/go/maths/v1"

type Handlers struct {
	CalculatorClient maths.CalculatorAPIClient
	CounterClient    maths.CounterAPIClient
}
