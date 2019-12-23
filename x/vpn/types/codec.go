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
	cdc.RegisterConcrete(MsgRemoveFreeClient{}, "x/vpn/MsgRemoveFreeClient", nil)
	cdc.RegisterConcrete(MsgRegisterVPNOnResolver{}, "x/vpn/MsgRegisterVPNOnResolver", nil)
	cdc.RegisterConcrete(MsgDeregisterVPNOnResolver{}, "x/vpn/MsgDeregisterVPNOnResolver", nil)
	cdc.RegisterConcrete(MsgDeregisterNode{}, "x/vpn/MsgDeregisterNode", nil)
	cdc.RegisterConcrete(MsgStartSubscription{}, "x/vpn/MsgStartSubscription", nil)
	cdc.RegisterConcrete(MsgEndSubscription{}, "x/vpn/MsgEndSubscription", nil)
	cdc.RegisterConcrete(MsgUpdateSessionInfo{}, "x/vpn/MsgUpdateSessionInfo", nil)
	cdc.RegisterConcrete(MsgRegisterResolver{}, "x/vpn/MsgRegisterResolver", nil)
	cdc.RegisterConcrete(MsgUpdateResolverInfo{}, "x/vpn/MsgUpdateResolverInfo", nil)
	cdc.RegisterConcrete(MsgDeregisterResolver{}, "x/vpn/MsgDeregisterResolver", nil)
}

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
