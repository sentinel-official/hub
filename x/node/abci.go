package node

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abcitypes.ValidatorUpdate {
	var (
		log              = k.Logger(ctx)
		inactiveDuration = k.InactiveDuration(ctx)
	)

	k.IterateInactiveNodesAt(ctx, ctx.BlockTime(), func(_ int, key []byte, item types.Node) bool {
		log.Info("inactive node", "key", key, "value", item)

		itemAddress := item.GetAddress()
		k.DeleteActiveNode(ctx, itemAddress)
		k.SetInactiveNode(ctx, itemAddress)

		if item.Provider != "" {
			itemProvider := item.GetProvider()
			k.DeleteActiveNodeForProvider(ctx, itemProvider, itemAddress)
			k.SetInactiveNodeForProvider(ctx, itemProvider, itemAddress)
		}

		k.DeleteInactiveNodeAt(ctx, item.StatusAt.Add(inactiveDuration), itemAddress)

		item.Status = hubtypes.StatusInactive
		item.StatusAt = ctx.BlockTime()
		k.SetNode(ctx, item)

		ctx.EventManager().EmitTypedEvent(&types.EventSetNodeStatus{
			From:    "",
			Address: item.Address,
			Status:  item.Status,
		})

		return false
	})

	return nil
}
