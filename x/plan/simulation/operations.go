package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sentinel-official/hub/x/plan/expected"
	"github.com/sentinel-official/hub/x/plan/keeper"
	"github.com/sentinel-official/hub/x/plan/types"
)

func WeightedOperations(params simulation.AppParams, cdc *codec.Codec, ak expected.AccountKeeper, pk expected.ProviderKeeper, nk expected.NodeKeeper, k keeper.Keeper) simulation.WeightedOperations {
	return simulation.WeightedOperations{
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, fmt.Sprintf("%s:weight_msg_add", types.ModuleName), &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: SimulateMsgAdd(ak, pk),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, fmt.Sprintf("%s:weight_msg_set_status", types.ModuleName), &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: SimulateMsgSetStatus(ak, pk, k),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, fmt.Sprintf("%s:weight_msg_add_node", types.ModuleName), &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: SimulateMsgAddNode(ak, pk, nk, k),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, fmt.Sprintf("%s:weight_msg_remove_node", types.ModuleName), &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: SimulateMsgRemoveNode(ak, pk, nk, k),
		},
	}
}
