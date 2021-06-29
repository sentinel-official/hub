package simulation

import (
	"fmt"
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	sdksimulation "github.com/cosmos/cosmos-sdk/types/simulation"
	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func RandomizedGenState(simState *module.SimulationState) *types.GenesisState {

	var inactiveDuration time.Duration

	inactiveDurationSim := func(r *rand.Rand) {
		inactiveDuration = getRandomDuration(r)
	}

	simState.AppParams.GetOrGenerate(simState.Cdc, string(types.KeyInactiveDuration), inactiveDuration, nil, inactiveDurationSim)

	params := types.NewParams(inactiveDuration)

	subscriptions := getRandomSubscriptions(simState.Rand)

	state := types.NewGenesisState(subscriptions, params)
	bz := simState.Cdc.MustMarshalJSON(&state.Params)

	fmt.Printf("selected randomly generated subscription parameters: %s\n", bz)
	return state
}

func getRandomSubscriptions(r *rand.Rand) types.GenesisSubscriptions {
	subscriptions := make([]types.GenesisSubscription, r.Intn(28)+4)

	for i := range subscriptions {
		subscriptions = append(subscriptions, types.GenesisSubscription{
			Subscription: types.Subscription{
				Id:       uint64(i + 1),
				Owner:    sdk.AccAddress(sdksimulation.RandStringOfLength(r, 10)).String(),
				Node:     hubtypes.NodeAddress(sdksimulation.RandStringOfLength(r, 10)).String(),
				Price:    getRandomPrice(r),
				Deposit:  getRandomPrice(r),
				Plan:     0,
				Denom:    "tsent",
				Expiry:   sdksimulation.RandTimestamp(r),
				Free:     getRandomFreeBytes(r),
				Status:   hubtypes.Status(r.Int31n(int32(4))),
				StatusAt: sdksimulation.RandTimestamp(r),
			},
			Quotas: getRandomQuotas(r),
		})
	}

	return subscriptions
}

func getRandomDuration(r *rand.Rand) time.Duration {
	return time.Duration(sdksimulation.RandIntBetween(r, 60, 60<<13))
}

func getRandomPrice(r *rand.Rand) sdk.Coin {
	return sdk.Coin{Denom:  "tsent", Amount: sdk.NewInt(r.Int63n(8<<12))}
}

func getRandomFreeBytes(r *rand.Rand) sdk.Int {
	return sdk.NewInt(r.Int63n(1024<<20))
}

func getRandomQuotas(r *rand.Rand) types.Quotas {
	quotas := make(types.Quotas, r.Intn(10)+2)

	allocated := getRandomFreeBytes(r)
	for range quotas {
		quotas = append(quotas, types.Quota{
			Address:   sdk.AccAddress("address").String(),
			Allocated: allocated,
			Consumed:  allocated.Quo(sdk.NewInt(10)),
		})
	}

	return quotas
}
