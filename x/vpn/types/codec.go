package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	crypto "github.com/cosmos/cosmos-sdk/crypto/codec"

	"github.com/sentinel-official/hub/x/deposit"
	"github.com/sentinel-official/hub/x/node"
	"github.com/sentinel-official/hub/x/plan"
	"github.com/sentinel-official/hub/x/provider"
	"github.com/sentinel-official/hub/x/session"
	"github.com/sentinel-official/hub/x/subscription"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	crypto.RegisterCrypto(amino)
	amino.Seal()
}

func RegisterLegacyAminoCodec(_ *codec.LegacyAmino) {}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	deposit.RegisterInterfaces(registry)
	provider.RegisterInterfaces(registry)
	node.RegisterInterfaces(registry)
	plan.RegisterInterfaces(registry)
	subscription.RegisterInterfaces(registry)
	session.RegisterInterfaces(registry)
}
