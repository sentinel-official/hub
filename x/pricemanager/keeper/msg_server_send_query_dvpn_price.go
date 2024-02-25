package keeper

import (
	"context"
	"time"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
	"github.com/sentinel-official/hub/v12/x/pricemanager/types"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// poolmanagertypes "github.com/osmosis-labs/osmosis/v23/x/poolmanager/types"
)

func (k msgServer) SendQueryDVPNPrice(goCtx context.Context, msg *types.MsgSendQueryDVPNPrice) (*types.MsgSendQueryDVPNPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chanCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(k.GetPort(ctx), msg.ChannelId))
	if !found {
		return nil, sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	// q := poolmanagertypes.QuerySpotPriceRequest{
	// 	PoolId:          msg.PoolId,
	// 	BaseAssetDenom:  msg.BaseAssetDenom,
	// 	QuoteAssetDenom: msg.QuoteAssetDenom,
	// 	Pagination:      msg.Pagination,
	// }
	reqs := []abcitypes.RequestQuery{
		{
			Path: "/osmosis.poolmanager.v1beta1.Query/SpotPrice",
			// Data: k.cdc.MustMarshal(&q),
		},
	}

	// timeoutTimestamp set to max value with the unsigned bit shifted to sastisfy hermes timestamp conversion
	// it is the responsibility of the auth module developer to ensure an appropriate timeout timestamp
	timeoutTimestamp := ctx.BlockTime().Add(time.Minute).UnixNano()
	seq, err := k.SendQuery(ctx, types.PortID, msg.ChannelId, chanCap, reqs, clienttypes.ZeroHeight(), uint64(timeoutTimestamp))
	if err != nil {
		return nil, err
	}

	k.SetQueryRequest(ctx, seq, q)

	return &types.MsgSendQueryDVPNPriceResponse{
		Sequence: seq,
	}, nil
}
