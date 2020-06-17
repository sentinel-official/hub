package provider

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/provider/keeper"
	"github.com/sentinel-official/hub/x/dvpn/provider/types"
)

func HandleRegisterProvider(ctx sdk.Context, k keeper.Keeper, msg types.MsgRegisterProvider) sdk.Result {
	_, found := k.GetProviderIDForAddress(ctx, msg.From)
	if found {
		return types.ErrorDuplicateProviderAddress().Result()
	}

	i := k.GetProvidersCount(ctx)
	provider := types.Provider{
		ID:          hub.NewProviderID(i),
		Address:     msg.From,
		Name:        msg.Name,
		Website:     msg.Website,
		Description: msg.Description,
	}

	k.SetProvidersCount(ctx, i+1)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetProvidersCount,
		sdk.NewAttribute(types.AttributeKeyProvidersCount, fmt.Sprintf("%d", i+1)),
	))

	k.SetProvider(ctx, provider)
	k.SetProviderIDForAddress(ctx, provider.Address, provider.ID)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetProvider,
		sdk.NewAttribute(types.AttributeKeyProviderID, provider.ID.String()),
		sdk.NewAttribute(types.AttributeKeyProviderAddress, provider.Address.String()),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func HandleUpdateProvider(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpdateProvider) sdk.Result {
	provider, found := k.GetProvider(ctx, msg.ID)
	if !found {
		return types.ErrorNoProviderFound().Result()
	}
	if !msg.From.Equals(provider.Address) {
		return types.ErrorUnauthorised().Result()
	}

	if len(msg.Name) > 0 {
		provider.Name = msg.Name
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
		sdk.NewAttribute(types.AttributeKeyProviderID, provider.ID.String()),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}
