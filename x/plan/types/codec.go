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
	cdc.RegisterConcrete(MsgAdd{}, fmt.Sprintf("x/%s/MsgAdd", ModuleName), nil)
	cdc.RegisterConcrete(MsgSetStatus{}, fmt.Sprintf("x/%s/MsgSetStatus", ModuleName), nil)
	cdc.RegisterConcrete(MsgAddNode{}, fmt.Sprintf("x/%s/MsgAddNode", ModuleName), nil)
	cdc.RegisterConcrete(MsgRemoveNode{}, fmt.Sprintf("x/%s/MsgRemoveNode", ModuleName), nil)
}
