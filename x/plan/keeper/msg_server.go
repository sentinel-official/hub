package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/plan/types"
)

var (
	_ types.MsgServiceServer = (*server)(nil)
)

type server struct {
	Keeper
}

func NewMsgServiceServer(keeper Keeper) types.MsgServiceServer {
	return &server{Keeper: keeper}
}

func (k *server) MsgAdd(c context.Context, msg *types.MsgAddRequest) (*types.MsgAddResponse, error) {
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
	)

	var (
		planProvider = plan.GetProvider()
	)

	k.SetPlan(ctx, plan)
	k.SetInactivePlan(ctx, plan.Id)
	k.SetInactivePlanForProvider(ctx, planProvider, plan.Id)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSet,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", plan.Id)),
		sdk.NewAttribute(types.AttributeKeyAddress, plan.Provider),
	))

	k.SetCount(ctx, count+1)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetCount,
		sdk.NewAttribute(types.AttributeKeyCount, fmt.Sprintf("%d", count+1)),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &types.MsgAddResponse{}, nil
}

func (k *server) MsgSetStatus(c context.Context, msg *types.MsgSetStatusRequest) (*types.MsgSetStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

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
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetStatus,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", plan.Id)),
		sdk.NewAttribute(types.AttributeKeyStatus, plan.Status.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &types.MsgSetStatusResponse{}, nil
}

func (k *server) MsgAddNode(c context.Context, msg *types.MsgAddNodeRequest) (*types.MsgAddNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

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

	var (
		nodeAddress = node.GetAddress()
	)

	k.SetNodeForPlan(ctx, plan.Id, nodeAddress)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAddNode,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", plan.Id)),
		sdk.NewAttribute(types.AttributeKeyAddress, node.Address),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &types.MsgAddNodeResponse{}, nil
}

func (k *server) MsgRemoveNode(c context.Context, msg *types.MsgRemoveNodeRequest) (*types.MsgRemoveNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

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

	k.DeleteNodeForPlan(ctx, plan.Id, msgAddress)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeRemoveNode,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", plan.Id)),
		sdk.NewAttribute(types.AttributeKeyAddress, msg.Address),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &types.MsgRemoveNodeResponse{}, nil
}
