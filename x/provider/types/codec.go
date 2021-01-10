package types

import (
	"fmt"

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
	cdc.RegisterConcrete(MsgRegister{}, fmt.Sprintf("x/%s/MsgRegister", ModuleName), nil)
	cdc.RegisterConcrete(MsgUpdate{}, fmt.Sprintf("x/%s/MsgUpdate", ModuleName), nil)
}
