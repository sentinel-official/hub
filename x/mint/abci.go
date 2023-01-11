package mint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/mint/keeper"
	"github.com/sentinel-official/hub/x/mint/types"
)

func BeginBlock(ctx sdk.Context, k keeper.Keeper) []abcitypes.ValidatorUpdate {
	k.IterateInflations(ctx, func(_ int, inflation types.Inflation) bool {
		if inflation.Timestamp.After(ctx.BlockTime()) {
			return true
		}

		params := k.GetParams(ctx)
		params.InflationMax = inflation.Max
		params.InflationMin = inflation.Min
		params.InflationRateChange = inflation.RateChange
		k.SetParams(ctx, params)

		minter := k.GetMinter(ctx)
		minter.Inflation = inflation.Min
		k.SetMinter(ctx, minter)

		k.DeleteInflation(ctx, inflation.Timestamp)
		return false
	})

	return nil
}
