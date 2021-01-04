package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sentinel-official/hub/x/provider/expected"
	"github.com/sentinel-official/hub/x/provider/keeper"
)

func WeightedOperations(params simulation.AppParams, cdc *codec.Codec, ak expected.AccountKeeper, k keeper.Keeper) simulation.WeightedOperations {
	return simulation.WeightedOperations{
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, "provider:weight_msg_register", &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: SimulateMsgRegister(ak),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, "provider:weight_msg_update", &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: SimulateMsgUpdate(ak, k),
		},
	}
}
