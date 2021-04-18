package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/types"
)

var (
	_ types.MsgServiceServer = server{}
)

type server struct {
	Keeper
}

func NewMsgServiceServer(keeper Keeper) types.MsgServiceServer {
	return &server{Keeper: keeper}
}

func (k server) MsgSubscribeToNode(c context.Context, msg *types.MsgSubscribeToNodeRequest) (*types.MsgSubscribeToNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	msgAddress, err := hub.NodeAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, msgAddress)
	if !found {
		return nil, types.ErrorNodeDoesNotExist
	}
	if node.Provider != "" {
		return nil, types.ErrorCanNotSubscribe
	}
	if !node.Status.Equal(hub.StatusActive) {
		return nil, types.ErrorInvalidNodeStatus
	}

	price, found := node.PriceForDenom(msg.Deposit.Denom)
	if !found {
		return nil, types.ErrorPriceDoesNotExist
	}

	msgFrom, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	if err := k.AddDeposit(ctx, msgFrom, msg.Deposit); err != nil {
		return nil, err
	}

	var (
		count        = k.GetCount(ctx)
		subscription = types.Subscription{
			Id:       count + 1,
			Owner:    msg.From,
			Node:     node.Address,
			Price:    price,
			Deposit:  msg.Deposit,
			Free:     sdk.ZeroInt(),
			Status:   hub.StatusActive,
			StatusAt: ctx.BlockTime(),
		}
		subscriptionNode = subscription.GetNode()
	)

	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForNode(ctx, subscriptionNode, subscription.Id)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSet,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", subscription.Id)),
		sdk.NewAttribute(types.AttributeKeyNode, subscription.Node),
		sdk.NewAttribute(types.AttributeKeyOwner, subscription.Owner),
	))

	var (
		bandwidth, _ = node.BytesForCoin(msg.Deposit)
		quota        = types.Quota{
			Address:   msg.From,
			Consumed:  sdk.ZeroInt(),
			Allocated: bandwidth,
		}
		quotaAddress = quota.GetAddress()
	)

	k.SetQuota(ctx, subscription.Id, quota)
	k.SetActiveSubscriptionForAddress(ctx, quotaAddress, subscription.Id)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAddQuota,
		sdk.NewAttribute(types.AttributeKeyAddress, quota.Address),
		sdk.NewAttribute(types.AttributeKeyConsumed, quota.Consumed.String()),
		sdk.NewAttribute(types.AttributeKeyAllocated, quota.Allocated.String()),
	))

	k.SetCount(ctx, count+1)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetCount,
		sdk.NewAttribute(types.AttributeKeyCount, fmt.Sprintf("%d", count+1)),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &types.MsgSubscribeToNodeResponse{}, nil
}

func (k server) MsgSubscribeToPlan(c context.Context, msg *types.MsgSubscribeToPlanRequest) (*types.MsgSubscribeToPlanResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	plan, found := k.GetPlan(ctx, msg.Id)
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

		msgFrom, err := sdk.AccAddressFromBech32(msg.From)
		if err != nil {
			return nil, err
		}

		var (
			planProvider = plan.GetProvider()
		)

		if err := k.SendCoin(ctx, msgFrom, planProvider.Bytes(), price); err != nil {
			return nil, err
		}
	}

	var (
		count        = k.GetCount(ctx)
		subscription = types.Subscription{
			Id:       count + 1,
			Owner:    msg.From,
			Plan:     plan.Id,
			Expiry:   ctx.BlockTime().Add(plan.Validity),
			Free:     sdk.ZeroInt(),
			Status:   hub.StatusActive,
			StatusAt: ctx.BlockTime(),
		}
	)

	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForPlan(ctx, subscription.Plan, subscription.Id)
	k.SetInactiveSubscriptionAt(ctx, subscription.Expiry, subscription.Id)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSet,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", subscription.Id)),
		sdk.NewAttribute(types.AttributeKeyPlan, fmt.Sprintf("%d", subscription.Plan)),
		sdk.NewAttribute(types.AttributeKeyOwner, subscription.Owner),
	))

	var (
		quota = types.Quota{
			Address:   msg.From,
			Consumed:  sdk.ZeroInt(),
			Allocated: plan.Bytes,
		}
		quotaAddress = quota.GetAddress()
	)

	k.SetQuota(ctx, subscription.Id, quota)
	k.SetActiveSubscriptionForAddress(ctx, quotaAddress, subscription.Id)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAddQuota,
		sdk.NewAttribute(types.AttributeKeyAddress, quota.Address),
		sdk.NewAttribute(types.AttributeKeyConsumed, quota.Consumed.String()),
		sdk.NewAttribute(types.AttributeKeyAllocated, quota.Allocated.String()),
	))

	k.SetCount(ctx, count+1)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetCount,
		sdk.NewAttribute(types.AttributeKeyCount, fmt.Sprintf("%d", count+1)),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &types.MsgSubscribeToPlanResponse{}, nil
}

func (k server) MsgCancel(c context.Context, msg *types.MsgCancelRequest) (*types.MsgCancelResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	subscription, found := k.GetSubscription(ctx, msg.Id)
	if !found {
		return nil, types.ErrorSubscriptionDoesNotExist
	}
	if msg.From != subscription.Owner {
		return nil, types.ErrorUnauthorized
	}
	if !subscription.Status.Equal(hub.StatusActive) {
		return nil, types.ErrorInvalidSubscriptionStatus
	}

	subscription.StatusAt = ctx.BlockTime()
	if subscription.Plan == 0 {
		subscription.Status = hub.StatusInactivePending
		k.SetInactiveSubscriptionAt(ctx, subscription.StatusAt.Add(k.InactiveDuration(ctx)), subscription.Id)
	} else {
		k.IterateQuotas(ctx, subscription.Id, func(_ int, quota types.Quota) bool {
			var (
				quotaAddress = quota.GetAddress()
			)

			k.DeleteActiveSubscriptionForAddress(ctx, quotaAddress, subscription.Id)
			k.SetInactiveSubscriptionForAddress(ctx, quotaAddress, subscription.Id)
			return false
		})

		subscription.Status = hub.StatusInactive
		k.DeleteInactiveSubscriptionAt(ctx, subscription.Expiry, subscription.Id)
	}

	k.SetSubscription(ctx, subscription)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeCancel,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", subscription.Id)),
		sdk.NewAttribute(types.AttributeKeyStatus, subscription.Status.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &types.MsgCancelResponse{}, nil
}

func (k server) MsgAddQuota(c context.Context, msg *types.MsgAddQuotaRequest) (*types.MsgAddQuotaResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	msgAddress, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	subscription, found := k.GetSubscription(ctx, msg.Id)
	if !found {
		return nil, types.ErrorSubscriptionDoesNotExist
	}
	if subscription.Plan == 0 {
		return nil, types.ErrorCanNotAddQuota
	}
	if msg.From != subscription.Owner {
		return nil, types.ErrorUnauthorized
	}
	if !subscription.Status.Equal(hub.StatusActive) {
		return nil, types.ErrorInvalidSubscriptionStatus
	}
	if k.HasQuota(ctx, subscription.Id, msgAddress) {
		return nil, types.ErrorDuplicateQuota
	}
	if msg.Bytes.GT(subscription.Free) {
		return nil, types.ErrorInvalidQuota
	}

	subscription.Free = subscription.Free.Sub(msg.Bytes)
	k.SetSubscription(ctx, subscription)

	var (
		quota = types.Quota{
			Address:   msg.Address,
			Consumed:  sdk.ZeroInt(),
			Allocated: msg.Bytes,
		}
		quotaAddress = quota.GetAddress()
	)

	k.SetQuota(ctx, subscription.Id, quota)
	k.SetActiveSubscriptionForAddress(ctx, quotaAddress, subscription.Id)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAddQuota,
		sdk.NewAttribute(types.AttributeKeyAddress, quota.Address),
		sdk.NewAttribute(types.AttributeKeyConsumed, quota.Consumed.String()),
		sdk.NewAttribute(types.AttributeKeyAllocated, quota.Allocated.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &types.MsgAddQuotaResponse{}, nil
}

func (k server) MsgUpdateQuota(c context.Context, msg *types.MsgUpdateQuotaRequest) (*types.MsgUpdateQuotaResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	subscription, found := k.GetSubscription(ctx, msg.Id)
	if !found {
		return nil, types.ErrorSubscriptionDoesNotExist
	}
	if msg.From != subscription.Owner {
		return nil, types.ErrorUnauthorized
	}
	if !subscription.Status.Equal(hub.StatusActive) {
		return nil, types.ErrorInvalidSubscriptionStatus
	}

	msgAddress, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	quota, found := k.GetQuota(ctx, subscription.Id, msgAddress)
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
	k.SetQuota(ctx, subscription.Id, quota)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUpdateQuota,
		sdk.NewAttribute(types.AttributeKeyAddress, quota.Address),
		sdk.NewAttribute(types.AttributeKeyConsumed, quota.Consumed.String()),
		sdk.NewAttribute(types.AttributeKeyAllocated, quota.Allocated.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &types.MsgUpdateQuotaResponse{}, nil
}
