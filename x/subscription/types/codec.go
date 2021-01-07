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
	cdc.RegisterConcrete(MsgSubscribeToPlan{}, fmt.Sprintf("x/%s/MsgSubscribeToPlan", ModuleName), nil)
	cdc.RegisterConcrete(MsgSubscribeToNode{}, fmt.Sprintf("x/%s/MsgSubscribeToNode", ModuleName), nil)
	cdc.RegisterConcrete(MsgCancel{}, fmt.Sprintf("x/%s/MsgCancel", ModuleName), nil)

	cdc.RegisterConcrete(MsgAddQuota{}, fmt.Sprintf("x/%s/MsgAddQuota", ModuleName), nil)
	cdc.RegisterConcrete(MsgUpdateQuota{}, fmt.Sprintf("x/%s/MsgUpdateQuota", ModuleName), nil)
}
