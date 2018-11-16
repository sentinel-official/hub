package vpn

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgRegisterNode{}, "x/vpn/msg_register_node", nil)
	cdc.RegisterConcrete(MsgUpdateNodeStatus{}, "x/vpn/msg_update_node_status", nil)
	cdc.RegisterConcrete(MsgPayVpnService{},"x/v/msg_pay_vpn_service", nil)
}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
}
