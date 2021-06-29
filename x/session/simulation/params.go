package simulation

import (
	"fmt"
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	sdksimulation "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sentinel-official/hub/x/session/types"
)

func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{

		sdksimulation.NewSimParamChange(types.ModuleName, string(types.KeyInactiveDuration), func(r *rand.Rand) string {
			return fmt.Sprintf("%d", getRandomInactiveDuration(r))
		}),

		sdksimulation.NewSimParamChange(types.ModuleName, string(types.KeyProofVerificationEnabled), func(r *rand.Rand) string {
			return fmt.Sprintf("%v", getRandomProofVerificationEnabled(r))
		}),
	}
}
