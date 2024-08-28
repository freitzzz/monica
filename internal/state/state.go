package state

import (
	"fmt"

	"github.com/freitzzz/monica/internal/schema"
)

// Holds the state of current advertised nodes usage.
var state = make(State)

type State = map[string]*nodeState
type nodeState struct {
	Info  schema.NodeInfo
	Usage schema.NodeUsage
}

// Starts tracking a node.
func Insert(info schema.NodeInfo) {
	state[info.ID] = &nodeState{
		Info: info,
	}
}

// Update node usage.
func Update(usage schema.NodeUsage) {
	state[usage.ID].Usage = usage
}

func Lookup(id string) (nodeState, error) {
	var ns nodeState

	if ns, ok := state[id]; ok {
		return *ns, nil
	}

	return ns, fmt.Errorf("node with id %s hasn't been registered yet", id)
}

func ToNodes() ([]schema.Node, error) {
	c := len(state)

	if c == 0 {
		return nil, fmt.Errorf("no nodes have advertised yet")
	}

	v := make([]schema.Node, c)

	i := 0
	for _, s := range state {
		if s == nil {
			continue
		}

		v[i] = schema.ToNode(s.Info, s.Usage)
		i++
	}

	return v, nil
}
