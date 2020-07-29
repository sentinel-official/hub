package session

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/session/keeper"
	"github.com/sentinel-official/hub/x/vpn/session/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	k.SetParams(ctx, state.Params)
	for _, session := range state.Sessions {
		k.SetSession(ctx, session)
		if session.Status.Equal(hub.StatusActive) {
			k.SetActiveSession(ctx, session.Subscription, session.Node, session.Address, session.ID)
		}

		k.SetSessionsCount(ctx, k.GetSessionsCount(ctx)+1)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.NewGenesisState(
		k.GetSessions(ctx),
		k.GetParams(ctx),
	)
}

func ValidateGenesis(state types.GenesisState) error {
	return nil
}
