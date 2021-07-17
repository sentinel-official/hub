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
	)

	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyDeposit),
		&deposit,
		state.Rand,
		func(r *rand.Rand) {
			deposit = sdk.NewInt64Coin(
				sdk.DefaultBondDenom,
				r.Int63n(MaxDepositAmount),
			)
		},
	)
	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyInactiveDuration),
		&inactiveDuration,
		state.Rand,
		func(r *rand.Rand) {
			inactiveDuration = time.Duration(r.Int63n(MaxInactiveDuration)) * time.Millisecond
		},
	)

	return types.NewGenesisState(
		RandomNodes(state.Rand, state.Accounts),
		types.NewParams(deposit, inactiveDuration),
	)
}
