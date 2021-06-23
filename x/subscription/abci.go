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

	k.IterateInactiveSubscriptions(ctx, ctx.BlockTime(), func(_ int, key []byte, item types.Subscription) bool {
		log.Info("inactive subscription", "key", key, "value", item)

		if item.Status.Equal(hubtypes.StatusActive) {
			k.DeleteInactiveSubscriptionAt(ctx, item.Expiry, item.Id)
			k.IterateQuotas(ctx, item.Id, func(_ int, quota types.Quota) bool {
				address := quota.GetAddress()
				k.DeleteActiveSubscriptionForAddress(ctx, address, item.Id)
				k.SetInactiveSubscriptionForAddress(ctx, address, item.Id)

				return false
			})

			item.Status = hubtypes.StatusInactivePending
			item.StatusAt = ctx.BlockTime()

			k.SetSubscription(ctx, item)
			k.SetInactiveSubscriptionAt(ctx, item.StatusAt.Add(inactiveDuration), item.Id)
			ctx.EventManager().EmitTypedEvent(
				&types.EventCancelSubscription{
					Id: item.Id,
				},
			)

			return false
		}

		if item.Plan == 0 {
			consumed := sdk.ZeroInt()
			k.IterateQuotas(ctx, item.Id, func(_ int, quota types.Quota) bool {
				consumed = consumed.Add(quota.Consumed)
				return false
			})

			amount := item.Deposit.Sub(item.Amount(consumed))
			log.Info("calculated refund of subscription", "id", item.Id,
				"consumed", consumed, "amount", amount)

			itemOwner := item.GetOwner()
			if err := k.SubtractDeposit(ctx, itemOwner, amount); err != nil {
				log.Error("failed to subtract the deposit", "cause", err)
			}
		}

		k.DeleteInactiveSubscriptionAt(ctx, item.StatusAt.Add(inactiveDuration), item.Id)

		item.Status = hubtypes.StatusInactive
		item.StatusAt = ctx.BlockTime()
		k.SetSubscription(ctx, item)

		return false
	})

	return nil
}
