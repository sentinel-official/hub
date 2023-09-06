package node

import (
	abcitypes "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v1/x/node/keeper"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abcitypes.ValidatorUpdate {
	return k.EndBlock(ctx)
}
