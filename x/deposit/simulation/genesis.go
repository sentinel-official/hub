// DO NOT COVER

package simulation

import (
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/hub/x/deposit/types"
)

func RandomizedGenesisState(_ *module.SimulationState) types.GenesisState {
	return types.NewGenesisState(
		nil,
	)
}
