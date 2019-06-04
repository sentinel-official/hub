package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeeper_SetSubscriptionsCount(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	keeper.SetSubscriptionsCount(ctx, 1)
	count := keeper.GetSubscriptionsCount(ctx)
	require.Equal(t, uint64(1), count)

	keeper.SetSubscriptionsCount(ctx, 0)
	count = keeper.GetSubscriptionsCount(ctx)
	require.Equal(t, uint64(0), count)

	keeper.SetSubscriptionsCount(ctx, 2)
	count = keeper.GetSubscriptionsCount(ctx)
	require.Equal(t, uint64(2), count)
}

func TestKeeper_GetSubscriptionsCount(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	count := keeper.GetSubscriptionsCount(ctx)
	require.Equal(t, uint64(0), count)

	keeper.SetSubscriptionsCount(ctx, 1)
	count = keeper.GetSubscriptionsCount(ctx)
	require.Equal(t, uint64(1), count)
}

func TestKeeper_SetSubscription(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	keeper.SetSubscription(ctx, TestSubscriptionValid)
	result, found := keeper.GetSubscription(ctx, TestNodeValid.ID)
	require.Equal(t, true, found)
	require.Equal(t, TestSubscriptionValid, result)

	keeper.SetSubscription(ctx, TestSubscriptionEmpty)
	result, found = keeper.GetSubscription(ctx, TestSubscriptionEmpty.ID)
	require.Equal(t, true, found)
}

func TestKeeper_GetSubscription(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	_, found := keeper.GetSubscription(ctx, TestIDPos)
	require.Equal(t, false, found)

	keeper.SetSubscription(ctx, TestSubscriptionValid)
	result, found := keeper.GetSubscription(ctx, TestSubscriptionValid.ID)
	require.Equal(t, true, found)
	require.Equal(t, TestSubscriptionValid, result)
}

func TestKeeper_SetSubscriptionsCountOfNode(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	keeper.SetSubscriptionsCountOfNode(ctx, TestIDPos, 1)
	count := keeper.GetSubscriptionsCountOfNode(ctx, TestIDPos)
	require.Equal(t, uint64(1), count)

	keeper.SetSubscriptionsCountOfNode(ctx, TestIDPos, 2)
	count = keeper.GetSubscriptionsCountOfNode(ctx, TestIDPos)
	require.Equal(t, uint64(2), count)
}

func TestKeeper_GetSubscriptionsCountOfNode(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	count := keeper.GetSubscriptionsCountOfNode(ctx, TestIDPos)
	require.Equal(t, uint64(0), count)

	keeper.SetSubscriptionsCountOfNode(ctx, TestIDPos, 10)
	count = keeper.GetSubscriptionsCountOfNode(ctx, TestIDPos)
	require.Equal(t, uint64(10), count)
}

func TestKeeper_SetSubscriptionIDByNodeID(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	keeper.SetSubscriptionIDByNodeID(ctx, TestIDPos, 1, TestIDPos)
	id, found := keeper.GetSubscriptionIDByNodeID(ctx, TestIDPos, 1)
	require.Equal(t, true, found)
	require.Equal(t, TestIDPos, id)

	keeper.SetSubscriptionIDByNodeID(ctx, TestIDPos, 2, TestIDPos)
	id, found = keeper.GetSubscriptionIDByNodeID(ctx, TestIDPos, 2)
	require.Equal(t, true, found)
	require.Equal(t, TestIDPos, id)
}

func TestKeeper_GetSubscriptionIDByNodeID(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	_, found := keeper.GetSubscriptionIDByNodeID(ctx, TestIDPos, 1)
	require.Equal(t, false, found)

	keeper.SetSubscriptionIDByNodeID(ctx, TestIDPos, 2, TestIDPos)
	id, found := keeper.GetSubscriptionIDByNodeID(ctx, TestIDPos, 2)
	require.Equal(t, true, found)
	require.Equal(t, TestIDPos, id)
}

func TestKeeper_SetSubscriptionsCountOfAddress(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	keeper.SetSubscriptionsCountOfAddress(ctx, TestAddress1, 1)
	count := keeper.GetSubscriptionsCountOfAddress(ctx, TestAddress1)
	require.Equal(t, uint64(1), count)

	keeper.SetSubscriptionsCountOfAddress(ctx, TestAddress1, 10)
	count = keeper.GetSubscriptionsCountOfAddress(ctx, TestAddress1)
	require.Equal(t, uint64(10), count)
}

func TestKeeper_GetSubscriptionsCountOfAddress(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	count := keeper.GetSubscriptionsCountOfAddress(ctx, TestAddress1)
	require.Equal(t, uint64(0), count)

	keeper.SetSubscriptionsCountOfAddress(ctx, TestAddress1, 1)
	count = keeper.GetSubscriptionsCountOfAddress(ctx, TestAddress1)
	require.Equal(t, uint64(1), count)
}

func TestKeeper_SetSubscriptionIDByAddress(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	keeper.SetSubscriptionIDByAddress(ctx, TestAddress1, 1, TestIDPos)
	id, found := keeper.GetSubscriptionIDByAddress(ctx, TestAddress1, 1)
	require.Equal(t, true, found)
	require.Equal(t, TestIDPos, id)

	keeper.SetSubscriptionIDByAddress(ctx, TestAddressEmpty, 1, TestIDZero)
	id, found = keeper.GetSubscriptionIDByAddress(ctx, TestAddressEmpty, 1)
	require.Equal(t, true, found)
	require.Equal(t, TestIDZero, id)
}

func TestKeeper_GetSubscriptionIDByAddress(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	_, found := keeper.GetSubscriptionIDByAddress(ctx, TestAddress1, 1)
	require.Equal(t, false, found)

	keeper.SetSubscriptionIDByAddress(ctx, TestAddress1, 1, TestIDPos)
	id, found := keeper.GetSubscriptionIDByAddress(ctx, TestAddress1, 1)
	require.Equal(t, true, found)
	require.Equal(t, TestIDPos, id)
}

func TestKeeper_GetSubscriptionsOfNode(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	subscriptions := keeper.GetSubscriptionsOfNode(ctx, TestIDPos)
	require.Equal(t, TestSubscriptionsEmpty, subscriptions)

	keeper.SetSubscription(ctx, TestSubscriptionValid)
	keeper.SetSubscriptionIDByAddress(ctx, TestAddress1, 0, TestSubscriptionValid.ID)
	keeper.SetSubscriptionIDByNodeID(ctx, TestSubscriptionValid.ID, 0, TestIDZero)
	keeper.SetSubscriptionsCountOfNode(ctx, TestIDZero, 1)

	subscriptions = keeper.GetSubscriptionsOfNode(ctx, TestIDZero)
	require.Equal(t, TestSubscriptionsValid, subscriptions)

}

func TestKeeper_GetSubscriptionsOfAddress(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	subscriptions := keeper.GetSubscriptionsOfAddress(ctx, TestAddress1)
	require.Equal(t, TestSubscriptionsEmpty, subscriptions)

	keeper.SetSubscription(ctx, TestSubscriptionValid)
	keeper.SetSubscriptionIDByAddress(ctx, TestAddress1, 0, TestSubscriptionValid.ID)
	keeper.SetSubscriptionsCountOfAddress(ctx, TestAddress1, 1)

	subscriptions = keeper.GetSubscriptionsOfAddress(ctx, TestAddress1)
	require.Equal(t, TestSubscriptionsValid, subscriptions)
}

func TestKeeper_GetAllSubscriptions(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	subscriptions := keeper.GetAllSubscriptions(ctx)
	require.Equal(t, TestSubscriptionsNil, subscriptions)

	keeper.SetSubscription(ctx, TestSubscriptionValid)
	subscriptions = keeper.GetAllSubscriptions(ctx)
	require.Equal(t, TestSubscriptionsValid, subscriptions)
}
