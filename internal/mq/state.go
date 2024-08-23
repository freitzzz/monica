package mq

import (
	"fmt"

	"github.com/freitzzz/monica/internal/schema"
)

// Holds the state of current advertised nodes statistics.
type State = map[string]schema.Stats

var state = make(State)

func Insert(id string, stats *schema.Stats) {
	if stats == nil {
		state[id] = schema.Stats{}
	} else {
		state[id] = *stats
	}
}

func Lookup(id string) (schema.Stats, error) {
	var stats schema.Stats

	if stats, ok := state[id]; ok {
		return stats, nil
	}

	return stats, fmt.Errorf("node with id %s hasn't been registered yet", id)
}
