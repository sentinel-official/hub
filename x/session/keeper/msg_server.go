package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/types"
	subscriptiontypes "github.com/sentinel-official/hub/x/subscription/types"
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
func NewMsgServiceServer(keeper Keeper) types.MsgServiceServer {
	return &msgServer{Keeper: keeper}
}

// MsgStart starts a new session for a subscription.
// It validates the start request, checks subscription and node status, and creates a new session.
func (k *msgServer) MsgStart(c context.Context, msg *types.MsgStartRequest) (*types.MsgStartResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Get the subscription from the Store based on the provided subscription ID (msg.ID).
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}

	// Check if the subscription is in an active state. If not, return an error.
	if !subscription.GetStatus().Equal(hubtypes.StatusActive) {
		return nil, types.NewErrorInvalidSubscriptionStatus(subscription.GetID(), subscription.GetStatus())
	}

	// Convert the `msg.Address` (node's address) from Bech32 format to a `hubtypes.NodeAddress`.
	nodeAddr, err := hubtypes.NodeAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	// Get the node from the Store based on the provided node address.
	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	// Check if the node is in an active state. If not, return an error.
	if !node.Status.Equal(hubtypes.StatusActive) {
		return nil, types.NewErrorInvalidNodeStatus(nodeAddr, node.Status)
	}

	// Validate the association between the subscription and the node.
	// Depending on the subscription type, it should either match the node address or be associated with the plan.
	switch s := subscription.(type) {
	case *subscriptiontypes.NodeSubscription:
		if node.Address != s.NodeAddress {
			return nil, types.NewErrorInvalidNode(node.Address)
		}
	case *subscriptiontypes.PlanSubscription:
		if !k.HasNodeForPlan(ctx, s.PlanID, nodeAddr) {
			return nil, types.NewErrorInvalidNode(node.Address)
		}
	default:
		return nil, types.NewErrorInvalidSubscriptionType(subscription.GetID(), subscription.Type().String())
	}

	// Convert the `msg.From` address (in Bech32 format) to an `sdk.AccAddress`.
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Check if an allocation check is required.
	checkAllocation := true
	if s, ok := subscription.(*subscriptiontypes.NodeSubscription); ok {
		if s.Hours != 0 {
			// If the subscription is for hourly usage, no allocation check is needed.
			checkAllocation = false
			if msg.From != s.Address {
				return nil, types.NewErrorUnauthorized(msg.From)
			}
		}
	}

	// If an allocation check is required, verify that the user has sufficient allocated bandwidth.
	if checkAllocation {
		alloc, found := k.GetAllocation(ctx, subscription.GetID(), accAddr)
		if !found {
			return nil, types.NewErrorAllocationNotFound(subscription.GetID(), accAddr)
		}
		if alloc.UtilisedBytes.GTE(alloc.GrantedBytes) {
			return nil, types.NewErrorInvalidAllocation(subscription.GetID(), accAddr)
		}
	}

	// Check if an active session already exists for the same subscription and user.
	// If found, return an error to prevent multiple active sessions for the same subscription and user.
	session, found := k.GetActiveSessionForAllocation(ctx, subscription.GetID(), accAddr)
	if found {
		return nil, types.NewErrorDuplicateActiveSession(subscription.GetID(), accAddr, session.ID)
	}

	// Create a new session based on the provided parameters.
	count := k.GetCount(ctx)
	session = types.Session{
		ID:             count + 1,
		SubscriptionID: subscription.GetID(),
		NodeAddress:    nodeAddr.String(),
		Address:        accAddr.String(),
		Bandwidth:      hubtypes.NewBandwidthFromInt64(0, 0),
		Duration:       0,
		InactiveAt: ctx.BlockTime().Add(
			k.InactivePendingDuration(ctx),
		),
		Status:   hubtypes.StatusActive,
		StatusAt: ctx.BlockTime(),
	}

	// Update the count in the Store and set the new session.
	k.SetCount(ctx, count+1)
	k.SetSession(ctx, session)
	k.SetSessionForAccount(ctx, accAddr, session.ID)
	k.SetSessionForNode(ctx, nodeAddr, session.ID)
	k.SetSessionForSubscription(ctx, subscription.GetID(), session.ID)
	k.SetSessionForAllocation(ctx, subscription.GetID(), accAddr, session.ID)
	k.SetSessionForInactiveAt(ctx, session.InactiveAt, session.ID)

	// Emit an event to notify that a new session has started.
	ctx.EventManager().EmitTypedEvent(
		&types.EventStart{
			ID:             session.ID,
			SubscriptionID: session.SubscriptionID,
			NodeAddress:    session.NodeAddress,
		},
	)

	return &types.MsgStartResponse{}, nil
}

// MsgUpdateDetails updates the details of an active session.
// It validates the update details request, verifies the signature if proof verification is enabled,
// and updates the bandwidth and duration of the session.
func (k *msgServer) MsgUpdateDetails(c context.Context, msg *types.MsgUpdateDetailsRequest) (*types.MsgUpdateDetailsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Get the session from the Store based on the provided session ID (msg.Proof.ID).
	session, found := k.GetSession(ctx, msg.Proof.ID)
	if !found {
		return nil, types.NewErrorSessionNotFound(msg.Proof.ID)
	}

	// Check if the session is in an active state. If not, return an error.
	if session.Status.Equal(hubtypes.StatusInactive) {
		return nil, types.NewErrorInvalidSessionStatus(session.ID, session.Status)
	}

	// Verify that the `msg.From` address matches the node address of the session. If not, return an error.
	if msg.From != session.NodeAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// If proof verification is enabled, verify the signature using the associated account address.
	if k.ProofVerificationEnabled(ctx) {
		accAddr := session.GetAddress()
		if err := k.VerifySignature(ctx, accAddr, msg.Proof, msg.Signature); err != nil {
			return nil, types.NewErrorInvalidSignature(msg.Signature)
		}
	}

	// If the session is currently active, delete it from the inactive time index.
	if session.Status.Equal(hubtypes.StatusActive) {
		k.DeleteSessionForInactiveAt(ctx, session.InactiveAt, session.ID)

		// Update the session's inactive time based on the inactive-pending duration.
		session.InactiveAt = ctx.BlockTime().Add(
			k.InactivePendingDuration(ctx),
		)
		k.SetSessionForInactiveAt(ctx, session.InactiveAt, session.ID)
	}

	// Update the bandwidth and duration of the session using the provided proof.
	session.Bandwidth = msg.Proof.Bandwidth
	session.Duration = msg.Proof.Duration

	// Update the session in the Store.
	k.SetSession(ctx, session)

	// Emit an event to notify that the session details have been updated.
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateDetails{
			ID:             session.ID,
			SubscriptionID: session.SubscriptionID,
			NodeAddress:    session.NodeAddress,
		},
	)

	return &types.MsgUpdateDetailsResponse{}, nil
}

// MsgEnd ends an active session.
// It validates the end request, updates the session status to inactive-pending, and sets the inactive time.
func (k *msgServer) MsgEnd(c context.Context, msg *types.MsgEndRequest) (*types.MsgEndResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Get the session from the Store based on the provided session ID (msg.ID).
	session, found := k.GetSession(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSessionNotFound(msg.ID)
	}

	// Check if the session is in an active state. If not, return an error.
	if !session.Status.Equal(hubtypes.StatusActive) {
		return nil, types.NewErrorInvalidSessionStatus(session.ID, session.Status)
	}

	// Verify that the `msg.From` address matches the user address of the session. If not, return an error.
	if msg.From != session.Address {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Delete the session from the inactive time index and update the inactive time based on the inactive-pending duration.
	k.DeleteSessionForInactiveAt(ctx, session.InactiveAt, session.ID)
	session.InactiveAt = ctx.BlockTime().Add(
		k.InactivePendingDuration(ctx),
	)
	k.SetSessionForInactiveAt(ctx, session.InactiveAt, session.ID)

	// Update the session status to inactive-pending and set the status timestamp.
	session.Status = hubtypes.StatusInactivePending
	session.StatusAt = ctx.BlockTime()

	// Update the session in the Store.
	k.SetSession(ctx, session)

	// Emit an event to notify that the session status has been updated.
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateStatus{
			ID:             session.ID,
			SubscriptionID: session.SubscriptionID,
			NodeAddress:    session.NodeAddress,
			Status:         session.Status,
		},
	)

	return &types.MsgEndResponse{}, nil
}
