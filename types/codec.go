package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*ID)(nil), nil)
	cdc.RegisterConcrete(ProviderID{}, "types/provider_id", nil)
	cdc.RegisterConcrete(NodeID{}, "types/node_id", nil)
	cdc.RegisterConcrete(SubscriptionID{}, "types/subscription_id", nil)
	cdc.RegisterConcrete(SessionID{}, "types/session_id", nil)
}
