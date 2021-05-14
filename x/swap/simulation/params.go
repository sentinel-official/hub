package simulation

import (
	"fmt"
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	sim "github.com/sentinel-official/hub/x/simulation"
	hubtypes "github.com/sentinel-official/hub/x/swap/types"
)

func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		sim.NewSimParamChange(hubtypes.ModuleName, string(hubtypes.KeySwapDenom), func(r *rand.Rand) string {
			return fmt.Sprintf("%s", GetRandomSwapDenom(r))
		}),
		sim.NewSimParamChange(hubtypes.ModuleName, string(hubtypes.KeySwapEnabled), func(r *rand.Rand) string {
			return fmt.Sprintf("%v", GetRandomSwapEnabled(r))
		}),
		sim.NewSimParamChange(hubtypes.ModuleName, string(hubtypes.KeyApproveBy), func(r *rand.Rand) string {
			return fmt.Sprintf("%s", GetRandomApprovedBy(r))
		}),
	}
}
