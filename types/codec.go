package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*ID)(nil), nil)
	cdc.RegisterConcrete(SubscriptionID{}, "types/SubscriptionID", nil)
	cdc.RegisterConcrete(SessionID{}, "types/SessionID", nil)
}
