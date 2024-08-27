package mq

import (
	"fmt"

	"github.com/freitzzz/monica/internal/logging"
	"github.com/freitzzz/monica/internal/schema"
	"github.com/pebbe/zmq4"
)

func RegisterHandlers(s *zmq4.Socket) {
	for {
		b, err := s.RecvBytes(0)

		if err != nil {
			logging.Aspirador.Error(fmt.Sprintf("s.Recv call failed: %v", err))

			LogReplyError(ReplyNOK(s))
			continue
		}

		if len(b) == 0 {
			logging.Aspirador.Error("empty message")

			LogReplyError(ReplyNOK(s))
			continue
		}

		rpm := handleMessage(
			b,
			onNodeAdvertisement,
			onNodeStats,
			onLookupNodes,
			onUnrecognizedMessage,
		)

		err = rpm.Error

		if err != nil {
			logging.Aspirador.Error(fmt.Sprintf("failed to process message: %v", err))

			LogReplyError(ReplyNOK(s))
			continue
		}

		LogReplyError(Reply(s, rpm))
	}
}

func onNodeAdvertisement(adv schema.NodeInfo) ReplyMessage {
	nid := adv.ID
	if _, err := Lookup(nid); err == nil {
		return ReplyMessage{
			Error: fmt.Errorf("node (%s) has been advertised before", nid),
		}
	}

	Insert(adv)

	return OkReplyMessage()
}

func onNodeStats(usage schema.NodeUsage) ReplyMessage {
	nid := usage.ID
	if _, err := Lookup(nid); err != nil {
		return ReplyMessage{
			Error: fmt.Errorf("node (%s) hasn't been advertised before", nid),
		}
	}

	Update(usage)

	return OkReplyMessage()
}

func onLookupNodes() ReplyMessage {
	nodes, err := ToNodes()

	if err != nil {
		return ErrorReplyMessage(err)
	}

	return JSONReplyMessage(nodes)
}

func onUnrecognizedMessage(m any) ReplyMessage {
	return ErrorReplyMessage(fmt.Errorf("did not recognize message: %v", m))
}
