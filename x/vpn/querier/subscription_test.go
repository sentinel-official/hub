package querier

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/sentinel-hub/x/vpn/keeper"
	hub "github.com/sentinel-official/sentinel-hub/x/vpn/types"
)

func TestNewQuerySubscriptionParams(t *testing.T) {
	params := NewQuerySubscriptionParams(hub.TestIDZero)
	require.Equal(t, TestSubscriptionParamsZero, params)

	params = NewQuerySubscriptionParams(hub.TestIDPos)
	require.Equal(t, TestSubscriptionParamsPos, params)
}

func Test_querySubscription(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var subscription hub.Subscription

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", hub.QuerierRoute, QuerySubscription),
		Data: []byte{},
	}

	res, _err := querySubscription(ctx, cdc, req, vpnKeeper)
	require.NotNil(t, _err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)

	vpnKeeper.SetSubscription(ctx, hub.TestSubscriptionValid)
	req.Data, err = cdc.MarshalJSON(NewQuerySubscriptionParams(hub.TestIDZero))
	require.Nil(t, err)

	res, _err = querySubscription(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &subscription)
	require.Nil(t, err)
	require.Equal(t, hub.TestSubscriptionValid, subscription)

	req.Data, err = cdc.MarshalJSON(NewQuerySubscriptionParams(hub.TestIDPos))
	require.Nil(t, err)

	res, _err = querySubscription(ctx, cdc, req, vpnKeeper)
	require.Nil(t, res)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)
}

func TestNewQuerySubscriptionsOfNodePrams(t *testing.T) {
	params := NewQuerySubscriptionsOfNodePrams(hub.TestIDZero)
	require.Equal(t, TestSubscriptionsOfNodeParamsZero, params)

	params = NewQuerySubscriptionsOfNodePrams(hub.TestIDPos)
	require.Equal(t, TestSubscriptionsOfNodeParamsPos, params)
}

func Test_querySubscriptionsOfNode(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var subscriptions []hub.Subscription

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", hub.QuerierRoute, QuerySubscriptionsOfNode),
		Data: []byte{},
	}

	res, _err := querySubscriptionsOfNode(ctx, cdc, req, vpnKeeper)
	require.NotNil(t, _err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)

	vpnKeeper.SetSubscription(ctx, hub.TestSubscriptionValid)
	vpnKeeper.SetSubscriptionsCountOfNode(ctx, hub.TestNodeValid.ID, 1)
	vpnKeeper.SetSubscriptionIDByNodeID(ctx, hub.TestNodeValid.ID, 0, hub.TestSubscriptionValid.ID)

	req.Data, err = cdc.MarshalJSON(NewQuerySubscriptionsOfNodePrams(hub.TestIDZero))
	require.Nil(t, err)

	res, _err = querySubscriptionsOfNode(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)

	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.Equal(t, hub.TestSubscriptionsValid, subscriptions)

	req.Data, err = cdc.MarshalJSON(NewQuerySubscriptionsOfNodePrams(hub.TestIDPos))
	require.Nil(t, err)

	res, _err = querySubscriptionsOfNode(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)

	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.NotEqual(t, hub.TestSubscriptionsValid, subscriptions)
}

func TestNewQuerySubscriptionsOfAddressParams(t *testing.T) {
	params := NewQuerySubscriptionsOfAddressParams(hub.TestAddressEmpty)
	require.Equal(t, TestSubscriptionsOfAddressParamsEmpty, params)

	params = NewQuerySubscriptionsOfAddressParams(hub.TestAddress1)
	require.Equal(t, TestSubscriptionsOfAddressParams1, params)

	params = NewQuerySubscriptionsOfAddressParams(hub.TestAddress2)
	require.Equal(t, TestSubscriptionsOfAddressParams2, params)
}

func Test_querySubscriptionsOfAddress(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var subscriptions []hub.Subscription

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", hub.QuerierRoute, QuerySubscriptionsOfAddress),
		Data: []byte{},
	}

	res, _err := querySubscriptionsOfAddress(ctx, cdc, req, vpnKeeper)
	require.NotNil(t, _err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)

	vpnKeeper.SetSubscription(ctx, hub.TestSubscriptionValid)
	vpnKeeper.SetSubscriptionsCountOfAddress(ctx, hub.TestAddress1, 1)
	vpnKeeper.SetSubscriptionIDByAddress(ctx, hub.TestAddress1, 0, hub.TestSubscriptionValid.ID)

	req.Data, err = cdc.MarshalJSON(NewQuerySubscriptionsOfAddressParams(hub.TestAddressEmpty))
	require.Nil(t, err)

	res, _err = querySubscriptionsOfAddress(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)

	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.NotEqual(t, hub.TestSubscriptionsValid, subscriptions)

	req.Data, err = cdc.MarshalJSON(NewQuerySubscriptionsOfAddressParams(hub.TestAddress1))
	require.Nil(t, err)

	res, _err = querySubscriptionsOfAddress(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)

	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.Equal(t, hub.TestSubscriptionsValid, subscriptions)

	req.Data, err = cdc.MarshalJSON(NewQuerySubscriptionsOfAddressParams(hub.TestAddress2))
	require.Nil(t, err)

	res, _err = querySubscriptionsOfAddress(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)

	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.NotEqual(t, hub.TestSubscriptionsValid, subscriptions)
}

func Test_queryAllSubscriptions(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var subscriptions []hub.Subscription

	res, _err := queryAllSubscriptions(ctx, cdc, vpnKeeper)
	require.Nil(t, _err)
	require.Equal(t, []byte("null"), res)

	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.NotEqual(t, hub.TestSubscriptionsValid, subscriptions)

	vpnKeeper.SetSubscription(ctx, hub.TestSubscriptionEmpty)
	res, _err = queryAllSubscriptions(ctx, cdc, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.NotEqual(t, hub.TestSubscriptionsValid, subscriptions)

	vpnKeeper.SetSubscription(ctx, hub.TestSubscriptionValid)
	require.Nil(t, err)

	res, _err = queryAllSubscriptions(ctx, cdc, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.Equal(t, hub.TestSubscriptionsValid, subscriptions)

	subscription := hub.TestSubscriptionValid
	subscription.ID = hub.TestIDPos
	vpnKeeper.SetSubscription(ctx, subscription)
	require.Nil(t, err)

	res, _err = queryAllSubscriptions(ctx, cdc, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.Equal(t, append(hub.TestSubscriptionsValid, subscription), subscriptions)
}

func TestNewQuerySessionsCountOfSubscriptionParams(t *testing.T) {
	params := NewQuerySessionsCountOfSubscriptionParams(hub.TestIDZero)
	require.Equal(t, TestSessionsCountOfSubscriptionParamsZero, params)

	params = NewQuerySessionsCountOfSubscriptionParams(hub.TestIDPos)
	require.Equal(t, TestSessionsCountOfSubscriptionParamsPos, params)
}

func Test_querySessionsCountOfSubscription(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var count uint64

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", hub.QuerierRoute, QuerySessionsCountOfSubscription),
		Data: []byte{},
	}

	res, _err := querySessionsCountOfSubscription(ctx, cdc, req, vpnKeeper)
	require.NotNil(t, _err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)

	vpnKeeper.SetSessionsCountOfSubscription(ctx, hub.TestIDZero, 1)
	req.Data, err = cdc.MarshalJSON(NewQuerySessionsCountOfSubscriptionParams(hub.TestIDZero))
	require.Nil(t, err)

	res, _err = querySessionsCountOfSubscription(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)

	err = cdc.UnmarshalJSON(res, &count)
	require.Nil(t, err)

	require.Equal(t, uint64(1), count)

	vpnKeeper.SetSessionsCountOfSubscription(ctx, hub.TestIDZero, 2)
	req.Data, err = cdc.MarshalJSON(NewQuerySessionsCountOfSubscriptionParams(hub.TestIDZero))
	require.Nil(t, err)

	res, _err = querySessionsCountOfSubscription(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)

	err = cdc.UnmarshalJSON(res, &count)
	require.Nil(t, err)
	require.Equal(t, uint64(2), count)

	vpnKeeper.SetSessionsCountOfSubscription(ctx, hub.TestIDPos, 1)
	req.Data, err = cdc.MarshalJSON(NewQuerySessionsCountOfSubscriptionParams(hub.TestIDPos))
	require.Nil(t, err)

	res, _err = querySessionsCountOfSubscription(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)

	err = cdc.UnmarshalJSON(res, &count)
	require.Nil(t, err)
	require.Equal(t, uint64(1), count)

	vpnKeeper.SetSessionsCountOfSubscription(ctx, hub.TestIDPos, 2)
	req.Data, err = cdc.MarshalJSON(NewQuerySessionsCountOfSubscriptionParams(hub.TestIDPos))
	require.Nil(t, err)

	res, _err = querySessionsCountOfSubscription(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)

	err = cdc.UnmarshalJSON(res, &count)
	require.Nil(t, err)
	require.Equal(t, uint64(2), count)

}
