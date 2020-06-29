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
	cdc.RegisterConcrete(MsgAddPlan{}, "sentinel/MsgAddPlan", nil)
	cdc.RegisterConcrete(MsgSetPlanStatus{}, "sentinel/MsgSetPlanStatus", nil)
	cdc.RegisterConcrete(MsgAddNodeForPlan{}, "sentinel/MsgAddNodeForPlan", nil)
	cdc.RegisterConcrete(MsgRemoveNodeForPlan{}, "sentinel/MsgRemoveNodeForPlan", nil)
}
