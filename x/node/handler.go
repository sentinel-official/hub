package node

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

func HandleRegister(ctx sdk.Context, k keeper.Keeper, msg types.MsgRegister) (*sdk.Result, error) {
	if k.HasNode(ctx, msg.From.Bytes()) {
		return nil, types.ErrorDuplicateNode
	}
	if !k.HasProvider(ctx, msg.Provider) {
		return nil, types.ErrorProviderDoesNotExist
	}

	node := types.Node{
		Address:   msg.From.Bytes(),
		Provider:  msg.Provider,
		Price:     msg.Price,
		RemoteURL: msg.RemoteURL,
		Status:    hub.StatusInactive,
		StatusAt:  ctx.BlockTime(),
	}

	k.SetNode(ctx, node)
	k.SetInactiveNode(ctx, node.Address)

	if node.Provider != nil {
		k.SetInactiveNodeForProvider(ctx, node.Provider, node.Address)
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSet,
		sdk.NewAttribute(types.AttributeKeyProvider, node.Provider.String()),
		sdk.NewAttribute(types.AttributeKeyAddress, node.Address.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func HandleUpdate(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpdate) (*sdk.Result, error) {
	node, found := k.GetNode(ctx, msg.From)
	if !found {
		return nil, types.ErrorNodeDoesNotExist
	}

	if node.Provider.Equals(msg.Provider) {
		msg.Provider = nil
	}

	if node.Provider != nil && (msg.Provider != nil || msg.Price != nil) {
		if node.Status.Equal(hub.StatusActive) {
			k.DeleteActiveNodeForProvider(ctx, node.Provider, node.Address)
		} else {
			k.DeleteInactiveNodeForProvider(ctx, node.Provider, node.Address)
		}

		plans := k.GetPlansForProvider(ctx, node.Provider)
		for _, plan := range plans {
			k.DeleteNodeForPlan(ctx, plan.ID, node.Address)
		}
	}

	if msg.Provider != nil {
		if !k.HasProvider(ctx, msg.Provider) {
			return nil, types.ErrorProviderDoesNotExist
		}

		node.Provider = msg.Provider
		node.Price = nil

		if node.Status.Equal(hub.StatusActive) {
			k.SetActiveNodeForProvider(ctx, node.Provider, node.Address)
		} else {
			k.SetInactiveNodeForProvider(ctx, node.Provider, node.Address)
		}
	}
	if msg.Price != nil {
		node.Provider = nil
		node.Price = msg.Price
	}
	if len(msg.RemoteURL) > 0 {
		node.RemoteURL = msg.RemoteURL
	}

	k.SetNode(ctx, node)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUpdate,
		sdk.NewAttribute(types.AttributeKeyAddress, node.Address.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func HandleSetStatus(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetStatus) (*sdk.Result, error) {
	node, found := k.GetNode(ctx, msg.From)
	if !found {
		return nil, types.ErrorNodeDoesNotExist
	}

	if node.Status.Equal(hub.StatusActive) {
		if msg.Status.Equal(hub.StatusInactive) {
			k.DeleteActiveNode(ctx, node.Address)
			k.SetInactiveNode(ctx, node.Address)

			if node.Provider != nil {
				k.DeleteActiveNodeForProvider(ctx, node.Provider, node.Address)
				k.SetInactiveNodeForProvider(ctx, node.Provider, node.Address)
			}
		}

		k.DeleteInactiveNodeAt(ctx, node.StatusAt, node.Address)
	} else {
		if msg.Status.Equal(hub.StatusActive) {
			k.DeleteInactiveNode(ctx, node.Address)
			k.SetActiveNode(ctx, node.Address)

			if node.Provider != nil {
				k.DeleteInactiveNodeForProvider(ctx, node.Provider, node.Address)
				k.SetActiveNodeForProvider(ctx, node.Provider, node.Address)
			}
		}
	}

	node.Status = msg.Status
	node.StatusAt = ctx.BlockTime()

	if node.Status.Equal(hub.StatusActive) {
		k.SetInactiveNodeAt(ctx, node.StatusAt, node.Address)
	}

	k.SetNode(ctx, node)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetStatus,
		sdk.NewAttribute(types.AttributeKeyAddress, node.Address.String()),
		sdk.NewAttribute(types.AttributeKeyStatus, node.Status.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
