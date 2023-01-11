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
		stakingShare sdk.Dec
	)

	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyDeposit),
		&deposit,
		state.Rand,
		func(r *rand.Rand) {
			deposit = sdk.NewInt64Coin(
				sdk.DefaultBondDenom,
				r.Int63n(MaxInt),
			)
		},
	)
	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyStakingShare),
		&stakingShare,
		state.Rand,
		func(r *rand.Rand) {
			stakingShare = sdk.NewDecWithPrec(
				r.Int63n(MaxInt),
				6,
			)
		},
	)

	return types.NewGenesisState(
		RandomProviders(state.Rand, state.Accounts),
		types.NewParams(deposit, stakingShare),
	)
}
