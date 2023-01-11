package hub

import (
	sdkstd "github.com/cosmos/cosmos-sdk/std"

	hubparams "github.com/sentinel-official/hub/params"
)

func MakeEncodingConfig() hubparams.EncodingConfig {
	config := hubparams.MakeEncodingConfig()
	sdkstd.RegisterLegacyAminoCodec(config.Amino)
	sdkstd.RegisterInterfaces(config.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(config.Amino)
	ModuleBasics.RegisterInterfaces(config.InterfaceRegistry)
	return config
}
