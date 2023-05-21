package session

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state *types.GenesisState) {
	k.SetParams(ctx, state.Params)

	for _, item := range state.Sessions {
		var (
			accAddr  = item.GetAddress()
			nodeAddr = item.GetNodeAddress()
		)

		k.SetSession(ctx, item)
		k.SetSessionForAccount(ctx, accAddr, item.ID)
		k.SetSessionForNode(ctx, nodeAddr, item.ID)
		k.SetSessionForSubscription(ctx, item.SubscriptionID, item.ID)
		k.SetSessionForAllocation(ctx, item.SubscriptionID, accAddr, item.ID)
		k.SetSessionForExpiryAt(ctx, item.ExpiryAt, item.ID)
	}

	count := uint64(0)
	for _, item := range state.Sessions {
		if item.ID > count {
			count = item.ID
		}
	}

	k.SetCount(ctx, count)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.NewGenesisState(
		k.GetSessions(ctx),
		k.GetParams(ctx),
	)
}
