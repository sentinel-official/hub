package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sentinel-official/hub/x/node/expected"
	"github.com/sentinel-official/hub/x/node/keeper"
)

func WeightedOperations(params simulation.AppParams, cdc *codec.Codec, ak expected.AccountKeeper, pk expected.ProviderKeeper, k keeper.Keeper) simulation.WeightedOperations {
	return simulation.WeightedOperations{
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, "node:weight_msg_register", &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: SimulateMsgRegister(ak, pk),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, "node:weight_msg_update", &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: SimulateMsgUpdate(ak, pk, k),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, "node:weight_msg_set_status", &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: SimulateMsgSetStatus(ak, k),
		},
	}
}
