package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgRegisterNode{}, "x/vpn/msg_register_node", nil)
	cdc.RegisterConcrete(MsgUpdateNodeDetails{}, "x/vpn/msg_update_node_details", nil)
	cdc.RegisterConcrete(MsgUpdateNodeStatus{}, "x/vpn/msg_update_node_status", nil)
	cdc.RegisterConcrete(MsgDeregisterNode{}, "x/vpn/msg_deregister_node", nil)

	cdc.RegisterConcrete(MsgAddSession{}, "x/vpn/msg_add_session", nil)
	cdc.RegisterConcrete(MsgUpdateSessionBandwidth{}, "x/vpn/msg_update_session_bandwidth", nil)
}

var vpnCdc = codec.New()

func init() {
	RegisterCodec(vpnCdc)
}
