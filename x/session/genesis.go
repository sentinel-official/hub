package session

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state *types.GenesisState) {
	k.SetParams(ctx, state.Params)
	for _, session := range state.Sessions {
		var (
			sessionNode    = session.GetNode()
			sessionAddress = session.GetAddress()
		)

		k.SetSession(ctx, session)
		k.SetSessionForSubscription(ctx, session.Subscription, session.Id)
		k.SetSessionForNode(ctx, sessionNode, session.Id)
		k.SetSessionForAddress(ctx, sessionAddress, session.Id)

		if session.Status.Equal(hubtypes.StatusActive) {
			k.SetActiveSessionForAddress(ctx, sessionAddress, session.Subscription, sessionNode, session.Id)
			k.SetActiveSessionAt(ctx, session.StatusAt, session.Id)
		}
	}

	k.SetCount(ctx, uint64(len(state.Sessions)))
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.NewGenesisState(
		k.GetSessions(ctx, 0, 0),
		k.GetParams(ctx),
	)
}

func ValidateGenesis(state *types.GenesisState) error {
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
		if sessions[item.Id] {
			return fmt.Errorf("duplicate session id %d", item.Id)
		}

		sessions[item.Id] = true
	}

	return nil
}
