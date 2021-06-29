package simulation

import (
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/plan/types"
)

func getRandomPlan(r *rand.Rand, plans types.Plans) types.Plan {
	if len(plans) == 0 {
		return types.Plan{
			Provider: hubtypes.ProvAddress([]byte("provider-1")).String(),
		}
	}

	index := r.Intn(len(plans)-1)
	return plans[index]
}

func getRandomNodes(r *rand.Rand) []string {

	nodes := make([]string, r.Intn(28)+4)

	for range nodes {
		bz := make([]byte, 20)
		if _, err := r.Read(bz); err != nil {
			panic(err)
		}
		nodeAddress := hubtypes.NodeAddress(bz)
		nodes = append(nodes, nodeAddress.String())
	}

	return nodes
}

func getRandomPlans(r *rand.Rand) types.Plans {

	plans := make(types.Plans, r.Intn(28)+4)

	for range plans {
		bz := make([]byte, 20)
		if _, err := r.Read(bz); err != nil {
			panic(err)
		}

		providerAddress := hubtypes.ProvAddress(bz)

		plans = append(plans, types.Plan{
			Id:       r.Uint64(),
			Provider: providerAddress.String(),
			Price: sdk.NewCoins(sdk.Coin{
				Denom:  simulation.RandStringOfLength(r, r.Intn(125)+3),
				Amount: sdk.NewInt(r.Int63n(8 << 12)),
			}),
			Validity: time.Duration(simulation.RandIntBetween(r, weekDurationInSeconds, monthDurationInSeconds)),
			Bytes:    sdk.NewInt(int64(simulation.RandIntBetween(r, gigabytes, terabytes))),
			Status:   hubtypes.Status(r.Intn(3)),
			StatusAt: simulation.RandTimestamp(r),
		})
	}

	return plans
}
