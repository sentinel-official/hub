package keeper

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

	_, found := k.GetSubscription(ctx, hub.NewSubscriptionID(0))
	require.Equal(t, false, found)

	k.SetSubscription(ctx, types.TestSubscription)
	result, found := k.GetSubscription(ctx, types.TestSubscription.ID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestSubscription, result)
}

func TestKeeper_GetSubscription(t *testing.T) {
	TestKeeper_SetSubscription(t)
}

func TestKeeper_SetSubscriptionsCountOfNode(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	count := k.GetSubscriptionsCountOfNode(ctx, hub.NewNodeID(0))
	require.Equal(t, uint64(0), count)

	k.SetSubscriptionsCountOfNode(ctx, hub.NewNodeID(0), 1)
	count = k.GetSubscriptionsCountOfNode(ctx, hub.NewNodeID(0))
	require.Equal(t, uint64(1), count)

	k.SetSubscriptionsCountOfNode(ctx, hub.NewNodeID(0), 2)
	count = k.GetSubscriptionsCountOfNode(ctx, hub.NewNodeID(0))
	require.Equal(t, uint64(2), count)

	k.SetSubscriptionsCountOfNode(ctx, hub.NewNodeID(0), 1)
	count = k.GetSubscriptionsCountOfNode(ctx, hub.NewNodeID(0))
	require.Equal(t, uint64(1), count)
}

func TestKeeper_GetSubscriptionsCountOfNode(t *testing.T) {
	TestKeeper_SetSubscriptionsCountOfNode(t)
}

func TestKeeper_SetSubscriptionIDByNodeID(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	_, found := k.GetSubscriptionIDByNodeID(ctx, hub.NewNodeID(0), 1)
	require.Equal(t, false, found)

	k.SetSubscriptionIDByNodeID(ctx, hub.NewNodeID(0), 1, hub.NewSubscriptionID(0))
	id, found := k.GetSubscriptionIDByNodeID(ctx, hub.NewNodeID(0), 1)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewSubscriptionID(0), id)

	k.SetSubscriptionIDByNodeID(ctx, hub.NewNodeID(0), 2, hub.NewSubscriptionID(1))
	id, found = k.GetSubscriptionIDByNodeID(ctx, hub.NewNodeID(0), 2)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewSubscriptionID(1), id)

	k.SetSubscriptionIDByNodeID(ctx, hub.NewNodeID(1), 1, hub.NewSubscriptionID(0))
	id, found = k.GetSubscriptionIDByNodeID(ctx, hub.NewNodeID(1), 1)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewSubscriptionID(0), id)

	k.SetSubscriptionIDByNodeID(ctx, hub.NewNodeID(1), 2, hub.NewSubscriptionID(1))
	id, found = k.GetSubscriptionIDByNodeID(ctx, hub.NewNodeID(1), 2)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewSubscriptionID(1), id)
}

func TestKeeper_GetSubscriptionIDByNodeID(t *testing.T) {
	TestKeeper_SetSubscriptionIDByNodeID(t)
}

func TestKeeper_SetSubscriptionsCountOfAddress(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	count := k.GetSubscriptionsCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(0), count)

	k.SetSubscriptionsCountOfAddress(ctx, []byte(""), 1)
	count = k.GetSubscriptionsCountOfAddress(ctx, []byte(""))
	require.Equal(t, uint64(1), count)

	k.SetSubscriptionsCountOfAddress(ctx, types.TestAddress1, 1)
	count = k.GetSubscriptionsCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(1), count)

	k.SetSubscriptionsCountOfAddress(ctx, types.TestAddress1, 2)
	count = k.GetSubscriptionsCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(2), count)

	k.SetSubscriptionsCountOfAddress(ctx, types.TestAddress1, 1)
	count = k.GetSubscriptionsCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(1), count)

	k.SetSubscriptionsCountOfAddress(ctx, types.TestAddress2, 1)
	count = k.GetSubscriptionsCountOfAddress(ctx, types.TestAddress2)
	require.Equal(t, uint64(1), count)
}

func TestKeeper_GetSubscriptionsCountOfAddress(t *testing.T) {
	TestKeeper_SetSubscriptionsCountOfAddress(t)
}

func TestKeeper_SetSubscriptionIDByAddress(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	_, found := k.GetSubscriptionIDByAddress(ctx, types.TestAddress1, 1)
	require.Equal(t, false, found)

	k.SetSubscriptionIDByAddress(ctx, []byte(""), 1, hub.NewSubscriptionID(0))
	id, found := k.GetSubscriptionIDByAddress(ctx, []byte(""), 1)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewSubscriptionID(0), id)

	k.SetSubscriptionIDByAddress(ctx, []byte(""), 2, hub.NewSubscriptionID(1))
	id, found = k.GetSubscriptionIDByAddress(ctx, []byte(""), 2)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewSubscriptionID(1), id)

	k.SetSubscriptionIDByAddress(ctx, types.TestAddress1, 1, hub.NewSubscriptionID(0))
	id, found = k.GetSubscriptionIDByAddress(ctx, types.TestAddress1, 1)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewSubscriptionID(0), id)

	k.SetSubscriptionIDByAddress(ctx, types.TestAddress1, 2, hub.NewSubscriptionID(1))
	id, found = k.GetSubscriptionIDByAddress(ctx, types.TestAddress1, 2)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewSubscriptionID(1), id)

	k.SetSubscriptionIDByAddress(ctx, types.TestAddress2, 1, hub.NewSubscriptionID(0))
	id, found = k.GetSubscriptionIDByAddress(ctx, types.TestAddress2, 1)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewSubscriptionID(0), id)

	k.SetSubscriptionIDByAddress(ctx, types.TestAddress2, 1, hub.NewSubscriptionID(1))
	id, found = k.GetSubscriptionIDByAddress(ctx, types.TestAddress1, 2)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewSubscriptionID(1), id)
}

func TestKeeper_GetSubscriptionIDByAddress(t *testing.T) {
	TestKeeper_SetSubscriptionIDByAddress(t)
}

func TestKeeper_GetSubscriptionsOfNode(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	subscriptions := k.GetSubscriptionsOfNode(ctx, hub.NewNodeID(0))
	require.Equal(t, []types.Subscription{}, subscriptions)

	k.SetSubscription(ctx, types.TestSubscription)
	k.SetSubscriptionsCountOfNode(ctx, hub.NewNodeID(0), 1)

	subscriptions = k.GetSubscriptionsOfNode(ctx, hub.NewNodeID(0))
	require.Equal(t, []types.Subscription{types.TestSubscription}, subscriptions)

	k.SetSubscription(ctx, types.TestSubscription)
	k.SetSubscriptionsCountOfNode(ctx, hub.NewNodeID(0), 2)

	subscriptions = k.GetSubscriptionsOfNode(ctx, hub.NewNodeID(0))
	require.Equal(t, append([]types.Subscription{types.TestSubscription}, types.TestSubscription), subscriptions)

}

func TestKeeper_GetSubscriptionsOfAddress(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	subscriptions := k.GetSubscriptionsOfAddress(ctx, types.TestAddress1)
	require.Equal(t, []types.Subscription{}, subscriptions)

	k.SetSubscription(ctx, types.TestSubscription)
	k.SetSubscriptionsCountOfNode(ctx, hub.NewNodeID(0), 1)

	subscriptions = k.GetSubscriptionsOfAddress(ctx, types.TestAddress2)
	require.Equal(t, []types.Subscription{}, subscriptions)

	k.SetSubscriptionsCountOfAddress(ctx, types.TestAddress2, 1)
	subscriptions = k.GetSubscriptionsOfAddress(ctx, types.TestAddress2)
	require.Equal(t, []types.Subscription{types.TestSubscription}, subscriptions)

	k.SetSubscriptionIDByAddress(ctx, types.TestAddress2, 1, hub.NewSubscriptionID(0))
	subscriptions = k.GetSubscriptionsOfAddress(ctx, types.TestAddress2)
	require.Equal(t, []types.Subscription{types.TestSubscription}, subscriptions)

	k.SetSubscription(ctx, types.TestSubscription)
	k.SetSubscriptionsCountOfNode(ctx, hub.NewNodeID(0), 2)
	k.SetSubscriptionsCountOfAddress(ctx, types.TestAddress2, 2)
	k.SetSubscriptionIDByAddress(ctx, types.TestAddress2, 2, hub.NewSubscriptionID(1))

	subscriptions = k.GetSubscriptionsOfAddress(ctx, types.TestAddress2)
	require.Equal(t, append([]types.Subscription{types.TestSubscription}, types.TestSubscription), subscriptions)
}

func TestKeeper_GetAllSubscriptions(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	subscriptions := k.GetAllSubscriptions(ctx)
	require.Equal(t, []types.Subscription(nil), subscriptions)

	k.SetSubscription(ctx, types.TestSubscription)
	subscriptions = k.GetAllSubscriptions(ctx)
	require.Equal(t, []types.Subscription{types.TestSubscription}, subscriptions)

	subscription := types.TestSubscription
	subscription.ID = hub.NewSubscriptionID(1)
	k.SetSubscription(ctx, subscription)
	subscriptions = k.GetAllSubscriptions(ctx)
	require.Equal(t, append([]types.Subscription{types.TestSubscription}, subscription), subscriptions)

	k.SetSubscription(ctx, types.TestSubscription)
	subscriptions = k.GetAllSubscriptions(ctx)
	require.Equal(t, append([]types.Subscription{types.TestSubscription}, subscription), subscriptions)
}
