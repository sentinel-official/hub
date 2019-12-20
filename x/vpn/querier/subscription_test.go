package querier

import (
	"fmt"
	"testing"
	
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	
	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func Test_querySubscription(t *testing.T) {
	ctx, k, _, _ := keeper.CreateTestInput(t, false)
	cdc := keeper.MakeTestCodec()
	var err error
	var subscription types.Subscription
	
	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySubscription),
		Data: []byte{},
	}
	
	res, _err := querySubscription(ctx, req, k)
	require.NotNil(t, _err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)
	
	k.SetSubscription(ctx, types.TestSubscription)
	req.Data, err = cdc.MarshalJSON(types.NewQuerySubscriptionParams(hub.NewSubscriptionID(0)))
	require.Nil(t, err)
	
	res, _err = querySubscription(ctx, req, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)
	
	err = cdc.UnmarshalJSON(res, &subscription)
	require.Nil(t, err)
	require.Equal(t, types.TestSubscription, subscription)
	
	req.Data, err = cdc.MarshalJSON(types.NewQuerySubscriptionParams(hub.NewSubscriptionID(1)))
	require.Nil(t, err)
	
	res, _err = querySubscription(ctx, req, k)
	require.Nil(t, res)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)
}

func Test_querySubscriptionsOfNode(t *testing.T) {
	ctx, k, _, _ := keeper.CreateTestInput(t, false)
	cdc := keeper.MakeTestCodec()
	var err error
	var subscriptions []types.Subscription
	
	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySubscriptionsOfNode),
		Data: []byte{},
	}
	
	res, _err := querySubscriptionsOfNode(ctx, req, k)
	require.NotNil(t, _err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)
	
	k.SetSubscription(ctx, types.TestSubscription)
	k.SetSubscriptionsCountOfNode(ctx, types.TestNode.ID, 1)
	k.SetSubscriptionIDByNodeID(ctx, types.TestNode.ID, 0, types.TestSubscription.ID)
	
	req.Data, err = cdc.MarshalJSON(types.NewQuerySubscriptionsOfNodePrams(hub.NewNodeID(0)))
	require.Nil(t, err)
	
	res, _err = querySubscriptionsOfNode(ctx, req, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	
	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.Equal(t, []types.Subscription{types.TestSubscription}, subscriptions)
	
	req.Data, err = cdc.MarshalJSON(types.NewQuerySubscriptionsOfNodePrams(hub.NewNodeID(1)))
	require.Nil(t, err)
	
	res, _err = querySubscriptionsOfNode(ctx, req, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	
	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.NotEqual(t, []types.Subscription{types.TestSubscription}, subscriptions)
}

func Test_querySubscriptionsOfAddress(t *testing.T) {
	ctx, k, _, _ := keeper.CreateTestInput(t, false)
	cdc := keeper.MakeTestCodec()
	var err error
	var subscriptions []types.Subscription
	
	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySubscriptionsOfAddress),
		Data: []byte{},
	}
	
	res, _err := querySubscriptionsOfAddress(ctx, req, k)
	require.NotNil(t, _err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)
	
	k.SetSubscription(ctx, types.TestSubscription)
	k.SetSubscriptionsCountOfAddress(ctx, types.TestAddress1, 1)
	k.SetSubscriptionIDByAddress(ctx, types.TestAddress1, 0, types.TestSubscription.ID)
	
	req.Data, err = cdc.MarshalJSON(types.NewQuerySubscriptionsOfAddressParams([]byte("")))
	require.Nil(t, err)
	
	res, _err = querySubscriptionsOfAddress(ctx, req, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	
	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.NotEqual(t, []types.Subscription{types.TestSubscription}, subscriptions)
	
	req.Data, err = cdc.MarshalJSON(types.NewQuerySubscriptionsOfAddressParams(types.TestAddress1))
	require.Nil(t, err)
	
	res, _err = querySubscriptionsOfAddress(ctx, req, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	
	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.Equal(t, []types.Subscription{types.TestSubscription}, subscriptions)
	
	req.Data, err = cdc.MarshalJSON(types.NewQuerySubscriptionsOfAddressParams(types.TestAddress2))
	require.Nil(t, err)
	
	res, _err = querySubscriptionsOfAddress(ctx, req, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	
	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.NotEqual(t, []types.Subscription{types.TestSubscription}, subscriptions)
}

func Test_queryAllSubscriptions(t *testing.T) {
	ctx, k, _, _ := keeper.CreateTestInput(t, false)
	cdc := keeper.MakeTestCodec()
	var err error
	var subscriptions []types.Subscription
	
	res, _err := queryAllSubscriptions(ctx, k)
	require.Nil(t, _err)
	require.Equal(t, []byte("null"), res)
	
	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.NotEqual(t, []types.Subscription{types.TestSubscription}, subscriptions)
	
	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.NotEqual(t, []types.Subscription{types.TestSubscription}, subscriptions)
	
	k.SetSubscription(ctx, types.TestSubscription)
	require.Nil(t, err)
	
	res, _err = queryAllSubscriptions(ctx, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)
	
	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.Equal(t, []types.Subscription{types.TestSubscription}, subscriptions)
	
	subscription := types.TestSubscription
	subscription.ID = hub.NewSubscriptionID(1)
	k.SetSubscription(ctx, subscription)
	require.Nil(t, err)
	
	res, _err = queryAllSubscriptions(ctx, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)
	
	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.Equal(t, append([]types.Subscription{types.TestSubscription}, subscription), subscriptions)
}

func Test_querySessionsCountOfSubscription(t *testing.T) {
	ctx, k, _, _ := keeper.CreateTestInput(t, false)
	cdc := keeper.MakeTestCodec()
	var err error
	var count uint64
	
	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySessionsCountOfSubscription),
		Data: []byte{},
	}
	
	res, _err := querySessionsCountOfSubscription(ctx, req, k)
	require.NotNil(t, _err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)
	
	k.SetSessionsCountOfSubscription(ctx, hub.NewSubscriptionID(0), 1)
	req.Data, err = cdc.MarshalJSON(types.NewQuerySessionsCountOfSubscriptionParams(hub.NewSubscriptionID(0)))
	require.Nil(t, err)
	
	res, _err = querySessionsCountOfSubscription(ctx, req, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	
	err = cdc.UnmarshalJSON(res, &count)
	require.Nil(t, err)
	
	require.Equal(t, uint64(1), count)
	
	k.SetSessionsCountOfSubscription(ctx, hub.NewSubscriptionID(0), 2)
	req.Data, err = cdc.MarshalJSON(types.NewQuerySessionsCountOfSubscriptionParams(hub.NewSubscriptionID(0)))
	require.Nil(t, err)
	
	res, _err = querySessionsCountOfSubscription(ctx, req, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	
	err = cdc.UnmarshalJSON(res, &count)
	require.Nil(t, err)
	require.Equal(t, uint64(2), count)
	
	k.SetSessionsCountOfSubscription(ctx, hub.NewSubscriptionID(1), 1)
	req.Data, err = cdc.MarshalJSON(types.NewQuerySessionsCountOfSubscriptionParams(hub.NewSubscriptionID(1)))
	require.Nil(t, err)
	
	res, _err = querySessionsCountOfSubscription(ctx, req, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	
	err = cdc.UnmarshalJSON(res, &count)
	require.Nil(t, err)
	require.Equal(t, uint64(1), count)
	
	k.SetSessionsCountOfSubscription(ctx, hub.NewSubscriptionID(1), 2)
	req.Data, err = cdc.MarshalJSON(types.NewQuerySessionsCountOfSubscriptionParams(hub.NewSubscriptionID(1)))
	require.Nil(t, err)
	
	res, _err = querySessionsCountOfSubscription(ctx, req, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	
	err = cdc.UnmarshalJSON(res, &count)
	require.Nil(t, err)
	require.Equal(t, uint64(2), count)
	
}
