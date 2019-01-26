package vpn

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgRegisterNode{}, "x/vpn/msg_register_node", nil)
}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
}
