package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

var (
	ModuleCdc *codec.Codec
)

func init() {
	ModuleCdc = codec.New()
	codec.RegisterCrypto(ModuleCdc)
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgRegisterNode{}, "x/dvpn/node/MsgRegisterNode", nil)
	cdc.RegisterConcrete(MsgUpdateNode{}, "x/dvpn/node/MsgUpdateNode", nil)
}
