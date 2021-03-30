package session

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/types"
)

func isAuthorized(ctx sdk.Context, k keeper.Keeper, plan, subscription uint64, node hub.NodeAddress) bool {
	if plan == 0 {
		return k.HasSubscriptionForNode(ctx, node, subscription)
	}

	return k.HasNodeForPlan(ctx, plan, node)
}

func HandleUpsert(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpsert) (*sdk.Result, error) {
	subscription, found := k.GetSubscription(ctx, msg.Proof.Subscription)
	if !found {
		return nil, types.ErrorSubscriptionDoesNotExit
	}
	if subscription.Status.Equal(hub.StatusInactive) {
		return nil, types.ErrorInvalidSubscriptionStatus
	}

	if !isAuthorized(ctx, k, subscription.Plan, subscription.ID, msg.Proof.Node) {
		return nil, types.ErrorUnauthorized
	}

	if !k.HasQuota(ctx, subscription.ID, msg.Address) {
		return nil, types.ErrorQuotaDoesNotExist
	}

	if k.ProofVerificationEnabled(ctx) {
		channel := k.GetChannel(ctx, msg.Address, msg.Proof.Subscription, msg.Proof.Node)
		if msg.Proof.Channel != channel {
			return nil, types.ErrorInvalidChannel
		}

		if err := k.VerifyProof(ctx, msg.Address, msg.Proof, msg.Signature); err != nil {
			return nil, types.ErrorFailedToVerifyProof
		}
	}

	session, found := k.GetActiveSessionForAddress(ctx, msg.Address, subscription.ID, msg.Proof.Node)
	if found {
		k.DeleteActiveSessionAt(ctx, session.StatusAt, session.ID)
	}

	if !found {
		count := k.GetCount(ctx)
		session = types.Session{
			ID:           count + 1,
			Subscription: subscription.ID,
			Node:         msg.Proof.Node,
			Address:      msg.Address,
			Duration:     0,
			Bandwidth:    hub.NewBandwidthFromInt64(0, 0),
			Status:       hub.StatusActive,
			StatusAt:     ctx.BlockTime(),
		}

		k.SetCount(ctx, count+1)
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeSetCount,
			sdk.NewAttribute(types.AttributeKeyCount, fmt.Sprintf("%d", count+1)),
		))

		k.SetSessionForSubscription(ctx, session.Subscription, session.ID)
		k.SetSessionForNode(ctx, session.Node, session.ID)
		k.SetSessionForAddress(ctx, session.Address, session.ID)
		k.SetActiveSessionForAddress(ctx, session.Address, session.Subscription, session.Node, session.ID)
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeSetActive,
			sdk.NewAttribute(types.AttributeKeySubscription, fmt.Sprintf("%d", session.Subscription)),
			sdk.NewAttribute(types.AttributeKeyAddress, session.Address.String()),
			sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", session.ID)),
		))
	}

	session.Duration = msg.Proof.Duration
	session.Bandwidth = msg.Proof.Bandwidth
	session.Status = hub.StatusActive
	session.StatusAt = ctx.BlockTime()

	k.SetSession(ctx, session)
	k.SetActiveSessionAt(ctx, session.StatusAt, session.ID)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUpdate,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", session.ID)),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
