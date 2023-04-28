package subscription

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abcitypes.ValidatorUpdate {
	var (
		log              = k.Logger(ctx)
		inactiveDuration = k.InactiveDuration(ctx)
	)

	k.IterateSubscriptionExpirys(ctx, ctx.BlockTime(), func(_ int, item types.Subscription) bool {
		log.Info("found an inactive subscription", "id", item.GetID())

		if item.GetStatus().Equal(hubtypes.StatusActive) {
			item.SetStatus(hubtypes.StatusInactivePending)
			item.SetStatusAt(ctx.BlockTime())
			k.SetSubscription(ctx, item)
			k.DeleteSubscriptionExpiryAt(ctx, item.GetExpiryAt(), item.GetID())

			statusAt := item.GetStatusAt().Add(inactiveDuration)
			k.SetSubscriptionExpiryAt(ctx, statusAt, item.GetID())
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
		k.DeleteSubscriptionExpiryAt(ctx, statusAt, item.GetID())

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
