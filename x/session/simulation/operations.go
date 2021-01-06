package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sentinel-official/hub/x/session/expected"
)

func WeightedOperations(params simulation.AppParams, cdc *codec.Codec, ak expected.AccountKeeper, nk expected.NodeKeeper, sk expected.SubscriptionKeeper) simulation.WeightedOperations {
	return simulation.WeightedOperations{
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, "session:weight_msg_upsert", &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: SimulateUpsert(ak, nk, sk),
		},
	}
}
