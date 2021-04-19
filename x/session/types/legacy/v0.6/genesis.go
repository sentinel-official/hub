package v0_6

import (
	"github.com/sentinel-official/hub/x/session/types"
	legacy "github.com/sentinel-official/hub/x/session/types/legacy/v0.5"
)

func MigrateGenesisState(state *legacy.GenesisState) *types.GenesisState {
	return types.NewGenesisState(
		MigrateSessions(state.Sessions),
		MigrateParams(state.Params),
	)
}
