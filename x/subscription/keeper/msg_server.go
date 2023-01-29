package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	hubutils "github.com/sentinel-official/hub/utils"
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

	nodeAddr, err := hubtypes.NodeAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, nodeAddr)
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

	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	if err := k.AddDeposit(ctx, fromAddr, msg.Deposit); err != nil {
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
	)

	k.SetCount(ctx, count+1)
	k.SetSubscription(ctx, subscription)
	ctx.EventManager().EmitTypedEvent(
		&types.EventSubscribe{
			Id:     subscription.Id,
			Node:   subscription.Node,
			Amount: msg.Deposit,
		},
	)

	var (
		bandwidth, _ = node.BytesForCoin(msg.Deposit)
		quota        = types.Quota{
			Address:   msg.From,
			Consumed:  sdk.ZeroInt(),
			Allocated: bandwidth,
		}
	)

	k.SetQuota(ctx, subscription.Id, quota)
	k.SetActiveSubscriptionForAddress(ctx, fromAddr, subscription.Id)
	ctx.EventManager().EmitTypedEvent(
		&types.EventAddQuota{
			Id:      subscription.Id,
			Address: quota.Address,
		},
	)

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

	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	var (
		stakingReward sdk.Coin
		payment       sdk.Coin
	)

	if plan.Price != nil {
		price, found := plan.PriceForDenom(msg.Denom)
		if !found {
			return nil, types.ErrorPriceDoesNotExist
		}

		stakingShare := k.provider.StakingShare(ctx)
		stakingReward = hubutils.GetProportionOfCoin(price, stakingShare)

		if err := k.SendCoinFromAccountToModule(ctx, fromAddr, k.feeCollectorName, stakingReward); err != nil {
			return nil, err
		}

		provAddr := plan.GetProvider()
		payment = price.Sub(stakingReward)

		if err := k.SendCoin(ctx, fromAddr, provAddr.Bytes(), payment); err != nil {
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

	ctx.EventManager().EmitTypedEvent(
		&types.EventStakingReward{
			Id:     subscription.Id,
			Node:   subscription.Node,
			Plan:   subscription.Plan,
			Amount: stakingReward,
		},
	)

	k.SetCount(ctx, count+1)
	k.SetSubscription(ctx, subscription)
	k.SetInactiveSubscriptionAt(ctx, subscription.Expiry, subscription.Id)
	ctx.EventManager().EmitTypedEvent(
		&types.EventSubscribe{
			Id:     subscription.Id,
			Plan:   subscription.Plan,
			Amount: payment,
		},
	)

	var (
		quota = types.Quota{
			Address:   msg.From,
			Consumed:  sdk.ZeroInt(),
			Allocated: plan.Bytes,
		}
	)

	k.SetQuota(ctx, subscription.Id, quota)
	k.SetActiveSubscriptionForAddress(ctx, fromAddr, subscription.Id)
	ctx.EventManager().EmitTypedEvent(
		&types.EventAddQuota{
			Id:      subscription.Id,
			Address: quota.Address,
		},
	)

	return &types.MsgSubscribeToPlanResponse{}, nil
}

func (k *msgServer) MsgCancel(c context.Context, msg *types.MsgCancelRequest) (*types.MsgCancelResponse, error) {
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

	activeSessionExist := false
	k.IterateQuotas(ctx, subscription.Id, func(_ int, quota types.Quota) bool {
		var (
			accAddr = quota.GetAddress()
			items   = k.GetActiveSessionsForAddress(ctx, accAddr, 0, 1)
		)

		if len(items) > 0 {
			activeSessionExist = true
		}

		return activeSessionExist
	})

	if activeSessionExist {
		return nil, types.ErrorCanNotCancel
	}

	inactiveDuration := k.InactiveDuration(ctx)
	if subscription.Plan > 0 {
		k.DeleteInactiveSubscriptionAt(ctx, subscription.Expiry, subscription.Id)
	}

	k.IterateQuotas(ctx, subscription.Id, func(_ int, quota types.Quota) bool {
		accAddr := quota.GetAddress()
		k.DeleteActiveSubscriptionForAddress(ctx, accAddr, subscription.Id)
		k.SetInactiveSubscriptionForAddress(ctx, accAddr, subscription.Id)

		return false
	})

	subscription.Status = hubtypes.StatusInactivePending
	subscription.StatusAt = ctx.BlockTime()

	k.SetSubscription(ctx, subscription)
	k.SetInactiveSubscriptionAt(ctx, subscription.StatusAt.Add(inactiveDuration), subscription.Id)
	ctx.EventManager().EmitTypedEvent(
		&types.EventSetStatus{
			Id:     subscription.Id,
			Status: subscription.Status,
		},
	)

	return &types.MsgCancelResponse{}, nil
}

func (k *msgServer) MsgAddQuota(c context.Context, msg *types.MsgAddQuotaRequest) (*types.MsgAddQuotaResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	accAddr, err := sdk.AccAddressFromBech32(msg.Address)
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
	if k.HasQuota(ctx, subscription.Id, accAddr) {
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
	)

	k.SetQuota(ctx, subscription.Id, quota)
	k.SetActiveSubscriptionForAddress(ctx, accAddr, subscription.Id)
	ctx.EventManager().EmitTypedEvent(
		&types.EventAddQuota{
			Id:      subscription.Id,
			Address: quota.Address,
		},
	)

	return &types.MsgAddQuotaResponse{}, nil
}

func (k *msgServer) MsgUpdateQuota(c context.Context, msg *types.MsgUpdateQuotaRequest) (*types.MsgUpdateQuotaResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	subscription, found := k.GetSubscription(ctx, msg.Id)
	if !found {
		return nil, types.ErrorSubscriptionDoesNotExist
	}
	if subscription.Plan == 0 {
		return nil, types.ErrorCanNotUpdateQuota
	}
	if msg.From != subscription.Owner {
		return nil, types.ErrorUnauthorized
	}
	if !subscription.Status.Equal(hubtypes.StatusActive) {
		return nil, types.ErrorInvalidSubscriptionStatus
	}

	accAddr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	quota, found := k.GetQuota(ctx, subscription.Id, accAddr)
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
			Id:      subscription.Id,
			Address: quota.Address,
		},
	)

	return &types.MsgUpdateQuotaResponse{}, nil
}
