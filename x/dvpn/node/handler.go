package node

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

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
		Address:          msg.From.Bytes(),
		Provider:         msg.Provider,
		PricePerGB:       msg.PricePerGB,
		RemoteURL:        msg.RemoteURL,
		Version:          msg.Version,
		BandwidthSpeed:   msg.BandwidthSpeed,
		Category:         msg.Category,
		Status:           types.StatusInactive,
		StatusModifiedAt: ctx.BlockHeight(),
	}

	k.SetNode(ctx, node)
	k.SetNodeAddressForProvider(ctx, node.Provider, node.Address)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetNode,
		sdk.NewAttribute(types.AttributeKeyProvider, node.Provider.String()),
		sdk.NewAttribute(types.AttributeKeyNodeAddress, node.Address.String()),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}
