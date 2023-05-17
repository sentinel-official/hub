// DO NOT COVER

package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/hub/x/provider/types"
)

func RandomizedGenesisState(state *module.SimulationState) *types.GenesisState {
	var (
		deposit      sdk.Coin
		revenueShare sdk.Dec
	)

	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyDeposit),
		&deposit,
		state.Rand,
		func(r *rand.Rand) {
			deposit = sdk.NewInt64Coin(
				sdk.DefaultBondDenom,
				r.Int63n(MaxDeposit),
			)
		},
	)
	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyRevenueShare),
		&revenueShare,
		state.Rand,
		func(r *rand.Rand) {
			revenueShare = sdk.NewDecWithPrec(
				r.Int63n(MaxRevenueShare),
				6,
			)
		},
	)

	return types.NewGenesisState(
		RandomProviders(state.Rand, state.Accounts),
		types.NewParams(deposit, revenueShare),
	)
}
