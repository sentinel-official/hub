package session

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/types"
)

func authorized(ctx sdk.Context, k keeper.Keeper, p uint64, n hub.NodeAddress, s uint64) bool {
	if p == 0 {
		return k.HasSubscriptionForNode(ctx, n, s)
	}

	return k.HasNodeForPlan(ctx, p, n)
}

func HandleUpsert(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpsert) sdk.Result {
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return types.ErrorSubscriptionDoesNotExit().Result()
	}
	if subscription.Status.Equal(hub.StatusInactive) {
		return types.ErrorInvalidSubscriptionStatus().Result()
	}

	if !authorized(ctx, k, subscription.Plan, msg.From, subscription.ID) {
		return types.ErrorUnauthorized().Result()
	}

	if !k.HasQuota(ctx, subscription.ID, msg.Address) {
		return types.ErrorQuotaDoesNotExist().Result()
	}

	session, found := k.GetOngoingSession(ctx, subscription.ID, msg.Address)
	if found {
		k.DeleteActiveSessionAt(ctx, session.StatusAt, session.ID)
		if !session.Node.Equals(msg.From) {
			session.Status = hub.StatusInactive
			session.StatusAt = ctx.BlockTime()
			k.SetSession(ctx, session)

			found = false
		}
	}

	if !found {
		count := k.GetSessionsCount(ctx)
		session = types.Session{
			ID:           count + 1,
			Subscription: subscription.ID,
			Node:         msg.From,
			Address:      msg.Address,
			Duration:     0,
			Bandwidth:    hub.NewBandwidthFromInt64(0, 0),
			Status:       hub.StatusActive,
			StatusAt:     ctx.BlockTime(),
		}

		k.SetSessionsCount(ctx, count+1)
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeSetCount,
			sdk.NewAttribute(types.AttributeKeyCount, fmt.Sprintf("%d", count+1)),
		))

		k.SetSessionForSubscription(ctx, session.Subscription, session.ID)
		k.SetSessionForNode(ctx, session.Node, session.ID)
		k.SetSessionForAddress(ctx, session.Address, session.ID)
		k.SetOngoingSession(ctx, session.Subscription, session.Address, session.ID)
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeSetActive,
			sdk.NewAttribute(types.AttributeKeySubscription, fmt.Sprintf("%d", session.Subscription)),
			sdk.NewAttribute(types.AttributeKeyAddress, session.Address.String()),
			sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", session.ID)),
		))
	}

	session.Duration = session.Duration + msg.Duration
	session.Bandwidth = session.Bandwidth.Add(msg.Bandwidth)
	session.Status = hub.StatusActive
	session.StatusAt = ctx.BlockTime()

	k.SetSession(ctx, session)
	k.SetActiveSessionAt(ctx, session.StatusAt, session.ID)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUpdate,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", session.ID)),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return sdk.Result{Events: ctx.EventManager().Events()}
}
