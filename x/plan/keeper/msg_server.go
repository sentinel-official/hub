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
		return nil, types.ErrorProviderDoesNotExist
	}

	var (
		count = k.GetCount(ctx)
		plan  = types.Plan{
			ID:              count + 1,
			ProviderAddress: provAddr.String(),
			Prices:          msg.Prices,
			Validity:        msg.Validity,
			Bytes:           msg.Bytes,
			Status:          hubtypes.StatusInactive,
			StatusAt:        ctx.BlockTime(),
		}
	)

	k.SetCount(ctx, count+1)
	k.SetInactivePlan(ctx, plan)
	k.SetInactivePlanForProvider(ctx, provAddr, plan.ID)
	ctx.EventManager().EmitTypedEvent(
		&types.EventCreate{
			ID:              plan.ID,
			ProviderAddress: plan.ProviderAddress,
		},
	)

	return &types.MsgCreateResponse{}, nil
}

func (k *msgServer) MsgUpdateStatus(c context.Context, msg *types.MsgUpdateStatusRequest) (*types.MsgUpdateStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.ErrorPlanDoesNotExist
	}
	if msg.From != plan.ProviderAddress {
		return nil, types.ErrorUnauthorized
	}

	provAddr := plan.GetProviderAddress()
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
			ID:              plan.ID,
			ProviderAddress: plan.ProviderAddress,
			Status:          plan.Status,
		},
	)

	return &types.MsgUpdateStatusResponse{}, nil
}

func (k *msgServer) MsgLinkNode(c context.Context, msg *types.MsgLinkNodeRequest) (*types.MsgLinkNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.ErrorPlanDoesNotExist
	}
	if msg.From != plan.ProviderAddress {
		return nil, types.ErrorUnauthorized
	}

	nodeAddr, err := hubtypes.NodeAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.ErrorNodeDoesNotExist
	}

	if k.HasNodeForPlan(ctx, plan.Id, nodeAddr) {
		return nil, types.DuplicateNodeForPlan
	}

	provAddr := plan.GetProvider()
	k.SetNodeForPlan(ctx, plan.Id, nodeAddr)
	k.IncreaseCountForNodeByProvider(ctx, provAddr, nodeAddr)
	ctx.EventManager().EmitTypedEvent(
		&types.EventLinkNode{
			Id:       plan.Id,
			Node:     nodeAddr.String(),
			Provider: plan.Provider,
		},
	)

	return &types.MsgLinkNodeResponse{}, nil
}

func (k *msgServer) MsgUnlinkNode(c context.Context, msg *types.MsgUnlinkNodeRequest) (*types.MsgUnlinkNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	plan, found := k.GetPlan(ctx, msg.Id)
	if !found {
		return nil, types.ErrorPlanDoesNotExist
	}

	if hubtypes.NodeAddress(fromAddr.Bytes()).String() != msg.Address {
		if hubtypes.ProvAddress(fromAddr.Bytes()).String() != plan.Provider {
			return nil, types.ErrorUnauthorized
		}
	}

	nodeAddr, err := hubtypes.NodeAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	if !k.HasNodeForPlan(ctx, plan.Id, nodeAddr) {
		return nil, types.ErrorNodeDoesNotExist
	}

	provAddr := plan.GetProvider()
	k.DeleteNodeForPlan(ctx, plan.Id, nodeAddr)
	k.DecreaseCountForNodeByProvider(ctx, provAddr, nodeAddr)
	ctx.EventManager().EmitTypedEvent(
		&types.EventUnlinkNode{
			Id:       plan.Id,
			Node:     nodeAddr.String(),
			Provider: plan.Provider,
		},
	)

	return &types.MsgUnlinkNodeResponse{}, nil
}

func (k *msgServer) MsgSubscribe(c context.Context, msg *types.MsgSubscribeRequest) (*types.MsgSubscribeResponse, error) {
	return &types.MsgSubscribeResponse{}, nil
}
