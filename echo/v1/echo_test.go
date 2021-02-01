package v1

import (
	"context"
	"reflect"
	"testing"

	echo "github.com/greenvine/hey-grpc-starter-interface/gen/go/echo/v1"
)

func TestEchoSend(t *testing.T) {
	type args struct {
		ctx context.Context
		req *echo.PingRequest
	}
	type testCase struct {
		name            string
		args            args
		want            *echo.PingResponse
		wantErr         bool
		ignoreTimestamp bool
	}

	testCases := []testCase{
		{
			name: "TestNonEmptyMessage",
			args: args{
				ctx: context.Background(),
				req: &echo.PingRequest{Message: "helloworld"},
			},
			want:            &echo.PingResponse{Message: "helloworld"},
			wantErr:         false,
			ignoreTimestamp: true,
		},
		{
			name: "TestEmptyMessage",
			args: args{
				ctx: context.Background(),
				req: &echo.PingRequest{},
			},
			want:            &echo.PingResponse{},
			wantErr:         false,
			ignoreTimestamp: true,
		},
	}

	for _, test := range testCases {
		test := test

		t.Run(test.name, func(t *testing.T) {
			s := &Handlers{}
			got, err := s.Send(test.args.ctx, test.args.req)

			if (err != nil) != test.wantErr {
				t.Errorf("Send() error = %v, wantErr: %v", err, test.wantErr)
				return
			}
			if test.ignoreTimestamp && got != nil {
				got.Timestamp = 0
				test.want.Timestamp = 0
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Send() got = %v, want = %v", got, test.want)
			}
		})
	}
}
