package subscription

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func HandleSubscribeToPlan(ctx sdk.Context, k keeper.Keeper, msg types.MsgSubscribeToPlan) sdk.Result {
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return types.ErrorPlanDoesNotExist().Result()
	}
	if !plan.Status.Equal(hub.StatusActive) {
		return types.ErrorInvalidPlanStatus().Result()
	}

	if plan.Price != nil {
		price, found := plan.PriceForDenom(msg.Denom)
		if !found {
			return types.ErrorPriceDoesNotExist().Result()
		}

		if err := k.SendCoin(ctx, msg.From, plan.Provider.Bytes(), price); err != nil {
			return err.Result()
		}
	}

	count := k.GetSubscriptionsCount(ctx)
	subscription := types.Subscription{
		ID:       count + 1,
		Owner:    msg.From,
		Plan:     plan.ID,
		Expiry:   ctx.BlockTime().Add(plan.Validity),
		Free:     hub.NewBandwidthFromInt64(0, 0),
		Status:   hub.StatusActive,
		StatusAt: ctx.BlockTime(),
	}

	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForPlan(ctx, subscription.Plan, subscription.ID)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSet,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", subscription.ID)),
		sdk.NewAttribute(types.AttributeKeyPlan, fmt.Sprintf("%d", subscription.Plan)),
		sdk.NewAttribute(types.AttributeKeyOwner, subscription.Owner.String()),
	))

	quota := types.Quota{
		Address:   msg.From,
		Consumed:  hub.NewBandwidthFromInt64(0, 0),
		Allocated: plan.Bandwidth,
	}

	k.SetQuota(ctx, subscription.ID, quota)
	k.SetSubscriptionForAddress(ctx, quota.Address, subscription.ID)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAddQuota,
		sdk.NewAttribute(types.AttributeKeyAddress, quota.Address.String()),
		sdk.NewAttribute(types.AttributeKeyConsumed, quota.Consumed.String()),
		sdk.NewAttribute(types.AttributeKeyAllocated, quota.Allocated.String()),
	))

	k.SetSubscriptionsCount(ctx, count+1)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetCount,
		sdk.NewAttribute(types.AttributeKeyCount, fmt.Sprintf("%d", count+1)),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleSubscribeToNode(ctx sdk.Context, k keeper.Keeper, msg types.MsgSubscribeToNode) sdk.Result {
	node, found := k.GetNode(ctx, msg.Address)
	if !found {
		return types.ErrorNodeDoesNotExist().Result()
	}
	if !node.Status.Equal(hub.StatusActive) {
		return types.ErrorInvalidNodeStatus().Result()
	}
	if node.Provider != nil {
		return types.ErrorCanNotSubscribe().Result()
	}

	price, found := node.PriceForDenom(msg.Deposit.Denom)
	if !found {
		return types.ErrorPriceDoesNotExist().Result()
	}

	if err := k.AddDeposit(ctx, msg.From, msg.Deposit); err != nil {
		return err.Result()
	}

	count := k.GetSubscriptionsCount(ctx)
	subscription := types.Subscription{
		ID:       count + 1,
		Owner:    msg.From,
		Node:     node.Address,
		Price:    price,
		Deposit:  msg.Deposit,
		Free:     hub.NewBandwidthFromInt64(0, 0),
		Status:   hub.StatusActive,
		StatusAt: ctx.BlockTime(),
	}

	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForNode(ctx, subscription.Node, subscription.ID)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSet,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", subscription.ID)),
		sdk.NewAttribute(types.AttributeKeyNode, subscription.Node.String()),
		sdk.NewAttribute(types.AttributeKeyOwner, subscription.Owner.String()),
	))

	bandwidth, _ := node.BandwidthForCoin(msg.Deposit)
	quota := types.Quota{
		Address:   msg.From,
		Consumed:  hub.NewBandwidthFromInt64(0, 0),
		Allocated: bandwidth,
	}

	k.SetQuota(ctx, subscription.ID, quota)
	k.SetSubscriptionForAddress(ctx, quota.Address, subscription.ID)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAddQuota,
		sdk.NewAttribute(types.AttributeKeyAddress, quota.Address.String()),
		sdk.NewAttribute(types.AttributeKeyConsumed, quota.Consumed.String()),
		sdk.NewAttribute(types.AttributeKeyAllocated, quota.Allocated.String()),
	))

	k.SetSubscriptionsCount(ctx, count+1)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetCount,
		sdk.NewAttribute(types.AttributeKeyCount, fmt.Sprintf("%d", count+1)),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleCancel(ctx sdk.Context, k keeper.Keeper, msg types.MsgCancel) sdk.Result {
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return types.ErrorSubscriptionDoesNotExist().Result()
	}
	if !msg.From.Equals(subscription.Owner) {
		return types.ErrorUnauthorized().Result()
	}
	if !subscription.Status.Equal(hub.StatusActive) {
		return types.ErrorInvalidSubscriptionStatus().Result()
	}

	subscription.Status = hub.StatusCancel
	subscription.StatusAt = ctx.BlockTime()

	k.SetSubscription(ctx, subscription)
	k.SetCancelSubscriptionAt(ctx, subscription.StatusAt, subscription.ID)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeCancel,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", subscription.ID)),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleAddQuota(ctx sdk.Context, k keeper.Keeper, msg types.MsgAddQuota) sdk.Result {
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return types.ErrorSubscriptionDoesNotExist().Result()
	}
	if !msg.From.Equals(subscription.Owner) {
		return types.ErrorUnauthorized().Result()
	}
	if !subscription.Status.Equal(hub.StatusActive) {
		return types.ErrorInvalidSubscriptionStatus().Result()
	}
	if msg.Bandwidth.IsAnyGT(subscription.Free) {
		return types.ErrorInvalidQuota().Result()
	}

	if k.HasQuota(ctx, subscription.ID, msg.Address) {
		return types.ErrorDuplicateQuota().Result()
	}

	subscription.Free = subscription.Free.Sub(msg.Bandwidth)
	k.SetSubscription(ctx, subscription)

	quota := types.Quota{
		Address:   msg.Address,
		Consumed:  hub.NewBandwidthFromInt64(0, 0),
		Allocated: msg.Bandwidth,
	}

	k.SetQuota(ctx, subscription.ID, quota)
	k.SetSubscriptionForAddress(ctx, quota.Address, subscription.ID)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAddQuota,
		sdk.NewAttribute(types.AttributeKeyAddress, quota.Address.String()),
		sdk.NewAttribute(types.AttributeKeyConsumed, quota.Consumed.String()),
		sdk.NewAttribute(types.AttributeKeyAllocated, quota.Allocated.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleUpdateQuota(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpdateQuota) sdk.Result {
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return types.ErrorSubscriptionDoesNotExist().Result()
	}
	if !msg.From.Equals(subscription.Owner) {
		return types.ErrorUnauthorized().Result()
	}
	if !subscription.Status.Equal(hub.StatusActive) {
		return types.ErrorInvalidSubscriptionStatus().Result()
	}

	quota, found := k.GetQuota(ctx, subscription.ID, msg.Address)
	if !found {
		return types.ErrorQuotaDoesNotExist().Result()
	}

	subscription.Free = subscription.Free.
		Add(quota.Allocated).Sub(quota.Consumed)
	if msg.Bandwidth.IsAnyGT(subscription.Free) ||
		quota.Consumed.IsAnyGT(msg.Bandwidth) {
		return types.ErrorInvalidQuota().Result()
	}

	subscription.Free = subscription.Free.
		Add(quota.Consumed).Sub(msg.Bandwidth)
	k.SetSubscription(ctx, subscription)

	quota.Allocated = msg.Bandwidth
	k.SetQuota(ctx, subscription.ID, quota)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUpdateQuota,
		sdk.NewAttribute(types.AttributeKeyAddress, quota.Address.String()),
		sdk.NewAttribute(types.AttributeKeyConsumed, quota.Consumed.String()),
		sdk.NewAttribute(types.AttributeKeyAllocated, quota.Allocated.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return sdk.Result{Events: ctx.EventManager().Events()}
}
