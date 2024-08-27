package mq

import (
	"fmt"
	"slices"

	"github.com/pebbe/zmq4"
)

func Send(s *zmq4.Socket, b []byte) (bool, error) {
	n, err := s.SendBytes(b, 0)

	if err != nil {
		return false, err
	}

	if len(b) != n {
		return false, fmt.Errorf("couldn't send all buffer bytes")
	}

	rb, err := s.RecvBytes(0)

	if err != nil {
		return false, err
	}

	return slices.Equal(rb, OK), nil
}

func SendRouteMessage[T any](s *zmq4.Socket, rid byte, d T) (bool, error) {
	mb, err := EncodeRouteMessage(rid, d)

	if err != nil {
		return false, err
	}

	return Send(s, mb)
}

func Reply(s *zmq4.Socket, b []byte) error {
	var n int
	var err error

	n, err = s.SendBytes(b, 0)

	if err != nil {
		return err
	}

	if len(b) != n {
		return fmt.Errorf("couldn't send all buffer bytes")
	}

	return nil
}

func ReplyOK(s *zmq4.Socket) error {
	return Reply(s, OK)
}

func ReplyNOK(s *zmq4.Socket) error {
	return Reply(s, NOK)
}
