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

func NewMsgServiceServer(k Keeper) types.MsgServiceServer {
	return &msgServer{k}
}

func (k *msgServer) MsgRegister(c context.Context, msg *types.MsgRegisterRequest) (*types.MsgRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	_, found := k.GetProvider(ctx, fromAddr.Bytes())
	if found {
		return nil, types.ErrorDuplicateProvider
	}

	deposit := k.Deposit(ctx)
	if err := k.FundCommunityPool(ctx, fromAddr, deposit); err != nil {
		return nil, err
	}

	var (
		provAddr = hubtypes.ProvAddress(fromAddr.Bytes())
		provider = types.Provider{
			Address:     provAddr.String(),
			Name:        msg.Name,
			Identity:    msg.Identity,
			Website:     msg.Website,
			Description: msg.Description,
			Status:      hubtypes.StatusInactive,
		}
	)

	k.SetProvider(ctx, provider)
	ctx.EventManager().EmitTypedEvent(
		&types.EventRegister{
			Address: provider.Address,
		},
	)

	return &types.MsgRegisterResponse{}, nil
}

func (k *msgServer) MsgUpdate(c context.Context, msg *types.MsgUpdateRequest) (*types.MsgUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	fromAddr, err := hubtypes.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	provider, found := k.GetProvider(ctx, fromAddr)
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
	if !msg.Status.Equal(hubtypes.StatusUnspecified) {
		switch provider.Status {
		case hubtypes.StatusActive:
			k.deleteActiveProvider(ctx, fromAddr)
		case hubtypes.StatusInactive:
			k.deleteInactiveProvider(ctx, fromAddr)
		default:
			return nil, types.ErrorInvalidStatus
		}

		provider.Status = msg.Status
	}

	k.SetProvider(ctx, provider)
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdate{
			Address: provider.Address,
		},
	)

	return &types.MsgUpdateResponse{}, nil
}
