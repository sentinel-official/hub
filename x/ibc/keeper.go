package ibc

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	hubTypes "github.com/ironman0x7b2/sentinel-hub/types"
)

type Keeper struct {
	IBCKey sdkTypes.StoreKey
	cdc    *codec.Codec
}

func (ibc Keeper) PostIBCPacket(ctx sdkTypes.Context, packet hubTypes.IBCPacket) sdkTypes.Error {
	store := ctx.KVStore(ibc.IBCKey)
	index := ibc.getEgressLength(store, packet.DestChain)
	bz, err := ibc.cdc.MarshalBinary(packet)

	if err != nil {
		panic(err)
	}

	store.Set(EgressKey(packet.DestChain, index), bz)
	bz, err = ibc.cdc.MarshalBinary(index + 1)

	if err != nil {
		panic(err)
	}

	store.Set(EgressLengthKey(packet.DestChain), bz)

	return nil
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

func (ibc Keeper) getEgressLength(store sdkTypes.KVStore, destChain string) int64 {
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
