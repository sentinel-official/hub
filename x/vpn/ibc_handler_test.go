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

func TestIBCHandler_UpdateNodeStatus(t *testing.T) {
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

	msg1 := msgIBCTransaction
	msg1.Sequence = 1
	result = handleUpdateNodeStatus(ctx, ibcKeeper, vpnKeeper, msg1)
	require.Equal(t, errorInvalidIBCSequence().Result(), result)

	msg1.Sequence = 0
	msg1.IBCPacket.Message = hub.MsgLockerStatus{
		vpnKey.Name() + "/" + accAddress1.String() + "/" + strconv.Itoa(1),
		sdkTypes.StatusLock,
	}
	result = handleUpdateNodeStatus(ctx, ibcKeeper, vpnKeeper, msg1)
	require.Equal(t, errorVPNNotExists().Result(), result)

	msg1.IBCPacket.Message = hub.MsgLockerStatus{
		vpnKey.Name() + "/" + accAddress1.String() + "/" + strconv.Itoa(0),
		"INVALID",
	}
	result = handleUpdateNodeStatus(ctx, ibcKeeper, vpnKeeper, msg1)
	require.Equal(t, errorInvalidLockStatus().Result(), result)

	msg1.IBCPacket.Message = hub.MsgLockerStatus{
		vpnKey.Name() + "/" + accAddress1.String() + "/" + strconv.Itoa(0),
		sdkTypes.StatusLock,
	}
	result = handleUpdateNodeStatus(ctx, ibcKeeper, vpnKeeper, msg1)
	require.Equal(t, csdkTypes.Result{}, result)

	msg1.Sequence = 1
	msg1.IBCPacket.Message = hub.MsgLockerStatus{
		vpnKey.Name() + "/" + accAddress1.String() + "/" + strconv.Itoa(0),
		sdkTypes.StatusRelease,
	}
	result = handleUpdateNodeStatus(ctx, ibcKeeper, vpnKeeper, msg1)
	require.Equal(t, csdkTypes.Result{}, result)
}

func TestIBCHandler_UpdateSessionStatus(t *testing.T) {
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
	result = handleSessionPayment(ctx, vpnKeeper, ibcKeeper, msg1)
	require.Equal(t, csdkTypes.Result{}, result)

	msg2 := msgIBCTransaction
	msg2.Sequence = 1
	msg2.IBCPacket.Message = hub.MsgLockerStatus{
		sessionKey.Name() + "/" + accAddress1.String() + "/" + strconv.Itoa(0),
		sdkTypes.StatusLock,
	}
	result = handleUpdateSessionStatus(ctx, ibcKeeper, vpnKeeper, msg2)
	require.Equal(t, errorInvalidIBCSequence().Result(), result)

	msg2.Sequence = 0
	msg2.IBCPacket.Message = hub.MsgLockerStatus{
		sessionKey.Name() + "/" + accAddress1.String() + "/" + strconv.Itoa(1),
		sdkTypes.StatusLock,
	}
	result = handleUpdateSessionStatus(ctx, ibcKeeper, vpnKeeper, msg2)
	require.Equal(t, errorSessionNotExists().Result(), result)

	msg2.IBCPacket.Message = hub.MsgLockerStatus{
		sessionKey.Name() + "/" + accAddress1.String() + "/" + strconv.Itoa(0),
		"INVALID",
	}
	result = handleUpdateSessionStatus(ctx, ibcKeeper, vpnKeeper, msg2)
	require.Equal(t, errorInvalidLockStatus().Result(), result)

	msg2.IBCPacket.Message = hub.MsgLockerStatus{
		sessionKey.Name() + "/" + accAddress1.String() + "/" + strconv.Itoa(0),
		sdkTypes.StatusLock,
	}
	result = handleUpdateSessionStatus(ctx, ibcKeeper, vpnKeeper, msg2)
	require.Equal(t, csdkTypes.Result{}, result)

	msg2.Sequence = 1
	msg2.IBCPacket.Message = hub.MsgLockerStatus{
		sessionKey.Name() + "/" + accAddress1.String() + "/" + strconv.Itoa(0),
		sdkTypes.StatusRelease,
	}
	result = handleUpdateSessionStatus(ctx, ibcKeeper, vpnKeeper, msg2)
	require.Equal(t, csdkTypes.Result{}, result)
}

func TestNewIBCVPNHandler(t *testing.T) {
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
	result = handleSessionPayment(ctx, vpnKeeper, ibcKeeper, msg1)
	require.Equal(t, csdkTypes.Result{}, result)

	msg2 := msgIBCTransaction
	msg2.IBCPacket.Message = hub.MsgLockerStatus{
		vpnKey.Name() + "/" + accAddress1.String() + "/" + strconv.Itoa(0),
		sdkTypes.StatusLock,
	}
	msg3 := msgIBCTransaction
	msg3.Sequence = 1
	msg3.IBCPacket.Message = hub.MsgLockerStatus{
		sessionKey.Name() + "/" + accAddress1.String() + "/" + strconv.Itoa(0),
		sdkTypes.StatusLock,
	}
	msg4 := csdkTypes.NewTestMsg(accAddress1)

	handler := NewIBCVPNHandler(ibcKeeper, vpnKeeper)
	result = handler(ctx, msg2)
	require.Equal(t, csdkTypes.Result{}, result)

	result = handler(ctx, msg3)
	require.Equal(t, csdkTypes.Result{}, result)

	msg2.IBCPacket.Message = hub.MsgLockerStatus{
		"unknown" + "/" + accAddress1.String() + "/" + strconv.Itoa(0),
		sdkTypes.StatusLock,
	}
	result = handler(ctx, msg2)
	require.Equal(t, csdkTypes.ErrUnknownRequest("Unrecognized locker id: " + msg2.IBCPacket.Message.(hub.MsgLockerStatus).LockerID).Result(), result)

	msg3.IBCPacket.Message = csdkTypes.TestMsg{}
	result = handler(ctx, msg3)
	require.Equal(t, csdkTypes.ErrUnknownRequest("Unrecognized IBC msg type: " + reflect.TypeOf(msg3.IBCPacket.Message).Name()).Result(), result)

	result = handler(ctx, msg4)
	require.Equal(t, csdkTypes.ErrUnknownRequest("Unrecognized msg type: " + reflect.TypeOf(msg4).Name()).Result(), result)
}
