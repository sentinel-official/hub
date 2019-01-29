package vpn

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgRegisterNode{}, "x/vpn/msg_register_node", nil)
	cdc.RegisterConcrete(MsgUpdateNode{}, "x/vpn/msg_update_node", nil)
	cdc.RegisterConcrete(MsgDeregisterNode{}, "x/vpn/msg_deregister_node", nil)
}

var vpnCdc = codec.New()

func init() {
	RegisterCodec(vpnCdc)
}
