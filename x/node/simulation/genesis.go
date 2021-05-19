package simulation

import (
	"fmt"
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
)

func getRandomDeposit(r *rand.Rand) int64 {
	return int64(r.Intn(100) + 1)
}

func getRandomInactiveDuration(r *rand.Rand) time.Duration {
	return time.Duration(simulation.RandIntBetween(r, 60, 60<<13))
}

func getNodeAddress() hubtypes.NodeAddress {
	bz := make([]byte, 20)
	_, err := rand.Read(bz)
	if err != nil {
		panic(err)
	}

	return hubtypes.NodeAddress(bz)
}

func RandomizedGenState(simState *module.SimulationState) {

	var (
		deposit          int64
		inactiveDuration time.Duration
	)

	depositSim := func(r *rand.Rand) {
		deposit = getRandomDeposit(r)
	}

	inactiveDurationSim := func(r *rand.Rand) {
		inactiveDuration = getRandomInactiveDuration(r)
	}

	simState.AppParams.GetOrGenerate(simState.Cdc, string(types.KeyDeposit), *&deposit, nil, depositSim)
	simState.AppParams.GetOrGenerate(simState.Cdc, string(types.KeyInactiveDuration), *&inactiveDuration, nil, inactiveDurationSim)

	params := types.NewParams(sdk.Coin{Denom: "sent", Amount: sdk.NewInt(deposit)}, inactiveDuration)

	var nodes types.Nodes

	nodes = append(nodes, types.Node{
		Address:  getNodeAddress().String(),
		Provider: "",
		Price: []sdk.Coin{
			{Denom: "sent", Amount: sdk.NewInt(1000)},
		},
		RemoteURL: "https://sentinel.co",
		Status:    hubtypes.Active,
		StatusAt:  time.Now(),
	})

	state := types.NewGenesisState(nodes, params)
	bz := simState.Cdc.MustMarshalJSON(&state.Params)

	fmt.Printf("selected randomly generated nodes parameters: %s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(state)
}
