package plan

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/plan/keeper"
	"github.com/sentinel-official/hub/x/plan/types"
)

func HandleAddPlan(ctx sdk.Context, k keeper.Keeper, msg types.MsgAddPlan) sdk.Result {
	if !k.HasProvider(ctx, msg.From) {
		return types.ErrorProviderDoesNotExist().Result()
	}

	count := k.GetPlansCount(ctx)
	plan := types.Plan{
		ID:        count + 1,
		Provider:  msg.From,
		Price:     msg.Price,
		Validity:  msg.Validity,
		Bandwidth: msg.Bandwidth,
		Status:    hub.StatusInactive,
		StatusAt:  ctx.BlockTime(),
	}

	k.SetPlan(ctx, plan)
	k.SetPlanForProvider(ctx, plan.Provider, plan.ID)
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

	if k.HasNodeForPlan(ctx, plan.ID, node.Address) {
		return types.ErrorDuplicateNode().Result()
	}

	k.SetNodeForPlan(ctx, plan.ID, node.Address)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAddNodeForPlan,
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

	if !k.HasNodeForPlan(ctx, plan.ID, msg.Address) {
		return types.ErrorNodeWasNotAdded().Result()
	}

	k.DeleteNodeForPlan(ctx, plan.ID, msg.Address)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeRemoveNodeForPlan,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", plan.ID)),
		sdk.NewAttribute(types.AttributeKeyAddress, msg.Address.String()),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}
