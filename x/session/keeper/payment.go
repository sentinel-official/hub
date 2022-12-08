package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	hubutils "github.com/sentinel-official/hub/utils"
	"github.com/sentinel-official/hub/x/session/types"
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
		consumed := hubtypes.NewBandwidth(
			session.Bandwidth.Sum(), sdk.ZeroInt(),
		).CeilTo(
			hubtypes.Gigabyte.Quo(subscription.Price.Amount),
		).Sum()
		if consumed.GT(available) {
			consumed = available
		}

		quota.Consumed = quota.Consumed.Add(consumed)
		k.SetQuota(ctx, session.Subscription, quota)

		var (
			amount        = subscription.Amount(consumed)
			nodeAddr      = session.GetNode()
			stakingShare  = k.node.StakingShare(ctx)
			stakingReward = hubutils.GetProportionOfCoin(amount, stakingShare)
		)

		if err := k.SendCoinFromDepositToModule(ctx, from, k.feeCollectorName, stakingReward); err != nil {
			return err
		}

		amount = amount.Sub(stakingReward)
		ctx.Logger().Info("processing the payment for session", "id", session.Id,
			"consumed", consumed, "to_address", nodeAddr, "amount", amount)

		ctx.EventManager().EmitTypedEvent(
			&types.EventPay{
				Id:           session.Id,
				Node:         session.Node,
				Subscription: session.Subscription,
				Amount:       amount,
			},
		)

		return k.SendCoinFromDepositToAccount(ctx, from, nodeAddr.Bytes(), amount)
	}

	consumed := session.Bandwidth.Sum()
	if consumed.GT(available) {
		consumed = available
	}

	quota.Consumed = quota.Consumed.Add(consumed)
	k.SetQuota(ctx, session.Subscription, quota)

	return nil
}
