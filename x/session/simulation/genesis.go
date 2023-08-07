// DO NOT COVER

package simulation

import (
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/hub/x/session/types"
)

func RandomizedGenesisState(state *module.SimulationState) *types.GenesisState {
	var (
		statusChangeDelay time.Duration
	)

	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyStatusChangeDelay),
		&statusChangeDelay,
		state.Rand,
		func(r *rand.Rand) {
			statusChangeDelay = time.Duration(r.Int63n(MaxStatusChangeDelay)) * time.Millisecond
		},
	)

	return types.NewGenesisState(
		nil,
		types.NewParams(statusChangeDelay, false),
	)
}
