package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/sentinel-official/hub/x/provider/types"
)

func RandomizedGenesisState(simState *module.SimulationState) *types.GenesisState {

	var deposit sdk.Coin

	depositSim := func(r *rand.Rand) {
		deposit = getRandomDeposit(r)
	}

	simState.AppParams.GetOrGenerate(simState.Cdc, string(types.KeyDeposit), &deposit, simState.Rand, depositSim)

	params := types.NewParams(deposit)
	providers := getRandomProviders(simState.Rand)

	state := types.NewGenesisState(providers, params)

	return state
}
