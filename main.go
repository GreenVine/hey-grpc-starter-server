package main

import (
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"

	echo "github.com/greenvine/hey-grpc-starter-interface/gen/go/echo/v1"
	maths "github.com/greenvine/hey-grpc-starter-interface/gen/go/maths/v1"

	echoImpl "github.com/greenvine/hey-grpc-starter-server/echo/v1"
	mathsImpl "github.com/greenvine/hey-grpc-starter-server/maths/v1"
)

func serve(proto, listenAddress string) error {
	s := grpc.NewServer()

	echo.RegisterEchoAPIServer(s, &echoImpl.Handlers{})
	maths.RegisterCalculatorAPIServer(s, &mathsImpl.Handlers{})
	maths.RegisterCounterAPIServer(s, &mathsImpl.Handlers{})

	listener, err := net.Listen(proto, listenAddress)
	if err != nil {
		return err
	}

	log.Printf("Server is listening on: %s", listener.Addr().String())

	return s.Serve(listener)
}

func main() {
	listenAddress := flag.String("address", "localhost:3000", "server listen address")
	flag.Parse()

	if err := serve("tcp", *listenAddress); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
