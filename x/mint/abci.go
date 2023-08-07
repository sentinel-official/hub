package mint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/mint/keeper"
	"github.com/sentinel-official/hub/x/mint/types"
)

func BeginBlock(ctx sdk.Context, k keeper.Keeper) []abcitypes.ValidatorUpdate {
	k.IterateInflations(ctx, func(_ int, item types.Inflation) bool {
		if item.Timestamp.After(ctx.BlockTime()) {
			return true
		}

		params := k.GetParams(ctx)
		params.InflationMax = item.Max
		params.InflationMin = item.Min
		params.InflationRateChange = item.RateChange
		k.SetParams(ctx, params)

		minter := k.GetMinter(ctx)
		minter.Inflation = item.Min
		k.SetMinter(ctx, minter)

		k.DeleteInflation(ctx, item.Timestamp)
		return false
	})

	return nil
}
