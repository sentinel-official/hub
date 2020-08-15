package vpn

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/node"
	"github.com/sentinel-official/hub/x/session"
	"github.com/sentinel-official/hub/x/subscription"
	"github.com/sentinel-official/hub/x/vpn/keeper"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	ctx, write := ctx.CacheContext()
	defer write()

	node.EndBlock(ctx, k.Node)
	subscription.EndBlock(ctx, k.Subscription)
	session.EndBlock(ctx, k.Session)

	return nil
}
