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
		return types.ErrorNoProviderFound().Result()
	}

	count := k.GetPlansCount(ctx, msg.From)
	plan := types.Plan{
		ID:           count,
		Provider:     msg.From,
		Price:        msg.Price,
		Duration:     msg.Duration,
		MaxBandwidth: msg.MaxBandwidth,
		MaxDuration:  msg.MaxDuration,
		Status:       hub.StatusInactive,
		StatusAt:     ctx.BlockHeight(),
	}

	k.SetPlan(ctx, plan)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetPlan,
		sdk.NewAttribute(types.AttributeKeyAddress, plan.Provider.String()),
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", plan.ID)),
	))

	k.SetPlansCount(ctx, plan.Provider, count+1)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetPlansCount,
		sdk.NewAttribute(types.AttributeKeyAddress, plan.Provider.String()),
		sdk.NewAttribute(types.AttributeKeyCount, fmt.Sprintf("%d", count+1)),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleSetPlanStatus(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetPlanStatus) sdk.Result {
	plan, found := k.GetPlan(ctx, msg.From, msg.ID)
	if !found {
		return types.ErrorNoPlanFound().Result()
	}

	plan.Status = msg.Status
	plan.StatusAt = ctx.BlockHeight()

	k.SetPlan(ctx, plan)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetPlanStatus,
		sdk.NewAttribute(types.AttributeKeyAddress, plan.Provider.String()),
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", plan.ID)),
		sdk.NewAttribute(types.AttributeKeyStatus, plan.Status.String()),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleAddNode(ctx sdk.Context, k keeper.Keeper, msg types.MsgAddNode) sdk.Result {
	plan, found := k.GetPlan(ctx, msg.From, msg.ID)
	if !found {
		return ErrorNoPlanFound().Result()
	}

	node, found := k.GetNode(ctx, msg.Address)
	if !found {
		return types.ErrorNoNodeFound().Result()
	}

	if !plan.Provider.Equals(node.Provider) {
		return types.ErrorUnauthorized().Result()
	}

	if k.HasNodeAddressForPlan(ctx, plan.Provider, plan.ID, node.Address) {
		return types.ErrorDuplicateNode().Result()
	}

	k.SetNodeAddressForPlan(ctx, plan.Provider, plan.ID, node.Address)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetNodeAddressForPlan,
		sdk.NewAttribute(types.AttributeKeyAddress, plan.Provider.String()),
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", plan.ID)),
		sdk.NewAttribute(types.AttributeKeyAddress, node.Address.String()),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleRemoveNode(ctx sdk.Context, k keeper.Keeper, msg types.MsgRemoveNode) sdk.Result {
	node, found := k.GetNode(ctx, msg.Address)
	if !found {
		return types.ErrorNoNodeFound().Result()
	}

	plan, found := k.GetPlan(ctx, msg.From, msg.ID)
	if !found {
		return ErrorNoPlanFound().Result()
	}

	if !plan.Provider.Equals(node.Provider) {
		return types.ErrorUnauthorized().Result()
	}

	if !k.HasNodeAddressForPlan(ctx, plan.Provider, plan.ID, node.Address) {
		return types.ErrorNoNodeAdded().Result()
	}

	k.DeleteNodeAddressForPlan(ctx, plan.Provider, plan.ID, node.Address)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDeleteNodeAddressForPlan,
		sdk.NewAttribute(types.AttributeKeyAddress, plan.Provider.String()),
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", plan.ID)),
		sdk.NewAttribute(types.AttributeKeyAddress, node.Address.String()),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}
