package vpn

import (
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	"testing"
)

func TestIBCHandler_UpdateNodeStatus(t *testing.T) {

	cdc := codec.New()
	codec.RegisterCrypto(cdc)
	sdkTypes.RegisterCodec(cdc)
	hub.RegisterCodec(cdc)

	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	ibcPacket := TestGetIBCPacket()

	res := handleUpdateNodeStatus(ctx, keeper, *ibcPacket)
	require.False(t, res.IsOK())

	vpnMsg := TestGetVPNDetails()
	keeper.SetVPNDetails(ctx, vpnID, vpnMsg)

	res1 := handleUpdateNodeStatus(ctx, keeper, *ibcPacket)
	require.True(t, res1.IsOK())

	vpnDetails, err := keeper.GetVPNDetails(ctx, vpnID)

	require.Nil(t, err)
	require.Equal(t, vpnDetails.Status, "ACTIVE")

}

func TestIBCHandler_UpdateSessionStatus(t *testing.T) {

	cdc := codec.New()
	codec.RegisterCrypto(cdc)
	sdkTypes.RegisterCodec(cdc)
	hub.RegisterCodec(cdc)

	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	ibcPacket := TestGetSessionIBCPacket()

	res := handleUpdateSessionStatus(ctx, keeper, *ibcPacket)
	require.False(t, res.IsOK())

	sessionMsg := TestGetSessionDetails()
	err := keeper.SetSessionDetails(ctx, sessionID1, sessionMsg)
	require.Nil(t, err)

	res1 := handleUpdateSessionStatus(ctx, keeper, *ibcPacket)
	require.True(t, res1.IsOK())

	sessionDetails, err := keeper.GetSessionDetails(ctx, sessionID1)
	require.Nil(t, err)
	require.Equal(t, sessionDetails.Status, "ACTIVE")

}