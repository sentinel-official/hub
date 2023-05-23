package subscription

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abcitypes.ValidatorUpdate {
	inactiveDuration := k.InactiveDuration(ctx)
	k.IterateSubscriptionsForExpiryAt(ctx, ctx.BlockTime(), func(_ int, item types.Subscription) bool {
		k.DeleteSubscriptionForExpiryAt(ctx, item.GetExpiryAt(), item.GetID())

		if item.GetStatus().Equal(hubtypes.StatusActive) {
			item.SetExpiryAt(
				item.GetExpiryAt().Add(inactiveDuration),
			)
			item.SetStatus(hubtypes.StatusInactivePending)
			item.SetStatusAt(ctx.BlockTime())

			k.SetSubscription(ctx, item)
			k.SetSubscriptionForExpiryAt(ctx, item.GetExpiryAt(), item.GetID())
			ctx.EventManager().EmitTypedEvent(
				&types.EventUpdateStatus{
					ID:     item.GetID(),
					Status: item.GetStatus(),
				},
			)

			payout, found := k.GetPayout(ctx, item.GetID())
			if !found {
				return false
			}

			k.DeletePayoutForTimestamp(ctx, payout.Timestamp, payout.ID)

			payout.Timestamp = time.Time{}
			k.SetPayout(ctx, payout)

			return false
		}

		// TODO: refund amount to account address

		k.DeleteSubscription(ctx, item.GetID())

		switch s := item.(type) {
		case *types.NodeSubscription:
			k.DeleteSubscriptionForNode(ctx, s.GetNodeAddress(), s.GetID())
		case *types.PlanSubscription:
			k.DeleteSubscriptionForPlan(ctx, s.PlanID, s.GetID())
		default:
			panic(fmt.Errorf("invalid type %T for the subscription %d", item, item.GetID()))
		}

		k.IterateAllocations(ctx, item.GetID(), func(_ int, alloc types.Allocation) bool {
			addr := alloc.GetAddress()
			k.DeleteAllocation(ctx, item.GetID(), addr)
			k.DeleteSubscriptionForAccount(ctx, addr, item.GetID())

			return false
		})

		ctx.EventManager().EmitTypedEvent(
			&types.EventUpdateStatus{
				ID:     item.GetID(),
				Status: hubtypes.StatusInactive,
			},
		)

		return false
	})

	k.IteratePayoutsForTimestamp(ctx, ctx.BlockTime(), func(_ int, item types.Payout) (stop bool) {
		k.DeletePayoutForTimestamp(ctx, item.Timestamp, item.ID)

		var (
			accAddr  = item.GetAddress()
			nodeAddr = item.GetNodeAddress()
		)

		if err := k.SendCoinFromDepositToAccount(ctx, accAddr, nodeAddr.Bytes(), item.Price); err != nil {
			panic(err)
		}

		item.Hours = item.Hours - 1
		if item.Hours > 0 {
			item.Timestamp = item.Timestamp.Add(time.Hour)
			k.SetPayoutForTimestamp(ctx, item.Timestamp, item.ID)
		}

		k.SetPayout(ctx, item)
		ctx.EventManager().EmitTypedEvent(
			&types.EventPayout{
				ID:          item.ID,
				FromAddress: item.Address,
				ToAddress:   item.NodeAddress,
			},
		)

		return false
	})

	return nil
}
