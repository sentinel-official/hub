// DO NOT COVER

package simulation

import (
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/hub/x/subscription/types"
)

func RandomizedGenesisState(state *module.SimulationState) *types.GenesisState {
	var (
		expiryDuration time.Duration
	)

	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyExpiryDuration),
		&expiryDuration,
		state.Rand,
		func(r *rand.Rand) {
			expiryDuration = time.Duration(r.Int63n(MaxExpiryDuration)) * time.Millisecond
		},
	)

	return types.NewGenesisState(
		nil,
		types.NewParams(expiryDuration),
	)
}
