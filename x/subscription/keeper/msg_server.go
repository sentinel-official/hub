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

func (k *msgServer) MsgShare(c context.Context, msg *types.MsgShareRequest) (*types.MsgShareResponse, error) {
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

	return &types.MsgShareResponse{}, nil
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
