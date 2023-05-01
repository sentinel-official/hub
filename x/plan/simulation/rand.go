package simulation

import (
	"math"
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	simulationhubtypes "github.com/sentinel-official/hub/types/simulation"
	"github.com/sentinel-official/hub/x/plan/types"
)

const (
	MaxPlanId          = 1 << 18
	MaxPlans           = 1 << 10
	MaxPlanPriceAmount = 1 << 10
	MaxPlanValidity    = 1 << 18
	MaxPlanBytes       = math.MaxInt64
)

func RandomPlans(r *rand.Rand) types.Plans {
	var (
		items      = make(types.Plans, 0, r.Intn(MaxPlans))
		duplicates = make(map[uint64]bool)
	)

	for len(items) < cap(items) {
		id := uint64(r.Int63n(MaxPlanId))
		if duplicates[id] {
			continue
		}

		var (
			prices = simulationhubtypes.RandomCoins(
				r,
				sdk.NewCoins(
					sdk.NewInt64Coin(
						sdk.DefaultBondDenom,
						MaxPlanPriceAmount,
					),
				),
			)
			validity = time.Duration(r.Int63n(MaxPlanValidity)) * time.Minute
			bytes    = sdk.NewInt(r.Int63n(MaxPlanBytes))
			status   = hubtypes.StatusActive
			statusAt = time.Now()
		)

		if rand.Intn(2) == 0 {
			status = hubtypes.StatusInactive
		}

		duplicates[id] = true
		items = append(
			items,
			types.Plan{
				ID:       id,
				Address:  "",
				Prices:   prices,
				Validity: validity,
				Bytes:    bytes,
				Status:   status,
				StatusAt: statusAt,
			},
		)
	}

	return items
}

func RandomGenesisPlans(r *rand.Rand) types.GenesisPlans {
	var (
		rItems = RandomPlans(r)
		items  = make(types.GenesisPlans, 0, len(rItems))
	)

	for _, item := range rItems {
		items = append(
			items,
			types.GenesisPlan{
				Plan:  item,
				Nodes: nil,
			},
		)
	}

	return items
}
