package node

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

func BeginBlock(ctx sdk.Context, k keeper.Keeper) {
	end := ctx.BlockTime().Add(-1 * k.InactiveDuration(ctx))
	k.IterateActiveNodes(ctx, end, func(_ int, node types.Node) (stop bool) {
		k.DeleteActiveNodeAt(ctx, node.StatusAt, node.Address)

		node.Status = hub.StatusInactive
		node.StatusAt = ctx.BlockTime()
		k.SetNode(ctx, node)

		return false
	})
}

func HandleRegisterNode(ctx sdk.Context, k keeper.Keeper, msg types.MsgRegisterNode) sdk.Result {
	if k.HasNode(ctx, msg.From.Bytes()) {
		return types.ErrorDuplicateNode().Result()
	}
	if !k.HasProvider(ctx, msg.Provider) {
		return types.ErrorProviderDoesNotExist().Result()
	}

	node := types.Node{
		Address:       msg.From.Bytes(),
		Provider:      msg.Provider,
		Price:         msg.Price,
		InternetSpeed: msg.InternetSpeed,
		RemoteURL:     msg.RemoteURL,
		Version:       msg.Version,
		Category:      msg.Category,
		Status:        hub.StatusInactive,
		StatusAt:      ctx.BlockTime(),
	}

	k.SetNode(ctx, node)
	if node.Provider != nil {
		k.SetNodeForProvider(ctx, node.Provider, node.Address)
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetNode,
		sdk.NewAttribute(types.AttributeKeyProvider, node.Provider.String()),
		sdk.NewAttribute(types.AttributeKeyAddress, node.Address.String()),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleUpdateNode(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpdateNode) sdk.Result {
	node, found := k.GetNode(ctx, msg.From)
	if !found {
		return types.ErrorNodeDoesNotExist().Result()
	}
	if node.Provider.Equals(msg.Provider) {
		return types.ErrorCanNotUpdate().Result()
	}

	if msg.Provider != nil || msg.Price != nil {
		k.DeleteNodeForProvider(ctx, node.Provider, node.Address)

		plans := k.GetPlansForProvider(ctx, node.Provider)
		for _, plan := range plans {
			k.DeleteNodeForPlan(ctx, plan.ID, node.Address)
		}
	}

	if msg.Provider != nil {
		if !k.HasProvider(ctx, msg.Provider) {
			return types.ErrorProviderDoesNotExist().Result()
		}

		node.Provider = msg.Provider
		node.Price = nil

		k.SetNodeForProvider(ctx, node.Provider, node.Address)
	}
	if msg.Price != nil {
		node.Provider = nil
		node.Price = msg.Price
	}
	if !msg.InternetSpeed.IsAnyZero() {
		node.InternetSpeed = msg.InternetSpeed
	}
	if len(msg.RemoteURL) > 0 {
		node.RemoteURL = msg.RemoteURL
	}
	if len(msg.Version) > 0 {
		node.Version = msg.Version
	}
	if msg.Category.IsValid() {
		node.Category = msg.Category
	}

	k.SetNode(ctx, node)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUpdateNode,
		sdk.NewAttribute(types.AttributeKeyAddress, node.Address.String()),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleSetNodeStatus(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetNodeStatus) sdk.Result {
	node, found := k.GetNode(ctx, msg.From)
	if !found {
		return types.ErrorNodeDoesNotExist().Result()
	}

	k.DeleteActiveNodeAt(ctx, node.StatusAt, node.Address)
	k.SetActiveNodeAt(ctx, ctx.BlockTime(), node.Address)

	node.Status = msg.Status
	node.StatusAt = ctx.BlockTime()

	k.SetNode(ctx, node)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetNodeStatus,
		sdk.NewAttribute(types.AttributeKeyAddress, node.Address.String()),
		sdk.NewAttribute(types.AttributeKeyStatus, node.Status.String()),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}
