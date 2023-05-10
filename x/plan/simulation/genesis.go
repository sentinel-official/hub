// DO NOT COVER

package simulation

import (
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/hub/x/plan/types"
)

func RandomizedGenesisState(state *module.SimulationState) types.GenesisState {
	return types.NewGenesisState(
		RandomGenesisPlans(state.Rand),
	)
}
