package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/types"
)

func (k Keeper) SetChannel(ctx sdk.Context, address sdk.AccAddress, subscription uint64, node hub.NodeAddress, channel uint64) {
	key := types.ChannelKey(address, subscription, node)
	value := k.cdc.MustMarshalBinaryBare(&protobuf.UInt64Value{Value: channel})

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetChannel(ctx sdk.Context, address sdk.AccAddress, subscription uint64, node hub.NodeAddress) uint64 {
	store := k.Store(ctx)

	key := types.ChannelKey(address, subscription, node)
	value := store.Get(key)
	if value == nil {
		return 0
	}

	var channel protobuf.UInt64Value
	k.cdc.MustUnmarshalBinaryBare(value, &channel)

	return channel.Value
}
