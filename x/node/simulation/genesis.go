package simulation

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/sentinel-official/hub/x/node/types"
)

func RandomizedGenesisState(cdc *codec.Codec) types.GenesisState {
	genesis := types.DefaultGenesisState()
	fmt.Printf("Selected randomly generated node parameters:\n%s\n", codec.MustMarshalJSONIndent(cdc, genesis.Params))
	return genesis
}
