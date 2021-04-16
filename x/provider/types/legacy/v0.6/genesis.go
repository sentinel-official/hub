package v0_6

import (
	"github.com/sentinel-official/hub/x/provider/types"
	legacy "github.com/sentinel-official/hub/x/provider/types/legacy/v0.5"
)

func MigrateGenesisState(state legacy.GenesisState) types.GenesisState {
	return types.NewGenesisState(
		MigrateProviders(state.Providers),
		MigrateParams(state.Params),
	)
}
