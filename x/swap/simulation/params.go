// DO NOT COVER

package simulation

import (
	"math/rand"

	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func ParamChanges(_ *rand.Rand) []simulationtypes.ParamChange {
	return nil
}
