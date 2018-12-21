package ibc

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgIBCTransaction{}, "x/ibc/msg_ibc_transaction", nil)
}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
}
