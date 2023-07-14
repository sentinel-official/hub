package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	hubutils "github.com/sentinel-official/hub/utils"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func (k *Keeper) HookEndSession(ctx sdk.Context, subscriptionID uint64, accAddr sdk.AccAddress, nodeAddr hubtypes.NodeAddress, bytes sdk.Int) error {
	subscription, found := k.GetSubscription(ctx, subscriptionID)
	if !found {
		return fmt.Errorf("subscription %d does not exist", subscriptionID)
	}

	alloc, found := k.GetAllocation(ctx, subscription.GetID(), accAddr)
	if !found {
		return fmt.Errorf("subscription allocation %d/%s does not exist", subscription.GetID(), accAddr)
	}

	var (
		gigabytePrice  sdk.Coin
		currentAmount  = sdk.ZeroInt()
		previousAmount = sdk.ZeroInt()
	)

	switch s := subscription.(type) {
	case *types.NodeSubscription:
		if s.Gigabytes != 0 {
			gigabytePrice = sdk.NewCoin(
				s.Deposit.Denom,
				s.Deposit.Amount.QuoRaw(s.Gigabytes),
			)
			previousAmount = hubtypes.AmountForBytes(gigabytePrice.Amount, alloc.UtilisedBytes)
		}
	}

	alloc.UtilisedBytes = alloc.UtilisedBytes.Add(bytes)
	if alloc.UtilisedBytes.GT(alloc.GrantedBytes) {
		alloc.UtilisedBytes = alloc.GrantedBytes
	}

	k.SetAllocation(ctx, alloc)

	switch s := subscription.(type) {
	case *types.NodeSubscription:
		if s.Gigabytes != 0 {
			currentAmount = hubtypes.AmountForBytes(gigabytePrice.Amount, alloc.UtilisedBytes)
		}
	}

	var (
		payment       = sdk.NewCoin(gigabytePrice.Denom, currentAmount.Sub(previousAmount))
		stakingShare  = k.node.StakingShare(ctx)
		stakingReward = hubutils.GetProportionOfCoin(payment, stakingShare)
	)

	if err := k.SendCoinFromDepositToModule(ctx, accAddr, k.feeCollectorName, stakingReward); err != nil {
		return err
	}

	payment = payment.Sub(stakingReward)
	return k.SendCoinFromDepositToAccount(ctx, accAddr, nodeAddr.Bytes(), payment)
}
