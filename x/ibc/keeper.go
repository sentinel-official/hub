package ibc

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type Keeper struct {
	IBCKey csdkTypes.StoreKey
	cdc    *codec.Codec
}

func NewKeeper(ibcKey csdkTypes.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		IBCKey: ibcKey,
		cdc:    cdc,
	}
}

func (ibc Keeper) PostIBCPacket(ctx csdkTypes.Context, packet sdkTypes.IBCPacket) csdkTypes.Error {
	store := ctx.KVStore(ibc.IBCKey)
	index := ibc.getEgressLength(store, packet.DestChainID)
	bz, err := ibc.cdc.MarshalBinaryLengthPrefixed(packet)

	if err != nil {
		panic(err)
	}

	store.Set(EgressKey(packet.DestChainID, index), bz)
	bz, err = ibc.cdc.MarshalBinaryLengthPrefixed(index + 1)

	if err != nil {
		panic(err)
	}

	store.Set(EgressLengthKey(packet.DestChainID), bz)

	return nil
}

func marshalBinaryPanic(cdc *codec.Codec, value interface{}) []byte {
	res, err := cdc.MarshalBinaryLengthPrefixed(value)

	if err != nil {
		panic(err)
	}

	return res
}

func unmarshalBinaryPanic(cdc *codec.Codec, bz []byte, ptr interface{}) {
	err := cdc.UnmarshalBinaryLengthPrefixed(bz, ptr)

	if err != nil {
		panic(err)
	}
}

func (ibc Keeper) getEgressLength(store csdkTypes.KVStore, destChain string) int64 {
	bz := store.Get(EgressLengthKey(destChain))

	if bz == nil {
		zero := marshalBinaryPanic(ibc.cdc, int64(0))
		store.Set(EgressLengthKey(destChain), zero)

		return 0
	}

	var res int64
	unmarshalBinaryPanic(ibc.cdc, bz, &res)

	return res
}

func EgressKey(destChain string, index int64) []byte {
	return []byte(fmt.Sprintf("egress/%s/%d", destChain, index))
}

func EgressLengthKey(destChain string) []byte {
	return []byte(fmt.Sprintf("egress/%s", destChain))
}

func IngressSequenceKey(srcChain string) []byte {
	return []byte(fmt.Sprintf("ingress/%s", srcChain))
}
