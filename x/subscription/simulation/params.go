package simulation

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sentinel-official/hub/x/subscription/types"
)

func RandomizedParams() []simulation.ParamChange {
	return []simulation.ParamChange{
		simulation.NewSimParamChange(types.ParamsSubspace, string(types.KeyInactiveDuration),
			func(r *rand.Rand) string {
				duration := time.Duration(simulation.RandIntBetween(r, 5*60, 60*60)) * time.Second
				return fmt.Sprintf(`"%d"`, duration)
			},
		),
	}
}
