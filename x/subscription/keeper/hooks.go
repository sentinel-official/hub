package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	hubutils "github.com/sentinel-official/hub/utils"
	"github.com/sentinel-official/hub/x/subscription/types"
)

// SessionInactiveHook is a function that handles the end of a session.
// It updates the allocation's utilized bytes, calculates and sends payments, and staking rewards.
func (k *Keeper) SessionInactiveHook(ctx sdk.Context, id uint64, accAddr sdk.AccAddress, nodeAddr hubtypes.NodeAddress, bytes sdk.Int) error {
	// Get the subscription associated with the provided subscription ID.
	subscription, found := k.GetSubscription(ctx, id)
	if !found {
		return fmt.Errorf("subscription %d does not exist", id)
	}

	// If the subscription is a NodeSubscription with non-zero duration (hours), no further action is needed.
	if s, ok := subscription.(*types.NodeSubscription); ok && s.Hours != 0 {
		return nil
	}

	// Get the allocation associated with the subscription and account address.
	alloc, found := k.GetAllocation(ctx, subscription.GetID(), accAddr)
	if !found {
		return fmt.Errorf("subscription allocation %d/%s does not exist", subscription.GetID(), accAddr)
	}

	var (
		gigabytePrice  sdk.Coin        // Gigabyte price based on the deposit amount and gigabytes (for NodeSubscription).
		previousAmount = sdk.ZeroInt() // Amount paid for previous utilization (for NodeSubscription).
	)

	// Based on the subscription type (NodeSubscription), calculate the payment amounts.
	if s, ok := subscription.(*types.NodeSubscription); ok && s.Gigabytes != 0 {
		gigabytePrice = sdk.NewCoin(
			s.Deposit.Denom,
			s.Deposit.Amount.QuoRaw(s.Gigabytes),
		)
		previousAmount = hubutils.AmountForBytes(gigabytePrice.Amount, alloc.UtilisedBytes)
	}

	// Update the allocation's utilized bytes by adding the provided bytes.
	alloc.UtilisedBytes = alloc.UtilisedBytes.Add(bytes)
	// Ensure that the utilized bytes don't exceed the granted bytes.
	if alloc.UtilisedBytes.GT(alloc.GrantedBytes) {
		alloc.UtilisedBytes = alloc.GrantedBytes
	}

	// Save the updated allocation to the store.
	k.SetAllocation(ctx, alloc)

	// Based on the subscription type (NodeSubscription), calculate the current payment amount.
	if s, ok := subscription.(*types.NodeSubscription); ok && s.Gigabytes != 0 {
		// Calculate the payment to be made for the current utilization.
		var (
			currentAmount = hubutils.AmountForBytes(gigabytePrice.Amount, alloc.UtilisedBytes)
			payment       = sdk.NewCoin(gigabytePrice.Denom, currentAmount.Sub(previousAmount))
			stakingShare  = k.node.StakingShare(ctx)
			stakingReward = hubutils.GetProportionOfCoin(payment, stakingShare)
		)

		// Move the staking reward from the deposit to the fee collector module account.
		if err := k.SendCoinFromDepositToModule(ctx, accAddr, k.feeCollectorName, stakingReward); err != nil {
			return err
		}

		// Subtract the staking reward from the payment to get the final payment amount.
		payment = payment.Sub(stakingReward)

		// Send the payment amount from the deposit to the node address.
		return k.SendCoinFromDepositToAccount(ctx, accAddr, nodeAddr.Bytes(), payment)
	}

	return nil
}
