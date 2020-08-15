package session

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/types"
)

func isAuthorized(ctx sdk.Context, k keeper.Keeper, p uint64, n hub.NodeAddress, s uint64) bool {
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

	// Check whether the msg.From is authorized to upsert the session or not
	// msg.From is authorized only if the node belongs to the subscription or to the plan.
	if !isAuthorized(ctx, k, subscription.Plan, msg.From, subscription.ID) {
		return types.ErrorUnauthorized().Result()
	}

	quota, found := k.GetQuota(ctx, subscription.ID, msg.Address)
	if !found {
		return types.ErrorQuotaDoesNotExist().Result()
	}

	quota.Consumed = quota.Consumed.Add(msg.Bandwidth)
	if quota.Consumed.IsAnyGT(quota.Allocated) {
		return types.ErrorInvalidBandwidth().Result()
	}

	k.SetQuota(ctx, subscription.ID, quota)

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

	k.SetSession(ctx, session)
	k.SetActiveSessionAt(ctx, session.StatusAt, session.ID)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUpdate,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", session.ID)),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return sdk.Result{Events: ctx.EventManager().Events()}
}
