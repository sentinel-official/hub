package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/plan/types"
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

func (k *msgServer) MsgCreate(c context.Context, msg *types.MsgCreateRequest) (*types.MsgCreateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	provAddr, err := hubtypes.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}
	if !k.HasProvider(ctx, provAddr) {
		return nil, types.NewErrorProviderNotFound(provAddr)
	}

	var (
		count = k.GetCount(ctx)
		plan  = types.Plan{
			ID:       count + 1,
			Address:  provAddr.String(),
			Prices:   msg.Prices,
			Validity: msg.Validity,
			Bytes:    msg.Bytes,
			Status:   hubtypes.StatusInactive,
			StatusAt: ctx.BlockTime(),
		}
	)

	k.SetCount(ctx, count+1)
	k.SetInactivePlan(ctx, plan)
	k.SetInactivePlanForProvider(ctx, provAddr, plan.ID)
	ctx.EventManager().EmitTypedEvent(
		&types.EventCreate{
			ID: plan.ID,
		},
	)

	return &types.MsgCreateResponse{}, nil
}

func (k *msgServer) MsgUpdateStatus(c context.Context, msg *types.MsgUpdateStatusRequest) (*types.MsgUpdateStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorPlanNotFound(msg.ID)
	}
	if msg.From != plan.Address {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	provAddr := plan.GetAddress()
	if plan.Status.Equal(hubtypes.StatusActive) {
		if msg.Status.Equal(hubtypes.StatusInactive) {
			k.DeleteActivePlan(ctx, plan.ID)
			k.DeleteActivePlanForProvider(ctx, provAddr, plan.ID)
			k.SetInactivePlanForProvider(ctx, provAddr, plan.ID)
		}
	} else {
		if msg.Status.Equal(hubtypes.StatusActive) {
			k.DeleteInactivePlan(ctx, plan.ID)
			k.DeleteInactivePlanForProvider(ctx, provAddr, plan.ID)
			k.SetActivePlanForProvider(ctx, provAddr, plan.ID)
		}
	}

	plan.Status = msg.Status
	plan.StatusAt = ctx.BlockTime()

	k.SetPlan(ctx, plan)
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateStatus{
			ID:     plan.ID,
			Status: plan.Status,
		},
	)

	return &types.MsgUpdateStatusResponse{}, nil
}

func (k *msgServer) MsgLinkNode(c context.Context, msg *types.MsgLinkNodeRequest) (*types.MsgLinkNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	nodeAddr, err := hubtypes.NodeAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorPlanNotFound(msg.ID)
	}
	if msg.From != plan.Address {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	k.SetNodeForPlan(ctx, plan.ID, nodeAddr)
	ctx.EventManager().EmitTypedEvent(
		&types.EventLinkNode{
			ID: plan.ID,
		},
	)

	return &types.MsgLinkNodeResponse{}, nil
}

func (k *msgServer) MsgUnlinkNode(c context.Context, msg *types.MsgUnlinkNodeRequest) (*types.MsgUnlinkNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	nodeAddr, err := hubtypes.NodeAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorPlanNotFound(msg.ID)
	}
	if msg.From != plan.Address {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	k.DeleteNodeForPlan(ctx, plan.ID, nodeAddr)
	ctx.EventManager().EmitTypedEvent(
		&types.EventUnlinkNode{
			ID: plan.ID,
		},
	)

	return &types.MsgUnlinkNodeResponse{}, nil
}

func (k *msgServer) MsgSubscribe(c context.Context, msg *types.MsgSubscribeRequest) (*types.MsgSubscribeResponse, error) {
	return &types.MsgSubscribeResponse{}, nil
}
