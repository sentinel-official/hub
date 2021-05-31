package simulation

import (
	"fmt"
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
)

func RandomizedGenState(simState *module.SimulationState) *types.GenesisState {

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

	simState.AppParams.GetOrGenerate(simState.Cdc, string(types.KeyDeposit), deposit, nil, depositSim)
	simState.AppParams.GetOrGenerate(simState.Cdc, string(types.KeyInactiveDuration), inactiveDuration, nil, inactiveDurationSim)

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
	return state
}
