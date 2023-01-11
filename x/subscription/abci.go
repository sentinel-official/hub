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

	k.IterateInactiveSubscriptions(ctx, ctx.BlockTime(), func(_ int, item types.Subscription) bool {
		log.Info("found an inactive subscription", "id", item.Id)

		if item.Status.Equal(hubtypes.StatusActive) {
			k.DeleteInactiveSubscriptionAt(ctx, item.Expiry, item.Id)
			k.IterateQuotas(ctx, item.Id, func(_ int, quota types.Quota) bool {
				accAddr := quota.GetAddress()
				k.DeleteActiveSubscriptionForAddress(ctx, accAddr, item.Id)
				k.SetInactiveSubscriptionForAddress(ctx, accAddr, item.Id)

				return false
			})

			item.Status = hubtypes.StatusInactivePending
			item.StatusAt = ctx.BlockTime()

			k.SetSubscription(ctx, item)
			k.SetInactiveSubscriptionAt(ctx, item.StatusAt.Add(inactiveDuration), item.Id)
			ctx.EventManager().EmitTypedEvent(
				&types.EventSetStatus{
					Id:     item.Id,
					Status: item.Status,
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

			var (
				amount    = item.Deposit.Sub(item.Amount(consumed))
				ownerAddr = item.GetOwner()
			)

			log.Info("releasing the amount for subscription", "id", item.Id,
				"consumed", consumed, "to_address", ownerAddr, "amount", amount)

			if err := k.SubtractDeposit(ctx, ownerAddr, amount); err != nil {
				log.Error("error occurred while releasing the amount", "cause", err)
			}
		}

		k.DeleteSubscription(ctx, item.Id)
		k.IterateQuotas(ctx, item.Id, func(_ int, quota types.Quota) bool {
			accAddr := quota.GetAddress()
			k.DeleteQuota(ctx, item.Id, accAddr)
			k.DeleteInactiveSubscriptionForAddress(ctx, accAddr, item.Id)

			return false
		})
		k.DeleteInactiveSubscriptionAt(ctx, item.StatusAt.Add(inactiveDuration), item.Id)

		ctx.EventManager().EmitTypedEvent(
			&types.EventSetStatus{
				Id:     item.Id,
				Status: item.Status,
			},
		)

		return false
	})

	return nil
}
