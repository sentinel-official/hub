package node

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/node/keeper"
	"github.com/sentinel-official/hub/x/dvpn/node/types"
)

func HandleRegisterNode(ctx sdk.Context, k keeper.Keeper, msg types.MsgRegisterNode) sdk.Result {
	_, found := k.GetProvider(ctx, msg.Provider)
	if !found {
		return types.ErrorNoProviderFound().Result()
	}

	_, found = k.GetNode(ctx, msg.From.Bytes())
	if found {
		return types.ErrorDuplicateNode().Result()
	}

	node := types.Node{
		Address:       msg.From.Bytes(),
		Provider:      msg.Provider,
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
		return ErrorNoNodeFound().Result()
	}

	if msg.Provider != nil {
		_, found := k.GetProvider(ctx, msg.Provider)
		if !found {
			return ErrorNoProviderFound().Result()
		}

		node.Provider = msg.Provider
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
		return ErrorNoNodeFound().Result()
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
