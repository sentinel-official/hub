package subscription

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func HandleSubscribeToPlan(ctx sdk.Context, k keeper.Keeper, msg types.MsgSubscribeToPlan) (*sdk.Result, error) {
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.ErrorPlanDoesNotExist
	}
	if !plan.Status.Equal(hub.StatusActive) {
		return nil, types.ErrorInvalidPlanStatus
	}

	if plan.Price != nil {
		price, found := plan.PriceForDenom(msg.Denom)
		if !found {
			return nil, types.ErrorPriceDoesNotExist
		}

		if err := k.SendCoin(ctx, msg.From, plan.Provider.Bytes(), price); err != nil {
			return nil, err
		}
	}

	count := k.GetSubscriptionsCount(ctx)
	subscription := types.Subscription{
		ID:       count + 1,
		Owner:    msg.From,
		Plan:     plan.ID,
		Expiry:   ctx.BlockTime().Add(plan.Validity),
		Free:     sdk.ZeroInt(),
		Status:   hub.StatusActive,
		StatusAt: ctx.BlockTime(),
	}

	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForPlan(ctx, subscription.Plan, subscription.ID)
	k.SetCancelSubscriptionAt(ctx, subscription.Expiry, subscription.ID)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSet,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", subscription.ID)),
		sdk.NewAttribute(types.AttributeKeyPlan, fmt.Sprintf("%d", subscription.Plan)),
		sdk.NewAttribute(types.AttributeKeyOwner, subscription.Owner.String()),
	))

	quota := types.Quota{
		Address:   msg.From,
		Consumed:  sdk.ZeroInt(),
		Allocated: plan.Bytes,
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
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func HandleSubscribeToNode(ctx sdk.Context, k keeper.Keeper, msg types.MsgSubscribeToNode) (*sdk.Result, error) {
	node, found := k.GetNode(ctx, msg.Address)
	if !found {
		return nil, types.ErrorNodeDoesNotExist
	}
	if node.Provider != nil {
		return nil, types.ErrorCanNotSubscribe
	}
	if !node.Status.Equal(hub.StatusActive) {
		return nil, types.ErrorInvalidNodeStatus
	}

	price, found := node.PriceForDenom(msg.Deposit.Denom)
	if !found {
		return nil, types.ErrorPriceDoesNotExist
	}

	if err := k.AddDeposit(ctx, msg.From, msg.Deposit); err != nil {
		return nil, err
	}

	count := k.GetSubscriptionsCount(ctx)
	subscription := types.Subscription{
		ID:       count + 1,
		Owner:    msg.From,
		Node:     node.Address,
		Price:    price,
		Deposit:  msg.Deposit,
		Free:     sdk.ZeroInt(),
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

	bandwidth, _ := node.BytesForCoin(msg.Deposit)
	quota := types.Quota{
		Address:   msg.From,
		Consumed:  sdk.ZeroInt(),
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
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func HandleCancel(ctx sdk.Context, k keeper.Keeper, msg types.MsgCancel) (*sdk.Result, error) {
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.ErrorSubscriptionDoesNotExist
	}
	if !msg.From.Equals(subscription.Owner) {
		return nil, types.ErrorUnauthorized
	}
	if !subscription.Status.Equal(hub.StatusActive) {
		return nil, types.ErrorInvalidSubscriptionStatus
	}

	if subscription.Plan == 0 {
		subscription.Status = hub.StatusCancel
		subscription.StatusAt = ctx.BlockTime().Add(k.CancelDuration(ctx))

		k.SetCancelSubscriptionAt(ctx, subscription.StatusAt, subscription.ID)
	} else {
		subscription.Status = hub.StatusInactive
		subscription.StatusAt = ctx.BlockTime()

		k.DeleteCancelSubscriptionAt(ctx, subscription.Expiry, subscription.ID)
	}

	k.SetSubscription(ctx, subscription)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeCancel,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", subscription.ID)),
		sdk.NewAttribute(types.AttributeKeyStatus, subscription.Status.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func HandleAddQuota(ctx sdk.Context, k keeper.Keeper, msg types.MsgAddQuota) (*sdk.Result, error) {
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.ErrorSubscriptionDoesNotExist
	}
	if subscription.Plan == 0 {
		return nil, types.ErrorCanNotAddQuota
	}
	if !msg.From.Equals(subscription.Owner) {
		return nil, types.ErrorUnauthorized
	}
	if !subscription.Status.Equal(hub.StatusActive) {
		return nil, types.ErrorInvalidSubscriptionStatus
	}
	if k.HasQuota(ctx, subscription.ID, msg.Address) {
		return nil, types.ErrorDuplicateQuota
	}
	if msg.Bytes.GT(subscription.Free) {
		return nil, types.ErrorInvalidQuota
	}

	subscription.Free = subscription.Free.Sub(msg.Bytes)
	k.SetSubscription(ctx, subscription)

	quota := types.Quota{
		Address:   msg.Address,
		Consumed:  sdk.ZeroInt(),
		Allocated: msg.Bytes,
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
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func HandleUpdateQuota(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpdateQuota) (*sdk.Result, error) {
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.ErrorSubscriptionDoesNotExist
	}
	if !msg.From.Equals(subscription.Owner) {
		return nil, types.ErrorUnauthorized
	}
	if !subscription.Status.Equal(hub.StatusActive) {
		return nil, types.ErrorInvalidSubscriptionStatus
	}

	quota, found := k.GetQuota(ctx, subscription.ID, msg.Address)
	if !found {
		return nil, types.ErrorQuotaDoesNotExist
	}

	subscription.Free = subscription.Free.Add(quota.Allocated)
	if msg.Bytes.LT(quota.Consumed) || msg.Bytes.GT(subscription.Free) {
		return nil, types.ErrorInvalidQuota
	}

	subscription.Free = subscription.Free.Sub(msg.Bytes)
	k.SetSubscription(ctx, subscription)

	quota.Allocated = msg.Bytes
	k.SetQuota(ctx, subscription.ID, quota)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUpdateQuota,
		sdk.NewAttribute(types.AttributeKeyAddress, quota.Address.String()),
		sdk.NewAttribute(types.AttributeKeyConsumed, quota.Consumed.String()),
		sdk.NewAttribute(types.AttributeKeyAllocated, quota.Allocated.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
