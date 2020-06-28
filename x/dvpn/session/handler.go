package session

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/session/keeper"
	"github.com/sentinel-official/hub/x/dvpn/session/types"
)

func isAuthorized(ctx sdk.Context, k keeper.Keeper, p, s uint64, n hub.NodeAddress) bool {
	return p == 0 && k.HasSubscriptionIDForNode(ctx, n, s) ||
		k.HasNodeAddressForPlan(ctx, p, n)
}

func HandleUpdateSession(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpdateSession) sdk.Result {
	subscription, found := k.GetSubscription(ctx, msg.Subscription)
	if !found {
		return types.ErrorSubscriptionDoesNotExit().Result()
	}
	if !subscription.Status.Equal(hub.StatusActive) {
		return types.ErrorInvalidSubscriptionStatus().Result()
	}
	if !isAuthorized(ctx, k, subscription.Plan, subscription.ID, msg.From) {
		return types.ErrorUnauthorized().Result()
	}
	if !k.HasAddressForSubscriptionID(ctx, subscription.ID, msg.Address) {
		return types.ErrorAddressWasNotAdded().Result()
	}

	subscription.Bandwidth = subscription.Bandwidth.Add(msg.Bandwidth)
	if subscription.Bandwidth.IsAnyGT(subscription.TotalBandwidth) {
		return types.ErrorInvalidBandwidth().Result()
	}

	session, found := k.GetActiveSession(ctx, subscription.ID, msg.From, msg.Address)
	if !found {
		count := k.GetSessionsCount(ctx)
		session = types.Session{
			ID:           count + 1,
			Subscription: subscription.ID,
			Node:         msg.From,
			Address:      msg.Address,
			Bandwidth:    hub.NewBandwidthFromInt64(0, 0),
			Status:       hub.StatusActive,
			StatusAt:     ctx.BlockTime(),
		}

		k.SetSessionsCount(ctx, count+1)
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeSetSessionsCount,
			sdk.NewAttribute(types.AttributeKeyCount, fmt.Sprintf("%d", count+1)),
		))

		k.SetActiveSessionID(ctx, session.Subscription, session.Node, session.Address, session.ID)
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeSetActiveSession,
			sdk.NewAttribute(types.AttributeKeySubscription, fmt.Sprintf("%d", session.Subscription)),
			sdk.NewAttribute(types.AttributeKeyNode, session.Node.String()),
			sdk.NewAttribute(types.AttributeKeyAddress, session.Address.String()),
			sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", session.ID)),
		))
	}

	session.Bandwidth = session.Bandwidth.Add(msg.Bandwidth)

	k.SetSession(ctx, session)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUpdateSession,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", session.ID)),
	))

	k.SetSubscription(ctx, subscription)
	return sdk.Result{Events: ctx.EventManager().Events()}
}
