// DO NOT COVER

package simulation

import (
	"fmt"
	"math/rand"
	"time"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sentinel-official/hub/x/session/types"
)

const (
	MaxStatusChangeDelay = 1 << 18
)

func ParamChanges(_ *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(
			types.ModuleName,
			string(types.KeyProofVerificationEnabled),
			func(r *rand.Rand) string {
				return fmt.Sprintf("%v", false)
			},
		),
		simulation.NewSimParamChange(
			types.ModuleName,
			string(types.KeyStatusChangeDelay),
			func(r *rand.Rand) string {
				return fmt.Sprintf(
					"%s",
					time.Duration(r.Int63n(MaxStatusChangeDelay))*time.Millisecond,
				)
			},
		),
	}
}
