package subscription

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	log := k.Logger(ctx)

	end := ctx.BlockTime().Add(-1 * k.CancelDuration(ctx))
	k.IterateCancelSubscriptions(ctx, end, func(_ int, item types.Subscription) bool {
		log.Info("Cancel subscription", "id", item.ID,
			"owner", item.Owner, "plan", item.Plan, "node", item.Node)

		if item.Plan == 0 {
			consumed := sdk.ZeroInt()
			k.IterateQuotas(ctx, item.ID, func(_ int, item types.Quota) bool {
				consumed = consumed.Add(item.Consumed)
				return false
			})

			amount := item.Deposit.Sub(item.Amount(consumed))
			log.Info("", "price", item.Price,
				"deposit", item.Deposit, "consumed", consumed, "amount", amount)

			if err := k.SubtractDeposit(ctx, item.Owner, amount); err != nil {
				panic(err)
			}

			k.DeleteCancelSubscriptionAt(ctx, item.StatusAt, item.ID)
		} else {
			k.DeleteCancelSubscriptionAt(ctx, item.Expiry, item.ID)
		}

		item.Status = hub.StatusInactive
		item.StatusAt = ctx.BlockTime()
		k.SetSubscription(ctx, item)

		return false
	})

	return nil
}
