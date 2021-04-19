package v0_6

import (
	"github.com/sentinel-official/hub/x/swap/types"
	legacy "github.com/sentinel-official/hub/x/swap/types/legacy/v0.5"
)

func MigrateGenesisState(state *legacy.GenesisState) *types.GenesisState {
	return types.NewGenesisState(
		MigrateSwaps(state.Swaps),
		MigrateParams(state.Params),
	)
}
