package state

import (
	"log"

	"github.com/denisbrodbeck/machineid"
)

var uniqueId string

func init() {
	id, err := machineid.ID()

	if err != nil {
		log.Fatalf("failed to get machined id. required to start a node. %v", err)
	}

	uniqueId = id
}

func Id() string {
	return uniqueId
}
