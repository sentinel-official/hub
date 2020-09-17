package node

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	log := k.Logger(ctx)

	end := ctx.BlockTime().Add(-1 * k.InactiveDuration(ctx))
	k.IterateActiveNodes(ctx, end, func(_ int, item types.Node) bool {
		log.Info("Inactive node", "address", item.Address, "provider", item.Provider)

		k.DeleteActiveNodeAt(ctx, item.StatusAt, item.Address)

		item.Status = hub.StatusInactive
		item.StatusAt = ctx.BlockTime()
		k.SetNode(ctx, item)

		return false
	})

	return nil
}
