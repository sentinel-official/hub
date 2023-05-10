// DO NOT COVER

package simulation

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/sentinel-official/hub/x/swap/types"
)

func RandomizedGenesisState(state *module.SimulationState) {
	var (
		account, _ = simulation.RandomAcc(state.Rand, state.Accounts)
		params     = types.NewParams(
			true,
			sdk.DefaultBondDenom,
			account.Address.String(),
		)
	)

	state.GenState[types.ModuleName] = state.Cdc.MustMarshalJSON(
		types.NewGenesisState(
			nil,
			params,
		),
	)
}
