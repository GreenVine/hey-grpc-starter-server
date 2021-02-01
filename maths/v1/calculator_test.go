package v1

import (
	"context"
	"reflect"
	"testing"

	maths "github.com/greenvine/hey-grpc-starter-interface/gen/go/maths/v1"
)

func TestCalculatorDivideIntegers(t *testing.T) {
	type args struct {
		ctx context.Context
		req *maths.DivideIntegersRequest
	}
	type testCase struct {
		name    string
		args    args
		want    *maths.DivideIntegersResponse
		wantErr bool
	}

	testCases := []testCase{
		{
			name: "TestDivisionWithoutRemainder",
			args: args{
				ctx: context.Background(),
				req: &maths.DivideIntegersRequest{
					Dividend: 10,
					Divisor:  2,
				},
			},
			want:    &maths.DivideIntegersResponse{Quotient: 5},
			wantErr: false,
		},
		{
			name: "TestDivisionWithRemainder",
			args: args{
				ctx: context.Background(),
				req: &maths.DivideIntegersRequest{
					Dividend: 11,
					Divisor:  -2,
				},
			},
			want:    &maths.DivideIntegersResponse{Quotient: -5},
			wantErr: false,
		},
		{
			name: "TestDivisionByZero",
			args: args{
				ctx: context.Background(),
				req: &maths.DivideIntegersRequest{
					Dividend: 10,
					Divisor:  0,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, test := range testCases {
		test := test

		t.Run(test.name, func(t *testing.T) {
			s := &Handlers{}
			got, err := s.DivideIntegers(test.args.ctx, test.args.req)

			if (err != nil) != test.wantErr {
				t.Errorf("DivideIntegers() error = %v, wantErr: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("DivideIntegers() got = %v, want = %v", got, test.want)
			}
		})
	}
}
