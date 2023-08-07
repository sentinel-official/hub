package subscription

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/subscription/keeper"
)

func BeginBlock(ctx sdk.Context, k keeper.Keeper) {
	k.BeginBlock(ctx)
}

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abcitypes.ValidatorUpdate {
	return k.EndBlock(ctx)
}
