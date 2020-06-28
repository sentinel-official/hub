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
	cdc.RegisterConcrete(MsgAddPlan{}, "x/dvpn/plan/MsgAddPlan", nil)
	cdc.RegisterConcrete(MsgSetPlanStatus{}, "x/dvpn/plan/MsgSetPlanStatus", nil)
	cdc.RegisterConcrete(MsgAddNodeForPlan{}, "x/dvpn/plan/MsgAddNodeForPlan", nil)
	cdc.RegisterConcrete(MsgRemoveNodeForPlan{}, "x/dvpn/plan/MsgRemoveNodeForPlan", nil)
}
