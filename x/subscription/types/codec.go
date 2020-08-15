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
	cdc.RegisterConcrete(MsgSubscribeToPlan{}, "x/subscription/MsgSubscribeToPlan", nil)
	cdc.RegisterConcrete(MsgSubscribeToNode{}, "x/subscription/MsgSubscribeToNode", nil)
	cdc.RegisterConcrete(MsgCancel{}, "x/subscription/MsgCancel", nil)

	cdc.RegisterConcrete(MsgAddQuota{}, "x/subscription/MsgAddQuota", nil)
	cdc.RegisterConcrete(MsgUpdateQuota{}, "x/subscription/MsgUpdateQuota", nil)
}
