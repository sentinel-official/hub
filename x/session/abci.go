package session

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	log := k.Logger(ctx)

	end := ctx.BlockTime().Add(-1 * k.InactiveDuration(ctx))
	k.IterateActiveSessions(ctx, end, func(_ int, item types.Session) bool {
		log.Info("Inactive session", "id", item.ID,
			"subscription", item.Subscription, "node", item.Node, "address", item.Address)

		s, _ := k.GetSubscription(ctx, item.Subscription)
		if s.Plan == 0 {
			var (
				amount    = s.Amount(item.Bandwidth)
				quota, _  = k.GetQuota(ctx, item.Subscription, item.Address)
				bandwidth = item.Bandwidth.CeilTo(hub.Gigabyte.Quo(s.Price.Amount))
			)

			log.Info("", "price", s.Price, "deposit", s.Deposit,
				"consumed", item.Bandwidth, "ceil", bandwidth, "amount", amount)

			if err := k.SendCoinsFromDepositToAccount(ctx, item.Address, item.Node.Bytes(), amount); err != nil {
				panic(err)
			}

			quota.Consumed = quota.Consumed.Sub(item.Bandwidth).Add(bandwidth)
			k.SetQuota(ctx, item.Subscription, quota)
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
