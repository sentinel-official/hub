package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/types"
	subscriptiontypes "github.com/sentinel-official/hub/x/subscription/types"
)

func (k *Keeper) ProcessPaymentAndUpdateQuota(ctx sdk.Context, session types.Session) error {
	from, err := sdk.AccAddressFromBech32(session.Address)
	if err != nil {
		return err
	}

	subscription, found := k.GetSubscription(ctx, session.Subscription)
	if !found {
		return types.ErrorSubscriptionDoesNotExit
	}
	if subscription.Status.Equal(hubtypes.StatusInactive) {
		return nil
	}

	quota, found := k.GetQuota(ctx, session.Subscription, from)
	if !found {
		return types.ErrorQuotaDoesNotExist
	}

	available := quota.Allocated.Sub(quota.Consumed)
	if !available.IsPositive() {
		return nil
	}

	if subscription.Plan == 0 {
		bandwidth := hubtypes.NewBandwidth(
			session.Bandwidth.Sum(), sdk.ZeroInt(),
		).CeilTo(
			hubtypes.Gigabyte.Quo(subscription.Price.Amount),
		).Sum()
		if bandwidth.GT(available) {
			bandwidth = available
		}

		quota.Consumed = quota.Consumed.Add(bandwidth)
		k.SetQuota(ctx, session.Subscription, quota)
		ctx.EventManager().EmitTypedEvent(
			&subscriptiontypes.EventUpdateQuota{
				Id:        session.Subscription,
				Address:   quota.Address,
				Consumed:  quota.Consumed,
				Allocated: quota.Allocated,
			},
		)

		amount := subscription.Amount(bandwidth)
		ctx.Logger().Info("calculated payment for session", "id", session.Id,
			"price", subscription.Price, "deposit", subscription.Deposit, "amount", amount,
			"consumed", session.Bandwidth.Sum(), "rounded", bandwidth)

		sessionNode := session.GetNode()
		if err := k.SendCoinsFromDepositToAccount(ctx, from, sessionNode.Bytes(), amount); err != nil {
			return err
		}

		ctx.EventManager().EmitTypedEvent(
			&types.EventPay{
				Id:           session.Id,
				Subscription: session.Subscription,
				Node:         session.Node,
				Payment:      amount,
			},
		)

		return nil
	}

	bandwidth := session.Bandwidth.Sum()
	if bandwidth.GT(available) {
		bandwidth = available
	}

	quota.Consumed = quota.Consumed.Add(bandwidth)
	k.SetQuota(ctx, session.Subscription, quota)
	ctx.EventManager().EmitTypedEvent(
		&subscriptiontypes.EventUpdateQuota{
			Id:        session.Subscription,
			Address:   quota.Address,
			Consumed:  quota.Consumed,
			Allocated: quota.Allocated,
		},
	)

	ctx.Logger().Info("calculated bandwidth for session", "id", session.Id,
		"plan", subscription.Plan, "consumed", session.Bandwidth.Sum(), "rounded", bandwidth)
	return nil
}
