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

func (k *msgServer) MsgStart(c context.Context, msg *types.MsgStartRequest) (*types.MsgStartResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}
	if !subscription.GetStatus().Equal(hubtypes.StatusActive) {
		return nil, types.NewErrorInvalidSubscriptionStatus(subscription.GetID(), subscription.GetStatus())
	}

	nodeAddr, err := hubtypes.NodeAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}
	if !node.Status.Equal(hubtypes.StatusActive) {
		return nil, types.NewErrorInvalidNodeStatus(nodeAddr, node.Status)
	}

	switch v := subscription.(type) {
	case *subscriptiontypes.NodeSubscription:
		if node.Address != v.NodeAddress {
			return nil, types.NewErrorInvalidNode(node.Address)
		}
	case *subscriptiontypes.PlanSubscription:
		if !k.HasNodeForPlan(ctx, v.PlanID, nodeAddr) {
			return nil, types.NewErrorInvalidNode(node.Address)
		}
	default:
		return nil, types.NewErrorInvalidSubscriptionType(subscription.GetID(), subscription.Type().String())
	}

	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	var id uint64
	k.IterateSessionsForQuota(ctx, subscription.GetID(), accAddr, func(_ int, item types.Session) bool {
		if item.Status.Equal(hubtypes.StatusActive) {
			id = item.ID
			return true
		}

		return false
	})

	if id > 0 {
		return nil, types.NewErrorDuplicateSession(subscription.GetID(), accAddr, id)
	}

	quota, found := k.GetQuota(ctx, subscription.GetID(), accAddr)
	if !found {
		return nil, types.NewErrorQuotaNotFound(subscription.GetID(), accAddr)
	}
	if quota.UtilisedBytes.GTE(quota.GrantedBytes) {
		return nil, types.NewErrorInvalidQuota(subscription.GetID(), accAddr)
	}

	var (
		count   = k.GetCount(ctx)
		session = types.Session{
			ID:             count + 1,
			SubscriptionID: subscription.GetID(),
			NodeAddress:    nodeAddr.String(),
			Address:        accAddr.String(),
			Bandwidth:      hubtypes.NewBandwidthFromInt64(0, 0),
			Duration:       0,
			ExpiryAt: ctx.BlockTime().Add(
				k.InactiveDuration(ctx),
			),
			Status:   hubtypes.StatusActive,
			StatusAt: ctx.BlockTime(),
		}
	)

	k.SetCount(ctx, count+1)
	k.SetSession(ctx, session)
	k.SetSessionForAccount(ctx, accAddr, session.ID)
	k.SetSessionForNode(ctx, nodeAddr, session.ID)
	k.SetSessionForSubscription(ctx, subscription.GetID(), session.ID)
	k.SetSessionForQuota(ctx, subscription.GetID(), accAddr, session.ID)
	k.SetSessionForExpiryAt(ctx, session.ExpiryAt, session.ID)
	ctx.EventManager().EmitTypedEvent(
		&types.EventStart{
			ID:             session.ID,
			SubscriptionID: session.SubscriptionID,
			NodeAddress:    session.NodeAddress,
		},
	)

	return &types.MsgStartResponse{}, nil
}

func (k *msgServer) MsgUpdateDetails(c context.Context, msg *types.MsgUpdateDetailsRequest) (*types.MsgUpdateDetailsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	session, found := k.GetSession(ctx, msg.Proof.ID)
	if !found {
		return nil, types.NewErrorSessionNotFound(msg.Proof.ID)
	}
	if !session.Status.IsOneOf(hubtypes.StatusActive, hubtypes.StatusInactivePending) {
		return nil, types.NewErrorInvalidSessionStatus(session.ID, session.Status)
	}
	if msg.From != session.NodeAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	if k.ProofVerificationEnabled(ctx) {
		accAddr := session.GetAddress()
		if err := k.VerifySignature(ctx, accAddr, msg.Proof, msg.Signature); err != nil {
			return nil, types.NewErrorInvalidSignature(msg.Signature)
		}
	}

	if session.Status.Equal(hubtypes.StatusActive) {
		k.DeleteSessionForExpiryAt(ctx, session.ExpiryAt, session.ID)

		session.ExpiryAt = ctx.BlockTime().Add(
			k.InactiveDuration(ctx),
		)
		k.SetSessionForExpiryAt(ctx, session.ExpiryAt, session.ID)
	}

	session.Bandwidth = msg.Proof.Bandwidth
	session.Duration = msg.Proof.Duration

	k.SetSession(ctx, session)
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateDetails{
			ID:             session.ID,
			SubscriptionID: session.SubscriptionID,
			NodeAddress:    session.NodeAddress,
		},
	)

	return &types.MsgUpdateDetailsResponse{}, nil
}

func (k *msgServer) MsgEnd(c context.Context, msg *types.MsgEndRequest) (*types.MsgEndResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	session, found := k.GetSession(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSessionNotFound(msg.ID)
	}
	if !session.Status.Equal(hubtypes.StatusActive) {
		return nil, types.NewErrorInvalidSessionStatus(session.ID, session.Status)
	}
	if msg.From != session.Address {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	k.DeleteSessionForExpiryAt(ctx, session.ExpiryAt, session.ID)

	session.ExpiryAt = ctx.BlockTime().Add(
		k.InactiveDuration(ctx),
	)
	k.SetSessionForExpiryAt(ctx, session.ExpiryAt, session.ID)

	session.Status = hubtypes.StatusInactivePending
	session.StatusAt = ctx.BlockTime()

	k.SetSession(ctx, session)
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
