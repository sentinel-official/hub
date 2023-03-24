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

	k.IterateInactiveSessionsAt(ctx, ctx.BlockTime(), func(_ int, item types.Session) bool {
		log.Info("found an inactive session", "id", item.Id)

		accAddr := item.GetAddress()
		if item.Status.Equal(hubtypes.StatusActive) {
			k.DeleteActiveSessionForAddress(ctx, accAddr, item.Id)
			k.DeleteInactiveSessionAt(ctx, item.StatusAt.Add(inactiveDuration), item.Id)

			item.Status = hubtypes.StatusInactivePending
			item.StatusAt = ctx.BlockTime()

			k.SetSession(ctx, item)
			k.SetInactiveSessionForAddress(ctx, accAddr, item.Id)
			k.SetInactiveSessionAt(ctx, item.StatusAt.Add(inactiveDuration), item.Id)
			ctx.EventManager().EmitTypedEvent(
				&types.EventSetStatus{
					Id:           item.Id,
					Node:         item.Node,
					Subscription: item.Subscription,
					Status:       item.Status,
				},
			)

			return false
		}

		if err := k.ProcessPaymentAndUpdateQuota(ctx, item); err != nil {
			log.Error("error occurred while processing the payment", "cause", err)
		}

		k.DeleteSession(ctx, item.Id)
		k.DeleteInactiveSessionForAddress(ctx, accAddr, item.Id)
		k.DeleteInactiveSessionAt(ctx, item.StatusAt.Add(inactiveDuration), item.Id)

		ctx.EventManager().EmitTypedEvent(
			&types.EventSetStatus{
				Id:           item.Id,
				Node:         item.Node,
				Subscription: item.Subscription,
				Status:       item.Status,
			},
		)

		return false
	})

	return nil
}
