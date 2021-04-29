package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/provider/types"
)

var (
	_ types.MsgServiceServer = (*msgServer)(nil)
)

type msgServer struct {
	Keeper
}

func NewMsgServiceServer(keeper Keeper) types.MsgServiceServer {
	return &msgServer{Keeper: keeper}
}

func (k *msgServer) MsgRegister(c context.Context, msg *types.MsgRegisterRequest) (*types.MsgRegisterResponse, error) {
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
		if err = k.FundCommunityPool(ctx, msgFrom, deposit); err != nil {
			return nil, err
		}
	}

	var (
		provAddress = hubtypes.ProvAddress(msgFrom.Bytes())
		provider    = types.Provider{
			Address:     provAddress.String(),
			Name:        msg.Name,
			Identity:    msg.Identity,
			Website:     msg.Website,
			Description: msg.Description,
		}
	)

	k.SetProvider(ctx, provider)
	ctx.EventManager().EmitTypedEvent(
		&types.EventRegisterProvider{
			From:        sdk.AccAddress(msgFrom.Bytes()).String(),
			Address:     provider.Address,
			Name:        provider.Name,
			Identity:    provider.Identity,
			Website:     provider.Website,
			Description: provider.Description,
		},
	)

	ctx.EventManager().EmitTypedEvent(&types.EventModuleName)
	return &types.MsgRegisterResponse{}, nil
}

func (k *msgServer) MsgUpdate(c context.Context, msg *types.MsgUpdateRequest) (*types.MsgUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	msgFrom, err := hubtypes.ProvAddressFromBech32(msg.From)
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
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateProvider{
			From:        sdk.AccAddress(msgFrom.Bytes()).String(),
			Address:     provider.Address,
			Name:        msg.Name,
			Identity:    msg.Identity,
			Website:     msg.Website,
			Description: msg.Description,
		},
	)

	ctx.EventManager().EmitTypedEvent(&types.EventModuleName)
	return &types.MsgUpdateResponse{}, nil
}
