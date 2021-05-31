package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sentinel-official/hub/x/provider/types"
)

func getRandomDeposit(r *rand.Rand) sdk.Coin {
	denom := simulation.RandStringOfLength(r, r.Intn(125)+3)
	amount := sdk.NewInt(r.Int63n(10<<12))

	return sdk.NewCoin(denom, amount)
}

func getRandomProviders(r *rand.Rand) types.Providers {
	var providers types.Providers

	for _, acc := range simulation.RandomAccounts(r, r.Intn(20)+4) {
		providers = append(providers, types.Provider{
			Address:     acc.Address.String(),
			Name:        simulation.RandStringOfLength(r, r.Intn(60)+4),
			Identity:    simulation.RandStringOfLength(r, r.Intn(60)+4),
			Website:     simulation.RandStringOfLength(r, r.Intn(60)+4),
			Description: simulation.RandStringOfLength(r, r.Intn(250)+6),
		})
	}

	return providers
}

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
