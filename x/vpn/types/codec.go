package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

var (
	ModuleCdc *codec.Codec
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgRegisterNode{}, "x/vpn/MsgRegisterNode", nil)
	cdc.RegisterConcrete(MsgUpdateNodeInfo{}, "x/vpn/MsgUpdateNodeInfo", nil)
	cdc.RegisterConcrete(MsgAddFreeClient{}, "x/vpn/MsgAddFreeClient", nil)
	cdc.RegisterConcrete(MsgDeregisterNode{}, "x/vpn/MsgDeregisterNode", nil)
	cdc.RegisterConcrete(MsgStartSubscription{}, "x/vpn/MsgStartSubscription", nil)
	cdc.RegisterConcrete(MsgEndSubscription{}, "x/vpn/MsgEndSubscription", nil)
	cdc.RegisterConcrete(MsgUpdateSessionInfo{}, "x/vpn/MsgUpdateSessionInfo", nil)
}

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
