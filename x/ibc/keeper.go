package ibc

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

func EgressKey(destChainID string, length int64) string {
	return fmt.Sprintf("egress/%s/%d", destChainID, length)
}

func EgressLengthKey(destChainID string) string {
	return fmt.Sprintf("egress/%s", destChainID)
}

func IngressKey(srcChainID string, length int64) string {
	return fmt.Sprintf("ingress/%s/%d", srcChainID, length)
}

func IngressLengthKey(srcChainID string) string {
	return fmt.Sprintf("ingress/%s", srcChainID)
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

func (k Keeper) SetIBCPacket(ctx csdkTypes.Context, packetID string, packet sdkTypes.IBCPacket) csdkTypes.Error {
	store := ctx.KVStore(k.IBCKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(packetID)

	if err != nil {
		return errorMarshal()
	}

	valueBytes, err := k.cdc.MarshalBinaryLengthPrefixed(packet)

	if err != nil {
		return errorMarshal()
	}

	store.Set(keyBytes, valueBytes)

	return nil
}

func (k Keeper) GetIBCPacket(ctx csdkTypes.Context, packetID string) (*sdkTypes.IBCPacket, csdkTypes.Error) {
	store := ctx.KVStore(k.IBCKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(packetID)

	if err != nil {
		return nil, errorMarshal()
	}

	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return nil, nil
	}

	var packet sdkTypes.IBCPacket

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(valueBytes, &packet); err != nil {
		return nil, errorUnmarshal()
	}

	return &packet, nil
}

func (k Keeper) SetEgressLength(ctx csdkTypes.Context, egressKey string, length int64) csdkTypes.Error {
	store := ctx.KVStore(k.IBCKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(egressKey)

	if err != nil {
		return errorMarshal()
	}

	valueBytes, err := k.cdc.MarshalBinaryLengthPrefixed(length)

	if err != nil {
		return errorMarshal()
	}

	store.Set(keyBytes, valueBytes)

	return nil
}

func (k Keeper) GetEgressLength(ctx csdkTypes.Context, egressKey string) (int64, csdkTypes.Error) {
	store := ctx.KVStore(k.IBCKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(egressKey)

	if err != nil {
		return 0, errorMarshal()
	}

	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return 0, nil
	}

	var length int64

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(valueBytes, &length); err != nil {
		return 0, errorUnmarshal()
	}

	return length, nil
}

func (k Keeper) SetIngressLength(ctx csdkTypes.Context, ingressKey string, length int64) csdkTypes.Error {
	store := ctx.KVStore(k.IBCKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(ingressKey)

	if err != nil {
		return errorMarshal()
	}

	valueBytes, err := k.cdc.MarshalBinaryLengthPrefixed(length)

	if err != nil {
		return errorMarshal()
	}

	store.Set(keyBytes, valueBytes)

	return nil
}

func (k Keeper) GetIngressLength(ctx csdkTypes.Context, ingressKey string) (int64, csdkTypes.Error) {
	store := ctx.KVStore(k.IBCKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(ingressKey)

	if err != nil {
		return 0, errorMarshal()
	}

	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return 0, nil
	}

	var length int64

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(valueBytes, &length); err != nil {
		return 0, errorUnmarshal()
	}

	return length, nil
}

func (k Keeper) PostIBCPacket(ctx csdkTypes.Context, packet sdkTypes.IBCPacket) csdkTypes.Error {
	egressLength, err := k.GetEgressLength(ctx, EgressLengthKey(packet.DestChainID))

	if err != nil {
		return err
	}

	if err := k.SetIBCPacket(ctx, EgressKey(packet.DestChainID, egressLength), packet); err != nil {
		return err
	}

	if err := k.SetEgressLength(ctx, EgressLengthKey(packet.DestChainID), egressLength+1); err != nil {
		return err
	}

	return nil
}
