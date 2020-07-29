package provider

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/vpn/provider/keeper"
	"github.com/sentinel-official/hub/x/vpn/provider/types"
)

func HandleRegisterProvider(ctx sdk.Context, k keeper.Keeper, msg types.MsgRegisterProvider) sdk.Result {
	_, found := k.GetProvider(ctx, msg.From.Bytes())
	if found {
		return types.ErrorDuplicateProvider().Result()
	}

	provider := types.Provider{
		Address:     msg.From.Bytes(),
		Name:        msg.Name,
		Identity:    msg.Identity,
		Website:     msg.Website,
		Description: msg.Description,
	}

	k.SetProvider(ctx, provider)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetProvider,
		sdk.NewAttribute(types.AttributeKeyAddress, provider.Address.String()),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleUpdateProvider(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpdateProvider) sdk.Result {
	provider, found := k.GetProvider(ctx, msg.From)
	if !found {
		return types.ErrorProviderDoesNotExist().Result()
	}

	if len(msg.Name) > 0 {
		provider.Name = msg.Name
	}
	if len(msg.Identity) > 0 {
		provider.Identity = msg.Identity
	}
	if len(msg.Website) > 0 {
		provider.Website = msg.Website
	}
	if len(msg.Description) > 0 {
		provider.Description = msg.Description
	}

	k.SetProvider(ctx, provider)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUpdateProvider,
		sdk.NewAttribute(types.AttributeKeyAddress, provider.Address.String()),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}
