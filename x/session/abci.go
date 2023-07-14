package session

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abcitypes.ValidatorUpdate {
	inactiveDuration := k.InactiveDuration(ctx)
	k.IterateSessionsForExpiryAt(ctx, ctx.BlockTime(), func(_ int, item types.Session) bool {
		k.DeleteSessionForExpiryAt(ctx, item.ExpiryAt, item.ID)

		if item.Status.Equal(hubtypes.StatusActive) {
			item.ExpiryAt = ctx.BlockTime().Add(inactiveDuration)
			k.SetSessionForExpiryAt(ctx, item.ExpiryAt, item.ID)

			item.Status = hubtypes.StatusInactivePending
			item.StatusAt = ctx.BlockTime()

			k.SetSession(ctx, item)
			ctx.EventManager().EmitTypedEvent(
				&types.EventUpdateStatus{
					ID:             item.ID,
					SubscriptionID: item.SubscriptionID,
					NodeAddress:    item.NodeAddress,
					Status:         item.Status,
				},
			)

			return false
		}

		var (
			accAddr  = item.GetAddress()
			nodeAddr = item.GetNodeAddress()
		)

		// TODO: call EndSessionHook

		k.DeleteSession(ctx, item.ID)
		k.DeleteSessionForAccount(ctx, accAddr, item.ID)
		k.DeleteSessionForNode(ctx, nodeAddr, item.ID)
		k.DeleteSessionForSubscription(ctx, item.SubscriptionID, item.ID)
		k.DeleteSessionForAllocation(ctx, item.SubscriptionID, accAddr, item.ID)
		ctx.EventManager().EmitTypedEvent(
			&types.EventUpdateStatus{
				ID:             item.ID,
				SubscriptionID: item.SubscriptionID,
				NodeAddress:    item.NodeAddress,
				Status:         item.Status,
			},
		)

		return false
	})

	return nil
}
