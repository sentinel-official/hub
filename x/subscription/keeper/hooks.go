package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	hubutils "github.com/sentinel-official/hub/utils"
	"github.com/sentinel-official/hub/x/subscription/types"
)

// HookEndSession is a function that handles the end of a session.
// It updates the allocation's utilized bytes, calculates and sends payments, and staking rewards.
func (k *Keeper) HookEndSession(ctx sdk.Context, subscriptionID uint64, accAddr sdk.AccAddress, nodeAddr hubtypes.NodeAddress, bytes sdk.Int) error {
	// Get the subscription associated with the provided subscription ID.
	subscription, found := k.GetSubscription(ctx, subscriptionID)
	if !found {
		return fmt.Errorf("subscription %d does not exist", subscriptionID)
	}

	// Get the allocation associated with the subscription and account address.
	alloc, found := k.GetAllocation(ctx, subscription.GetID(), accAddr)
	if !found {
		return fmt.Errorf("subscription allocation %d/%s does not exist", subscription.GetID(), accAddr)
	}

	var (
		gigabytePrice  sdk.Coin        // Gigabyte price based on the deposit amount and gigabytes (for NodeSubscription).
		currentAmount  = sdk.ZeroInt() // Amount to be paid for the current utilization (for NodeSubscription).
		previousAmount = sdk.ZeroInt() // Amount paid for previous utilization (for NodeSubscription).
	)

	// Based on the subscription type (NodeSubscription), calculate the payment amounts.
	switch s := subscription.(type) {
	case *types.NodeSubscription:
		if s.Gigabytes != 0 {
			gigabytePrice = sdk.NewCoin(
				s.Deposit.Denom,
				s.Deposit.Amount.QuoRaw(s.Gigabytes),
			)
			previousAmount = hubutils.AmountForBytes(gigabytePrice.Amount, alloc.UtilisedBytes)
		}
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
	switch s := subscription.(type) {
	case *types.NodeSubscription:
		if s.Gigabytes != 0 {
			currentAmount = hubutils.AmountForBytes(gigabytePrice.Amount, alloc.UtilisedBytes)
		}
	}

	// Calculate the payment to be made for the current utilization.
	var (
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
