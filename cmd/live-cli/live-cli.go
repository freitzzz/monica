package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/freitzzz/monica/internal/logging"
	"github.com/freitzzz/monica/internal/mq"
	"github.com/freitzzz/monica/internal/schema"
	_ "github.com/joho/godotenv/autoload"
	"github.com/pebbe/zmq4"
)

func main() {
	ctx, err := mq.NewContext()

	if err != nil {
		log.Fatal(err)
	}

	s, err := mq.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer mq.Close(s)
	observe(s)
}

func observe(s *zmq4.Socket) {
	for {
		d, err := mq.SendRawRouteMessage(s, 99, 0)

		if err != nil {
			logging.Aspirador.Error(fmt.Sprintf("failed to receive data from central-server: %s", err))
		} else {
			var nodes []schema.Node
			err = json.Unmarshal(d, &nodes)

			if err != nil {
				logging.Aspirador.Error(fmt.Sprintf("failed to unmarshall json: %s", err))
			}

			for _, n := range nodes {
				println(
					fmt.Sprintf(
						"Node\nHostname: %s\nDistribution: %s\nHardware: %s\nType: %s\nCPU: %f\nRAM: %f",
						n.OS.Hostname,
						n.OS.Distribution,
						n.OS.Hardware,
						n.OS.Type,
						n.Usage.CPU,
						n.Usage.RAM,
					),
				)
			}
		}

		time.Sleep(time.Duration(1) * time.Second)
	}
}
