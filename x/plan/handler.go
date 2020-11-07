package plan

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/plan/keeper"
	"github.com/sentinel-official/hub/x/plan/types"
)

func HandleAdd(ctx sdk.Context, k keeper.Keeper, msg types.MsgAdd) sdk.Result {
	if !k.HasProvider(ctx, msg.From) {
		return types.ErrorProviderDoesNotExist().Result()
	}

	count := k.GetCount(ctx)
	plan := types.Plan{
		ID:       count + 1,
		Provider: msg.From,
		Price:    msg.Price,
		Validity: msg.Validity,
		Bytes:    msg.Bytes,
		Status:   hub.StatusInactive,
		StatusAt: ctx.BlockTime(),
	}

	k.SetPlan(ctx, plan)
	k.SetInactivePlan(ctx, plan.ID)
	k.SetInactivePlanForProvider(ctx, plan.Provider, plan.ID)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSet,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", plan.ID)),
		sdk.NewAttribute(types.AttributeKeyAddress, plan.Provider.String()),
	))

	k.SetCount(ctx, count+1)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetCount,
		sdk.NewAttribute(types.AttributeKeyCount, fmt.Sprintf("%d", count+1)),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleSetStatus(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetStatus) sdk.Result {
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return types.ErrorPlanDoesNotExist().Result()
	}
	if !msg.From.Equals(plan.Provider) {
		return types.ErrorUnauthorized().Result()
	}

	if plan.Status.Equal(hub.StatusActive) {
		if msg.Status.Equal(hub.StatusInactive) {
			k.DeleteActivePlan(ctx, plan.ID)
			k.DeleteActivePlanForProvider(ctx, plan.Provider, plan.ID)

			k.SetInactivePlan(ctx, plan.ID)
			k.SetInactivePlanForProvider(ctx, plan.Provider, plan.ID)
		}
	} else {
		if msg.Status.Equal(hub.StatusActive) {
			k.DeleteInactivePlan(ctx, plan.ID)
			k.DeleteInactivePlanForProvider(ctx, plan.Provider, plan.ID)

			k.SetActivePlan(ctx, plan.ID)
			k.SetActivePlanForProvider(ctx, plan.Provider, plan.ID)
		}
	}

	plan.Status = msg.Status
	plan.StatusAt = ctx.BlockTime()

	k.SetPlan(ctx, plan)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetStatus,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", plan.ID)),
		sdk.NewAttribute(types.AttributeKeyStatus, plan.Status.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleAddNode(ctx sdk.Context, k keeper.Keeper, msg types.MsgAddNode) sdk.Result {
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

	k.SetNodeForPlan(ctx, plan.ID, node.Address)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAddNode,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", plan.ID)),
		sdk.NewAttribute(types.AttributeKeyAddress, node.Address.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleRemoveNode(ctx sdk.Context, k keeper.Keeper, msg types.MsgRemoveNode) sdk.Result {
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return types.ErrorPlanDoesNotExist().Result()
	}
	if !msg.From.Equals(plan.Provider) {
		return types.ErrorUnauthorized().Result()
	}

	k.DeleteNodeForPlan(ctx, plan.ID, msg.Address)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeRemoveNode,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", plan.ID)),
		sdk.NewAttribute(types.AttributeKeyAddress, msg.Address.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return sdk.Result{Events: ctx.EventManager().Events()}
}
