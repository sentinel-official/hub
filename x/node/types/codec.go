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
	cdc.RegisterConcrete(MsgRegisterNode{}, "sentinel/MsgRegisterNode", nil)
	cdc.RegisterConcrete(MsgUpdateNode{}, "sentinel/MsgUpdateNode", nil)
	cdc.RegisterConcrete(MsgSetNodeStatus{}, "sentinel/MsgSetNodeStatus", nil)
}
