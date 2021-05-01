package hub

import (
	"github.com/cosmos/cosmos-sdk/std"

	"github.com/sentinel-official/hub/params"
)

func MakeEncodingConfig() params.EncodingConfig {
	config := params.MakeEncodingConfig()
	std.RegisterLegacyAminoCodec(config.Amino)
	std.RegisterInterfaces(config.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(config.Amino)
	ModuleBasics.RegisterInterfaces(config.InterfaceRegistry)
	return config
}
