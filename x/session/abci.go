package session

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	end := ctx.BlockTime().Add(-1 * k.InactiveDuration(ctx))
	k.IterateActiveSessions(ctx, end, func(_ int, item types.Session) bool {
		k.Logger(ctx).Info("Inactive session", "id", item.ID,
			"subscription", item.Subscription, "node", item.Node, "address", item.Address)

		if err := k.Pay(ctx, item); err != nil {
			panic(err)
		}

		k.DeleteOngoingSession(ctx, item.Subscription, item.Address)
		k.DeleteActiveSessionAt(ctx, item.StatusAt, item.ID)

		item.Status = hub.StatusInactive
		item.StatusAt = ctx.BlockTime()
		k.SetSession(ctx, item)

		return false
	})

	return nil
}
