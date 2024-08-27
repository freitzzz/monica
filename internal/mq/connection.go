package mq

import (
	"fmt"
	"os"

	"github.com/pebbe/zmq4"
)

const (
	serverHostEnvKey = "server_host"
	serverPortEnvKey = "server_port"
)

var (
	serverHost = os.Getenv(serverHostEnvKey)
	serverPort = os.Getenv(serverPortEnvKey)
)

func NewContext() (*zmq4.Context, error) {
	return zmq4.NewContext()
}

func Start(ctx *zmq4.Context) (*zmq4.Socket, error) {
	addr := ServerAddress()

	s, err := ctx.NewSocket(zmq4.REP)

	if err == nil {
		err = s.Bind(addr)
	}

	return s, err
}

func Connect(ctx *zmq4.Context) (*zmq4.Socket, error) {
	addr := ServerAddress()

	s, err := ctx.NewSocket(zmq4.REQ)

	if err == nil {
		err = s.Connect(addr)
	}

	return s, err
}

func Close(s *zmq4.Socket) error {
	return s.Close()
}

func ServerAddress() string {
	return fmt.Sprintf("tcp://%s:%s", serverHost, serverPort)
}
