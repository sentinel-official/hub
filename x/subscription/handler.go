package subscription

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func startPlanSubscription(ctx sdk.Context, k keeper.Keeper, from sdk.AccAddress, id uint64, denom string) sdk.Result {
	plan, found := k.GetPlan(ctx, id)
	if !found {
		return types.ErrorPlanDoesNotExist().Result()
	}
	if !plan.Status.Equal(hub.StatusActive) {
		return types.ErrorInvalidPlanStatus().Result()
	}

	if plan.Price != nil {
		price, found := plan.PriceForDenom(denom)
		if !found {
			return types.ErrorPriceDoesNotExist().Result()
		}

		if err := k.SendCoin(ctx, from, plan.Provider.Bytes(), price); err != nil {
			return err.Result()
		}
	}

	count := k.GetSubscriptionsCount(ctx)
	subscription := types.Subscription{
		ID:          count + 1,
		Address:     from,
		Plan:        plan.ID,
		ExpiresAt:   ctx.BlockTime().Add(plan.Validity),
		Unallocated: hub.NewBandwidthFromInt64(0, 0),
		Status:      hub.StatusActive,
		StatusAt:    ctx.BlockTime(),
	}

	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForPlan(ctx, subscription.Plan, subscription.ID)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetSubscription,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", subscription.ID)),
		sdk.NewAttribute(types.AttributeKeyPlan, fmt.Sprintf("%d", subscription.Plan)),
		sdk.NewAttribute(types.AttributeKeyAddress, subscription.Address.String()),
	))

	quota := types.Quota{
		Address:   from,
		Consumed:  hub.NewBandwidthFromInt64(0, 0),
		Allocated: plan.Bandwidth,
	}

	k.SetQuotaForSubscription(ctx, subscription.ID, quota)
	k.SetSubscriptionForAddress(ctx, quota.Address, subscription.ID)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAddQuotaForSubscription,
		sdk.NewAttribute(types.AttributeKeyAddress, quota.Address.String()),
		sdk.NewAttribute(types.AttributeKeyConsumed, quota.Consumed.String()),
		sdk.NewAttribute(types.AttributeKeyAllocated, quota.Allocated.String()),
	))

	k.SetSubscriptionsCount(ctx, count+1)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetSubscriptionsCount,
		sdk.NewAttribute(types.AttributeKeyCount, fmt.Sprintf("%d", count+1)),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func startNodeSubscription(ctx sdk.Context, k keeper.Keeper, from sdk.AccAddress, address hub.NodeAddress, deposit sdk.Coin) sdk.Result {
	node, found := k.GetNode(ctx, address)
	if !found {
		return types.ErrorNodeDoesNotExist().Result()
	}
	if !node.Status.Equal(hub.StatusActive) {
		return types.ErrorInvalidNodeStatus().Result()
	}
	if node.Provider != nil {
		return types.ErrorCanNotSubscribe().Result()
	}

	price, found := node.PriceForDenom(deposit.Denom)
	if !found {
		return types.ErrorPriceDoesNotExist().Result()
	}

	if err := k.AddDeposit(ctx, from, deposit); err != nil {
		return err.Result()
	}

	count := k.GetSubscriptionsCount(ctx)
	subscription := types.Subscription{
		ID:          count + 1,
		Address:     from,
		Node:        address,
		Price:       price,
		Deposit:     deposit,
		Unallocated: hub.NewBandwidthFromInt64(0, 0),
		Status:      hub.StatusActive,
		StatusAt:    ctx.BlockTime(),
	}

	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForNode(ctx, subscription.Node, subscription.ID)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetSubscription,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", subscription.ID)),
		sdk.NewAttribute(types.AttributeKeyNode, subscription.Node.String()),
		sdk.NewAttribute(types.AttributeKeyAddress, subscription.Address.String()),
	))

	bandwidth, _ := node.BandwidthForCoin(deposit)
	quota := types.Quota{
		Address:   from,
		Consumed:  hub.NewBandwidthFromInt64(0, 0),
		Allocated: bandwidth,
	}

	k.SetQuotaForSubscription(ctx, subscription.ID, quota)
	k.SetSubscriptionForAddress(ctx, quota.Address, subscription.ID)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAddQuotaForSubscription,
		sdk.NewAttribute(types.AttributeKeyAddress, quota.Address.String()),
		sdk.NewAttribute(types.AttributeKeyConsumed, quota.Consumed.String()),
		sdk.NewAttribute(types.AttributeKeyAllocated, quota.Allocated.String()),
	))

	k.SetSubscriptionsCount(ctx, count+1)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetSubscriptionsCount,
		sdk.NewAttribute(types.AttributeKeyCount, fmt.Sprintf("%d", count+1)),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleStartSubscription(ctx sdk.Context, k keeper.Keeper, msg types.MsgStartSubscription) sdk.Result {
	if msg.ID == 0 {
		return startNodeSubscription(ctx, k, msg.From, msg.Address, msg.Deposit)
	}

	return startPlanSubscription(ctx, k, msg.From, msg.ID, msg.Denom)
}

func HandleAddQuotaForSubscription(ctx sdk.Context, k keeper.Keeper, msg types.MsgAddQuotaForSubscription) sdk.Result {
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return types.ErrorSubscriptionDoesNotExist().Result()
	}
	if !msg.From.Equals(subscription.Address) {
		return types.ErrorUnauthorized().Result()
	}
	if msg.Bandwidth.IsAnyGT(subscription.Unallocated) {
		return types.ErrorInvalidQuota().Result()
	}
	if !subscription.Status.Equal(hub.StatusActive) {
		return types.ErrorInvalidSubscriptionStatus().Result()
	}

	if k.HasQuotaForSubscription(ctx, subscription.ID, msg.Address) {
		return types.ErrorDuplicateQuota().Result()
	}

	subscription.Unallocated = subscription.Unallocated.Sub(msg.Bandwidth)
	k.SetSubscription(ctx, subscription)

	quota := types.Quota{
		Address:   msg.Address,
		Consumed:  hub.NewBandwidthFromInt64(0, 0),
		Allocated: msg.Bandwidth,
	}

	k.SetQuotaForSubscription(ctx, subscription.ID, quota)
	k.SetSubscriptionForAddress(ctx, quota.Address, subscription.ID)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAddQuotaForSubscription,
		sdk.NewAttribute(types.AttributeKeyAddress, quota.Address.String()),
		sdk.NewAttribute(types.AttributeKeyConsumed, quota.Consumed.String()),
		sdk.NewAttribute(types.AttributeKeyAllocated, quota.Allocated.String()),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleUpdateQuotaForSubscription(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpdateQuotaForSubscription) sdk.Result {
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return types.ErrorSubscriptionDoesNotExist().Result()
	}
	if !msg.From.Equals(subscription.Address) {
		return types.ErrorUnauthorized().Result()
	}
	if !subscription.Status.Equal(hub.StatusActive) {
		return types.ErrorInvalidSubscriptionStatus().Result()
	}

	quota, found := k.GetQuotaForSubscription(ctx, subscription.ID, msg.Address)
	if !found {
		return types.ErrorQuotaDoesNotExist().Result()
	}

	subscription.Unallocated = subscription.Unallocated.
		Add(quota.Allocated).Sub(quota.Consumed)
	if msg.Bandwidth.IsAnyGT(subscription.Unallocated) {
		return types.ErrorInvalidQuota().Result()
	}

	subscription.Unallocated = subscription.Unallocated.Sub(msg.Bandwidth)
	k.SetSubscription(ctx, subscription)

	quota.Allocated = msg.Bandwidth
	k.SetQuotaForSubscription(ctx, subscription.ID, quota)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAddQuotaForSubscription,
		sdk.NewAttribute(types.AttributeKeyAddress, quota.Address.String()),
		sdk.NewAttribute(types.AttributeKeyConsumed, quota.Consumed.String()),
		sdk.NewAttribute(types.AttributeKeyAllocated, quota.Allocated.String()),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleEndSubscription(ctx sdk.Context, k keeper.Keeper, msg types.MsgEndSubscription) sdk.Result {
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return types.ErrorSubscriptionDoesNotExist().Result()
	}
	if !msg.From.Equals(subscription.Address) {
		return types.ErrorUnauthorized().Result()
	}
	if !subscription.Status.Equal(hub.StatusActive) {
		return types.ErrorInvalidSubscriptionStatus().Result()
	}

	if subscription.Plan == 0 {
		consumed := hub.NewBandwidthFromInt64(0, 0)
		k.IterateQuotasForSubscription(ctx, subscription.ID, func(_ int, item types.Quota) bool {
			consumed = consumed.Add(item.Consumed)
			return false
		})

		amount := subscription.Deposit.Sub(subscription.Amount(consumed))
		if err := k.SubtractDeposit(ctx, subscription.Address, amount); err != nil {
			return err.Result()
		}
	}

	subscription.Status = hub.StatusInactive
	subscription.StatusAt = ctx.BlockTime()

	k.SetSubscription(ctx, subscription)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeEndSubscription,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", subscription.ID)),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}
