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
	cdc.RegisterConcrete(MsgRegister{}, "x/node/MsgRegister", nil)
	cdc.RegisterConcrete(MsgUpdate{}, "x/node/MsgUpdate", nil)
	cdc.RegisterConcrete(MsgSetStatus{}, "x/node/MsgSetStatus", nil)
}
