package simulation

import (
	"math/rand"

	"github.com/sentinel-official/hub/x/node/types"
)

func RandomNode(r *rand.Rand, nodes types.Nodes) types.Node {
	if len(nodes) == 0 {
		return types.Node{
			Address: []byte("address"),
		}
	}

	return nodes[r.Intn(
		len(nodes),
	)]
}
