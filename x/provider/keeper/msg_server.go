package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/provider/types"
)

var (
	_ types.MsgServiceServer = server{}
)

type server struct {
	Keeper
}

func NewMsgServiceServer(keeper Keeper) types.MsgServiceServer {
	return &server{Keeper: keeper}
}

func (k server) MsgRegister(c context.Context, msg *types.MsgRegisterRequest) (*types.MsgRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	msgFrom, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	_, found := k.GetProvider(ctx, msgFrom.Bytes())
	if found {
		return nil, types.ErrorDuplicateProvider
	}

	deposit := k.Deposit(ctx)
	if deposit.IsPositive() {
		if err := k.FundCommunityPool(ctx, msgFrom, deposit); err != nil {
			return nil, err
		}
	}

	var (
		provAddress = hub.ProvAddress(msgFrom.Bytes())
		provider    = types.Provider{
			Address:     provAddress.String(),
			Name:        msg.Name,
			Identity:    msg.Identity,
			Website:     msg.Website,
			Description: msg.Description,
		}
	)

	k.SetProvider(ctx, provider)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSet,
		sdk.NewAttribute(types.AttributeKeyAddress, provider.Address),
		sdk.NewAttribute(types.AttributeKeyDeposit, deposit.String()),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &types.MsgRegisterResponse{}, nil
}

func (k server) MsgUpdate(c context.Context, msg *types.MsgUpdateRequest) (*types.MsgUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	msgFrom, err := hub.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	provider, found := k.GetProvider(ctx, msgFrom)
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
		sdk.NewAttribute(types.AttributeKeyAddress, provider.Address),
	))

	ctx.EventManager().EmitEvent(types.EventModuleName)
	return &types.MsgUpdateResponse{}, nil
}
