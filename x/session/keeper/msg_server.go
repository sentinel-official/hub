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
	// Unwrap the SDK context from the standard context.
	ctx := sdk.UnwrapSDKContext(c)

	// Get the subscription from the store using the provided subscription ID.
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		// If the subscription is not found, return an error indicating that the subscription was not found.
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}

	// Check if the subscription status is 'Active' as only active subscriptions can start sessions.
	if !subscription.GetStatus().Equal(hubtypes.StatusActive) {
		// If the subscription status is not 'Active', return an error indicating that the subscription status is invalid for starting a session.
		return nil, types.NewErrorInvalidSubscriptionStatus(subscription.GetID(), subscription.GetStatus())
	}

	// Parse the node address from the Bech32 encoded address provided in the message.
	nodeAddr, err := hubtypes.NodeAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	// Get the node from the store using the parsed node address.
	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		// If the node is not found, return an error indicating that the node was not found.
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	// Check if the node status is 'Active' as only active nodes can be used for starting a session.
	if !node.Status.Equal(hubtypes.StatusActive) {
		// If the node status is not 'Active', return an error indicating that the node status is invalid for starting a session.
		return nil, types.NewErrorInvalidNodeStatus(nodeAddr, node.Status)
	}

	// Based on the type of subscription, perform additional checks on the node and subscription relationship.
	switch s := subscription.(type) {
	case *subscriptiontypes.NodeSubscription:
		// For node-level subscriptions, ensure that the node address in the subscription matches the provided node address.
		if node.Address != s.NodeAddress {
			return nil, types.NewErrorInvalidNode(node.Address)
		}
	case *subscriptiontypes.PlanSubscription:
		// For plan-level subscriptions, check if the node is associated with the plan.
		if !k.HasNodeForPlan(ctx, s.PlanID, nodeAddr) {
			return nil, types.NewErrorInvalidNode(node.Address)
		}
	default:
		// If the subscription type is not recognized, return an error indicating an invalid subscription type.
		return nil, types.NewErrorInvalidSubscriptionType(subscription.GetID(), subscription.Type().String())
	}

	// Parse the account address from the Bech32 encoded address provided in the message.
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Determine if an allocation check is required based on the subscription type.
	checkAllocation := true
	if s, ok := subscription.(*subscriptiontypes.NodeSubscription); ok {
		if s.Hours != 0 {
			// If the subscription's duration is specified in hours (non-zero), no allocation check is needed.
			checkAllocation = false
			// Additionally, check if the message sender matches the subscription's address to prevent unauthorized session starts.
			if msg.From != s.Address {
				return nil, types.NewErrorUnauthorized(msg.From)
			}
		}
	}

	if checkAllocation {
		// If an allocation check is required, get the allocation associated with the subscription and account.
		alloc, found := k.GetAllocation(ctx, subscription.GetID(), accAddr)
		if !found {
			// If the allocation is not found, return an error indicating that the allocation was not found for the given subscription and account.
			return nil, types.NewErrorAllocationNotFound(subscription.GetID(), accAddr)
		}

		// Check if the allocation's utilized bandwidth exceeds the granted bandwidth.
		if alloc.UtilisedBytes.GTE(alloc.GrantedBytes) {
			// If the allocation's bandwidth is fully utilized, return an error indicating an invalid allocation.
			return nil, types.NewErrorInvalidAllocation(subscription.GetID(), accAddr)
		}
	}

	// Check if there is already an active session for the given subscription and account.
	session, found := k.GetLatestSessionForAllocation(ctx, subscription.GetID(), accAddr)
	if found && session.Status.Equal(hubtypes.StatusActive) {
		// If an active session already exists, return an error indicating a duplicate active session.
		return nil, types.NewErrorDuplicateActiveSession(subscription.GetID(), accAddr, session.ID)
	}

	// Increment the session count to assign a new session ID.
	count := k.GetCount(ctx)
	session = types.Session{
		ID:             count + 1,
		SubscriptionID: subscription.GetID(),
		NodeAddress:    nodeAddr.String(),
		Address:        accAddr.String(),
		Bandwidth:      hubtypes.NewBandwidthFromInt64(0, 0),
		Duration:       0,
		InactiveAt: ctx.BlockTime().Add(
			k.StatusChangeDelay(ctx),
		),
		Status:   hubtypes.StatusActive,
		StatusAt: ctx.BlockTime(),
	}

	// Save the new session to the store.
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

	// Return an empty MsgStartResponse, indicating the successful completion of the message.
	return &types.MsgStartResponse{}, nil
}

// MsgUpdateDetails updates the details of an active session.
// It validates the update details request, verifies the signature if proof verification is enabled,
// and updates the bandwidth and duration of the session.
func (k *msgServer) MsgUpdateDetails(c context.Context, msg *types.MsgUpdateDetailsRequest) (*types.MsgUpdateDetailsResponse, error) {
	// Unwrap the SDK context from the standard context.
	ctx := sdk.UnwrapSDKContext(c)

	// Get the session from the store using the provided session ID.
	session, found := k.GetSession(ctx, msg.Proof.ID)
	if !found {
		// If the session is not found, return an error indicating that the session was not found.
		return nil, types.NewErrorSessionNotFound(msg.Proof.ID)
	}

	// Check if the session status is 'Inactive' as only active or inactive-pending sessions can be updated.
	if session.Status.Equal(hubtypes.StatusInactive) {
		// If the session status is 'Inactive', return an error indicating that the session status is invalid for updating details.
		return nil, types.NewErrorInvalidSessionStatus(session.ID, session.Status)
	}

	// Ensure that the message sender (msg.From) is authorized to update the session details.
	if msg.From != session.NodeAddress {
		// If the message sender is not authorized, return an error indicating unauthorized access.
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// If proof verification is enabled, verify the signature of the message using the account address associated with the session.
	if k.ProofVerificationEnabled(ctx) {
		accAddr := session.GetAddress()
		if err := k.VerifySignature(ctx, accAddr, msg.Proof, msg.Signature); err != nil {
			// If the signature verification fails, return an error indicating an invalid signature.
			return nil, types.NewErrorInvalidSignature(msg.Signature)
		}
	}

	// If the session status is 'Active', update the session's InactiveAt value based on the status change delay.
	if session.Status.Equal(hubtypes.StatusActive) {
		// Delete the session's entry from the InactiveAt index before updating the InactiveAt value.
		k.DeleteSessionForInactiveAt(ctx, session.InactiveAt, session.ID)

		// Calculate the new InactiveAt value by adding the status change delay to the current block time.
		session.InactiveAt = ctx.BlockTime().Add(
			k.StatusChangeDelay(ctx),
		)

		// Update the session entry in the InactiveAt index with the new InactiveAt value.
		k.SetSessionForInactiveAt(ctx, session.InactiveAt, session.ID)
	}

	// Update the session's bandwidth and duration using the details from the provided proof.
	session.Bandwidth = msg.Proof.Bandwidth
	session.Duration = msg.Proof.Duration

	// Save the updated session to the store.
	k.SetSession(ctx, session)

	// Emit an event to notify that the session details have been updated.
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateDetails{
			ID:             session.ID,
			SubscriptionID: session.SubscriptionID,
			NodeAddress:    session.NodeAddress,
		},
	)

	// Return an empty MsgUpdateDetailsResponse, indicating the successful completion of the message.
	return &types.MsgUpdateDetailsResponse{}, nil
}

// MsgEnd ends an active session.
// It validates the end request, updates the session status to inactive-pending, and sets the inactive time.
func (k *msgServer) MsgEnd(c context.Context, msg *types.MsgEndRequest) (*types.MsgEndResponse, error) {
	// Unwrap the SDK context from the standard context.
	ctx := sdk.UnwrapSDKContext(c)

	// Get the session from the store using the provided session ID.
	session, found := k.GetSession(ctx, msg.ID)
	if !found {
		// If the session is not found, return an error indicating that the session was not found.
		return nil, types.NewErrorSessionNotFound(msg.ID)
	}

	// Check if the session status is 'Active' as only active sessions can be ended.
	if !session.Status.Equal(hubtypes.StatusActive) {
		// If the session status is not 'Active', return an error indicating that the session status is invalid.
		return nil, types.NewErrorInvalidSessionStatus(session.ID, session.Status)
	}

	// Ensure that the message sender (msg.From) is authorized to end the session.
	if msg.From != session.Address {
		// If the message sender is not authorized, return an error indicating unauthorized access.
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Delete the session's entry from the InactiveAt index before updating the InactiveAt value.
	k.DeleteSessionForInactiveAt(ctx, session.InactiveAt, session.ID)

	// Calculate the new InactiveAt value by adding the status change delay to the current block time.
	session.InactiveAt = ctx.BlockTime().Add(
		k.StatusChangeDelay(ctx),
	)

	// Update the session entry in the InactiveAt index with the new InactiveAt value.
	k.SetSessionForInactiveAt(ctx, session.InactiveAt, session.ID)

	// Set the session status to 'InactivePending' to mark it for an upcoming status update.
	session.Status = hubtypes.StatusInactivePending

	// Record the time of the status update in 'StatusAt' field.
	session.StatusAt = ctx.BlockTime()

	// Update the session entry in the store with the new status and status update time.
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

	// Return an empty MsgEndResponse, indicating the successful completion of the message.
	return &types.MsgEndResponse{}, nil
}
