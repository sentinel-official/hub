package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/types"
	subscriptiontypes "github.com/sentinel-official/hub/x/subscription/types"
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

// MsgStart handles the MsgStartRequest and starts a new session for a given subscription and account.
// It takes a context (c) and a MsgStartRequest (msg) as input parameters.
// It returns a MsgStartResponse and an error.
func (k *msgServer) MsgStart(c context.Context, msg *types.MsgStartRequest) (*types.MsgStartResponse, error) {
	// Unwrap the SDK context.
	ctx := sdk.UnwrapSDKContext(c)

	// Retrieve the subscription using the given ID from the message.
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}

	// Check if the subscription status is active.
	if !subscription.GetStatus().Equal(hubtypes.StatusActive) {
		return nil, types.NewErrorInvalidSubscriptionStatus(subscription.GetID(), subscription.GetStatus())
	}

	// Parse the provided node address from Bech32 to NodeAddress format.
	nodeAddr, err := hubtypes.NodeAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	// Retrieve the node using the parsed node address.
	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	// Check if the node status is active.
	if !node.Status.Equal(hubtypes.StatusActive) {
		return nil, types.NewErrorInvalidNodeStatus(nodeAddr, node.Status)
	}

	// Check the subscription type and perform necessary validations accordingly.
	switch s := subscription.(type) {
	case *subscriptiontypes.NodeSubscription:
		// If the subscription is a NodeSubscription, verify that the provided node address matches.
		if node.Address != s.NodeAddress {
			return nil, types.NewErrorInvalidNode(node.Address)
		}
	case *subscriptiontypes.PlanSubscription:
		// If the subscription is a PlanSubscription, check if the node is associated with the plan.
		if !k.HasNodeForPlan(ctx, s.PlanID, nodeAddr) {
			return nil, types.NewErrorInvalidNode(node.Address)
		}
	default:
		// Return an error if the subscription type is unknown or unsupported.
		return nil, types.NewErrorInvalidSubscriptionType(subscription.GetID(), subscription.Type().String())
	}

	// Parse the account address from Bech32 format.
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Initialize a flag to determine if allocation should be checked.
	checkAllocation := true

	// If the subscription is a NodeSubscription and has a non-zero hours value,
	// then no allocation check is required. Also, the from address should match the subscription creator address.
	if s, ok := subscription.(*subscriptiontypes.NodeSubscription); ok {
		if s.Hours != 0 {
			checkAllocation = false
			if msg.From != s.Address {
				return nil, types.NewErrorUnauthorized(msg.From)
			}
		}
	}

	// If allocation check is required, verify that the allocation has enough bytes to start the session.
	if checkAllocation {
		alloc, found := k.GetAllocation(ctx, subscription.GetID(), accAddr)
		if !found {
			return nil, types.NewErrorAllocationNotFound(subscription.GetID(), accAddr)
		}
		if alloc.UtilisedBytes.GTE(alloc.GrantedBytes) {
			return nil, types.NewErrorInvalidAllocation(subscription.GetID(), accAddr)
		}
	}

	// Check if there is an active session for the given subscription and account.
	// If an active session is found, return an error as only one active session is allowed per account and subscription.
	session, found := k.GetActiveSessionForAllocation(ctx, subscription.GetID(), accAddr)
	if found {
		return nil, types.NewErrorDuplicateActiveSession(subscription.GetID(), accAddr, session.ID)
	}

	// Generate a new session ID by incrementing the count.
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

	// Increment the count for future sessions.
	k.SetCount(ctx, count+1)

	// Set the newly created session in the store.
	k.SetSession(ctx, session)

	// Associate the session with the account and node in the store.
	k.SetSessionForAccount(ctx, accAddr, session.ID)
	k.SetSessionForNode(ctx, nodeAddr, session.ID)

	// Associate the session with the subscription and allocation in the store.
	k.SetSessionForSubscription(ctx, subscription.GetID(), session.ID)
	k.SetSessionForAllocation(ctx, subscription.GetID(), accAddr, session.ID)

	// Associate the session with the InactiveAt time in the store.
	k.SetSessionForInactiveAt(ctx, session.InactiveAt, session.ID)

	// Emit a Start event for the newly created session using the context's EventManager.
	ctx.EventManager().EmitTypedEvent(
		&types.EventStart{
			ID:             session.ID,
			SubscriptionID: session.SubscriptionID,
			NodeAddress:    session.NodeAddress,
		},
	)

	// Return an empty response and no error to indicate a successful end of the session.
	return &types.MsgStartResponse{}, nil
}

// MsgUpdateDetails updates the details of a session based on the provided message.
// It handles session verification, updates session details, and emits an event for the update.
func (k *msgServer) MsgUpdateDetails(c context.Context, msg *types.MsgUpdateDetailsRequest) (*types.MsgUpdateDetailsResponse, error) {
	// Unwrap the SDK context.
	ctx := sdk.UnwrapSDKContext(c)

	// Get the session based on the provided session ID from the message.
	session, found := k.GetSession(ctx, msg.Proof.ID)
	if !found {
		return nil, types.NewErrorSessionNotFound(msg.Proof.ID)
	}

	// Check if the session is in the "Inactive" state, if so, return an error.
	if session.Status.Equal(hubtypes.StatusInactive) {
		return nil, types.NewErrorInvalidSessionStatus(session.ID, session.Status)
	}

	// Verify that the sender of the message is authorized to update the session.
	if msg.From != session.NodeAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// If proof verification is enabled, verify the signature provided in the message.
	if k.ProofVerificationEnabled(ctx) {
		// Get the account address associated with the session.
		accAddr := session.GetAddress()

		// Verify the signature using the account address, provided proof, and signature.
		if err := k.VerifySignature(ctx, accAddr, msg.Proof, msg.Signature); err != nil {
			return nil, types.NewErrorInvalidSignature(msg.Signature)
		}
	}

	// If the session is in the "Active" state, update the session's inactive time.
	if session.Status.Equal(hubtypes.StatusActive) {
		// Delete the session from the old inactive time.
		k.DeleteSessionForInactiveAt(ctx, session.InactiveAt, session.ID)

		// Calculate the new inactive time based on the current block time and the pending duration.
		session.InactiveAt = ctx.BlockTime().Add(
			k.InactivePendingDuration(ctx),
		)

		// Set the session for the new inactive time.
		k.SetSessionForInactiveAt(ctx, session.InactiveAt, session.ID)
	}

	// Update the session's bandwidth and duration.
	session.Bandwidth = msg.Proof.Bandwidth
	session.Duration = msg.Proof.Duration

	// Save the updated session to the state.
	k.SetSession(ctx, session)

	// Emit a typed event to signal the session details update.
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateDetails{
			ID:             session.ID,
			SubscriptionID: session.SubscriptionID,
			NodeAddress:    session.NodeAddress,
		},
	)

	// Return an empty response and no error to indicate a successful update.
	return &types.MsgUpdateDetailsResponse{}, nil
}

// MsgEnd ends an active session based on the provided message.
// It changes the session status to "InactivePending", sets a new inactive time,
// and emits an event for the status update.
func (k *msgServer) MsgEnd(c context.Context, msg *types.MsgEndRequest) (*types.MsgEndResponse, error) {
	// Unwrap the SDK context.
	ctx := sdk.UnwrapSDKContext(c)

	// Get the session based on the provided session ID from the message.
	session, found := k.GetSession(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSessionNotFound(msg.ID)
	}

	// Check if the session is in the "Active" state, if not, return an error.
	if !session.Status.Equal(hubtypes.StatusActive) {
		return nil, types.NewErrorInvalidSessionStatus(session.ID, session.Status)
	}

	// Verify that the sender of the message is authorized to end the session.
	if msg.From != session.Address {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Delete the session from the old inactive time.
	k.DeleteSessionForInactiveAt(ctx, session.InactiveAt, session.ID)

	// Calculate the new inactive time based on the current block time and the pending duration.
	session.InactiveAt = ctx.BlockTime().Add(
		k.InactivePendingDuration(ctx),
	)

	// Set the session for the new inactive time.
	k.SetSessionForInactiveAt(ctx, session.InactiveAt, session.ID)

	// Update the session status to "InactivePending".
	session.Status = hubtypes.StatusInactivePending

	// Set the status change timestamp to the current block time.
	session.StatusAt = ctx.BlockTime()

	// Save the updated session to the state.
	k.SetSession(ctx, session)

	// Emit a typed event to signal the session status update.
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateStatus{
			ID:             session.ID,
			SubscriptionID: session.SubscriptionID,
			NodeAddress:    session.NodeAddress,
			Status:         session.Status,
		},
	)

	// Return an empty response and no error to indicate a successful end of the session.
	return &types.MsgEndResponse{}, nil
}
