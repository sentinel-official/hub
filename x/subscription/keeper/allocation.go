package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	hubutils "github.com/sentinel-official/hub/utils"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func (k *Keeper) SetAllocation(ctx sdk.Context, alloc types.Allocation) {
	var (
		store = k.Store(ctx)
		key   = types.AllocationKey(alloc.ID, alloc.GetAddress())
		value = k.cdc.MustMarshal(&alloc)
	)

	store.Set(key, value)
}

func (k *Keeper) GetAllocation(ctx sdk.Context, id uint64, addr sdk.AccAddress) (alloc types.Allocation, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.AllocationKey(id, addr)
		value = store.Get(key)
	)

	if value == nil {
		return alloc, false
	}

	k.cdc.MustUnmarshal(value, &alloc)
	return alloc, true
}

func (k *Keeper) HasAllocation(ctx sdk.Context, id uint64, addr sdk.AccAddress) bool {
	var (
		store = k.Store(ctx)
		key   = types.AllocationKey(id, addr)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteAllocation(ctx sdk.Context, id uint64, addr sdk.AccAddress) {
	var (
		store = k.Store(ctx)
		key   = types.AllocationKey(id, addr)
	)

	store.Delete(key)
}

func (k *Keeper) GetAllocations(ctx sdk.Context, id uint64) (items types.Allocations) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.GetAllocationKeyPrefix(id))
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Allocation
		k.cdc.MustUnmarshal(iter.Value(), &item)
		items = append(items, item)
	}

	return items
}

func (k *Keeper) IterateAllocations(ctx sdk.Context, id uint64, fn func(index int, item types.Allocation) (stop bool)) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.GetAllocationKeyPrefix(id))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var alloc types.Allocation
		k.cdc.MustUnmarshal(iter.Value(), &alloc)

		if stop := fn(i, alloc); stop {
			break
		}
		i++
	}
}

func (k *Keeper) HookEndSession(ctx sdk.Context, id uint64, accAddr sdk.AccAddress, nodeAddr hubtypes.NodeAddress, bytes sdk.Int) error {
	subscription, found := k.GetSubscription(ctx, id)
	if !found {
		return fmt.Errorf("subscription %d does not exist", id)
	}

	alloc, found := k.GetAllocation(ctx, id, accAddr)
	if !found {
		return fmt.Errorf("subscription allocation %d/%s does not exist", id, accAddr)
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
