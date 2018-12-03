package vpn

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
	"github.com/ironman0x7b2/sentinel-sdk/x/ibc"
)

func TestHandler_RegisterNode(t *testing.T) {
	ms, ibcKey, vpnKey, sessionKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abciTypes.Header{}, false, log.NewNopLogger())

	cdc := codec.New()

	codec.RegisterCrypto(cdc)
	sdkTypes.RegisterCodec(cdc)
	ibc.RegisterCodec(cdc)
	hub.RegisterCodec(cdc)

	vpnKeeper := NewKeeper(cdc, vpnKey, sessionKey)
	ibcKeeper := ibc.NewKeeper(ibcKey, cdc)

	var result csdkTypes.Result

	msg := *msgRegisterNode
	msg.LockerID = sdkTypes.KeyVPN + "/" + accAddress1.String() + "/" + strconv.Itoa(1)
	msg.Signature = getMsgLockCoinsSignature(sdkTypes.KeyVPN+"/"+accAddress1.String()+"/"+strconv.Itoa(1), csdkTypes.Coins{coin(100, "x")})
	result = handleRegisterNode(ctx, vpnKeeper, ibcKeeper, msg)
	require.Equal(t, errorLockerIDMismatch().Result(), result)

	msg.LockerID = vpnKey.Name() + "/" + accAddress1.String() + "/" + strconv.Itoa(0)
	msg.Signature = getMsgLockCoinsSignature(msg.LockerID, csdkTypes.Coins{coin(100, "x")})
	result = handleRegisterNode(ctx, vpnKeeper, ibcKeeper, msg)
	require.Equal(t, csdkTypes.Result{}, result)
}

func TestHandler_SessionPayment(t *testing.T) {
	ms, ibcKey, vpnKey, sessionKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abciTypes.Header{}, false, log.NewNopLogger())

	cdc := codec.New()

	codec.RegisterCrypto(cdc)
	sdkTypes.RegisterCodec(cdc)
	ibc.RegisterCodec(cdc)
	hub.RegisterCodec(cdc)

	vpnKeeper := NewKeeper(cdc, vpnKey, sessionKey)
	ibcKeeper := ibc.NewKeeper(ibcKey, cdc)

	var result csdkTypes.Result

	msg0 := *msgRegisterNode
	result = handleRegisterNode(ctx, vpnKeeper, ibcKeeper, msg0)
	require.Equal(t, csdkTypes.Result{}, result)

	msg1 := *msgPayVPNService
	msg1.LockerID = sdkTypes.KeySession + "/" + accAddress1.String() + "/" + strconv.Itoa(1)
	msg1.Signature = getMsgLockCoinsSignature(sdkTypes.KeySession+"/"+accAddress1.String()+"/"+strconv.Itoa(1), csdkTypes.Coins{coin(100, "x")})
	result = handleSessionPayment(ctx, vpnKeeper, ibcKeeper, msg1)
	require.Equal(t, errorLockerIDMismatch().Result(), result)

	msg1.LockerID = sessionKey.Name() + "/" + accAddress1.String() + "/" + strconv.Itoa(0)
	msg1.Signature = getMsgLockCoinsSignature(msg1.LockerID, csdkTypes.Coins{coin(100, "x")})
	result = handleSessionPayment(ctx, vpnKeeper, ibcKeeper, msg1)
	require.Equal(t, csdkTypes.Result{}, result)

	msg1.VPNID = accAddress1.String() + "/" + strconv.Itoa(1)
	msg1.LockerID = sessionKey.Name() + "/" + accAddress1.String() + "/" + strconv.Itoa(1)
	msg1.Signature = getMsgLockCoinsSignature(msg1.LockerID, csdkTypes.Coins{coin(100, "x")})
	result = handleSessionPayment(ctx, vpnKeeper, ibcKeeper, msg1)
	require.Equal(t, errorVPNNotExists().Result(), result)
}

func TestHandler_DeregisterNode(t *testing.T) {
	ms, ibcKey, vpnKey, sessionKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abciTypes.Header{}, false, log.NewNopLogger())

	cdc := codec.New()

	codec.RegisterCrypto(cdc)
	sdkTypes.RegisterCodec(cdc)
	ibc.RegisterCodec(cdc)
	hub.RegisterCodec(cdc)

	vpnKeeper := NewKeeper(cdc, vpnKey, sessionKey)
	ibcKeeper := ibc.NewKeeper(ibcKey, cdc)

	var result csdkTypes.Result

	msg0 := *msgRegisterNode
	result = handleRegisterNode(ctx, vpnKeeper, ibcKeeper, msg0)
	require.Equal(t, csdkTypes.Result{}, result)

	msg1 := *msgDeregisterNode
	msg1.VPNID = accAddress1.String() + "/" + strconv.Itoa(1)
	msg1.Signature = getMsgReleaseCoinsSignature(sdkTypes.KeyVPN + "/" + accAddress1.String() + "/" + strconv.Itoa(1))
	result = handleDeregisterNode(ctx, vpnKeeper, ibcKeeper, msg1)
	require.Equal(t, errorVPNNotExists().Result(), result)

	msg1.From = accAddress2
	msg1.VPNID = accAddress1.String() + "/" + strconv.Itoa(0)
	result = handleDeregisterNode(ctx, vpnKeeper, ibcKeeper, msg1)
	require.Equal(t, errorInvalidNodeOwnerAddress().Result(), result)

	msg1.From = accAddress1
	msg1.LockerID = vpnKey.Name() + "/" + accAddress1.String() + "/" + strconv.Itoa(1)
	result = handleDeregisterNode(ctx, vpnKeeper, ibcKeeper, msg1)
	require.Equal(t, errorLockerIDMismatch().Result(), result)

	msg1.LockerID = vpnKey.Name() + "/" + accAddress1.String() + "/" + strconv.Itoa(0)

	result = handleDeregisterNode(ctx, vpnKeeper, ibcKeeper, msg1)
	require.Equal(t, csdkTypes.Result{}, result)
}

func TestNewHandler(t *testing.T) {
	ms, ibcKey, vpnKey, sessionKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abciTypes.Header{}, false, log.NewNopLogger())

	cdc := codec.New()

	codec.RegisterCrypto(cdc)
	sdkTypes.RegisterCodec(cdc)
	ibc.RegisterCodec(cdc)
	hub.RegisterCodec(cdc)

	vpnKeeper := NewKeeper(cdc, vpnKey, sessionKey)
	ibcKeeper := ibc.NewKeeper(ibcKey, cdc)

	var result csdkTypes.Result

	msg0 := *msgRegisterNode
	msg1 := *msgPayVPNService
	msg2 := *msgDeregisterNode
	msg3 := csdkTypes.NewTestMsg(accAddress1)

	handler := NewHandler(vpnKeeper, ibcKeeper)
	result = handler(ctx, msg0)
	require.Equal(t, csdkTypes.Result{}, result)

	result = handler(ctx, msg1)
	require.Equal(t, csdkTypes.Result{}, result)

	result = handler(ctx, msg2)
	require.Equal(t, csdkTypes.Result{}, result)

	result = handler(ctx, msg3)
	require.Equal(t, csdkTypes.ErrUnknownRequest("Unrecognized msg type: " + reflect.TypeOf(msg3).Name()).Result(), result)
}
