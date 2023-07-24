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
		inactivePendingDuration time.Duration
	)

	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyInactivePendingDuration),
		&inactivePendingDuration,
		state.Rand,
		func(r *rand.Rand) {
			inactivePendingDuration = time.Duration(r.Int63n(MaxInactivePendingDuration)) * time.Millisecond
		},
	)

	return types.NewGenesisState(
		nil,
		types.NewParams(inactivePendingDuration, false),
	)
}
