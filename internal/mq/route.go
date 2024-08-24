package mq

import "github.com/freitzzz/monica/internal/schema"

const (
	nodeAdvertisement = 0
	nodeStats         = 1
)

// Handles a message based on the route id that is extracted from the message.
// Message = [<route_id>, <encoded struct bitstream>...]
func handleMessage(
	b []byte,
	onNodeAdvertisement func(schema.Advertisement) error,
	onNodeStats func(schema.Stats) error,
	onUnrecognizedMessage func(any) error,
) error {
	rid := b[0]
	b = b[1:]

	switch rid {
	case nodeAdvertisement:
		return decodeAndCallback(b, onNodeAdvertisement)
	case nodeStats:
		return decodeAndCallback(b, onNodeStats)
	default:
		return decodeAndCallback(b, onUnrecognizedMessage)
	}
}

// Wraps message [Decode] call and if it doesn't fail, passes the decoded struct to a callback.
func decodeAndCallback[T any](b []byte, cb func(T) error) error {
	d, err := Decode[T](b)

	if err == nil {
		return cb(d)
	}

	return err
}
