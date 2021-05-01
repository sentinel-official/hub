package v0_6

import (
	deposit "github.com/sentinel-official/hub/x/deposit/types/legacy/v0.6"
	node "github.com/sentinel-official/hub/x/node/types/legacy/v0.6"
	plan "github.com/sentinel-official/hub/x/plan/types/legacy/v0.6"
	provider "github.com/sentinel-official/hub/x/provider/types/legacy/v0.6"
	session "github.com/sentinel-official/hub/x/session/types/legacy/v0.6"
	subscription "github.com/sentinel-official/hub/x/subscription/types/legacy/v0.6"
	"github.com/sentinel-official/hub/x/vpn/types"
	legacy "github.com/sentinel-official/hub/x/vpn/types/legacy/v0.5"
)

func MigrateGenesisState(state *legacy.GenesisState) *types.GenesisState {
	return types.NewGenesisState(
		deposit.MigrateGenesisState(state.Deposits),
		provider.MigrateGenesisState(state.Providers),
		node.MigrateGenesisState(state.Nodes),
		plan.MigrateGenesisState(state.Plans),
		subscription.MigrateGenesisState(state.Subscriptions),
		session.MigrateGenesisState(state.Sessions),
	)
}
