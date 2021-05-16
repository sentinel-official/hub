package v06

import (
	"github.com/sentinel-official/hub/x/node/types"
	legacy "github.com/sentinel-official/hub/x/node/types/legacy/v0.5"
)

func MigrateGenesisState(state *legacy.GenesisState) *types.GenesisState {
	return types.NewGenesisState(
		MigrateNodes(state.Nodes),
		MigrateParams(state.Params),
	)
}
