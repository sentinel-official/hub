package provider

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/provider/keeper"
	"github.com/sentinel-official/hub/x/provider/types"
)

func HandleRegister(ctx sdk.Context, k keeper.Keeper, msg types.MsgRegister) (*sdk.Result, error) {
	_, found := k.GetProvider(ctx, msg.From.Bytes())
	if found {
		return nil, types.ErrorDuplicateProvider
	}

	deposit := k.Deposit(ctx)
	if deposit.IsPositive() {
		if err := k.FundCommunityPool(ctx, msg.From, deposit); err != nil {
			return nil, err
		}
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
		types.EventTypeSet,
		sdk.NewAttribute(types.AttributeKeyAddress, provider.Address.String()),
		sdk.NewAttribute(types.AttributeKeyDeposit, deposit.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func HandleUpdate(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpdate) (*sdk.Result, error) {
	provider, found := k.GetProvider(ctx, msg.From)
	if !found {
		return nil, types.ErrorProviderDoesNotExist
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
		types.EventTypeUpdate,
		sdk.NewAttribute(types.AttributeKeyAddress, provider.Address.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
