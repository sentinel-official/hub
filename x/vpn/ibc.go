package vpn

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/ibc"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	vpnTypes "github.com/ironman0x7b2/sentinel-hub/types"
)

func PostIBCPacket(ctx sdkTypes.Context, k Keeper, ibcm ibc.Mapper, packet vpnTypes.VpnIBCPacket) sdkTypes.Error {
	// write everything into the state
	cdc := codec.New()
	store := ctx.KVStore(k.IbcStoreKey)
	index := getEgressLength(cdc, store, packet.DestChain)
	bz, err := cdc.MarshalBinary(packet)
	if err != nil {
		panic(err)
	}

	store.Set(ibc.EgressKey(packet.DestChain, index), bz)
	bz, err = cdc.MarshalBinary(index + 1)
	if err != nil {
		panic(err)
	}
	store.Set(ibc.EgressLengthKey(packet.DestChain), bz)

	return nil
}

func getEgressLength(cdc *codec.Codec, store sdkTypes.KVStore, destChain string) int64 {

	bz := store.Get(ibc.EgressLengthKey(destChain))
	if bz == nil {
		zero := marshalBinaryPanic(cdc, int64(0))
		store.Set(ibc.EgressLengthKey(destChain), zero)
		return 0
	}
	var res int64
	unmarshalBinaryPanic(cdc, bz, &res)
	return res
}

func marshalBinaryPanic(cdc *codec.Codec, value interface{}) []byte {
	res, err := cdc.MarshalBinary(value)
	if err != nil {
		panic(err)
	}
	return res
}

func unmarshalBinaryPanic(cdc *codec.Codec, bz []byte, ptr interface{}) {
	err := cdc.UnmarshalBinary(bz, ptr)
	if err != nil {
		panic(err)
	}
}


