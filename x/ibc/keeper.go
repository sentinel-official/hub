package ibc

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

func EgressKey(destChain string, length int64) []byte {
	return []byte(fmt.Sprintf("egress/%s/%d", destChain, length))
}

func EgressLengthKey(destChain string) []byte {
	return []byte(fmt.Sprintf("egress/%s", destChain))
}

func IngressKey(srcChain string, length int64) []byte {
	return []byte(fmt.Sprintf("ingress/%s/%d", srcChain, length))
}

func IngressLengthKey(srcChain string) []byte {
	return []byte(fmt.Sprintf("ingress/%s", srcChain))
}

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

func (k Keeper) SetIBCPacket(ctx csdkTypes.Context, packetID string, packet sdkTypes.IBCPacket) {
	store := ctx.KVStore(k.IBCKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(packetID)

	if err != nil {
		panic(err)
	}

	valueBytes, err := k.cdc.MarshalBinaryLengthPrefixed(packet)

	if err != nil {
		panic(err)
	}

	store.Set(keyBytes, valueBytes)
}

func (k Keeper) GetIBCPacket(ctx csdkTypes.Context, packetID string) *sdkTypes.IBCPacket {
	store := ctx.KVStore(k.IBCKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(packetID)

	if err != nil {
		panic(err)
	}

	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return nil
	}

	var packet sdkTypes.IBCPacket

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(valueBytes, &packet); err != nil {
		panic(err)
	}

	return &packet
}

func (k Keeper) SetEgressLength(ctx csdkTypes.Context, egressKey string, length int64) {
	store := ctx.KVStore(k.IBCKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(egressKey)

	if err != nil {
		panic(err)
	}

	valueBytes, err := k.cdc.MarshalBinaryLengthPrefixed(length)

	if err != nil {
		panic(err)
	}

	store.Set(keyBytes, valueBytes)
}

func (k Keeper) GetEgressLength(ctx csdkTypes.Context, egressKey string) int64 {
	store := ctx.KVStore(k.IBCKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(egressKey)

	if err != nil {
		panic(err)
	}

	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return 0
	}

	var length int64

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(valueBytes, &length); err != nil {
		panic(err)
	}

	return length
}

func (k Keeper) SetIngressLength(ctx csdkTypes.Context, ingressKey string, length int64) {
	store := ctx.KVStore(k.IBCKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(ingressKey)

	if err != nil {
		panic(err)
	}

	valueBytes, err := k.cdc.MarshalBinaryLengthPrefixed(length)

	if err != nil {
		panic(err)
	}

	store.Set(keyBytes, valueBytes)
}

func (k Keeper) GetIngressLength(ctx csdkTypes.Context, ingressKey string) int64 {
	store := ctx.KVStore(k.IBCKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(ingressKey)

	if err != nil {
		panic(err)
	}

	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return 0
	}

	var length int64

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(valueBytes, &length); err != nil {
		panic(err)
	}

	return length
}

func (k Keeper) PostIBCPacket(ctx csdkTypes.Context, packet sdkTypes.IBCPacket) csdkTypes.Error {
	egressLength := k.GetEgressLength(ctx, string(EgressLengthKey(packet.DestChainID)))
	k.SetIBCPacket(ctx, string(EgressKey(packet.DestChainID, egressLength)), packet)
	k.SetEgressLength(ctx, string(EgressLengthKey(packet.DestChainID)), egressLength+1)

	return nil
}
