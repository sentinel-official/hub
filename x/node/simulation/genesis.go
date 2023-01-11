package simulation

import (
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/hub/x/node/types"
)

func RandomizedGenesisState(state *module.SimulationState) *types.GenesisState {
	var (
		deposit          sdk.Coin
		inactiveDuration time.Duration
		maxPrice         sdk.Coins
		minPrice         sdk.Coins
		stakingShare     sdk.Dec
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
		string(types.KeyInactiveDuration),
		&inactiveDuration,
		state.Rand,
		func(r *rand.Rand) {
			inactiveDuration = time.Duration(r.Int63n(MaxInt)) * time.Millisecond
		},
	)
	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyMaxPrice),
		&maxPrice,
		state.Rand,
		func(r *rand.Rand) {
			maxPrice = sdk.NewCoins(
				sdk.NewInt64Coin(
					sdk.DefaultBondDenom,
					r.Int63n(MaxInt),
				),
			)
		},
	)
	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyMinPrice),
		&minPrice,
		state.Rand,
		func(r *rand.Rand) {
			minPrice = sdk.NewCoins(
				sdk.NewInt64Coin(
					sdk.DefaultBondDenom,
					r.Int63n(MaxInt),
				),
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
		RandomNodes(state.Rand, state.Accounts),
		types.NewParams(deposit, inactiveDuration, maxPrice, minPrice, stakingShare),
	)
}
