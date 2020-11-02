package session

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	k.SetParams(ctx, state.Params)
	for _, session := range state.Sessions {
		k.SetSession(ctx, session)
		k.SetSessionForSubscription(ctx, session.Subscription, session.ID)
		k.SetSessionForNode(ctx, session.Node, session.ID)
		k.SetSessionForAddress(ctx, session.Address, session.ID)

		if session.Status.Equal(hub.StatusActive) {
			k.SetOngoingSession(ctx, session.Subscription, session.Address, session.ID)
			k.SetActiveSessionAt(ctx, session.StatusAt, session.ID)
		}
	}

	k.SetSessionsCount(ctx, uint64(len(state.Sessions)))
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.NewGenesisState(
		k.GetSessions(ctx, 0, 0),
		k.GetParams(ctx),
	)
}

func ValidateGenesis(state types.GenesisState) error {
	if err := state.Params.Validate(); err != nil {
		return err
	}

	for _, session := range state.Sessions {
		if err := session.Validate(); err != nil {
			return err
		}
	}

	sessions := make(map[uint64]bool)
	for _, item := range state.Sessions {
		id := item.ID
		if sessions[id] {
			return fmt.Errorf("duplicate session id %d", id)
		}

		sessions[id] = true
	}

	return nil
}
