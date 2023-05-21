// DO NOT COVER

package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sentinel-official/hub/x/provider/types"
)

const (
	MaxDeposit      = 1 << 8
	MaxRevenueShare = 1 << 8
)

func ParamChanges(_ *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(
			types.ModuleName,
			string(types.KeyDeposit),
			func(r *rand.Rand) string {
				return sdk.NewInt64Coin(
					sdk.DefaultBondDenom,
					r.Int63n(MaxDeposit),
				).String()
			},
		),
		simulation.NewSimParamChange(
			types.ModuleName,
			string(types.KeyRevenueShare),
			func(r *rand.Rand) string {
				return sdk.NewDecWithPrec(
					r.Int63n(MaxRevenueShare),
					6,
				).String()
			},
		),
	}
}
