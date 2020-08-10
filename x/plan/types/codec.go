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
	cdc.RegisterConcrete(MsgAdd{}, "x/plan/MsgAdd", nil)
	cdc.RegisterConcrete(MsgSetStatus{}, "x/plan/MsgSetStatus", nil)
	cdc.RegisterConcrete(MsgAddNode{}, "x/plan/MsgAddNode", nil)
	cdc.RegisterConcrete(MsgRemoveNode{}, "x/plan/MsgRemoveNode", nil)
}
