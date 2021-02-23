package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sentinel-official/hub/x/subscription/expected"
	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func WeightedOperations(params simulation.AppParams, cdc *codec.Codec, ak expected.AccountKeeper, pk expected.PlanKeeper, nk expected.NodeKeeper, k keeper.Keeper) simulation.WeightedOperations {
	return simulation.WeightedOperations{
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, fmt.Sprintf("%s:weight_msg_subscribe_to_node", types.ModuleName), &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: SimulateMsgSubscribeToNode(ak, nk),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, fmt.Sprintf("%s:weight_msg_subscribe_to_plan", types.ModuleName), &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: SimulateMsgSubscribeToPlan(ak, pk),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, fmt.Sprintf("%s:weight_msg_cancel", types.ModuleName), &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: SimulateMsgCancel(ak, k),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, fmt.Sprintf("%s:weight_msg_add_quota", types.ModuleName), &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: SimulateMsgAddQuota(ak, k),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, fmt.Sprintf("%s:weight_msg_update_quota", types.ModuleName), &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: SimulateMsgUpdateQuota(ak, k),
		},
	}
}
