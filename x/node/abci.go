package node

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	end := ctx.BlockTime().Add(-1 * k.InactiveDuration(ctx))
	k.IterateActiveNodes(ctx, end, func(_ int, node types.Node) bool {
		k.DeleteActiveNodeAt(ctx, node.StatusAt, node.Address)

		node.Status = hub.StatusInactive
		node.StatusAt = ctx.BlockTime()
		k.SetNode(ctx, node)

		return false
	})

	return nil
}
