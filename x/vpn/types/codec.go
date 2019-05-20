package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgRegisterNode{}, "x/vpn/MsgRegisterNode", nil)
	cdc.RegisterConcrete(MsgUpdateNodeInfo{}, "x/vpn/MsgUpdateNodeInfo", nil)
	cdc.RegisterConcrete(MsgUpdateNodeStatus{}, "x/vpn/MsgUpdateNodeStatus", nil)
	cdc.RegisterConcrete(MsgDeregisterNode{}, "x/vpn/MsgDeregisterNode", nil)

	cdc.RegisterConcrete(MsgStartSubscription{}, "x/vpn/MsgStartSubscription", nil)
	cdc.RegisterConcrete(MsgEndSubscription{}, "x/vpn/MsgEndSubscription", nil)

	cdc.RegisterConcrete(MsgUpdateSessionInfo{}, "x/vpn/MsgUpdateSessionInfo", nil)
}

var cdc = codec.New()

func init() {
	RegisterCodec(cdc)
}
