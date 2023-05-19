package subscription

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abcitypes.ValidatorUpdate {
	inactiveDuration := k.InactiveDuration(ctx)
	k.IterateSubscriptionsForExpiryAt(ctx, ctx.BlockTime(), func(_ int, item types.Subscription) bool {
		if item.GetStatus().Equal(hubtypes.StatusActive) {
			item.SetStatus(hubtypes.StatusInactivePending)
			item.SetStatusAt(ctx.BlockTime())
			k.SetSubscription(ctx, item)
			k.DeleteSubscriptionForExpiryAt(ctx, item.GetExpiryAt(), item.GetID())

			statusAt := item.GetStatusAt().Add(inactiveDuration)
			k.SetSubscriptionForExpiryAt(ctx, statusAt, item.GetID())
			ctx.EventManager().EmitTypedEvent(
				&types.EventUpdateStatus{
					ID:     item.GetID(),
					Status: item.GetStatus(),
				},
			)

			return false
		}

		k.DeleteSubscription(ctx, item.GetID())
		k.IterateQuotas(ctx, item.GetID(), func(_ int, quota types.Quota) bool {
			addr := quota.GetAddress()
			k.DeleteQuota(ctx, item.GetID(), addr)
			k.DeleteSubscriptionForAccount(ctx, addr, item.GetID())

			return false
		})

		statusAt := item.GetStatusAt().Add(inactiveDuration)
		k.DeleteSubscriptionForExpiryAt(ctx, statusAt, item.GetID())

		ctx.EventManager().EmitTypedEvent(
			&types.EventUpdateStatus{
				ID:     item.GetID(),
				Status: hubtypes.StatusInactive,
			},
		)

		return false
	})

	return nil
}
