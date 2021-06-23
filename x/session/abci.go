package session

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abcitypes.ValidatorUpdate {
	var (
		log              = k.Logger(ctx)
		inactiveDuration = k.InactiveDuration(ctx)
	)

	k.IterateInactiveSessionsAt(ctx, ctx.BlockTime(), func(_ int, key []byte, item types.Session) bool {
		log.Info("inactive session", "key", key, "value", item)

		itemAddress := item.GetAddress()
		if item.Status.Equal(hubtypes.Active) {
			k.DeleteActiveSessionForAddress(ctx, itemAddress, item.Id)
			k.DeleteInactiveSessionAt(ctx, item.StatusAt.Add(inactiveDuration), item.Id)

			item.Status = hubtypes.StatusInactivePending
			item.StatusAt = ctx.BlockTime()

			k.SetSession(ctx, item)
			k.SetInactiveSessionForAddress(ctx, itemAddress, item.Id)
			k.SetInactiveSessionAt(ctx, item.StatusAt.Add(inactiveDuration), item.Id)
			ctx.EventManager().EmitTypedEvent(
				&types.EventEndSession{
					Id:           item.Id,
					Subscription: item.Subscription,
					Node:         item.Node,
				},
			)

			return false
		}

		if err := k.ProcessPaymentAndUpdateQuota(ctx, item); err != nil {
			log.Error("failed to process the payment", "cause", err)
		}

		k.DeleteInactiveSessionAt(ctx, item.StatusAt.Add(inactiveDuration), item.Id)

		item.Status = hubtypes.StatusInactive
		item.StatusAt = ctx.BlockTime()
		k.SetSession(ctx, item)

		return false
	})

	return nil
}
