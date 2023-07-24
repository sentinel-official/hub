package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	hubtypes "github.com/sentinel-official/hub/types"
	hubutils "github.com/sentinel-official/hub/utils"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func (k *Keeper) BeginBlock(ctx sdk.Context) {
	k.IteratePayoutsForNextAt(ctx, ctx.BlockTime(), func(_ int, item types.Payout) (stop bool) {
		k.DeletePayoutForNextAt(ctx, item.NextAt, item.ID)

		var (
			accAddr       = item.GetAddress()
			nodeAddr      = item.GetNodeAddress()
			stakingShare  = k.node.StakingShare(ctx)
			stakingReward = hubutils.GetProportionOfCoin(item.Price, stakingShare)
		)

		if err := k.SendCoinFromDepositToModule(ctx, accAddr, k.feeCollectorName, stakingReward); err != nil {
			panic(err)
		}

		payment := item.Price.Sub(stakingReward)
		if err := k.SendCoinFromDepositToAccount(ctx, accAddr, nodeAddr.Bytes(), payment); err != nil {
			panic(err)
		}

		item.Hours = item.Hours - 1
		if item.Hours > 0 {
			item.NextAt = item.NextAt.Add(time.Hour)
			k.SetPayoutForNextAt(ctx, item.NextAt, item.ID)
		}

		k.SetPayout(ctx, item)
		ctx.EventManager().EmitTypedEvent(
			&types.EventPayout{
				ID:          item.ID,
				Address:     item.Address,
				NodeAddress: item.NodeAddress,
			},
		)

		return false
	})
}

func (k *Keeper) EndBlock(ctx sdk.Context) []abcitypes.ValidatorUpdate {
	expiryDuration := k.ExpiryDuration(ctx)
	k.IterateSubscriptionsForExpiryAt(ctx, ctx.BlockTime(), func(_ int, item types.Subscription) bool {
		k.DeleteSubscriptionForExpiryAt(ctx, item.GetExpiryAt(), item.GetID())

		if item.GetStatus().Equal(hubtypes.StatusActive) {
			item.SetExpiryAt(
				item.GetExpiryAt().Add(expiryDuration),
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

			k.DeletePayoutForNextAt(ctx, payout.NextAt, payout.ID)

			payout.NextAt = time.Time{}
			k.SetPayout(ctx, payout)

			return false
		}

		switch s := item.(type) {
		case *types.NodeSubscription:
			if s.Gigabytes != 0 {
				var (
					accAddr       = item.GetAddress()
					gigabytePrice = sdk.NewCoin(
						s.Deposit.Denom,
						s.Deposit.Amount.QuoRaw(s.Gigabytes),
					)
				)

				alloc, found := k.GetAllocation(ctx, item.GetID(), accAddr)
				if !found {
					panic(fmt.Errorf("subscription allocation %d/%s does not exist", item.GetID(), accAddr))
				}

				var (
					paidAmount = hubtypes.AmountForBytes(gigabytePrice.Amount, alloc.UtilisedBytes)
					refund     = sdk.NewCoin(
						s.Deposit.Denom,
						s.Deposit.Amount.Sub(paidAmount),
					)
				)

				if err := k.DepositSubtract(ctx, accAddr, refund); err != nil {
					panic(err)
				}
			}
		}

		k.IterateAllocations(ctx, item.GetID(), func(_ int, alloc types.Allocation) bool {
			addr := alloc.GetAddress()
			k.DeleteAllocation(ctx, item.GetID(), addr)
			k.DeleteSubscriptionForAccount(ctx, addr, item.GetID())

			return false
		})

		switch s := item.(type) {
		case *types.NodeSubscription:
			k.DeleteSubscriptionForNode(ctx, s.GetNodeAddress(), s.GetID())
		case *types.PlanSubscription:
			k.DeleteSubscriptionForPlan(ctx, s.PlanID, s.GetID())
		default:
			panic(fmt.Errorf("invalid subscription %d with type %T", item.GetID(), item))
		}

		k.DeleteSubscription(ctx, item.GetID())
		ctx.EventManager().EmitTypedEvent(
			&types.EventUpdateStatus{
				ID:     item.GetID(),
				Status: hubtypes.StatusInactive,
			},
		)

		payout, found := k.GetPayout(ctx, item.GetID())
		if !found {
			return false
		}

		k.DeletePayout(ctx, payout.ID)
		k.DeletePayoutForAccount(ctx, payout.GetAddress(), payout.ID)
		k.DeletePayoutForNode(ctx, payout.GetNodeAddress(), payout.ID)

		return false
	})

	return nil
}
