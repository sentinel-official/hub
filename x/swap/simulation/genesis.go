package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	hubtypes "github.com/sentinel-official/hub/x/swap/types"
)

const (
	swapEnabledKey = "swap_enabled"
	approvedByKey  = "approve_by"
	swapDenomKey   = "swap_denom"
)

func GetRandomSwapDenom(r *rand.Rand) string {
	return simulation.RandStringOfLength(r, r.Intn(8))
}

func GetRandomApprovedBy(r *rand.Rand) string {
	return simulation.RandStringOfLength(r, 32)
}

func GetRandomSwapEnabled(r *rand.Rand) bool {
	return r.Intn(2) == 1
}

func GetRandomizedGenesisState(ss *module.SimulationState) {
	var (
		swapEnabled bool
		approvedBy  string
	)

	swapEnabledSim := func(r *rand.Rand) {
		swapEnabled = GetRandomSwapEnabled(r)
	}

	approvedbySim := func(r *rand.Rand) {
		approvedBy = GetRandomApprovedBy(r)
	}
	ss.AppParams.GetOrGenerate(ss.Cdc, swapEnabledKey, &swapEnabled, ss.Rand, swapEnabledSim)
	ss.AppParams.GetOrGenerate(ss.Cdc, approvedByKey, &approvedBy, ss.Rand, approvedbySim)

	params := hubtypes.NewParams(swapEnabled, hubtypes.DefaultSwapDenom, approvedBy)

	var swaps hubtypes.Swaps

	swaps = append(swaps, hubtypes.Swap{
		TxHash:   []byte("b37c4b36298f20ac7fc5de2b4ac1167d7a044402a40f2b5f1e1bb7f0ee5f6346"),
		Receiver: "sent1grdunxx5jxd0ja75wt508sn6v39p70hhw53zs8",
		Amount:   sdk.Coin{Denom: "sent", Amount: sdk.NewInt(1000)},
	})

	swapGenesis := hubtypes.NewGenesisState(swaps, params)
	bz, err := json.MarshalIndent(&swapGenesis.Params, "", "")
	if err != nil {
		panic(err)
	}

	fmt.Printf("selected randomly generated swap parameters: %s\n", bz)
	ss.GenState[hubtypes.ModuleName] = ss.Cdc.MustMarshalJSON(swapGenesis)
}
