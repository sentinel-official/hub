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

// BeginBlock is a function that gets called at the beginning of every block.
// It processes the payouts scheduled to be made and performs the necessary actions accordingly.
func (k *Keeper) BeginBlock(ctx sdk.Context) {
	// Iterate over all payouts that are scheduled to happen at the current block time.
	k.IteratePayoutsForNextAt(ctx, ctx.BlockTime(), func(_ int, item types.Payout) (stop bool) {
		// Delete the payout from the NextAt index before updating the NextAt value.
		k.DeletePayoutForNextAt(ctx, item.NextAt, item.ID)

		// Calculate the staking reward for the payout.
		var (
			accAddr       = item.GetAddress()
			nodeAddr      = item.GetNodeAddress()
			stakingShare  = k.node.StakingShare(ctx)
			stakingReward = hubutils.GetProportionOfCoin(item.Price, stakingShare)
		)

		// Move the staking reward from the deposit to the fee collector module account.
		if err := k.SendCoinFromDepositToModule(ctx, accAddr, k.feeCollectorName, stakingReward); err != nil {
			panic(err)
		}

		// Calculate the payment amount to be sent to the node address by subtracting the staking reward from the payout price.
		payment := item.Price.Sub(stakingReward)

		// Send the payment amount from the deposit to the node address.
		if err := k.SendCoinFromDepositToAccount(ctx, accAddr, nodeAddr.Bytes(), payment); err != nil {
			panic(err)
		}

		// Decrement the remaining payout duration (in hours) by 1 and update the NextAt value.
		item.Hours = item.Hours - 1
		item.NextAt = item.NextAt.Add(time.Hour)

		// If the payout duration has reached 0, set the NextAt value to an empty time.
		if item.Hours == 0 {
			item.NextAt = time.Time{}
		}

		// Update the payout in the store with the updated duration and NextAt value.
		k.SetPayout(ctx, item)

		// If the payout still has remaining duration (hours), update the NextAt index.
		if item.Hours > 0 {
			k.SetPayoutForNextAt(ctx, item.NextAt, item.ID)
		}

		return false
	})
}

// EndBlock is a function that gets called at the end of every block.
// It processes the subscriptions that have become inactive and performs the necessary actions accordingly.
func (k *Keeper) EndBlock(ctx sdk.Context) []abcitypes.ValidatorUpdate {
	// Get the status change delay from the store.
	statusChangeDelay := k.StatusChangeDelay(ctx)

	// Iterate over all subscriptions that have become inactive at the current block time.
	k.IterateSubscriptionsForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item types.Subscription) bool {
		// Delete the subscription from the InactiveAt index before updating the InactiveAt value.
		k.DeleteSubscriptionForInactiveAt(ctx, item.GetInactiveAt(), item.GetID())

		// If the subscription status is 'Active', update its InactiveAt value and set it to 'InactivePending'.
		if item.GetStatus().Equal(hubtypes.StatusActive) {
			// Run the SubscriptionInactivePendingHook to perform custom actions before setting the subscription to inactive pending state.
			if err := k.SubscriptionInactivePendingHook(ctx, item.GetID()); err != nil {
				panic(err)
			}

			item.SetInactiveAt(
				ctx.BlockTime().Add(statusChangeDelay),
			)
			item.SetStatus(hubtypes.StatusInactivePending)
			item.SetStatusAt(ctx.BlockTime())

			// Save the updated subscription to the store and update the InactiveAt index.
			k.SetSubscription(ctx, item)
			k.SetSubscriptionForInactiveAt(ctx, item.GetInactiveAt(), item.GetID())

			// Emit an event to notify that the subscription status has been updated.
			ctx.EventManager().EmitTypedEvent(
				&types.EventUpdateStatus{
					ID:     item.GetID(),
					Status: hubtypes.StatusInactivePending,
				},
			)

			// If the subscription is a NodeSubscription and the duration is specified in hours (non-zero), update the associated payout.
			if s, ok := item.(*types.NodeSubscription); ok && s.Hours != 0 {
				payout, found := k.GetPayout(ctx, s.GetID())
				if !found {
					panic(fmt.Errorf("payout for subscription %d does not exist", s.GetID()))
				}

				var (
					accAddr  = payout.GetAddress()
					nodeAddr = payout.GetNodeAddress()
				)

				// Delete the payout from the Store for the given account and node.
				k.DeletePayoutForAccountByNode(ctx, accAddr, nodeAddr, payout.ID)
				k.DeletePayoutForNextAt(ctx, payout.NextAt, payout.ID)

				// Reset the `NextAt` field of the payout and update it in the Store.
				payout.NextAt = time.Time{}
				k.SetPayout(ctx, payout)
			}

			return false
		}

		// If the subscription status is not 'Active', handle the different types of subscriptions based on their attributes.
		if s, ok := item.(*types.NodeSubscription); ok {
			// Check if it has a non-zero bandwidth (Gigabytes != 0).
			if s.Gigabytes != 0 {
				// Calculate the gigabyte price based on the deposit amount and gigabytes.
				var (
					accAddr       = item.GetAddress()
					gigabytePrice = sdk.NewCoin(
						s.Deposit.Denom,
						s.Deposit.Amount.QuoRaw(s.Gigabytes),
					)
				)

				// Get the allocation associated with the subscription and account.
				alloc, found := k.GetAllocation(ctx, item.GetID(), accAddr)
				if !found {
					panic(fmt.Errorf("subscription allocation %d/%s does not exist", item.GetID(), accAddr))
				}

				// Calculate the amount paid based on the gigabyte price and utilized bandwidth.
				var (
					paidAmount = hubutils.AmountForBytes(gigabytePrice.Amount, alloc.UtilisedBytes)
					refund     = sdk.NewCoin(
						s.Deposit.Denom,
						s.Deposit.Amount.Sub(paidAmount),
					)
				)

				// Refund the difference between the deposit and the amount paid to the node's account.
				if err := k.SubtractDeposit(ctx, accAddr, refund); err != nil {
					panic(err)
				}
			}

			// Check if the number of hours for the subscription is not zero.
			if s.Hours != 0 {
				// Get the payout information associated with the subscription ID.
				payout, found := k.GetPayout(ctx, item.GetID())
				if !found {
					panic(fmt.Errorf("payout for subscription %d does not exist", item.GetID()))
				}

				// Calculate the refund amount by multiplying the payout price with the number of remaining hours.
				var (
					accAddr = payout.GetAddress()
					refund  = sdk.NewCoin(
						payout.Price.Denom,
						payout.Price.Amount.MulRaw(payout.Hours),
					)
				)

				// Subtract the refund amount from the account's deposit balance.
				if err := k.SubtractDeposit(ctx, accAddr, refund); err != nil {
					panic(err)
				}
			}
		}

		// Iterate over all allocations associated with the subscription and delete them from the store.
		k.IterateAllocationsForSubscription(ctx, item.GetID(), func(_ int, alloc types.Allocation) bool {
			accAddr := alloc.GetAddress()
			k.DeleteAllocation(ctx, item.GetID(), accAddr)
			k.DeleteSubscriptionForAccount(ctx, accAddr, item.GetID())

			return false
		})

		// Based on the subscription type, perform additional cleanup actions.
		switch s := item.(type) {
		case *types.NodeSubscription:
			// For node-level subscriptions, delete the subscription from the NodeAddress index.
			k.DeleteSubscriptionForNode(ctx, s.GetNodeAddress(), s.GetID())
		case *types.PlanSubscription:
			// For plan-level subscriptions, delete the subscription from the PlanID index.
			k.DeleteSubscriptionForPlan(ctx, s.PlanID, s.GetID())
		default:
			// If the subscription type is not recognized, panic with an error indicating an invalid subscription type.
			panic(fmt.Errorf("invalid subscription %d with type %T", item.GetID(), item))
		}

		// Finally, delete the subscription from the store and emit an event to notify its status change to 'Inactive'.
		k.DeleteSubscription(ctx, item.GetID())
		ctx.EventManager().EmitTypedEvent(
			&types.EventUpdateStatus{
				ID:     item.GetID(),
				Status: hubtypes.StatusInactive,
			},
		)

		// If the subscription is a NodeSubscription and the duration is specified in hours (non-zero),
		// delete the payout from the store and its associated indexes.
		if s, ok := item.(*types.NodeSubscription); ok && s.Hours != 0 {
			payout, found := k.GetPayout(ctx, item.GetID())
			if !found {
				// If the payout is not found, panic with an error indicating the missing payout.
				panic(fmt.Errorf("payout for subscription %d does not exist", item.GetID()))
			}

			// Delete the payout and its associated indexes from the Store.
			k.DeletePayout(ctx, payout.ID)
			k.DeletePayoutForAccount(ctx, payout.GetAddress(), payout.ID)
			k.DeletePayoutForNode(ctx, payout.GetNodeAddress(), payout.ID)
		}

		return false
	})

	// Return an empty ValidatorUpdate slice as no validator updates are needed for the end block.
	return nil
}
