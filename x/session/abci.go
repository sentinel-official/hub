package session

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abcitypes.ValidatorUpdate {
	end := ctx.BlockTime().Add(-1 * k.InactiveDuration(ctx))
	k.IterateActiveSessionsAt(ctx, end, func(_ int, item types.Session) bool {
		k.Logger(ctx).Info("Inactive session", "id", item.Id,
			"subscription", item.Subscription, "node", item.Node, "address", item.Address)

		itemAddress := item.GetAddress()
		if err := k.ProcessPaymentAndUpdateQuota(ctx, item); err != nil {
			panic(err)
		}

		k.DeleteActiveSessionForAddress(ctx, itemAddress, item.Id)
		k.DeleteActiveSessionAt(ctx, item.StatusAt, item.Id)

		item.Status = hubtypes.StatusInactive
		item.StatusAt = ctx.BlockTime()

		k.SetSession(ctx, item)
		k.SetInactiveSessionForAddress(ctx, itemAddress, item.Id)

		return false
	})

	return nil
}
