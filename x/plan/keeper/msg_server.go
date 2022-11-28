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

func NewMsgServiceServer(keeper Keeper) types.MsgServiceServer {
	return &msgServer{Keeper: keeper}
}

func (k *msgServer) MsgAdd(c context.Context, msg *types.MsgAddRequest) (*types.MsgAddResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	fromAddr, err := hubtypes.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}
	if !k.HasProvider(ctx, fromAddr) {
		return nil, types.ErrorProviderDoesNotExist
	}

	var (
		count = k.GetCount(ctx)
		plan  = types.Plan{
			Id:       count + 1,
			Provider: msg.From,
			Price:    msg.Price,
			Validity: msg.Validity,
			Bytes:    msg.Bytes,
			Status:   hubtypes.StatusInactive,
			StatusAt: ctx.BlockTime(),
		}
		provAddr = plan.GetProvider()
	)

	k.SetCount(ctx, count+1)
	k.SetPlan(ctx, plan)
	k.SetInactivePlan(ctx, plan.Id)
	k.SetInactivePlanForProvider(ctx, provAddr, plan.Id)
	ctx.EventManager().EmitTypedEvent(
		&types.EventAdd{
			Id:       plan.Id,
			Provider: plan.Provider,
		},
	)

	return &types.MsgAddResponse{}, nil
}

func (k *msgServer) MsgSetStatus(c context.Context, msg *types.MsgSetStatusRequest) (*types.MsgSetStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	plan, found := k.GetPlan(ctx, msg.Id)
	if !found {
		return nil, types.ErrorPlanDoesNotExist
	}
	if msg.From != plan.Provider {
		return nil, types.ErrorUnauthorized
	}

	provAddr := plan.GetProvider()
	if plan.Status.Equal(hubtypes.StatusActive) {
		if msg.Status.Equal(hubtypes.StatusInactive) {
			k.DeleteActivePlan(ctx, plan.Id)
			k.DeleteActivePlanForProvider(ctx, provAddr, plan.Id)

			k.SetInactivePlan(ctx, plan.Id)
			k.SetInactivePlanForProvider(ctx, provAddr, plan.Id)
		}
	} else {
		if msg.Status.Equal(hubtypes.StatusActive) {
			k.DeleteInactivePlan(ctx, plan.Id)
			k.DeleteInactivePlanForProvider(ctx, provAddr, plan.Id)

			k.SetActivePlan(ctx, plan.Id)
			k.SetActivePlanForProvider(ctx, provAddr, plan.Id)
		}
	}

	plan.Status = msg.Status
	plan.StatusAt = ctx.BlockTime()

	k.SetPlan(ctx, plan)
	ctx.EventManager().EmitTypedEvent(
		&types.EventSetStatus{
			Id:       plan.Id,
			Provider: plan.Provider,
			Status:   plan.Status,
		},
	)

	return &types.MsgSetStatusResponse{}, nil
}

func (k *msgServer) MsgAddNode(c context.Context, msg *types.MsgAddNodeRequest) (*types.MsgAddNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	plan, found := k.GetPlan(ctx, msg.Id)
	if !found {
		return nil, types.ErrorPlanDoesNotExist
	}
	if msg.From != plan.Provider {
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
	if msg.From != node.Provider {
		return nil, types.ErrorUnauthorized
	}

	if k.HasNodeForPlan(ctx, plan.Id, nodeAddr) {
		return nil, types.DuplicateNodeForPlan
	}

	provAddr := plan.GetProvider()
	k.SetNodeForPlan(ctx, plan.Id, nodeAddr)
	k.IncreaseCountForNodeByProvider(ctx, provAddr, nodeAddr)
	ctx.EventManager().EmitTypedEvent(
		&types.EventAddNode{
			Id:       plan.Id,
			Node:     nodeAddr.String(),
			Provider: plan.Provider,
		},
	)

	return &types.MsgAddNodeResponse{}, nil
}

func (k *msgServer) MsgRemoveNode(c context.Context, msg *types.MsgRemoveNodeRequest) (*types.MsgRemoveNodeResponse, error) {
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
		&types.EventRemoveNode{
			Id:       plan.Id,
			Node:     nodeAddr.String(),
			Provider: plan.Provider,
		},
	)

	return &types.MsgRemoveNodeResponse{}, nil
}
