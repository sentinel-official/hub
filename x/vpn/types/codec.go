package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"

	deposittypes "github.com/sentinel-official/hub/x/deposit/types"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
	plantypes "github.com/sentinel-official/hub/x/plan/types"
	providertypes "github.com/sentinel-official/hub/x/provider/types"
	sessiontypes "github.com/sentinel-official/hub/x/session/types"
	subscriptiontypes "github.com/sentinel-official/hub/x/subscription/types"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

func RegisterLegacyAminoCodec(_ *codec.LegacyAmino) {}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	deposittypes.RegisterInterfaces(registry)
	providertypes.RegisterInterfaces(registry)
	nodetypes.RegisterInterfaces(registry)
	plantypes.RegisterInterfaces(registry)
	subscriptiontypes.RegisterInterfaces(registry)
	sessiontypes.RegisterInterfaces(registry)
}
