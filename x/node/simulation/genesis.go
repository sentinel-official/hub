// DO NOT COVER

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
		deposit           sdk.Coin
		inactiveDuration  time.Duration
		maxGigabytePrices sdk.Coins
		minGigabytePrices sdk.Coins
		maxHourlyPrices   sdk.Coins
		minHourlyPrices   sdk.Coins
		maxLeaseHours     int64
		minLeaseHours     int64
		maxLeaseGigabytes int64
		minLeaseGigabytes int64
		revenueShare      sdk.Dec
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
		string(types.KeyMaxGigabytePrices),
		&maxGigabytePrices,
		state.Rand,
		func(r *rand.Rand) {
			maxGigabytePrices = sdk.NewCoins(
				sdk.NewInt64Coin(
					sdk.DefaultBondDenom,
					r.Int63n(MaxInt),
				),
			)
		},
	)
	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyMinGigabytePrices),
		&minGigabytePrices,
		state.Rand,
		func(r *rand.Rand) {
			minGigabytePrices = sdk.NewCoins(
				sdk.NewInt64Coin(
					sdk.DefaultBondDenom,
					r.Int63n(MaxInt),
				),
			)
		},
	)
	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyMaxHourlyPrices),
		&maxHourlyPrices,
		state.Rand,
		func(r *rand.Rand) {
			maxHourlyPrices = sdk.NewCoins(
				sdk.NewInt64Coin(
					sdk.DefaultBondDenom,
					r.Int63n(MaxInt),
				),
			)
		},
	)
	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyMinHourlyPrices),
		&minHourlyPrices,
		state.Rand,
		func(r *rand.Rand) {
			minHourlyPrices = sdk.NewCoins(
				sdk.NewInt64Coin(
					sdk.DefaultBondDenom,
					r.Int63n(MaxInt),
				),
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
				r.Int63n(MaxInt),
				6,
			)
		},
	)

	return types.NewGenesisState(
		RandomNodes(state.Rand, state.Accounts),
		types.NewParams(
			deposit, inactiveDuration, maxGigabytePrices, minGigabytePrices,
			maxHourlyPrices, minHourlyPrices, maxLeaseHours, minLeaseHours,
			maxLeaseGigabytes, minLeaseGigabytes, revenueShare,
		),
	)
}
