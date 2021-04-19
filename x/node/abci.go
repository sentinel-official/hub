package node

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	var (
		log = k.Logger(ctx)
		end = ctx.BlockTime().Add(-1 * k.InactiveDuration(ctx))
	)

	k.IterateInactiveNodesAt(ctx, end, func(_ int, item types.Node) bool {
		log.Info("Inactive node", "address", item.Address, "provider", item.Provider)

		var (
			itemAddress = item.GetAddress()
		)

		k.DeleteActiveNode(ctx, itemAddress)
		k.SetInactiveNode(ctx, itemAddress)

		if item.Provider != "" {
			var (
				itemProvider = item.GetProvider()
			)

			k.DeleteActiveNodeForProvider(ctx, itemProvider, itemAddress)
			k.SetInactiveNodeForProvider(ctx, itemProvider, itemAddress)
		}

		k.DeleteInactiveNodeAt(ctx, item.StatusAt, itemAddress)

		item.Status = hubtypes.StatusInactive
		item.StatusAt = ctx.BlockTime()
		k.SetNode(ctx, item)

		return false
	})

	return nil
}
