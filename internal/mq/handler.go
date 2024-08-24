package mq

import (
	"fmt"

	"github.com/freitzzz/monica/internal/logging"
	"github.com/freitzzz/monica/internal/schema"
	"github.com/pebbe/zmq4"
)

func Handle(s *zmq4.Socket) {
	for {
		b, err := s.RecvBytes(0)

		if err != nil {
			logging.Aspirador.Error(fmt.Sprintf("s.Recv call failed: %v", err))
			replyNok(s)
			continue
		}

		if len(b) == 0 {
			logging.Aspirador.Error("empty message")
			replyNok(s)
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
			replyNok(s)
			continue
		}

		replyOk(s)
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

func replyOk(s *zmq4.Socket) (int, error) {
	return s.Send("ok", 0)
}

func replyNok(s *zmq4.Socket) (int, error) {
	return s.Send("nok", 0)
}
