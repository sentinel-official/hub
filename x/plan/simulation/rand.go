// DO NOT COVER

package simulation

import (
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	simulationhubtypes "github.com/sentinel-official/hub/types/simulation"
	"github.com/sentinel-official/hub/x/plan/types"
)

const (
	MaxCount      = 1 << 10
	MaxID         = 1 << 10
	MaxBytes      = 1 << 10
	MaxDuration   = 1 << 10
	MaxCoinAmount = 1 << 10
)

func RandomPlans(r *rand.Rand) types.Plans {
	var (
		items = make(types.Plans, 0, r.Intn(MaxCount))
		m     = make(map[uint64]bool)
	)

	for len(items) < cap(items) {
		id := uint64(r.Int63n(MaxID))
		if m[id] {
			continue
		}

		var (
			bytes    = sdk.NewInt(r.Int63n(MaxBytes))
			duration = time.Duration(r.Int63n(MaxDuration)) * time.Minute
			prices   = simulationhubtypes.RandomCoins(
				r,
				sdk.NewCoins(
					sdk.NewInt64Coin(
						sdk.DefaultBondDenom,
						MaxCoinAmount,
					),
				),
			)
			status   = hubtypes.StatusActive
			statusAt = time.Now()
		)

		if rand.Intn(2) == 0 {
			status = hubtypes.StatusInactive
		}

		m[id] = true
		items = append(
			items,
			types.Plan{
				ID:       id,
				Address:  "",
				Bytes:    bytes,
				Duration: duration,
				Prices:   prices,
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
