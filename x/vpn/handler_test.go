package vpn

import (
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
	"github.com/ironman0x7b2/sentinel-sdk/x/ibc"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	"testing"
)

func TestHandler_RegisterNode(t *testing.T) {
	cdc := codec.New()
	codec.RegisterCrypto(cdc)
	sdkTypes.RegisterCodec(cdc)
	ibc.RegisterCodec(cdc)
	hub.RegisterCodec(cdc)

	ms, ibcStoreKey, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)
	ibcKeeper := ibc.NewKeeper(ibcStoreKey, cdc)

	require.NotNil(t, ibcKeeper)
	msg := TestGetMsgRegisterNode()

	res := handleRegisterNode(ctx, keeper, ibcKeeper, *msg)
	require.True(t, res.IsOK())
}

func TestHandler_SessionPayment(t *testing.T) {
	cdc := codec.New()
	codec.RegisterCrypto(cdc)
	sdkTypes.RegisterCodec(cdc)
	ibc.RegisterCodec(cdc)
	hub.RegisterCodec(cdc)

	ms, ibcStoreKey, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)
	ibcKeeper := ibc.NewKeeper(ibcStoreKey, cdc)

	require.NotNil(t, ibcKeeper)
	msg := TestGetMsgPayVpnService()

	res := handleSessionPayment(ctx, keeper, ibcKeeper, *msg)
	require.False(t, res.IsOK())

	vpnMsg := TestGetVPNDetails()
	keeper.SetVPNDetails(ctx, vpnID, vpnMsg)

	res1 := handleSessionPayment(ctx, keeper, ibcKeeper, *msg)
	require.True(t, res1.IsOK())
}

func TestHandler_DeregisterNode(t *testing.T) {
	cdc := codec.New()
	codec.RegisterCrypto(cdc)
	sdkTypes.RegisterCodec(cdc)
	ibc.RegisterCodec(cdc)
	hub.RegisterCodec(cdc)

	ms, ibcStoreKey, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)
	ibcKeeper := ibc.NewKeeper(ibcStoreKey, cdc)

	require.NotNil(t, ibcKeeper)
	msg := TestGetMsgDeregisterNode()

	res := handleDeregisterNode(ctx, keeper, ibcKeeper, *msg)
	require.False(t, res.IsOK())

	vpnMsg := TestGetVPNDetails()
	keeper.SetVPNDetails(ctx, vpnID, vpnMsg)

	res1 := handleDeregisterNode(ctx, keeper, ibcKeeper, *msg)
	require.True(t, res1.IsOK())
}
