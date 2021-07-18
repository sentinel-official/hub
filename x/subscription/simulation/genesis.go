package simulation

import (
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/hub/x/subscription/types"
)

func RandomizedGenesisState(state *module.SimulationState) *types.GenesisState {
	var (
		inactiveDuration time.Duration
	)

	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyInactiveDuration),
		&inactiveDuration,
		state.Rand,
		func(r *rand.Rand) {
			inactiveDuration = time.Duration(r.Int63n(MaxInactiveDuration)) * time.Millisecond
		},
	)

	return types.NewGenesisState(
		nil,
		types.NewParams(inactiveDuration),
	)
}
