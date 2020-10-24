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
	k.SetInActiveNode(ctx, node.Address)

	if node.Provider != nil {
		k.SetInActiveNodeForProvider(ctx, node.Provider, node.Address)
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

	if node.Provider != nil && (msg.Provider != nil || msg.Price != nil) {
		if node.Status.Equal(hub.StatusActive) {
			k.DeleteActiveNodeForProvider(ctx, node.Provider, node.Address)
		} else {
			k.DeleteInActiveNodeForProvider(ctx, node.Provider, node.Address)
		}

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

		if node.Status.Equal(hub.StatusActive) {
			k.SetActiveNodeForProvider(ctx, node.Provider, node.Address)
		} else {
			k.SetInActiveNodeForProvider(ctx, node.Provider, node.Address)
		}
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
		if msg.Status.Equal(hub.StatusInactive) {
			k.DeleteActiveNode(ctx, node.Address)
			k.SetInActiveNode(ctx, node.Address)

			if node.Provider != nil {
				k.DeleteActiveNodeForProvider(ctx, node.Provider, node.Address)
				k.SetInActiveNodeForProvider(ctx, node.Provider, node.Address)
			}
		}

		k.DeleteInActiveNodeAt(ctx, node.StatusAt, node.Address)
	} else {
		if msg.Status.Equal(hub.StatusActive) {
			k.DeleteInActiveNode(ctx, node.Address)
			k.SetActiveNode(ctx, node.Address)

			if node.Provider != nil {
				k.DeleteInActiveNodeForProvider(ctx, node.Provider, node.Address)
				k.SetActiveNodeForProvider(ctx, node.Provider, node.Address)
			}
		}
	}

	node.Status = msg.Status
	node.StatusAt = ctx.BlockTime()

	if node.Status.Equal(hub.StatusActive) {
		k.SetInActiveNodeAt(ctx, node.StatusAt, node.Address)
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
