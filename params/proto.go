package params

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
)

func MakeEncodingConfig() EncodingConfig {
	var (
		amino             = codec.NewLegacyAmino()
		interfaceRegistry = codectypes.NewInterfaceRegistry()
		marshaler         = codec.NewProtoCodec(interfaceRegistry)
		txConfig          = authtx.NewTxConfig(marshaler, authtx.DefaultSignModes)
	)

	return EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txConfig,
		Amino:             amino,
	}
}
