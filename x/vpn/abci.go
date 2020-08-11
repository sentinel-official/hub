package vpn

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/node"
	"github.com/sentinel-official/hub/x/subscription"
	"github.com/sentinel-official/hub/x/vpn/keeper"
)

func BeginBlock(ctx sdk.Context, k keeper.Keeper) {
	ctx, write := ctx.CacheContext()
	defer write()

	node.BeginBlock(ctx, k.Node)
}

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	ctx, write := ctx.CacheContext()
	defer write()

	subscription.EndBlock(ctx, k.Subscription)
	return nil
}
