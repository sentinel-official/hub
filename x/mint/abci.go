package mint

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/mint/keeper"
	"github.com/sentinel-official/hub/x/mint/types"
)

const (
	fix1Height = 2_642_500
)

func BeginBlock(ctx sdk.Context, k keeper.Keeper) []abcitypes.ValidatorUpdate {
	if ctx.BlockHeight() == fix1Height {
		k.SetInflation(ctx, types.Inflation{
			Max:        sdk.NewDecWithPrec(49, 2),
			Min:        sdk.NewDecWithPrec(43, 2),
			RateChange: sdk.NewDecWithPrec(6, 2),
			Timestamp:  time.Date(2021, 9, 27, 12, 0, 0, 0, time.UTC),
		})
	}

	k.IterateInflations(ctx, func(_ int, inflation types.Inflation) bool {
		if inflation.Timestamp.After(ctx.BlockTime()) {
			return true
		}

		params := k.GetParams(ctx)
		params.InflationMax = inflation.Max
		params.InflationMin = inflation.Min
		params.InflationRateChange = inflation.RateChange
		k.SetParams(ctx, params)

		if ctx.BlockHeight() >= fix1Height {
			minter := k.GetMinter(ctx)
			minter.Inflation = inflation.Min
			k.SetMinter(ctx, minter)
		}

		k.DeleteInflation(ctx, inflation.Timestamp)
		return false
	})

	return nil
}
