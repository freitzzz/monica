package mq

import (
	"fmt"
	"slices"

	"github.com/freitzzz/monica/internal/logging"
	"github.com/pebbe/zmq4"
)

func Send(s *zmq4.Socket, b []byte) (bool, error) {
	rb, err := sendRaw(s, b)

	if err != nil {
		return false, err
	}

	return slices.Equal(rb, OK), nil
}

func sendRaw(s *zmq4.Socket, b []byte) ([]byte, error) {
	n, err := s.SendBytes(b, 0)

	if err != nil {
		return nil, err
	}

	if len(b) != n {
		return nil, fmt.Errorf("couldn't send all buffer bytes")
	}

	return s.RecvBytes(0)
}

func SendRouteMessage[T any](s *zmq4.Socket, rid byte, d T) (bool, error) {
	mb, err := EncodeRouteMessage(rid, d)

	if err != nil {
		return false, err
	}

	return Send(s, mb)
}

func SendRawRouteMessage[T any](s *zmq4.Socket, rid byte, d T) ([]byte, error) {
	mb, err := EncodeRouteMessage(rid, d)

	if err != nil {
		return nil, err
	}

	return sendRaw(s, mb)
}

func Reply(s *zmq4.Socket, m ReplyMessage) error {
	d := m.Data

	switch t := d.(type) {
	case bool:
		if t {
			return ReplyOK(s)
		} else {
			return ReplyNOK(s)
		}
	case []byte:
		return replyRaw(s, t)
	default:
		b, err := Encode(d)

		if err != nil {
			logging.Aspirador.Error(fmt.Sprintf("failed to encode data %v, error: %s", d, err))
			return replyRaw(s, ERROR)
		}

		return replyRaw(s, b)
	}
}

func ReplyOK(s *zmq4.Socket) error {
	return replyRaw(s, OK)
}

func ReplyNOK(s *zmq4.Socket) error {
	return replyRaw(s, NOK)
}

func replyRaw(s *zmq4.Socket, b []byte) error {
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
