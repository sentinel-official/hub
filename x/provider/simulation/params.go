package simulation

import (
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	sdksimulation "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sentinel-official/hub/x/provider/types"
)

func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		sdksimulation.NewSimParamChange(types.ModuleName, string(types.KeyDeposit), func(r *rand.Rand) string {
			return getRandomDeposit(r).String()
		}),
	}
}
