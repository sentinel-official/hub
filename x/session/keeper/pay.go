package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/types"
)

func (k Keeper) Pay(ctx sdk.Context, session types.Session) sdk.Error {
	subscription, found := k.GetSubscription(ctx, session.Subscription)
	if !found {
		return types.ErrorSubscriptionDoesNotExit()
	}
	if subscription.Plan > 0 || subscription.Status.Equal(hub.StatusInactive) {
		return nil
	}

	quota, found := k.GetQuota(ctx, session.Subscription, session.Address)
	if !found {
		return types.ErrorQuotaDoesNotExist()
	}

	var (
		free      = quota.Allocated.Sub(quota.Consumed)
		bandwidth = hub.NewBandwidth(session.Bandwidth.Sum(), sdk.ZeroInt()).
				CeilTo(hub.Gigabyte.Quo(subscription.Price.Amount)).Sum()
	)

	if free.IsZero() {
		return nil
	}
	if bandwidth.GT(free) {
		bandwidth = free
	}

	quota.Consumed = quota.Consumed.Add(bandwidth)
	k.SetQuota(ctx, session.Subscription, quota)

	amount := subscription.Amount(bandwidth)
	ctx.Logger().Info("", "price", subscription.Price, "deposit", subscription.Deposit,
		"consumed", session.Bandwidth.Sum(), "ceil", bandwidth, "amount", amount)

	return k.SendCoinsFromDepositToAccount(ctx, session.Address, session.Node.Bytes(), amount)
}
