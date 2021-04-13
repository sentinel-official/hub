package params

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
)

func MakeEncodingConfig() EncodingConfig {
	var (
		amino             = codec.NewLegacyAmino()
		interfaceRegistry = types.NewInterfaceRegistry()
		marshaler         = codec.NewProtoCodec(interfaceRegistry)
		txConfig          = tx.NewTxConfig(marshaler, tx.DefaultSignModes)
	)

	return EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txConfig,
		Amino:             amino,
	}
}
