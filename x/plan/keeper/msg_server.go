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

	msgFrom, err := hubtypes.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}
	if !k.HasProvider(ctx, msgFrom) {
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
		planProvider = plan.GetProvider()
	)

	k.SetPlan(ctx, plan)
	k.SetInactivePlan(ctx, plan.Id)
	k.SetInactivePlanForProvider(ctx, planProvider, plan.Id)
	ctx.EventManager().EmitTypedEvent(
		&types.EventAddPlan{
			From:     sdk.AccAddress(msgFrom.Bytes()).String(),
			Provider: plan.Provider,
			Price:    plan.Price,
			Validity: plan.Validity,
			Bytes:    plan.Bytes,
		},
	)

	k.SetCount(ctx, count+1)
	ctx.EventManager().EmitTypedEvent(
		&types.EventSetPlanCount{
			Count: count + 1,
		},
	)

	ctx.EventManager().EmitTypedEvent(&types.EventModuleName)
	return &types.MsgAddResponse{}, nil
}

func (k *msgServer) MsgSetStatus(c context.Context, msg *types.MsgSetStatusRequest) (*types.MsgSetStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	msgFrom, err := hubtypes.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	plan, found := k.GetPlan(ctx, msg.Id)
	if !found {
		return nil, types.ErrorPlanDoesNotExist
	}
	if msg.From != plan.Provider {
		return nil, types.ErrorUnauthorized
	}

	var (
		planProvider = plan.GetProvider()
	)

	if plan.Status.Equal(hubtypes.StatusActive) {
		if msg.Status.Equal(hubtypes.StatusInactive) {
			k.DeleteActivePlan(ctx, plan.Id)
			k.DeleteActivePlanForProvider(ctx, planProvider, plan.Id)

			k.SetInactivePlan(ctx, plan.Id)
			k.SetInactivePlanForProvider(ctx, planProvider, plan.Id)
		}
	} else {
		if msg.Status.Equal(hubtypes.StatusActive) {
			k.DeleteInactivePlan(ctx, plan.Id)
			k.DeleteInactivePlanForProvider(ctx, planProvider, plan.Id)

			k.SetActivePlan(ctx, plan.Id)
			k.SetActivePlanForProvider(ctx, planProvider, plan.Id)
		}
	}

	plan.Status = msg.Status
	plan.StatusAt = ctx.BlockTime()

	k.SetPlan(ctx, plan)
	ctx.EventManager().EmitTypedEvent(
		&types.EventSetPlanStatus{
			From:     sdk.AccAddress(msgFrom.Bytes()).String(),
			Provider: plan.Provider,
			Id:       plan.Id,
			Status:   plan.Status,
		},
	)

	ctx.EventManager().EmitTypedEvent(&types.EventModuleName)
	return &types.MsgSetStatusResponse{}, nil
}

func (k *msgServer) MsgAddNode(c context.Context, msg *types.MsgAddNodeRequest) (*types.MsgAddNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	msgFrom, err := hubtypes.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	plan, found := k.GetPlan(ctx, msg.Id)
	if !found {
		return nil, types.ErrorPlanDoesNotExist
	}
	if msg.From != plan.Provider {
		return nil, types.ErrorUnauthorized
	}

	msgAddress, err := hubtypes.NodeAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, msgAddress)
	if !found {
		return nil, types.ErrorNodeDoesNotExist
	}
	if msg.From != node.Provider {
		return nil, types.ErrorUnauthorized
	}

	if k.HasNodeForPlan(ctx, plan.Id, msgAddress) {
		return nil, types.DuplicateNodeForPlan
	}

	var (
		planProvider = plan.GetProvider()
		nodeAddress  = node.GetAddress()
	)

	k.SetNodeForPlan(ctx, plan.Id, nodeAddress)
	k.IncreaseCountForNodeByProvider(ctx, planProvider, nodeAddress)
	ctx.EventManager().EmitTypedEvent(
		&types.EventAddNodeForPlan{
			From:     sdk.AccAddress(msgFrom.Bytes()).String(),
			Provider: plan.Provider,
			Id:       plan.Id,
			Address:  msg.Address,
		},
	)

	ctx.EventManager().EmitTypedEvent(&types.EventModuleName)
	return &types.MsgAddNodeResponse{}, nil
}

func (k *msgServer) MsgRemoveNode(c context.Context, msg *types.MsgRemoveNodeRequest) (*types.MsgRemoveNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	msgFrom, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	plan, found := k.GetPlan(ctx, msg.Id)
	if !found {
		return nil, types.ErrorPlanDoesNotExist
	}

	planProvider := plan.GetProvider()
	if hubtypes.NodeAddress(msgFrom.Bytes()).String() != msg.Address {
		if hubtypes.ProvAddress(msgFrom.Bytes()).String() != plan.Provider {
			return nil, types.ErrorUnauthorized
		}
	}

	msgAddress, err := hubtypes.NodeAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	if !k.HasNodeForPlan(ctx, plan.Id, msgAddress) {
		return nil, types.ErrorNodeDoesNotExist
	}

	k.DeleteNodeForPlan(ctx, plan.Id, msgAddress)
	k.DecreaseCountForNodeByProvider(ctx, planProvider, msgAddress)
	ctx.EventManager().EmitTypedEvent(
		&types.EventRemoveNodeForPlan{
			From:    sdk.AccAddress(msgFrom.Bytes()).String(),
			Id:      plan.Id,
			Address: msg.Address,
		},
	)

	ctx.EventManager().EmitTypedEvent(&types.EventModuleName)
	return &types.MsgRemoveNodeResponse{}, nil
}
