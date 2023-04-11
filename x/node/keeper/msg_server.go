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

	if msg.GigabytePrices != nil {
		if !k.IsValidGigabytePrices(ctx, msg.GigabytePrices) {
			return nil, types.ErrorInvalidGigabytePrices
		}
	}
	if msg.HourlyPrices != nil {
		if !k.IsValidHourlyPrices(ctx, msg.HourlyPrices) {
			return nil, types.ErrorInvalidHourlyPrices
		}
	}

	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}
	if k.HasNode(ctx, fromAddr.Bytes()) {
		return nil, types.ErrorDuplicateNode
	}

	deposit := k.Deposit(ctx)
	if err = k.FundCommunityPool(ctx, fromAddr, deposit); err != nil {
		return nil, err
	}

	var (
		nodeAddr = hubtypes.NodeAddress(fromAddr.Bytes())
		node     = types.Node{
			Address:        nodeAddr.String(),
			GigabytePrices: msg.GigabytePrices,
			HourlyPrices:   msg.HourlyPrices,
			RemoteURL:      msg.RemoteURL,
			Status:         hubtypes.StatusInactive,
			StatusAt:       ctx.BlockTime(),
		}
	)

	k.SetNode(ctx, node)
	ctx.EventManager().EmitTypedEvent(
		&types.EventRegister{
			Address: node.Address,
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

	if msg.GigabytePrices != nil {
		if !k.IsValidGigabytePrices(ctx, msg.GigabytePrices) {
			return nil, types.ErrorInvalidGigabytePrices
		}

		node.GigabytePrices = msg.GigabytePrices
	}
	if msg.HourlyPrices != nil {
		if !k.IsValidHourlyPrices(ctx, msg.HourlyPrices) {
			return nil, types.ErrorInvalidHourlyPrices
		}

		node.HourlyPrices = msg.HourlyPrices
	}
	if msg.RemoteURL != "" {
		node.RemoteURL = msg.RemoteURL
	}

	k.SetNode(ctx, node)
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdate{
			Address: node.Address,
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
		inactiveDuration = k.InactiveDuration(ctx)
	)

	if node.Status.Equal(hubtypes.StatusActive) {
		if msg.Status.Equal(hubtypes.StatusInactive) {
			k.DeleteActiveNode(ctx, nodeAddr)
		}

		inactiveAt := node.StatusAt.Add(inactiveDuration)
		k.DeleteInactiveNodeAt(ctx, inactiveAt, nodeAddr)
	} else {
		if msg.Status.Equal(hubtypes.StatusActive) {
			k.DeleteInactiveNode(ctx, nodeAddr)
		}
	}

	node.Status = msg.Status
	node.StatusAt = ctx.BlockTime()

	if node.Status.Equal(hubtypes.StatusActive) {
		inactiveAt := node.StatusAt.Add(inactiveDuration)
		k.SetInactiveNodeAt(ctx, inactiveAt, nodeAddr)
	}

	k.SetNode(ctx, node)
	ctx.EventManager().EmitTypedEvent(
		&types.EventSetStatus{
			Address: node.Address,
			Status:  node.Status,
		},
	)

	return &types.MsgSetStatusResponse{}, nil
}
