package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/types"
)

var (
	_ types.MsgServiceServer = (*msgServer)(nil)
)

type msgServer struct {
	Keeper
}

func NewMsgServiceServer(keeper Keeper) types.MsgServiceServer {
	return &msgServer{Keeper: keeper}
}

func (k *msgServer) MsgSubscribeToNode(c context.Context, msg *types.MsgSubscribeToNodeRequest) (*types.MsgSubscribeToNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	msgAddress, err := hubtypes.NodeAddressFromBech32(msg.Address)
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
	if !node.Status.Equal(hubtypes.StatusActive) {
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
			Status:   hubtypes.StatusActive,
			StatusAt: ctx.BlockTime(),
		}
		subscriptionNode = subscription.GetNode()
	)

	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForNode(ctx, subscriptionNode, subscription.Id)
	ctx.EventManager().EmitTypedEvent(
		&types.EventSubscribeToNode{
			Id:      subscription.Id,
			From:    sdk.AccAddress(msgFrom.Bytes()).String(),
			Owner:   subscription.Owner,
			Node:    subscription.Node,
			Price:   subscription.Price,
			Deposit: subscription.Deposit,
			Free:    subscription.Free,
		},
	)

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
	ctx.EventManager().EmitTypedEvent(
		&types.EventAddQuota{
			From:      subscription.Owner,
			Id:        subscription.Id,
			Address:   quota.Address,
			Consumed:  quota.Consumed,
			Allocated: quota.Allocated,
		},
	)

	k.SetCount(ctx, count+1)
	ctx.EventManager().EmitTypedEvent(
		&types.EventSetSubscriptionCount{
			Count: count + 1,
		},
	)

	ctx.EventManager().EmitTypedEvent(&types.EventModuleName)
	return &types.MsgSubscribeToNodeResponse{}, nil
}

func (k *msgServer) MsgSubscribeToPlan(c context.Context, msg *types.MsgSubscribeToPlanRequest) (*types.MsgSubscribeToPlanResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	plan, found := k.GetPlan(ctx, msg.Id)
	if !found {
		return nil, types.ErrorPlanDoesNotExist
	}
	if !plan.Status.Equal(hubtypes.StatusActive) {
		return nil, types.ErrorInvalidPlanStatus
	}

	msgFrom, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	if plan.Price != nil {
		price, found := plan.PriceForDenom(msg.Denom)
		if !found {
			return nil, types.ErrorPriceDoesNotExist
		}

		planProvider := plan.GetProvider()
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
			Denom:    msg.Denom,
			Expiry:   ctx.BlockTime().Add(plan.Validity),
			Free:     sdk.ZeroInt(),
			Status:   hubtypes.StatusActive,
			StatusAt: ctx.BlockTime(),
		}
	)

	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForPlan(ctx, subscription.Plan, subscription.Id)
	k.SetInactiveSubscriptionAt(ctx, subscription.Expiry, subscription.Id)
	ctx.EventManager().EmitTypedEvent(
		&types.EventSubscribeToPlan{
			Id:     subscription.Id,
			From:   sdk.AccAddress(msgFrom.Bytes()).String(),
			Owner:  subscription.Owner,
			Plan:   subscription.Plan,
			Denom:  subscription.Denom,
			Expiry: subscription.Expiry,
			Free:   subscription.Free,
		},
	)

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
	ctx.EventManager().EmitTypedEvent(
		&types.EventAddQuota{
			From:      subscription.Owner,
			Id:        subscription.Id,
			Address:   quota.Address,
			Consumed:  quota.Consumed,
			Allocated: quota.Allocated,
		},
	)

	k.SetCount(ctx, count+1)
	ctx.EventManager().EmitTypedEvent(
		&types.EventSetSubscriptionCount{
			Count: count + 1,
		},
	)

	ctx.EventManager().EmitTypedEvent(&types.EventModuleName)
	return &types.MsgSubscribeToPlanResponse{}, nil
}

func (k *msgServer) MsgCancel(c context.Context, msg *types.MsgCancelRequest) (*types.MsgCancelResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	msgFrom, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	subscription, found := k.GetSubscription(ctx, msg.Id)
	if !found {
		return nil, types.ErrorSubscriptionDoesNotExist
	}
	if msg.From != subscription.Owner {
		return nil, types.ErrorUnauthorized
	}
	if !subscription.Status.Equal(hubtypes.StatusActive) {
		return nil, types.ErrorInvalidSubscriptionStatus
	}

	subscription.StatusAt = ctx.BlockTime()
	if subscription.Plan == 0 {
		subscription.Status = hubtypes.StatusInactivePending
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

		subscription.Status = hubtypes.StatusInactive
		k.DeleteInactiveSubscriptionAt(ctx, subscription.Expiry, subscription.Id)
	}

	k.SetSubscription(ctx, subscription)
	ctx.EventManager().EmitTypedEvent(
		&types.EventCancelSubscription{
			From: sdk.AccAddress(msgFrom.Bytes()).String(),
			Id:   subscription.Id,
		},
	)

	ctx.EventManager().EmitTypedEvent(&types.EventModuleName)
	return &types.MsgCancelResponse{}, nil
}

func (k *msgServer) MsgAddQuota(c context.Context, msg *types.MsgAddQuotaRequest) (*types.MsgAddQuotaResponse, error) {
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
	if !subscription.Status.Equal(hubtypes.StatusActive) {
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
	ctx.EventManager().EmitTypedEvent(
		&types.EventAddQuota{
			From:      subscription.Owner,
			Id:        subscription.Id,
			Address:   quota.Address,
			Consumed:  quota.Consumed,
			Allocated: quota.Allocated,
		},
	)

	ctx.EventManager().EmitTypedEvent(&types.EventModuleName)
	return &types.MsgAddQuotaResponse{}, nil
}

func (k *msgServer) MsgUpdateQuota(c context.Context, msg *types.MsgUpdateQuotaRequest) (*types.MsgUpdateQuotaResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	subscription, found := k.GetSubscription(ctx, msg.Id)
	if !found {
		return nil, types.ErrorSubscriptionDoesNotExist
	}
	if msg.From != subscription.Owner {
		return nil, types.ErrorUnauthorized
	}
	if !subscription.Status.Equal(hubtypes.StatusActive) {
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
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateQuota{
			From:      subscription.Owner,
			Id:        subscription.Id,
			Address:   quota.Address,
			Consumed:  quota.Consumed,
			Allocated: quota.Allocated,
		},
	)

	ctx.EventManager().EmitTypedEvent(&types.EventModuleName)
	return &types.MsgUpdateQuotaResponse{}, nil
}
