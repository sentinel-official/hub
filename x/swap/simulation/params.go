// DO NOT COVER

package simulation

import (
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func ParamChanges(_ *rand.Rand) []simtypes.ParamChange {
	return nil
}
