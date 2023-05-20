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
		deposit                  sdk.Coin
		expiryDuration           time.Duration
		maxGigabytePrices        sdk.Coins
		minGigabytePrices        sdk.Coins
		maxHourlyPrices          sdk.Coins
		minHourlyPrices          sdk.Coins
		maxSubscriptionGigabytes int64
		minSubscriptionGigabytes int64
		maxSubscriptionHours     int64
		minSubscriptionHours     int64
		revenueShare             sdk.Dec
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
		string(types.KeyExpiryDuration),
		&expiryDuration,
		state.Rand,
		func(r *rand.Rand) {
			expiryDuration = time.Duration(r.Int63n(MaxInt)) * time.Millisecond
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
			deposit, expiryDuration, maxGigabytePrices, minGigabytePrices,
			maxHourlyPrices, minHourlyPrices, maxSubscriptionGigabytes, minSubscriptionGigabytes,
			maxSubscriptionHours, minSubscriptionHours, revenueShare,
		),
	)
}
