package ibc

import (
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	"testing"
)

func TestIBC(t *testing.T) {

	//var egressKey int64
	//var ingressKey int64

	cdc := codec.New()
	cdc.RegisterConcrete(vpn.MsgRegisterNode{},"test/ibc/msg_register_node",nil)
	multiStore, vpnKey, _, ibcKey, sessionKey := DefaultSetup()
	ctx := csdkTypes.NewContext(multiStore, abci.Header{}, false, log.NewNopLogger())

	vpnKeeper := vpn.NewKeeper(cdc, vpnKey, sessionKey)
	ibcKeeper := NewKeeper(ibcKey, cdc)

	vpnHandler := vpn.NewHandler(vpnKeeper, ibcKeeper)

	msgRegisterNode := TestNewMsgRegisterNode()
	msgRegisterNodeRes := vpnHandler(ctx, msgRegisterNode)

	require.True(t, msgRegisterNodeRes.IsOK())
}
