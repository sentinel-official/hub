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
	cdc.RegisterConcrete(MsgAddPlan{}, "x/dvpn/subscription/MsgAddPlan", nil)
	cdc.RegisterConcrete(MsgSetPlanStatus{}, "x/dvpn/subscription/MsgSetPlanStatus", nil)
	cdc.RegisterConcrete(MsgAddNode{}, "x/dvpn/subscription/MsgAddNode", nil)
	cdc.RegisterConcrete(MsgRemoveNode{}, "x/dvpn/subscription/MsgRemoveNode", nil)
}
