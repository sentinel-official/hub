package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	hubtypes "github.com/sentinel-official/hub/types"
	hubutils "github.com/sentinel-official/hub/utils"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func (k *Keeper) SetSubscription(ctx sdk.Context, subscription types.Subscription) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionKey(subscription.GetID())
	)

	value, err := k.cdc.MarshalInterface(subscription)
	if err != nil {
		panic(err)
	}

	store.Set(key, value)
}

func (k *Keeper) GetSubscription(ctx sdk.Context, id uint64) (subscription types.Subscription, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return subscription, false
	}
	if err := k.cdc.UnmarshalInterface(value, &subscription); err != nil {
		panic(err)
	}

	return subscription, true
}

func (k *Keeper) DeleteSubscription(ctx sdk.Context, id uint64) {
	key := types.SubscriptionKey(id)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k *Keeper) GetSubscriptions(ctx sdk.Context) (items types.Subscriptions) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.SubscriptionKeyPrefix)
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Subscription
		if err := k.cdc.UnmarshalInterface(iter.Value(), &item); err != nil {
			panic(err)
		}

		items = append(items, item)
	}

	return items
}

func (k *Keeper) SetSubscriptionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForAccountKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HasSubscriptionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForAccountKey(addr, id)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteSubscriptionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForAccountKey(addr, id)
	)

	store.Delete(key)
}

func (k *Keeper) GetSubscriptionsForAccount(ctx sdk.Context, addr sdk.AccAddress) (items types.Subscriptions) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.GetSubscriptionForAccountKeyPrefix(addr))
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		item, found := k.GetSubscription(ctx, types.IDFromSubscriptionForAccountKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("subscription for account key %X does not exist", iter.Key()))
		}

		items = append(items, item)
	}

	return items
}

func (k *Keeper) SetSubscriptionForNode(ctx sdk.Context, addr hubtypes.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForNodeKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HashSubscriptionForNode(ctx sdk.Context, addr hubtypes.NodeAddress, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForNodeKey(addr, id)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteSubscriptionForNode(ctx sdk.Context, addr hubtypes.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForNodeKey(addr, id)
	)

	store.Delete(key)
}

func (k *Keeper) GetSubscriptionsForNode(ctx sdk.Context, addr hubtypes.NodeAddress) (items types.Subscriptions) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.GetSubscriptionForNodeKeyPrefix(addr))
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		item, found := k.GetSubscription(ctx, types.IDFromSubscriptionForNodeKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("subscription for node key %X does not exist", iter.Key()))
		}

		items = append(items, item)
	}

	return items
}

func (k *Keeper) SetSubscriptionForPlan(ctx sdk.Context, planID, subscriptionID uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForPlanKey(planID, subscriptionID)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HashSubscriptionForPlan(ctx sdk.Context, planID, subscriptionID uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForPlanKey(planID, subscriptionID)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteSubscriptionForPlan(ctx sdk.Context, planID, subscriptionID uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForPlanKey(planID, subscriptionID)
	)

	store.Delete(key)
}

func (k *Keeper) GetSubscriptionsForPlan(ctx sdk.Context, id uint64) (items types.Subscriptions) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.GetSubscriptionForPlanKeyPrefix(id))
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		item, found := k.GetSubscription(ctx, types.IDFromSubscriptionForPlanKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("subscription for plan key %X does not exist", iter.Key()))
		}

		items = append(items, item)
	}

	return items
}

func (k *Keeper) SetSubscriptionForInactiveAt(ctx sdk.Context, at time.Time, id uint64) {
	key := types.SubscriptionForInactiveAtKey(at, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) DeleteSubscriptionForInactiveAt(ctx sdk.Context, at time.Time, id uint64) {
	key := types.SubscriptionForInactiveAtKey(at, id)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k *Keeper) IterateSubscriptionsForInactiveAt(ctx sdk.Context, endTime time.Time, fn func(index int, item types.Subscription) (stop bool)) {
	store := k.Store(ctx)

	iter := store.Iterator(types.SubscriptionForInactiveAtKeyPrefix, sdk.PrefixEndBytes(types.GetSubscriptionForInactiveAtKeyPrefix(endTime)))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		subscription, found := k.GetSubscription(ctx, types.IDFromSubscriptionForInactiveAtKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("subscription for inactive at key %X does not exist", iter.Key()))
		}

		if stop := fn(i, subscription); stop {
			break
		}
		i++
	}
}

// CreateSubscriptionForNode creates a new NodeSubscription for a specific node and account.
func (k *Keeper) CreateSubscriptionForNode(ctx sdk.Context, accAddr sdk.AccAddress, nodeAddr hubtypes.NodeAddress, gigabytes, hours int64, denom string) (*types.NodeSubscription, error) {
	// Check if the node exists and is in an active status.
	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}
	if !node.Status.Equal(hubtypes.StatusActive) {
		return nil, types.NewErrorInvalidNodeStatus(nodeAddr, node.Status)
	}

	// Retrieve the current count and create a new NodeSubscription.
	count := k.GetCount(ctx)
	subscription := &types.NodeSubscription{
		BaseSubscription: &types.BaseSubscription{
			ID:         count + 1,
			Address:    accAddr.String(),
			InactiveAt: time.Time{},
			Status:     hubtypes.StatusActive,
			StatusAt:   ctx.BlockTime(),
		},
		NodeAddress: nodeAddr.String(),
		Gigabytes:   gigabytes,
		Hours:       hours,
		Deposit:     sdk.Coin{},
	}

	// Based on the provided gigabytes and hours, calculate the deposit and set the InactiveAt time.
	if gigabytes != 0 {
		price, found := node.GigabytePrice(denom)
		if !found {
			return nil, types.NewErrorPriceNotFound(denom)
		}
		subscription.InactiveAt = ctx.BlockTime().Add(types.Year) // TODO: move to params
		subscription.Deposit = sdk.NewCoin(
			price.Denom,
			hubtypes.AmountForBytes(price.Amount, hubtypes.Gigabyte.MulRaw(gigabytes)),
		)
	}
	if hours != 0 {
		price, found := node.HourlyPrice(denom)
		if !found {
			return nil, types.NewErrorPriceNotFound(denom)
		}
		subscription.InactiveAt = ctx.BlockTime().Add(
			time.Duration(hours) * time.Hour,
		)
		subscription.Deposit = sdk.NewCoin(
			price.Denom,
			price.Amount.MulRaw(hours),
		)
	}

	// Add the required deposit to the account's balance.
	if err := k.DepositAdd(ctx, accAddr, subscription.Deposit); err != nil {
		return nil, err
	}

	// Save the new NodeSubscription to the store and update the count.
	k.SetCount(ctx, count+1)
	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForAccount(ctx, accAddr, subscription.GetID())
	k.SetSubscriptionForNode(ctx, nodeAddr, subscription.GetID())
	k.SetSubscriptionForInactiveAt(ctx, subscription.GetInactiveAt(), subscription.GetID())

	// If the subscription is based on gigabytes, create an allocation and emit an event.
	if gigabytes != 0 {
		alloc := types.Allocation{
			ID:            subscription.GetID(),
			Address:       accAddr.String(),
			GrantedBytes:  hubtypes.Gigabyte.MulRaw(gigabytes),
			UtilisedBytes: sdk.ZeroInt(),
		}

		k.SetAllocation(ctx, alloc)
		ctx.EventManager().EmitTypedEvent(
			&types.EventAllocate{
				ID:      subscription.GetID(),
				Address: alloc.Address,
				Bytes:   alloc.GrantedBytes,
			},
		)
	}

	// If the subscription is based on hours, create a payout and emit an event.
	if hours != 0 {
		payout := types.Payout{
			ID:          subscription.GetID(),
			Address:     accAddr.String(),
			NodeAddress: nodeAddr.String(),
			Hours:       hours,
			Price: sdk.NewCoin(
				subscription.Deposit.Denom,
				subscription.Deposit.Amount.QuoRaw(hours),
			),
			NextAt: ctx.BlockTime(),
		}

		k.SetPayout(ctx, payout)
		k.SetPayoutForAccount(ctx, accAddr, payout.ID)
		k.SetPayoutForNode(ctx, nodeAddr, payout.ID)
		k.SetPayoutForNextAt(ctx, payout.NextAt, payout.ID)
		ctx.EventManager().EmitTypedEvent(
			&types.EventPayout{
				ID:          payout.ID,
				Address:     payout.Address,
				NodeAddress: payout.NodeAddress,
			},
		)
	}

	return subscription, nil
}

// CreateSubscriptionForPlan creates a new PlanSubscription for a specific plan and account.
func (k *Keeper) CreateSubscriptionForPlan(ctx sdk.Context, accAddr sdk.AccAddress, id uint64, denom string) (*types.PlanSubscription, error) {
	// Check if the plan exists and is in an active status.
	plan, found := k.GetPlan(ctx, id)
	if !found {
		return nil, types.NewErrorPlanNotFound(id)
	}
	if !plan.Status.Equal(hubtypes.StatusActive) {
		return nil, types.NewErrorInvalidPlanStatus(plan.ID, plan.Status)
	}

	// Get the price of the plan in the specified denomination.
	price, found := plan.Price(denom)
	if !found {
		return nil, types.NewErrorPriceNotFound(denom)
	}

	// Calculate the staking reward based on the plan price and staking share.
	var (
		stakingShare  = k.provider.StakingShare(ctx)
		stakingReward = hubutils.GetProportionOfCoin(price, stakingShare)
	)

	// Move the staking reward from the account to the fee collector module account.
	if err := k.SendCoinFromAccountToModule(ctx, accAddr, k.feeCollectorName, stakingReward); err != nil {
		return nil, err
	}

	// Calculate the payment amount after deducting the staking reward.
	var (
		provAddr = plan.GetAddress()
		payment  = price.Sub(stakingReward)
	)

	// Send the payment amount from the account to the plan provider address.
	if err := k.SendCoin(ctx, accAddr, provAddr.Bytes(), payment); err != nil {
		return nil, err
	}

	// Retrieve the current count and create a new PlanSubscription.
	count := k.GetCount(ctx)
	subscription := &types.PlanSubscription{
		BaseSubscription: &types.BaseSubscription{
			ID:         count + 1,
			Address:    accAddr.String(),
			InactiveAt: ctx.BlockTime().Add(plan.Duration),
			Status:     hubtypes.StatusActive,
			StatusAt:   ctx.BlockTime(),
		},
		PlanID: plan.ID,
		Denom:  price.Denom,
	}

	// Save the new PlanSubscription to the store and update the count.
	k.SetCount(ctx, count+1)
	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForAccount(ctx, accAddr, subscription.GetID())
	k.SetSubscriptionForPlan(ctx, plan.ID, subscription.GetID())
	k.SetSubscriptionForInactiveAt(ctx, subscription.GetInactiveAt(), subscription.GetID())

	// Create an allocation for the plan subscription and emit an event.
	alloc := types.Allocation{
		ID:            subscription.GetID(),
		Address:       accAddr.String(),
		GrantedBytes:  hubtypes.Gigabyte.MulRaw(plan.Gigabytes),
		UtilisedBytes: sdk.ZeroInt(),
	}

	k.SetAllocation(ctx, alloc)
	ctx.EventManager().EmitTypedEvent(
		&types.EventAllocate{
			ID:      subscription.GetID(),
			Address: alloc.Address,
			Bytes:   alloc.GrantedBytes,
		},
	)

	return subscription, nil
}
