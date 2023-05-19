// DO NOT COVER

package simulation

import (
	"fmt"
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sentinel-official/hub/x/node/types"
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
			string(types.KeyExpiryDuration),
			func(r *rand.Rand) string {
				return fmt.Sprintf(
					"%s",
					time.Duration(r.Int63n(MaxInt))*time.Millisecond,
				)
			},
		),
		simulation.NewSimParamChange(
			types.ModuleName,
			string(types.KeyMaxGigabytePrices),
			func(r *rand.Rand) string {
				return sdk.NewCoins(
					sdk.NewInt64Coin(
						sdk.DefaultBondDenom,
						r.Int63n(MaxInt),
					),
				).String()
			},
		),
		simulation.NewSimParamChange(
			types.ModuleName,
			string(types.KeyMinGigabytePrices),
			func(r *rand.Rand) string {
				return sdk.NewCoins(
					sdk.NewInt64Coin(
						sdk.DefaultBondDenom,
						r.Int63n(MaxInt),
					),
				).String()
			},
		),
		simulation.NewSimParamChange(
			types.ModuleName,
			string(types.KeyMaxHourlyPrices),
			func(r *rand.Rand) string {
				return sdk.NewCoins(
					sdk.NewInt64Coin(
						sdk.DefaultBondDenom,
						r.Int63n(MaxInt),
					),
				).String()
			},
		),
		simulation.NewSimParamChange(
			types.ModuleName,
			string(types.KeyMinHourlyPrices),
			func(r *rand.Rand) string {
				return sdk.NewCoins(
					sdk.NewInt64Coin(
						sdk.DefaultBondDenom,
						r.Int63n(MaxInt),
					),
				).String()
			},
		),
		simulation.NewSimParamChange(
			types.ModuleName,
			string(types.KeyRevenueShare),
			func(r *rand.Rand) string {
				return sdk.NewDecWithPrec(
					MaxInt,
					6,
				).String()
			},
		),
	}
}
