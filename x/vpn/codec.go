package vpn

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgRegisterVpn{}, "x/vpn/msg_register_vpn", nil)
	cdc.RegisterConcrete(MsgNodeStatus{}, "x/vpn/msg_alive_node", nil)
}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
}
