package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/simulation"
	nodesim "github.com/sentinel-official/hub/x/node/simulation"
	sessionsim "github.com/sentinel-official/hub/x/session/simulation"
	subscriptionsim "github.com/sentinel-official/hub/x/subscription/simulation"
)

func RandomizedParams(r *rand.Rand) []simulation.ParamChange {
	var params []simulation.ParamChange

	params = append(params, nodesim.ParamChanges(r)...)
	params = append(params, subscriptionsim.ParamChanges(r)...)
	params = append(params, sessionsim.ParamChanges(r)...)

	return params
}
