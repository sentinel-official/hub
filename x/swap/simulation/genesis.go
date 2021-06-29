package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/sentinel-official/hub/x/swap/types"
)

func RandomizedGenesisState(simState *module.SimulationState) {
	var (
		swapEnabled bool
		approveBy   string
		swapDenom   string
	)

	swapEnabledSim := func(r *rand.Rand) {
		swapEnabled = GetRandomSwapEnabled(r)
	}

	approveBySim := func(r *rand.Rand) {
		approveBy = GetRandomApproveBy(r)
	}

	swapDenomSim := func(r *rand.Rand) {
		swapDenom = GetRandomSwapDenom(r)
	}

	simState.AppParams.GetOrGenerate(simState.Cdc, string(types.KeySwapEnabled), &swapEnabled, simState.Rand, swapEnabledSim)
	simState.AppParams.GetOrGenerate(simState.Cdc, string(types.KeyApproveBy), &approveBy, simState.Rand, approveBySim)
	simState.AppParams.GetOrGenerate(simState.Cdc, string(types.KeySwapDenom), &swapDenom, simState.Rand, swapDenomSim)

	params := types.NewParams(swapEnabled, swapDenom, approveBy)

	var swaps types.Swaps

	swaps = append(swaps, types.Swap{
		TxHash:   []byte("b37c4b36298f20ac7fc5de2b4ac1167d7a044402a40f2b5f1e1bb7f0ee5f6346"),
		Receiver: "sent1grdunxx5jxd0ja75wt508sn6v39p70hhw53zs8",
		Amount:   sdk.Coin{Denom: "sent", Amount: sdk.NewInt(1000)},
	})

	swapGenesis := types.NewGenesisState(swaps, params)
	bz, err := json.MarshalIndent(&swapGenesis.Params, "", "")
	if err != nil {
		panic(err)
	}

	fmt.Printf("selected randomly generated swap parameters: %s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(swapGenesis)
}
