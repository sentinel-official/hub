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

	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}
	if k.HasNode(ctx, fromAddr.Bytes()) {
		return nil, types.ErrorDuplicateNode
	}

	if msg.Provider != "" {
		provAddr, err := hubtypes.ProvAddressFromBech32(msg.Provider)
		if err != nil {
			return nil, err
		}
		if !k.HasProvider(ctx, provAddr) {
			return nil, types.ErrorProviderDoesNotExist
		}
	}
	if msg.Price != nil {
		if !k.IsValidPrice(ctx, msg.Price) {
			return nil, types.ErrorInvalidPrice
		}
	}

	deposit := k.Deposit(ctx)
	if err := k.FundCommunityPool(ctx, fromAddr, deposit); err != nil {
		return nil, err
	}

	var (
		nodeAddr = hubtypes.NodeAddress(fromAddr.Bytes())
		node     = types.Node{
			Address:   nodeAddr.String(),
			Provider:  msg.Provider,
			Price:     msg.Price,
			RemoteURL: msg.RemoteURL,
			Status:    hubtypes.StatusInactive,
			StatusAt:  ctx.BlockTime(),
		}
		provAddr = node.GetProvider()
	)

	k.SetNode(ctx, node)
	k.SetInactiveNode(ctx, nodeAddr)

	if provAddr != nil {
		k.SetInactiveNodeForProvider(ctx, provAddr, nodeAddr)
	}

	ctx.EventManager().EmitTypedEvent(
		&types.EventRegisterNode{
			From:      msg.From,
			Address:   node.Address,
			Provider:  node.Provider,
			Price:     node.Price,
			RemoteURL: node.RemoteURL,
		},
	)

	return &types.MsgRegisterResponse{}, nil
}

func (k *msgServer) MsgUpdate(c context.Context, msg *types.MsgUpdateRequest) (*types.MsgUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	fromAddr, err := hubtypes.NodeAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, fromAddr)
	if !found {
		return nil, types.ErrorNodeDoesNotExist
	}

	if msg.Provider == node.Provider {
		msg.Provider = ""
	}

	if (msg.Provider != "" || msg.Price != nil) && node.Provider != "" {
		var (
			nodeAddr = node.GetAddress()
			provAddr = node.GetProvider()
		)

		if k.GetPlanCountForNodeByProvider(ctx, provAddr, nodeAddr) > 0 {
			return nil, types.ErrorInvalidPlanCount
		}

		if node.Status.Equal(hubtypes.StatusActive) {
			k.DeleteActiveNodeForProvider(ctx, provAddr, nodeAddr)
		} else {
			k.DeleteInactiveNodeForProvider(ctx, provAddr, nodeAddr)
		}
	}

	if msg.Provider != "" {
		provAddr, err := hubtypes.ProvAddressFromBech32(msg.Provider)
		if err != nil {
			return nil, err
		}
		if !k.HasProvider(ctx, provAddr) {
			return nil, types.ErrorProviderDoesNotExist
		}

		node.Price = nil
		node.Provider = msg.Provider

		nodeAddr := node.GetAddress()
		if node.Status.Equal(hubtypes.StatusActive) {
			k.SetActiveNodeForProvider(ctx, provAddr, nodeAddr)
		} else {
			k.SetInactiveNodeForProvider(ctx, provAddr, nodeAddr)
		}
	}
	if msg.Price != nil {
		if !k.IsValidPrice(ctx, msg.Price) {
			return nil, types.ErrorInvalidPrice
		}

		node.Provider = ""
		node.Price = msg.Price
	}
	if msg.RemoteURL != "" {
		node.RemoteURL = msg.RemoteURL
	}

	k.SetNode(ctx, node)
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateNode{
			From:      msg.From,
			Address:   node.Address,
			Provider:  node.Provider,
			Price:     node.Price,
			RemoteURL: node.RemoteURL,
		},
	)

	return &types.MsgUpdateResponse{}, nil
}

func (k *msgServer) MsgSetStatus(c context.Context, msg *types.MsgSetStatusRequest) (*types.MsgSetStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	fromAddr, err := hubtypes.NodeAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, fromAddr)
	if !found {
		return nil, types.ErrorNodeDoesNotExist
	}

	var (
		nodeAddr         = node.GetAddress()
		provAddr         = node.GetProvider()
		inactiveDuration = k.InactiveDuration(ctx)
	)

	if node.Status.Equal(hubtypes.StatusActive) {
		if msg.Status.Equal(hubtypes.StatusInactive) {
			k.DeleteActiveNode(ctx, nodeAddr)
			k.SetInactiveNode(ctx, nodeAddr)

			if node.Provider != "" {
				k.DeleteActiveNodeForProvider(ctx, provAddr, nodeAddr)
				k.SetInactiveNodeForProvider(ctx, provAddr, nodeAddr)
			}
		}

		k.DeleteInactiveNodeAt(ctx, node.StatusAt.Add(inactiveDuration), nodeAddr)
	} else {
		if msg.Status.Equal(hubtypes.StatusActive) {
			k.DeleteInactiveNode(ctx, nodeAddr)
			k.SetActiveNode(ctx, nodeAddr)

			if node.Provider != "" {
				k.DeleteInactiveNodeForProvider(ctx, provAddr, nodeAddr)
				k.SetActiveNodeForProvider(ctx, provAddr, nodeAddr)
			}
		}
	}

	node.Status = msg.Status
	node.StatusAt = ctx.BlockTime()

	if node.Status.Equal(hubtypes.StatusActive) {
		k.SetInactiveNodeAt(ctx, node.StatusAt.Add(inactiveDuration), nodeAddr)
	}

	k.SetNode(ctx, node)
	ctx.EventManager().EmitTypedEvent(
		&types.EventSetNodeStatus{
			From:    msg.From,
			Address: node.Address,
			Status:  node.Status,
		},
	)

	return &types.MsgSetStatusResponse{}, nil
}
