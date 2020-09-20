package session

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	end := ctx.BlockTime().Add(-1 * k.InactiveDuration(ctx))
	k.IterateActiveSessions(ctx, end, func(_ int, item types.Session) bool {
		k.Logger(ctx).Info("Inactive session", "id", item.ID,
			"subscription", item.Subscription, "node", item.Node, "address", item.Address)

		if err := process(ctx, k, item); err != nil {
			panic(err)
		}

		k.DeleteOngoingSession(ctx, item.Subscription, item.Address)
		k.DeleteActiveSessionAt(ctx, item.StatusAt, item.ID)

		item.Status = hub.StatusInactive
		item.StatusAt = ctx.BlockTime()
		k.SetSession(ctx, item)

		return false
	})

	return nil
}

func process(ctx sdk.Context, k keeper.Keeper, session types.Session) sdk.Error {
	subscription, found := k.GetSubscription(ctx, session.Subscription)
	if !found {
		return types.ErrorSubscriptionDoesNotExit()
	}

	quota, found := k.GetQuota(ctx, session.Subscription, session.Address)
	if !found {
		return types.ErrorQuotaDoesNotExist()
	}

	free := quota.Allocated.Sub(quota.Consumed)
	if !free.IsPositive() {
		return nil
	}

	if subscription.Plan == 0 {
		if subscription.Status.Equal(hub.StatusInactive) {
			return nil
		}

		bandwidth := hub.NewBandwidth(
			session.Bandwidth.Sum(), sdk.ZeroInt(),
		).CeilTo(
			hub.Gigabyte.Quo(subscription.Price.Amount),
		).Sum()
		if bandwidth.GT(free) {
			bandwidth = free
		}

		quota.Consumed = quota.Consumed.Add(bandwidth)
		k.SetQuota(ctx, session.Subscription, quota)

		amount := subscription.Amount(bandwidth)
		ctx.Logger().Info("", "price", subscription.Price, "deposit", subscription.Deposit,
			"consumed", session.Bandwidth.Sum(), "rounded", bandwidth, "amount", amount)

		return k.SendCoinsFromDepositToAccount(ctx, session.Address, session.Node.Bytes(), amount)
	}

	bandwidth := session.Bandwidth.Sum()
	if bandwidth.GT(free) {
		bandwidth = free
	}

	quota.Consumed = quota.Consumed.Add(bandwidth)
	k.SetQuota(ctx, session.Subscription, quota)

	ctx.Logger().Info("", "plan", subscription.Plan,
		"consumed", session.Bandwidth.Sum(), "rounded", bandwidth)

	return nil
}
