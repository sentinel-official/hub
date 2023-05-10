// DO NOT COVER

package simulation

import (
	"fmt"
	"math/rand"
	"time"

	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sentinel-official/hub/x/subscription/types"
)

const (
	MaxInactiveDuration = 1 << 18
)

func ParamChanges(_ *rand.Rand) []simulationtypes.ParamChange {
	return []simulationtypes.ParamChange{
		simulation.NewSimParamChange(
			types.ModuleName,
			string(types.KeyInactiveDuration),
			func(r *rand.Rand) string {
				return fmt.Sprintf(
					"%s",
					time.Duration(r.Int63n(MaxInactiveDuration))*time.Millisecond,
				)
			},
		),
	}
}
