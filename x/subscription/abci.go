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

		if item.Plan == 0 {
			consumed := sdk.ZeroInt()
			k.IterateQuotas(ctx, item.Id, func(_ int, quota types.Quota) bool {
				consumed = consumed.Add(quota.Consumed)
				return false
			})

			amount := item.Deposit.Sub(item.Amount(consumed))
			log.Info("calculated refund of subscription", "id", item.Id,
				"consumed", consumed, "amount", amount)

			itemOwner, err := sdk.AccAddressFromBech32(item.Owner)
			if err != nil {
				panic(err)
			}

			if err := k.SubtractDeposit(ctx, itemOwner, amount); err != nil {
				panic(err)
			}
		} else {
			if item.Status.Equal(hubtypes.StatusActive) {
				item.Status = hubtypes.StatusInactivePending
				item.StatusAt = item.Expiry

				k.SetSubscription(ctx, item)
				k.DeleteInactiveSubscriptionAt(ctx, item.Expiry, item.Id)
				k.SetInactiveSubscriptionAt(ctx, item.Expiry.Add(inactiveDuration), item.Id)

				return false
			}
		}

		k.DeleteInactiveSubscriptionAt(ctx, item.StatusAt.Add(inactiveDuration), item.Id)
		k.IterateQuotas(ctx, item.Id, func(_ int, quota types.Quota) bool {
			quotaAddress := quota.GetAddress()
			k.DeleteActiveSubscriptionForAddress(ctx, quotaAddress, item.Id)
			k.SetInactiveSubscriptionForAddress(ctx, quotaAddress, item.Id)

			return false
		})

		item.Status = hubtypes.StatusInactive
		item.StatusAt = ctx.BlockTime()
		k.SetSubscription(ctx, item)

		return false
	})

	return nil
}
