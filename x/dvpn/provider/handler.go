package provider

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/provider/keeper"
	"github.com/sentinel-official/hub/x/dvpn/provider/types"
)

func HandleRegisterProvider(ctx sdk.Context, k keeper.Keeper, msg types.MsgRegisterProvider) sdk.Result {
	i := k.GetProvidersCount(ctx) + 1
	provider := types.Provider{
		ID:          hub.NewProviderID(i),
		Name:        msg.Name,
		Website:     msg.Website,
		Description: msg.Description,
	}

	k.SetProvidersCount(ctx, i)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetProvidersCount,
		sdk.NewAttribute(types.AttributeKeyProvidersCount, fmt.Sprintf("%d", i)),
	))

	k.SetProvider(ctx, provider)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSetProvider,
		sdk.NewAttribute(types.AttributeKeyProviderID, provider.ID.String()),
	))

	return sdk.Result{Events: ctx.EventManager().Events()}
}
