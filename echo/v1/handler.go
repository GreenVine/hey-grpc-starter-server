package v1

import echo "github.com/greenvine/hey-grpc-starter-interface/gen/go/echo/v1"

type Handlers struct {
	EchoClient echo.EchoAPIClient
}
