// DO NOT COVER

package simulation

import (
	"fmt"
	"math/rand"
	"time"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sentinel-official/hub/x/subscription/types"
)

const (
	MaxInactivePendingDuration = 1 << 18
)

func ParamChanges(_ *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(
			types.ModuleName,
			string(types.KeyInactivePendingDuration),
			func(r *rand.Rand) string {
				return fmt.Sprintf(
					"%s",
					time.Duration(r.Int63n(MaxInactivePendingDuration))*time.Millisecond,
				)
			},
		),
	}
}
