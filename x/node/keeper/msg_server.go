package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
)

// The following line asserts that the `msgServer` type implements the `types.MsgServiceServer` interface.
var (
	_ types.MsgServiceServer = (*msgServer)(nil)
)

// msgServer is a message server that implements the `types.MsgServiceServer` interface.
type msgServer struct {
	Keeper // Keeper is an instance of the main keeper for the module.
}

// NewMsgServiceServer creates a new instance of `types.MsgServiceServer` using the provided Keeper.
func NewMsgServiceServer(k Keeper) types.MsgServiceServer {
	return &msgServer{k}
}

// MsgRegister registers a new node in the network.
// It validates the registration request, checks prices, and creates a new node.
func (k *msgServer) MsgRegister(c context.Context, msg *types.MsgRegisterRequest) (*types.MsgRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Check if the provided GigabytePrices are valid, if not, return an error.
	if !k.IsValidGigabytePrices(ctx, msg.GigabytePrices) {
		return nil, types.NewErrorInvalidPrices(msg.GigabytePrices)
	}

	// Check if the provided HourlyPrices are valid, if not, return an error.
	if !k.IsValidHourlyPrices(ctx, msg.HourlyPrices) {
		return nil, types.NewErrorInvalidPrices(msg.HourlyPrices)
	}

	// Convert the `msg.From` address (in Bech32 format) to an `sdk.AccAddress`.
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Convert the account address to a `hubtypes.NodeAddress`.
	nodeAddr := hubtypes.NodeAddress(accAddr.Bytes())

	// Check if the node already exists in the network. If found, return an error to prevent duplicate registration.
	if k.HasNode(ctx, nodeAddr) {
		return nil, types.NewErrorDuplicateNode(nodeAddr)
	}

	// Get the required deposit for registering a new node.
	deposit := k.Deposit(ctx)

	// Fund the community pool with the required deposit from the registrant's account.
	if err = k.FundCommunityPool(ctx, accAddr, deposit); err != nil {
		return nil, err
	}

	// Create a new node with the provided information and set its status to `Inactive`.
	node := types.Node{
		Address:        nodeAddr.String(),
		GigabytePrices: msg.GigabytePrices,
		HourlyPrices:   msg.HourlyPrices,
		RemoteURL:      msg.RemoteURL,
		InactiveAt:     time.Time{},
		Status:         hubtypes.StatusInactive,
		StatusAt:       ctx.BlockTime(),
	}

	// Save the new node in the Store.
	k.SetNode(ctx, node)

	// Emit an event to notify that a new node has been registered.
	ctx.EventManager().EmitTypedEvent(
		&types.EventRegister{
			Address: node.Address,
		},
	)

	return &types.MsgRegisterResponse{}, nil
}

// MsgUpdateDetails updates the details of a registered node.
// It validates the update details request, checks prices, and updates the node information.
func (k *msgServer) MsgUpdateDetails(c context.Context, msg *types.MsgUpdateDetailsRequest) (*types.MsgUpdateDetailsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Check if the provided GigabytePrices are valid, if not, return an error.
	if msg.GigabytePrices != nil {
		if !k.IsValidGigabytePrices(ctx, msg.GigabytePrices) {
			return nil, types.NewErrorInvalidPrices(msg.GigabytePrices)
		}
	}

	// Check if the provided HourlyPrices are valid, if not, return an error.
	if msg.HourlyPrices != nil {
		if !k.IsValidHourlyPrices(ctx, msg.HourlyPrices) {
			return nil, types.NewErrorInvalidPrices(msg.HourlyPrices)
		}
	}

	// Convert the `msg.From` address (in Bech32 format) to a `hubtypes.NodeAddress`.
	nodeAddr, err := hubtypes.NodeAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Get the node from the Store based on the provided node address.
	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	// Update the node's GigabytePrices, HourlyPrices, and RemoteURL with the provided values.
	if msg.GigabytePrices != nil {
		node.GigabytePrices = msg.GigabytePrices
	}
	if msg.HourlyPrices != nil {
		node.HourlyPrices = msg.HourlyPrices
	}
	if msg.RemoteURL != "" {
		node.RemoteURL = msg.RemoteURL
	}

	// Save the updated node in the Store.
	k.SetNode(ctx, node)

	// Emit an event to notify that the node details have been updated.
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateDetails{
			Address: node.Address,
		},
	)

	return &types.MsgUpdateDetailsResponse{}, nil
}

// MsgUpdateStatus updates the status of a registered node.
// It validates the update status request, checks the node's current status, and updates the status and inactive time accordingly.
func (k *msgServer) MsgUpdateStatus(c context.Context, msg *types.MsgUpdateStatusRequest) (*types.MsgUpdateStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Convert the `msg.From` address (in Bech32 format) to a `hubtypes.NodeAddress`.
	nodeAddr, err := hubtypes.NodeAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Get the node from the Store based on the provided node address.
	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	// If the current status of the node is `Active`, handle the necessary updates for changing to `Inactive` status.
	if node.Status.Equal(hubtypes.StatusActive) {
		k.DeleteNodeForInactiveAt(ctx, node.InactiveAt, nodeAddr)
		if msg.Status.Equal(hubtypes.StatusInactive) {
			k.DeleteActiveNode(ctx, nodeAddr)
		}
	}

	// If the current status of the node is `Inactive`, handle the necessary updates for changing to `Active` status.
	if node.Status.Equal(hubtypes.StatusInactive) {
		if msg.Status.Equal(hubtypes.StatusActive) {
			k.DeleteInactiveNode(ctx, nodeAddr)
		}
	}

	// If the new status is `Active`, update the node's inactive time based on the active duration.
	if msg.Status.Equal(hubtypes.StatusActive) {
		node.InactiveAt = ctx.BlockTime().Add(
			k.ActiveDuration(ctx),
		)
		k.SetNodeForInactiveAt(ctx, node.InactiveAt, nodeAddr)
	}

	// If the new status is `Inactive`, set the node's inactive time to zero.
	if msg.Status.Equal(hubtypes.StatusInactive) {
		node.InactiveAt = time.Time{}
	}

	// Update the node's status and status timestamp.
	node.Status = msg.Status
	node.StatusAt = ctx.BlockTime()

	// Save the updated node in the Store.
	k.SetNode(ctx, node)

	// Emit an event to notify that the node status has been updated.
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateStatus{
			Status:  node.Status,
			Address: node.Address,
		},
	)

	return &types.MsgUpdateStatusResponse{}, nil
}

// MsgSubscribe subscribes to a node for a specific amount of gigabytes or hours.
// It validates the subscription request and creates a new subscription for the provided node and user account.
func (k *msgServer) MsgSubscribe(c context.Context, msg *types.MsgSubscribeRequest) (*types.MsgSubscribeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Check if the provided Gigabytes value is valid, if not, return an error.
	if msg.Gigabytes != 0 {
		if !k.IsValidSubscriptionGigabytes(ctx, msg.Gigabytes) {
			return nil, types.NewErrorInvalidGigabytes(msg.Gigabytes)
		}
	}

	// Check if the provided Hours value is valid, if not, return an error.
	if msg.Hours != 0 {
		if !k.IsValidSubscriptionHours(ctx, msg.Hours) {
			return nil, types.NewErrorInvalidHours(msg.Hours)
		}
	}

	// Convert the `msg.From` address (in Bech32 format) to an `sdk.AccAddress`.
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Convert the `msg.NodeAddress` (node address) to a `hubtypes.NodeAddress`.
	nodeAddr, err := hubtypes.NodeAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return nil, err
	}

	// Create a new subscription for the provided node, user account, gigabytes, hours, and denom.
	subscription, err := k.CreateSubscriptionForNode(ctx, accAddr, nodeAddr, msg.Gigabytes, msg.Hours, msg.Denom)
	if err != nil {
		return nil, err
	}

	// Emit an event to notify that a new subscription has been created.
	ctx.EventManager().EmitTypedEvent(
		&types.EventCreateSubscription{
			Address:     subscription.Address,
			NodeAddress: subscription.NodeAddress,
			ID:          subscription.ID,
		},
	)

	return &types.MsgSubscribeResponse{}, nil
}
