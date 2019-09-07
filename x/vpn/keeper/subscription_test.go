package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func TestKeeper_SetSubscriptionsCount(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	count := k.GetSubscriptionsCount(ctx)
	require.Equal(t, uint64(0), count)

	k.SetSubscriptionsCount(ctx, 1)
	count = k.GetSubscriptionsCount(ctx)
	require.Equal(t, uint64(1), count)

	k.SetSubscriptionsCount(ctx, 2)
	count = k.GetSubscriptionsCount(ctx)
	require.Equal(t, uint64(2), count)

	k.SetSubscriptionsCount(ctx, 3)
	count = k.GetSubscriptionsCount(ctx)
	require.Equal(t, uint64(3), count)
}

func TestKeeper_GetSubscriptionsCount(t *testing.T) {
	TestKeeper_SetSubscriptionsCount(t)
}

func TestKeeper_SetSubscription(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	_, found := k.GetSubscription(ctx, hub.NewIDFromUInt64(0))
	require.Equal(t, false, found)

	k.SetSubscription(ctx, types.Subscription{})
	result, found := k.GetSubscription(ctx, types.Subscription{}.ID)
	require.Equal(t, true, found)

	k.SetSubscription(ctx, subscriptionValid)
	result, found = k.GetSubscription(ctx, nodeValid.ID)
	require.Equal(t, true, found)
	require.Equal(t, subscriptionValid, result)
}

func TestKeeper_GetSubscription(t *testing.T) {
	TestKeeper_SetSubscription(t)
}

func TestKeeper_SetSubscriptionsCountOfNode(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	count := k.GetSubscriptionsCountOfNode(ctx, hub.NewIDFromUInt64(0))
	require.Equal(t, uint64(0), count)

	k.SetSubscriptionsCountOfNode(ctx, hub.NewIDFromUInt64(0), 1)
	count = k.GetSubscriptionsCountOfNode(ctx, hub.NewIDFromUInt64(0))
	require.Equal(t, uint64(1), count)

	k.SetSubscriptionsCountOfNode(ctx, hub.NewIDFromUInt64(0), 2)
	count = k.GetSubscriptionsCountOfNode(ctx, hub.NewIDFromUInt64(0))
	require.Equal(t, uint64(2), count)

	k.SetSubscriptionsCountOfNode(ctx, hub.NewIDFromUInt64(0), 1)
	count = k.GetSubscriptionsCountOfNode(ctx, hub.NewIDFromUInt64(0))
	require.Equal(t, uint64(1), count)
}

func TestKeeper_GetSubscriptionsCountOfNode(t *testing.T) {
	TestKeeper_SetSubscriptionsCountOfNode(t)
}

func TestKeeper_SetSubscriptionIDByNodeID(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	_, found := k.GetSubscriptionIDByNodeID(ctx, hub.NewIDFromUInt64(0), 1)
	require.Equal(t, false, found)

	k.SetSubscriptionIDByNodeID(ctx, hub.NewIDFromUInt64(0), 1, hub.NewIDFromUInt64(0))
	id, found := k.GetSubscriptionIDByNodeID(ctx, hub.NewIDFromUInt64(0), 1)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewIDFromUInt64(0), id)

	k.SetSubscriptionIDByNodeID(ctx, hub.NewIDFromUInt64(0), 2, hub.NewIDFromUInt64(1))
	id, found = k.GetSubscriptionIDByNodeID(ctx, hub.NewIDFromUInt64(0), 2)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewIDFromUInt64(1), id)

	k.SetSubscriptionIDByNodeID(ctx, hub.NewIDFromUInt64(1), 1, hub.NewIDFromUInt64(0))
	id, found = k.GetSubscriptionIDByNodeID(ctx, hub.NewIDFromUInt64(1), 1)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewIDFromUInt64(0), id)

	k.SetSubscriptionIDByNodeID(ctx, hub.NewIDFromUInt64(1), 2, hub.NewIDFromUInt64(1))
	id, found = k.GetSubscriptionIDByNodeID(ctx, hub.NewIDFromUInt64(1), 2)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewIDFromUInt64(1), id)
}

func TestKeeper_GetSubscriptionIDByNodeID(t *testing.T) {
	TestKeeper_SetSubscriptionIDByNodeID(t)
}

func TestKeeper_SetSubscriptionsCountOfAddress(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	count := k.GetSubscriptionsCountOfAddress(ctx, address1)
	require.Equal(t, uint64(0), count)

	k.SetSubscriptionsCountOfAddress(ctx, []byte(""), 1)
	count = k.GetSubscriptionsCountOfAddress(ctx, []byte(""))
	require.Equal(t, uint64(1), count)

	k.SetSubscriptionsCountOfAddress(ctx, address1, 1)
	count = k.GetSubscriptionsCountOfAddress(ctx, address1)
	require.Equal(t, uint64(1), count)

	k.SetSubscriptionsCountOfAddress(ctx, address1, 2)
	count = k.GetSubscriptionsCountOfAddress(ctx, address1)
	require.Equal(t, uint64(2), count)

	k.SetSubscriptionsCountOfAddress(ctx, address1, 1)
	count = k.GetSubscriptionsCountOfAddress(ctx, address1)
	require.Equal(t, uint64(1), count)

	k.SetSubscriptionsCountOfAddress(ctx, address2, 1)
	count = k.GetSubscriptionsCountOfAddress(ctx, address2)
	require.Equal(t, uint64(1), count)
}

func TestKeeper_GetSubscriptionsCountOfAddress(t *testing.T) {
	TestKeeper_SetSubscriptionsCountOfAddress(t)
}

func TestKeeper_SetSubscriptionIDByAddress(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	_, found := k.GetSubscriptionIDByAddress(ctx, address1, 1)
	require.Equal(t, false, found)

	k.SetSubscriptionIDByAddress(ctx, []byte(""), 1, hub.NewIDFromUInt64(0))
	id, found := k.GetSubscriptionIDByAddress(ctx, []byte(""), 1)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewIDFromUInt64(0), id)

	k.SetSubscriptionIDByAddress(ctx, []byte(""), 2, hub.NewIDFromUInt64(1))
	id, found = k.GetSubscriptionIDByAddress(ctx, []byte(""), 2)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewIDFromUInt64(1), id)

	k.SetSubscriptionIDByAddress(ctx, address1, 1, hub.NewIDFromUInt64(0))
	id, found = k.GetSubscriptionIDByAddress(ctx, address1, 1)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewIDFromUInt64(0), id)

	k.SetSubscriptionIDByAddress(ctx, address1, 2, hub.NewIDFromUInt64(1))
	id, found = k.GetSubscriptionIDByAddress(ctx, address1, 2)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewIDFromUInt64(1), id)

	k.SetSubscriptionIDByAddress(ctx, address2, 1, hub.NewIDFromUInt64(0))
	id, found = k.GetSubscriptionIDByAddress(ctx, address2, 1)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewIDFromUInt64(0), id)

	k.SetSubscriptionIDByAddress(ctx, address2, 1, hub.NewIDFromUInt64(1))
	id, found = k.GetSubscriptionIDByAddress(ctx, address1, 2)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewIDFromUInt64(1), id)
}

func TestKeeper_GetSubscriptionIDByAddress(t *testing.T) {
	TestKeeper_SetSubscriptionIDByAddress(t)
}

func TestKeeper_GetSubscriptionsOfNode(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	subscriptions := k.GetSubscriptionsOfNode(ctx, hub.NewIDFromUInt64(0))
	require.Equal(t, []types.Subscription{}, subscriptions)

	k.SetSubscription(ctx, subscriptionValid)
	k.SetSubscriptionsCountOfNode(ctx, hub.NewIDFromUInt64(0), 1)

	subscriptions = k.GetSubscriptionsOfNode(ctx, hub.NewIDFromUInt64(0))
	require.Equal(t, []types.Subscription{subscriptionValid}, subscriptions)

	k.SetSubscription(ctx, subscriptionValid)
	k.SetSubscriptionsCountOfNode(ctx, hub.NewIDFromUInt64(0), 2)

	subscriptions = k.GetSubscriptionsOfNode(ctx, hub.NewIDFromUInt64(0))
	require.Equal(t, append([]types.Subscription{subscriptionValid}, subscriptionValid), subscriptions)

}

func TestKeeper_GetSubscriptionsOfAddress(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	subscriptions := k.GetSubscriptionsOfNode(ctx, hub.NewIDFromUInt64(0))
	require.Equal(t, []types.Subscription{}, subscriptions)

	k.SetSubscription(ctx, subscriptionValid)
	k.SetSubscriptionsCountOfNode(ctx, hub.NewIDFromUInt64(0), 1)

	subscriptions = k.GetSubscriptionsOfNode(ctx, hub.NewIDFromUInt64(0))
	require.Equal(t, []types.Subscription{subscriptionValid}, subscriptions)

	k.SetSubscription(ctx, subscriptionValid)
	k.SetSubscriptionsCountOfNode(ctx, hub.NewIDFromUInt64(0), 2)

	subscriptions = k.GetSubscriptionsOfNode(ctx, hub.NewIDFromUInt64(0))
	require.Equal(t, append([]types.Subscription{subscriptionValid}, subscriptionValid), subscriptions)
}

func TestKeeper_GetAllSubscriptions(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	subscriptions := k.GetAllSubscriptions(ctx)
	require.Equal(t, []types.Subscription(nil), subscriptions)

	k.SetSubscription(ctx, subscriptionValid)
	subscriptions = k.GetAllSubscriptions(ctx)
	require.Equal(t, []types.Subscription{subscriptionValid}, subscriptions)

	subscription := subscriptionValid
	subscription.ID = hub.NewIDFromUInt64(1)
	k.SetSubscription(ctx, subscription)
	subscriptions = k.GetAllSubscriptions(ctx)
	require.Equal(t, append([]types.Subscription{subscriptionValid}, subscription), subscriptions)

	k.SetSubscription(ctx, subscriptionValid)
	subscriptions = k.GetAllSubscriptions(ctx)
	require.Equal(t, append([]types.Subscription{subscriptionValid}, subscription), subscriptions)
}
