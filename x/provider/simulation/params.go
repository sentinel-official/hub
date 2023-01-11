package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sentinel-official/hub/x/provider/types"
)

const (
	MaxInt = 1 << 18
)

func ParamChanges(_ *rand.Rand) []simulationtypes.ParamChange {
	return []simulationtypes.ParamChange{
		simulation.NewSimParamChange(
			types.ModuleName,
			string(types.KeyDeposit),
			func(r *rand.Rand) string {
				return sdk.NewInt64Coin(
					sdk.DefaultBondDenom,
					r.Int63n(MaxInt),
				).String()
			},
		),
		simulation.NewSimParamChange(
			types.ModuleName,
			string(types.KeyStakingShare),
			func(r *rand.Rand) string {
				return sdk.NewDecWithPrec(
					r.Int63n(MaxInt),
					6,
				).String()
			},
		),
	}
}
