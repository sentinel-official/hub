package ibc

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

func TestKeeper(t *testing.T) {
	cdc := codec.New()

	sdkTypes.RegisterCodec(cdc)

	multiStore, ibcKey := DefaultSetup()
	ctx := csdkTypes.NewContext(multiStore, abciTypes.Header{}, false, log.NewNopLogger())
	ibcKeeper := NewKeeper(ibcKey, cdc)

	var err csdkTypes.Error

	ibcPacket1 := sdkTypes.IBCPacket{
		"src_chain_id",
		"dest_chain_id",
		nil,
	}

	ingressLength0, err := ibcKeeper.GetIngressLength(ctx, sdkTypes.IngressLengthKey(ibcPacket1.SrcChainID))
	require.Nil(t, err)
	require.Equal(t, uint64(0), ingressLength0)

	egressLength0, err := ibcKeeper.GetEgressLength(ctx, sdkTypes.EgressLengthKey(ibcPacket1.DestChainID))
	require.Nil(t, err)
	require.Equal(t, uint64(0), egressLength0)

	getIBCPacket0, err := ibcKeeper.GetIBCPacket(ctx, sdkTypes.EgressKey(ibcPacket1.DestChainID, uint64(0)))
	require.Nil(t, err)
	require.Nil(t, getIBCPacket0)

	err = ibcKeeper.SetIngressLength(ctx, sdkTypes.IngressLengthKey(ibcPacket1.SrcChainID), uint64(1))
	require.Nil(t, err)

	ingressLength1, err := ibcKeeper.GetIngressLength(ctx, sdkTypes.IngressLengthKey(ibcPacket1.SrcChainID))
	require.Nil(t, err)
	require.Equal(t, uint64(1), ingressLength1)

	ibcPacketRes1 := ibcKeeper.PostIBCPacket(ctx, ibcPacket1)
	require.Nil(t, ibcPacketRes1)

	getIBCPacket1, err := ibcKeeper.GetIBCPacket(ctx, sdkTypes.EgressKey(ibcPacket1.DestChainID, uint64(0)))
	require.Nil(t, err)
	require.Equal(t, getIBCPacket1, &ibcPacket1)

	egressLength1, err := ibcKeeper.GetEgressLength(ctx, sdkTypes.EgressLengthKey(ibcPacket1.DestChainID))
	require.Nil(t, err)
	require.Equal(t, uint64(1), egressLength1)

	err = ibcKeeper.SetEgressLength(ctx, sdkTypes.EgressLengthKey(ibcPacket1.DestChainID), uint64(2))
	require.Nil(t, err)

	egressLength2, err := ibcKeeper.GetEgressLength(ctx, sdkTypes.EgressLengthKey(ibcPacket1.DestChainID))
	require.Nil(t, err)
	require.Equal(t, uint64(2), egressLength2)
}
