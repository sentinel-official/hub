package subscription

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	var (
		log = k.Logger(ctx)
		end = ctx.BlockTime().Add(-1 * k.InactiveDuration(ctx))
	)

	k.IterateInactiveSubscriptions(ctx, end, func(_ int, item types.Subscription) bool {
		log.Info("Inactive subscription", "id", item.Id,
			"owner", item.Owner, "plan", item.Plan, "node", item.Node)

		if item.Plan == 0 {
			consumed := sdk.ZeroInt()
			k.IterateQuotas(ctx, item.Id, func(_ int, quota types.Quota) bool {
				consumed = consumed.Add(quota.Consumed)
				return false
			})

			amount := item.Deposit.Sub(item.Amount(consumed))
			log.Info("", "price", item.Price,
				"deposit", item.Deposit, "consumed", consumed, "amount", amount)

			itemOwner, err := sdk.AccAddressFromBech32(item.Owner)
			if err != nil {
				panic(err)
			}

			if err := k.SubtractDeposit(ctx, itemOwner, amount); err != nil {
				panic(err)
			}

			k.DeleteInactiveSubscriptionAt(ctx, item.StatusAt.Add(k.InactiveDuration(ctx)), item.Id)
		} else {
			k.DeleteInactiveSubscriptionAt(ctx, item.Expiry, item.Id)
		}

		k.IterateQuotas(ctx, item.Id, func(_ int, quota types.Quota) bool {
			var (
				quotaAddress = quota.GetAddress()
			)

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
