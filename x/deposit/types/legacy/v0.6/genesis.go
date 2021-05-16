package v06

import (
	"github.com/sentinel-official/hub/x/deposit/types"
	legacy "github.com/sentinel-official/hub/x/deposit/types/legacy/v0.5"
)

func MigrateGenesisState(state legacy.GenesisState) types.GenesisState {
	return types.NewGenesisState(
		MigrateDeposits(legacy.Deposits(state)),
	)
}
