package v1

import (
	"context"
	"time"

	echo "github.com/greenvine/hey-grpc-starter-interface/gen/go/echo/v1"
)

var _ echo.EchoAPIServer = (*Handlers)(nil)

func (s *Handlers) Send(ctx context.Context, req *echo.PingRequest) (*echo.PingResponse, error) {
	return &echo.PingResponse{
		Message:   req.GetMessage(),
		Timestamp: time.Now().Unix(),
	}, nil
}

func (s *Handlers) Subscribe(req *echo.PingRequest, stream echo.EchoAPI_SubscribeServer) error {
	counter := 0

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	ctx := stream.Context()

	for {
		select {
		case timestamp := <-ticker.C:
			response := &echo.PingResponse{
				Message:   req.GetMessage(),
				Timestamp: timestamp.Unix(),
			}

			if err := stream.Send(response); err != nil {
				return err
			}

			counter++
			if counter >= 5 {
				return nil // stop after repeating for 5 times
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
