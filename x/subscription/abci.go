package subscription

import (
	abcitypes "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/subscription/keeper"
)

func BeginBlock(ctx sdk.Context, k keeper.Keeper) {
	k.BeginBlock(ctx)
}

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abcitypes.ValidatorUpdate {
	return k.EndBlock(ctx)
}
