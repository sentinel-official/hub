package simulation

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/sentinel-official/hub/x/provider/types"
)

func RandomizedGenesisState(_ *codec.Codec) types.GenesisState {
	genesis := types.DefaultGenesisState()
	return genesis
}
