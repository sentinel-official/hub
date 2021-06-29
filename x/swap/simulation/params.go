package simulation

import (
	"fmt"
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	sdksimulation "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sentinel-official/hub/x/swap/types"
)

func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		sdksimulation.NewSimParamChange(types.ModuleName, string(types.KeySwapDenom), func(r *rand.Rand) string {
			return GetRandomSwapDenom(r)
		}),
		sdksimulation.NewSimParamChange(types.ModuleName, string(types.KeySwapEnabled), func(r *rand.Rand) string {
			return fmt.Sprintf("%v", GetRandomSwapEnabled(r))
		}),
		sdksimulation.NewSimParamChange(types.ModuleName, string(types.KeyApproveBy), func(r *rand.Rand) string {
			return GetRandomApproveBy(r)
		}),
	}
}
