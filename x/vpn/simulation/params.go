// DO NOT COVER

package simulation

import (
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	nodesimulation "github.com/sentinel-official/hub/x/node/simulation"
	providersimulation "github.com/sentinel-official/hub/x/provider/simulation"
	sessionsimulation "github.com/sentinel-official/hub/x/session/simulation"
	subscriptionsimulation "github.com/sentinel-official/hub/x/subscription/simulation"
)

func RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	var params []simtypes.ParamChange
	params = append(params, providersimulation.ParamChanges(r)...)
	params = append(params, nodesimulation.ParamChanges(r)...)
	params = append(params, subscriptionsimulation.ParamChanges(r)...)
	params = append(params, sessionsimulation.ParamChanges(r)...)

	return params
}
