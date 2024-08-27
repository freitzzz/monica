package mq

import "github.com/freitzzz/monica/internal/schema"

const (
	publishNodeInformationRoute = 0
	publishNodeStatsRoute       = 1
	lookupAllStatsRoute         = 99
)

// Handles a message based on the route id that is extracted from the message.
// Message = [<route_id>, <encoded struct bitstream>...]
func handleMessage(
	b []byte,
	onPublishNodeInformation func(schema.NodeInfo) ReplyMessage,
	onPublishNodeStats func(schema.NodeUsage) ReplyMessage,
	onLookupAllStats func() ReplyMessage,
	onUnrecognizedMessage func(any) ReplyMessage,
) ReplyMessage {
	rid := b[0]
	b = b[1:]

	switch rid {
	case publishNodeInformationRoute:
		return decodeAndCallback(b, onPublishNodeInformation)
	case publishNodeStatsRoute:
		return decodeAndCallback(b, onPublishNodeStats)
	case lookupAllStatsRoute:
		return onLookupAllStats()
	default:
		return decodeAndCallback(b, onUnrecognizedMessage)
	}
}

// Wraps message [Decode] call and if it doesn't fail, passes the decoded struct to a callback.
func decodeAndCallback[T any](b []byte, cb func(T) ReplyMessage) ReplyMessage {
	d, err := Decode[T](b)

	if err == nil {
		return cb(d)
	}

	return ReplyMessage{
		Error: err,
	}
}
