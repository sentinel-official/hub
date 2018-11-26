package vpn

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgRegisterNode{}, "x/vpn/msg_register_node", nil)
	cdc.RegisterConcrete(MsgPayVPNService{}, "x/vpn/msg_pay_vpn_service", nil)
	cdc.RegisterConcrete(MsgUpdateNodeStatus{}, "x/vpn/msg_update_node_status", nil)
	cdc.RegisterConcrete(MsgUpdateSessionStatus{}, "x/vpn/msg_update_session_status", nil)
	cdc.RegisterConcrete(MsgDeregisterNode{}, "x/vpn/msg_deregister_node", nil)
}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
}
