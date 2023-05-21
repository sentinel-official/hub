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

func NewMsgServiceServer(k Keeper) types.MsgServiceServer {
	return &msgServer{k}
}

func (k *msgServer) MsgCancel(c context.Context, msg *types.MsgCancelRequest) (*types.MsgCancelResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}
	if !subscription.GetStatus().Equal(hubtypes.StatusActive) {
		return nil, types.NewErrorInvalidSubscriptionStatus(subscription.GetID(), subscription.GetStatus())
	}
	if msg.From != subscription.GetAddress() {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	k.DeleteSubscriptionForExpiryAt(ctx, subscription.GetExpiryAt(), subscription.GetID())

	inactiveDuration := k.InactiveDuration(ctx)
	subscription.SetExpiryAt(ctx.BlockTime().Add(inactiveDuration))
	k.SetSubscriptionForExpiryAt(ctx, subscription.GetExpiryAt(), subscription.GetID())

	subscription.SetStatus(hubtypes.StatusInactivePending)
	subscription.SetStatusAt(ctx.BlockTime())

	k.SetSubscription(ctx, subscription)
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateStatus{
			ID:     subscription.GetID(),
			Status: subscription.GetStatus(),
		},
	)

	return &types.MsgCancelResponse{}, nil
}

func (k *msgServer) MsgAllocate(c context.Context, msg *types.MsgAllocateRequest) (*types.MsgAllocateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}
	if subscription.Type() != types.TypePlan {
		return nil, types.NewErrorInvalidSubscriptionType(subscription.GetID(), subscription.Type())
	}
	if msg.From != subscription.GetAddress() {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	fromAllocation, found := k.GetAllocation(ctx, subscription.GetID(), fromAddr)
	if !found {
		return nil, types.NewErrorAllocationNotFound(subscription.GetID(), fromAddr)
	}

	toAddr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	toAllocation, found := k.GetAllocation(ctx, subscription.GetID(), toAddr)
	if !found {
		toAllocation = types.Allocation{
			Address:       toAddr.String(),
			GrantedBytes:  sdk.ZeroInt(),
			UtilisedBytes: sdk.ZeroInt(),
		}

		k.SetSubscriptionForAccount(ctx, toAddr, subscription.GetID())
	}

	var (
		granted   = fromAllocation.GrantedBytes.Add(toAllocation.GrantedBytes)
		utilised  = fromAllocation.UtilisedBytes.Add(toAllocation.UtilisedBytes)
		available = granted.Sub(utilised)
	)

	if msg.Bytes.GT(available) {
		return nil, types.NewErrorInsufficientBytes(subscription.GetID(), msg.Bytes)
	}

	fromAllocation.GrantedBytes = available.Sub(msg.Bytes)
	if fromAllocation.GrantedBytes.LT(fromAllocation.UtilisedBytes) {
		return nil, types.NewErrorInvalidAllocation(subscription.GetID(), fromAddr)
	}

	k.SetAllocation(ctx, subscription.GetID(), fromAllocation)
	ctx.EventManager().EmitTypedEvent(
		&types.EventAllocate{
			ID:      subscription.GetID(),
			Address: fromAllocation.Address,
			Bytes:   fromAllocation.GrantedBytes,
		},
	)

	toAllocation.GrantedBytes = msg.Bytes
	if toAllocation.GrantedBytes.LT(toAllocation.UtilisedBytes) {
		return nil, types.NewErrorInvalidAllocation(subscription.GetID(), toAddr)
	}

	k.SetAllocation(ctx, subscription.GetID(), toAllocation)
	ctx.EventManager().EmitTypedEvent(
		&types.EventAllocate{
			ID:      subscription.GetID(),
			Address: toAllocation.Address,
			Bytes:   toAllocation.GrantedBytes,
		},
	)

	return &types.MsgAllocateResponse{}, nil
}
