package vpn

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/node"
	"github.com/sentinel-official/hub/x/session"
	"github.com/sentinel-official/hub/x/subscription"
	"github.com/sentinel-official/hub/x/vpn/keeper"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) abcitypes.ValidatorUpdates {
	ctx, write := ctx.CacheContext()
	defer write()

	node.EndBlock(ctx, k.Node)
	session.EndBlock(ctx, k.Session)
	subscription.EndBlock(ctx, k.Subscription)

	return nil
}
