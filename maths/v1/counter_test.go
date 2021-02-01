package v1

import (
	"context"
	"reflect"
	"testing"

	maths "github.com/greenvine/hey-grpc-starter-interface/gen/go/maths/v1"
)

func TestHandlers_Increment(t *testing.T) {
	type args struct {
		ctx context.Context
		req *maths.IncrementCounterRequest
	}
	type testCase struct {
		name    string
		args    args
		want    *maths.IncrementCounterResponse
		wantErr bool
	}

	// tests must be executed in order for checks to be successful
	testCases := []testCase{
		{
			name: "TestIncrementByZero",
			args: args{
				ctx: context.Background(),
				req: &maths.IncrementCounterRequest{Step: 0},
			},
			want:    &maths.IncrementCounterResponse{Value: 0},
			wantErr: false,
		},
		{
			name: "TestIncrementByOne",
			args: args{
				ctx: context.Background(),
				req: &maths.IncrementCounterRequest{Step: 1},
			},
			want:    &maths.IncrementCounterResponse{Value: 1},
			wantErr: false,
		},
		{
			name: "TestIncrementByTwo",
			args: args{
				ctx: context.Background(),
				req: &maths.IncrementCounterRequest{Step: 2},
			},
			want:    &maths.IncrementCounterResponse{Value: 3},
			wantErr: false,
		},
	}

	for _, test := range testCases {
		test := test

		t.Run(test.name, func(t *testing.T) {
			s := &Handlers{}
			got, err := s.Increment(test.args.ctx, test.args.req)

			if (err != nil) != test.wantErr {
				t.Errorf("Increment() error = %v, wantErr: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Increment() got = %v, want = %v", got, test.want)
			}
		})
	}
}
