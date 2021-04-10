package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
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

func (k server) MsgRegister(c context.Context, msg *types.MsgRegisterRequest) (*types.MsgRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	msgFrom, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}
	if k.HasNode(ctx, msgFrom.Bytes()) {
		return nil, types.ErrorDuplicateNode
	}

	msgProvider, err := hub.ProvAddressFromBech32(msg.Provider)
	if err != nil {
		return nil, err
	}
	if !k.HasProvider(ctx, msgProvider) {
		return nil, types.ErrorProviderDoesNotExist
	}

	deposit := k.Deposit(ctx)
	if deposit.IsPositive() {
		if err := k.FundCommunityPool(ctx, msgFrom, deposit); err != nil {
			return nil, err
		}
	}

	var (
		nodeAddress  = hub.NodeAddress(msgFrom.Bytes())
		nodeProvider = hub.ProvAddress(msgProvider.Bytes())
		node         = types.Node{
			Address:   nodeAddress.String(),
			Provider:  msg.Provider,
			Price:     msg.Price,
			RemoteUrl: msg.RemoteUrl,
			Status:    hub.StatusInactive,
			StatusAt:  ctx.BlockTime(),
		}
	)

	k.SetNode(ctx, node)
	k.SetInactiveNode(ctx, nodeAddress)

	if nodeProvider != nil {
		k.SetInactiveNodeForProvider(ctx, nodeProvider, nodeAddress)
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSet,
		sdk.NewAttribute(types.AttributeKeyProvider, node.Provider),
		sdk.NewAttribute(types.AttributeKeyAddress, node.Address),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &types.MsgRegisterResponse{}, nil
}

func (k server) MsgUpdate(c context.Context, msg *types.MsgUpdateRequest) (*types.MsgUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	msgFrom, err := hub.NodeAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, msgFrom)
	if !found {
		return nil, types.ErrorNodeDoesNotExist
	}

	if node.Provider == msg.Provider {
		msg.Provider = ""
	}

	if node.Provider != "" && (msg.Provider != "" || msg.Price != nil) {
		var (
			nodeAddress  = node.GetAddress()
			nodeProvider = node.GetProvider()
		)

		if node.Status.Equal(hub.StatusActive) {
			k.DeleteActiveNodeForProvider(ctx, nodeProvider, nodeAddress)
		} else {
			k.DeleteInactiveNodeForProvider(ctx, nodeProvider, nodeAddress)
		}

		// TODO: Remove or optimize this?
		plans := k.GetPlansForProvider(ctx, nodeProvider)
		for _, plan := range plans {
			k.DeleteNodeForPlan(ctx, plan.Id, nodeAddress)
		}
	}

	if msg.Provider != "" {
		msgProvider, err := hub.ProvAddressFromBech32(msg.Provider)
		if err != nil {
			return nil, err
		}
		if !k.HasProvider(ctx, msgProvider) {
			return nil, types.ErrorProviderDoesNotExist
		}

		node.Price = nil
		node.Provider = msg.Provider

		var (
			nodeAddress  = node.GetAddress()
			nodeProvider = node.GetProvider()
		)

		if node.Status.Equal(hub.StatusActive) {
			k.SetActiveNodeForProvider(ctx, nodeProvider, nodeAddress)
		} else {
			k.SetInactiveNodeForProvider(ctx, nodeProvider, nodeAddress)
		}
	}
	if msg.Price != nil {
		node.Provider = ""
		node.Price = msg.Price
	}
	if msg.RemoteUrl != "" {
		node.RemoteUrl = msg.RemoteUrl
	}

	k.SetNode(ctx, node)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUpdate,
		sdk.NewAttribute(types.AttributeKeyAddress, node.Address),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &types.MsgUpdateResponse{}, nil
}

func (k server) MsgSetStatus(c context.Context, msg *types.MsgSetStatusRequest) (*types.MsgSetStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	msgFrom, err := hub.NodeAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, msgFrom)
	if !found {
		return nil, types.ErrorNodeDoesNotExist
	}

	var (
		nodeAddress  = node.GetAddress()
		nodeProvider = node.GetProvider()
	)

	if node.Status.Equal(hub.StatusActive) {
		if msg.Status.Equal(hub.StatusInactive) {
			k.DeleteActiveNode(ctx, nodeAddress)
			k.SetInactiveNode(ctx, nodeAddress)

			if node.Provider != "" {
				k.DeleteActiveNodeForProvider(ctx, nodeProvider, nodeAddress)
				k.SetInactiveNodeForProvider(ctx, nodeProvider, nodeAddress)
			}
		}

		k.DeleteInactiveNodeAt(ctx, node.StatusAt, nodeAddress)
	} else {
		if msg.Status.Equal(hub.StatusActive) {
			k.DeleteInactiveNode(ctx, nodeAddress)
			k.SetActiveNode(ctx, nodeAddress)

			if node.Provider != "" {
				k.DeleteInactiveNodeForProvider(ctx, nodeProvider, nodeAddress)
				k.SetActiveNodeForProvider(ctx, nodeProvider, nodeAddress)
			}
		}
	}

	node.Status = msg.Status
	node.StatusAt = ctx.BlockTime()

	if node.Status.Equal(hub.StatusActive) {
		k.SetInactiveNodeAt(ctx, node.StatusAt, nodeAddress)
	}

	k.SetNode(ctx, node)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetStatus,
		sdk.NewAttribute(types.AttributeKeyAddress, node.Address),
		sdk.NewAttribute(types.AttributeKeyStatus, node.Status.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &types.MsgSetStatusResponse{}, nil
}
