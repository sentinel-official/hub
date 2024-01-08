package vpn

import (
	abcitypes "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/node"
	"github.com/sentinel-official/hub/v12/x/session"
	"github.com/sentinel-official/hub/v12/x/subscription"
	"github.com/sentinel-official/hub/v12/x/vpn/keeper"
)

func cacheContext(c sdk.Context) (cc sdk.Context, writeCache func()) {
	cms := c.MultiStore().CacheMultiStore()
	cc = c.WithMultiStore(cms)
	return cc, cms.Write
}

func BeginBlock(ctx sdk.Context, k keeper.Keeper) {
	ctx, write := cacheContext(ctx)
	defer write()

	subscription.BeginBlock(ctx, k.Subscription)
}

func EndBlock(ctx sdk.Context, k keeper.Keeper) abcitypes.ValidatorUpdates {
	ctx, write := cacheContext(ctx)
	defer write()

	node.EndBlock(ctx, k.Node)
	session.EndBlock(ctx, k.Session)
	subscription.EndBlock(ctx, k.Subscription)

	return nil
}
