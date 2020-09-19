package node

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

func HandleRegister(ctx sdk.Context, k keeper.Keeper, msg types.MsgRegister) sdk.Result {
	if k.HasNode(ctx, msg.From.Bytes()) {
		return types.ErrorDuplicateNode().Result()
	}
	if !k.HasProvider(ctx, msg.Provider) {
		return types.ErrorProviderDoesNotExist().Result()
	}

	node := types.Node{
		Moniker:       msg.Moniker,
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
		types.EventTypeSet,
		sdk.NewAttribute(types.AttributeKeyProvider, node.Provider.String()),
		sdk.NewAttribute(types.AttributeKeyAddress, node.Address.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleUpdate(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpdate) sdk.Result {
	node, found := k.GetNode(ctx, msg.From)
	if !found {
		return types.ErrorNodeDoesNotExist().Result()
	}

	if node.Provider.Equals(msg.Provider) {
		msg.Provider = nil
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
	if len(msg.Moniker) > 0 {
		node.Moniker = msg.Moniker
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
		types.EventTypeUpdate,
		sdk.NewAttribute(types.AttributeKeyAddress, node.Address.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleSetStatus(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetStatus) sdk.Result {
	node, found := k.GetNode(ctx, msg.From)
	if !found {
		return types.ErrorNodeDoesNotExist().Result()
	}

	if node.Status.Equal(hub.StatusActive) {
		k.DeleteActiveNodeAt(ctx, node.StatusAt, node.Address)
	}

	node.Status = msg.Status
	node.StatusAt = ctx.BlockTime()

	if node.Status.Equal(hub.StatusActive) {
		k.SetActiveNodeAt(ctx, node.StatusAt, node.Address)
	}

	k.SetNode(ctx, node)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetStatus,
		sdk.NewAttribute(types.AttributeKeyAddress, node.Address.String()),
		sdk.NewAttribute(types.AttributeKeyStatus, node.Status.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return sdk.Result{Events: ctx.EventManager().Events()}
}
