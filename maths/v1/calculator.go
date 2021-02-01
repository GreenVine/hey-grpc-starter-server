package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	maths "github.com/greenvine/hey-grpc-starter-interface/gen/go/maths/v1"
)

var _ maths.CalculatorAPIServer = (*Handlers)(nil)

func (s *Handlers) DivideIntegers(
	ctx context.Context, req *maths.DivideIntegersRequest) (*maths.DivideIntegersResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	quotient := req.GetDividend() / req.GetDivisor()
	return &maths.DivideIntegersResponse{Quotient: quotient}, nil
}
