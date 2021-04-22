package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
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

func (k *msgServer) MsgRegister(c context.Context, msg *types.MsgRegisterRequest) (*types.MsgRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	msgFrom, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}
	if k.HasNode(ctx, msgFrom.Bytes()) {
		return nil, types.ErrorDuplicateNode
	}

	msgProvider, err := hubtypes.ProvAddressFromBech32(msg.Provider)
	if err != nil {
		return nil, err
	}
	if msg.Provider != "" {
		if !k.HasProvider(ctx, msgProvider) {
			return nil, types.ErrorProviderDoesNotExist
		}
	}

	deposit := k.Deposit(ctx)
	if deposit.IsPositive() {
		if err := k.FundCommunityPool(ctx, msgFrom, deposit); err != nil {
			return nil, err
		}
	}

	var (
		nodeAddress  = hubtypes.NodeAddress(msgFrom.Bytes())
		nodeProvider = hubtypes.ProvAddress(msgProvider.Bytes())
		node         = types.Node{
			Address:   nodeAddress.String(),
			Provider:  msg.Provider,
			Price:     msg.Price,
			RemoteURL: msg.RemoteURL,
			Status:    hubtypes.StatusInactive,
			StatusAt:  ctx.BlockTime(),
		}
	)

	k.SetNode(ctx, node)
	k.SetInactiveNode(ctx, nodeAddress)

	if nodeProvider != nil {
		k.SetInactiveNodeForProvider(ctx, nodeProvider, nodeAddress)
	}

	ctx.EventManager().EmitTypedEvent(
		&types.EventRegisterNode{
			From:      sdk.AccAddress(msgFrom.Bytes()).String(),
			Address:   node.Address,
			Provider:  node.Provider,
			Price:     node.Price,
			RemoteURL: node.RemoteURL,
		},
	)

	ctx.EventManager().EmitTypedEvent(&types.EventModuleName)
	return &types.MsgRegisterResponse{}, nil
}

func (k *msgServer) MsgUpdate(c context.Context, msg *types.MsgUpdateRequest) (*types.MsgUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	msgFrom, err := hubtypes.NodeAddressFromBech32(msg.From)
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

		if node.Status.Equal(hubtypes.StatusActive) {
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
		msgProvider, err := hubtypes.ProvAddressFromBech32(msg.Provider)
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

		if node.Status.Equal(hubtypes.StatusActive) {
			k.SetActiveNodeForProvider(ctx, nodeProvider, nodeAddress)
		} else {
			k.SetInactiveNodeForProvider(ctx, nodeProvider, nodeAddress)
		}
	}
	if msg.Price != nil {
		node.Provider = ""
		node.Price = msg.Price
	}
	if msg.RemoteURL != "" {
		node.RemoteURL = msg.RemoteURL
	}

	k.SetNode(ctx, node)
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateNode{
			From:      sdk.AccAddress(msgFrom.Bytes()).String(),
			Address:   node.Address,
			Provider:  msg.Provider,
			Price:     msg.Price,
			RemoteURL: msg.RemoteURL,
		},
	)

	ctx.EventManager().EmitTypedEvent(&types.EventModuleName)
	return &types.MsgUpdateResponse{}, nil
}

func (k *msgServer) MsgSetStatus(c context.Context, msg *types.MsgSetStatusRequest) (*types.MsgSetStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	msgFrom, err := hubtypes.NodeAddressFromBech32(msg.From)
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

	if node.Status.Equal(hubtypes.StatusActive) {
		if msg.Status.Equal(hubtypes.StatusInactive) {
			k.DeleteActiveNode(ctx, nodeAddress)
			k.SetInactiveNode(ctx, nodeAddress)

			if node.Provider != "" {
				k.DeleteActiveNodeForProvider(ctx, nodeProvider, nodeAddress)
				k.SetInactiveNodeForProvider(ctx, nodeProvider, nodeAddress)
			}
		}

		k.DeleteInactiveNodeAt(ctx, node.StatusAt, nodeAddress)
	} else {
		if msg.Status.Equal(hubtypes.StatusActive) {
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

	if node.Status.Equal(hubtypes.StatusActive) {
		k.SetInactiveNodeAt(ctx, node.StatusAt, nodeAddress)
	}

	k.SetNode(ctx, node)
	ctx.EventManager().EmitTypedEvent(
		&types.EventSetNodeStatus{
			From:    sdk.AccAddress(msgFrom.Bytes()).String(),
			Address: node.Address,
			Status:  node.Status,
		},
	)

	ctx.EventManager().EmitTypedEvent(&types.EventModuleName)
	return &types.MsgSetStatusResponse{}, nil
}
