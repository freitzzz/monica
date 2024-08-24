package mq

import (
	"bytes"
	"encoding/gob"
)

// Nodes should use this function to encode messages
// to send to central-server.
func EncodeRouteMessage[T any](rid byte, d T) ([]byte, error) {
	b, err := Encode(d)

	if err != nil {
		return b, err
	}

	return append([]byte{rid}, b...), nil
}

func Encode(v any) ([]byte, error) {
	var buf bytes.Buffer
	var b []byte

	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)

	if err == nil {
		b = buf.Bytes()
	}

	return b, err
}

func Decode[T any](b []byte) (T, error) {
	var buf bytes.Buffer
	var v T

	buf.Write(b)

	enc := gob.NewDecoder(&buf)
	err := enc.Decode(&v)

	return v, err
}
