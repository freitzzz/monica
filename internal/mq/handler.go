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

			ReplyNOK(s)
			continue
		}

		if len(b) == 0 {
			logging.Aspirador.Error("empty message")

			ReplyNOK(s)
			continue
		}

		err = handleMessage(
			b,
			onNodeAdvertisement,
			onNodeStats,
			onUnrecognizedMessage,
		)

		if err != nil {
			logging.Aspirador.Error(fmt.Sprintf("failed to process message: %v", err))

			ReplyNOK(s)
			continue
		}

		ReplyOK(s)
	}
}

func onNodeAdvertisement(adv schema.Advertisement) error {
	nid := adv.ID
	if _, err := Lookup(nid); err == nil {
		return fmt.Errorf("node (%s) has been advertised before", nid)
	}

	Insert(nid, nil)

	return nil
}

func onNodeStats(stats schema.Stats) error {
	Insert(stats.ID, &stats)

	return nil
}

func onUnrecognizedMessage(m any) error {
	return fmt.Errorf("did not recognize message: %v", m)
}
