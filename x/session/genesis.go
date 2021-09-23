package session

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state *types.GenesisState) {
	k.SetParams(ctx, state.Params)

	inactiveDuration := k.InactiveDuration(ctx)
	for _, session := range state.Sessions {
		address := session.GetAddress()
		k.SetSession(ctx, session)

		if session.Status.Equal(hubtypes.StatusActive) {
			k.SetActiveSessionForAddress(ctx, address, session.Id)
		} else {
			k.SetInactiveSessionForAddress(ctx, address, session.Id)
		}

		if session.Status.Equal(hubtypes.StatusActive) {
			k.SetInactiveSessionAt(ctx, session.StatusAt.Add(inactiveDuration), session.Id)
		}
		if session.Status.Equal(hubtypes.StatusInactivePending) {
			k.SetInactiveSessionAt(ctx, session.StatusAt.Add(inactiveDuration), session.Id)
		}
	}

	count := uint64(0)
	for _, item := range state.Sessions {
		if item.Id > count {
			count = item.Id
		}
	}

	k.SetCount(ctx, count)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.NewGenesisState(
		k.GetSessions(ctx, 0, 0),
		k.GetParams(ctx),
	)
}
