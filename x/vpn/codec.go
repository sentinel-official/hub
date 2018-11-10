package vpn

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ironman0x7b2/sentinel-hub/types"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgRegisterVpn{}, "vpn/register", nil)
	cdc.RegisterConcrete(MsgAliveNode{}, "vpn/AliveNode", nil)
	cdc.RegisterConcrete(types.IBCMsgRegisterVpn{},"vpn/ibc",nil)
}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
}
