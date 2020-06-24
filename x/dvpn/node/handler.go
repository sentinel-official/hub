package node

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/node/keeper"
	"github.com/sentinel-official/hub/x/dvpn/node/types"
)

func HandleRegisterNode(ctx sdk.Context, k keeper.Keeper, msg types.MsgRegisterNode) sdk.Result {
	if msg.Provider != nil {
		_, found := k.GetProvider(ctx, msg.Provider)
		if !found {
			return types.ErrorNoProviderFound().Result()
		}
	}

	_, found := k.GetNode(ctx, msg.From.Bytes())
	if found {
		return types.ErrorDuplicateNode().Result()
	}

	node := types.Node{
		Address:       msg.From.Bytes(),
		Provider:      msg.Provider,
		PricePerGB:    msg.PricePerGB,
		InternetSpeed: msg.InternetSpeed,
		RemoteURL:     msg.RemoteURL,
		Version:       msg.Version,
		Category:      msg.Category,
		Status:        hub.StatusInactive,
		StatusAt:      ctx.BlockHeight(),
	}

	k.SetNode(ctx, node)
	k.SetNodeAddressForProvider(ctx, node.Provider, node.Address)
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
		return types.ErrorNoNodeFound().Result()
	}

	if msg.Provider != nil && !msg.Provider.Equals(node.Provider) {
		if node.Provider != nil {
			k.DeleteNodeAddressForProvider(ctx, node.Provider, node.Address)

			plans := k.GetPlansForProvider(ctx, node.Provider)
			for _, plan := range plans {
				k.DeleteNodeAddressForPlan(ctx, plan.ID, node.Address)
			}
		}

		if msg.Provider.Equals(hub.EmptyProviderAddress) {
			node.Provider = nil
		} else {
			_, found := k.GetProvider(ctx, msg.Provider)
			if !found {
				return types.ErrorNoProviderFound().Result()
			}

			node.Provider = msg.Provider
		}

		if node.Provider != nil {
			k.SetNodeAddressForProvider(ctx, node.Provider, node.Address)
		}
	}
	if msg.PricePerGB != nil {
		node.PricePerGB = msg.PricePerGB

		if hub.AreEmptyCoins(msg.PricePerGB) {
			node.PricePerGB = nil
		}
	}
	if node.Provider != nil {
		node.PricePerGB = nil
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
		sdk.NewAttribute(types.AttributeKeyProvider, node.Provider.String()),
		sdk.NewAttribute(types.AttributeKeyAddress, node.Address.String()),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleSetNodeStatus(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetNodeStatus) sdk.Result {
	node, found := k.GetNode(ctx, msg.From)
	if !found {
		return types.ErrorNoNodeFound().Result()
	}

	node.Status = msg.Status
	node.StatusAt = ctx.BlockHeight()

	k.SetNode(ctx, node)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetNodeStatus,
		sdk.NewAttribute(types.AttributeKeyStatus, node.Status.String()),
		sdk.NewAttribute(types.AttributeKeyStatusAt, fmt.Sprintf("%d", node.StatusAt)),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}
