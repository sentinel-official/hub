package subscription

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/subscription/keeper"
	"github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

func HandleAddPlan(ctx sdk.Context, k keeper.Keeper, msg types.MsgAddPlan) sdk.Result {
	_, found := k.GetProvider(ctx, msg.From)
	if !found {
		return types.ErrorProviderDoesNotExist().Result()
	}

	count := k.GetPlansCount(ctx)
	plan := types.Plan{
		ID:        count + 1,
		Provider:  msg.From,
		Price:     msg.Price,
		Validity:  msg.Validity,
		Bandwidth: msg.Bandwidth,
		Duration:  msg.Duration,
		Status:    hub.StatusInactive,
		StatusAt:  ctx.BlockTime(),
	}

	k.SetPlan(ctx, plan)
	k.SetPlanIDForProvider(ctx, plan.Provider, plan.ID)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetPlan,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", plan.ID)),
		sdk.NewAttribute(types.AttributeKeyAddress, plan.Provider.String()),
	))

	k.SetPlansCount(ctx, count+1)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetPlansCount,
		sdk.NewAttribute(types.AttributeKeyCount, fmt.Sprintf("%d", count+1)),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleSetPlanStatus(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetPlanStatus) sdk.Result {
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return types.ErrorPlanDoesNotExist().Result()
	}
	if !msg.From.Equals(plan.Provider) {
		return types.ErrorUnauthorized().Result()
	}

	plan.Status = msg.Status
	plan.StatusAt = ctx.BlockTime()

	k.SetPlan(ctx, plan)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetPlanStatus,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", plan.ID)),
		sdk.NewAttribute(types.AttributeKeyStatus, plan.Status.String()),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleAddNodeForPlan(ctx sdk.Context, k keeper.Keeper, msg types.MsgAddNodeForPlan) sdk.Result {
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return types.ErrorPlanDoesNotExist().Result()
	}
	if !msg.From.Equals(plan.Provider) {
		return types.ErrorUnauthorized().Result()
	}

	node, found := k.GetNode(ctx, msg.Address)
	if !found {
		return types.ErrorNodeDoesNotExist().Result()
	}
	if !msg.From.Equals(node.Provider) {
		return types.ErrorUnauthorized().Result()
	}

	if k.HasNodeAddressForPlan(ctx, plan.ID, node.Address) {
		return types.ErrorDuplicateNode().Result()
	}

	k.SetNodeAddressForPlan(ctx, plan.ID, node.Address)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetNodeAddressForPlan,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", plan.ID)),
		sdk.NewAttribute(types.AttributeKeyAddress, node.Address.String()),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleRemoveNodeForPlan(ctx sdk.Context, k keeper.Keeper, msg types.MsgRemoveNodeForPlan) sdk.Result {
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return types.ErrorPlanDoesNotExist().Result()
	}
	if !msg.From.Equals(plan.Provider) {
		return types.ErrorUnauthorized().Result()
	}

	if !k.HasNodeAddressForPlan(ctx, plan.ID, msg.Address) {
		return types.ErrorNodeWasNotAdded().Result()
	}

	k.DeleteNodeAddressForPlan(ctx, plan.ID, msg.Address)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDeleteNodeAddressForPlan,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", plan.ID)),
		sdk.NewAttribute(types.AttributeKeyAddress, msg.Address.String()),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleStartSubscription(ctx sdk.Context, k keeper.Keeper, msg types.MsgStartSubscription) sdk.Result {
	if msg.ID > 0 {
		return handleStartPlanSubscription(ctx, k, msg.From, msg.ID, msg.Denom)
	}
	return handleStartNodeSubscription(ctx, k, msg.From, msg.Address, msg.Deposit)
}

func handleStartPlanSubscription(ctx sdk.Context, k keeper.Keeper,
	from sdk.AccAddress, id uint64, denom string) sdk.Result {
	plan, found := k.GetPlan(ctx, id)
	if !found {
		return types.ErrorPlanDoesNotExist().Result()
	}
	if !plan.Status.Equal(hub.StatusActive) {
		return types.ErrorInvalidPlanStatus().Result()
	}

	price, found := plan.GetPriceForDenom(denom)
	if !found {
		return types.ErrorPriceDoesNotExist().Result()
	}

	if err := k.SendCoin(ctx, from, plan.Provider.Bytes(), price); err != nil {
		return err.Result()
	}

	count := k.GetSubscriptionsCount(ctx)
	subscription := types.Subscription{
		ID:        count + 1,
		Address:   from,
		Plan:      plan.ID,
		Duration:  plan.Duration,
		ExpiresAt: ctx.BlockTime().Add(plan.Validity),
		Node:      nil,
		Price:     sdk.Coin{},
		Deposit:   sdk.Coin{},
		Bandwidth: plan.Bandwidth,
		Status:    hub.StatusActive,
		StatusAt:  ctx.BlockTime(),
	}

	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionIDForAddress(ctx, subscription.Address, subscription.ID)
	k.SetSubscriptionIDForPlan(ctx, subscription.Plan, subscription.ID)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetSubscription,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", subscription.ID)),
		sdk.NewAttribute(types.AttributeKeyPlan, fmt.Sprintf("%d", subscription.Plan)),
		sdk.NewAttribute(types.AttributeKeyAddress, subscription.Address.String()),
	))

	k.SetSubscriptionsCount(ctx, count+1)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetSubscriptionsCount,
		sdk.NewAttribute(types.AttributeKeyCount, fmt.Sprintf("%d", count+1)),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleStartNodeSubscription(ctx sdk.Context, k keeper.Keeper,
	from sdk.AccAddress, address hub.NodeAddress, deposit sdk.Coin) sdk.Result {
	node, found := k.GetNode(ctx, address)
	if !found {
		return types.ErrorNodeDoesNotExist().Result()
	}
	if !node.Status.Equal(hub.StatusActive) {
		return types.ErrorInvalidNodeStatus().Result()
	}

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleEndSubscription(ctx sdk.Context, k keeper.Keeper, msg types.MsgEndSubscription) sdk.Result {
	return sdk.Result{Events: ctx.EventManager().Events()}
}
