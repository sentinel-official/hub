package querier

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"
	sdk "github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func TestNewQuerySubscriptionParams(t *testing.T) {
	params := NewQuerySubscriptionParams(sdk.TestIDZero)
	require.Equal(t, TestSubscriptionParamsZero, params)

	params = NewQuerySubscriptionParams(sdk.TestIDPos)
	require.Equal(t, TestSubscriptionParamsPos, params)
}

func Test_querySubscription(t *testing.T) {
	ctx, _, vpnkeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var subscription sdk.Subscription

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", sdk.QuerierRoute, QuerySubscription),
		Data: []byte{},
	}

	res, _err := querySubscription(ctx, cdc, req, vpnkeeper)
	require.NotNil(t, _err)
	require.Len(t, res, 0)

	vpnkeeper.SetSubscription(ctx, sdk.TestSubscriptionValid)
	req.Data, err = cdc.MarshalJSON(NewQuerySubscriptionParams(sdk.TestIDZero))
	require.Nil(t, err)

	res, _err = querySubscription(ctx, cdc, req, vpnkeeper)
	require.Nil(t, _err)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &subscription)
	require.Nil(t, err)
	require.Equal(t, sdk.TestSubscriptionValid, subscription)

	req.Data, err = cdc.MarshalJSON(NewQuerySubscriptionParams(sdk.TestIDPos))
	require.Nil(t, err)

	res, _err = querySubscription(ctx, cdc, req, vpnkeeper)
	require.Nil(t, res)
	require.Len(t, res, 0)
}

func TestNewQuerySubscriptionsOfNodePrams(t *testing.T) {
	params := NewQuerySubscriptionsOfNodePrams(sdk.TestIDZero)
	require.Equal(t, TestSubscriptionsOfNodeParamsZero, params)

	params = NewQuerySubscriptionsOfNodePrams(sdk.TestIDPos)
	require.Equal(t, TestSubscriptionsOfNodeParamsPos, params)
}

func Test_querySubscriptionsOfNode(t *testing.T) {
	ctx, _, vpnkeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var subscriptions []sdk.Subscription

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", sdk.QuerierRoute, QuerySubscriptionsOfNode),
		Data: []byte{},
	}

	res, _err := querySubscriptionsOfNode(ctx, cdc, req, vpnkeeper)
	require.NotNil(t, _err)
	require.Len(t, res, 0)

	vpnkeeper.SetSubscription(ctx, sdk.TestSubscriptionValid)
	vpnkeeper.SetSubscriptionsCountOfNode(ctx, sdk.TestNodeValid.ID, 1)
	vpnkeeper.SetSubscriptionIDByNodeID(ctx, sdk.TestNodeValid.ID, 0, sdk.TestSubscriptionValid.ID)

	req.Data, err = cdc.MarshalJSON(NewQuerySubscriptionsOfNodePrams(sdk.TestIDZero))
	require.Nil(t, err)

	res, _err = querySubscriptionsOfNode(ctx, cdc, req, vpnkeeper)
	require.Nil(t, _err)

	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.Equal(t, sdk.TestSubscriptionsValid, subscriptions)

	req.Data, err = cdc.MarshalJSON(NewQuerySubscriptionsOfNodePrams(sdk.TestIDPos))
	require.Nil(t, err)

	res, _err = querySubscriptionsOfNode(ctx, cdc, req, vpnkeeper)
	require.Nil(t, _err)

	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.NotEqual(t, sdk.TestSubscriptionsValid, subscriptions)
}

func TestNewQuerySubscriptionsOfAddressParams(t *testing.T) {
	params := NewQuerySubscriptionsOfAddressParams(sdk.TestAddressEmpty)
	require.Equal(t, TestSubscriptionsOfAddressParamsEmpty, params)

	params = NewQuerySubscriptionsOfAddressParams(sdk.TestAddress1)
	require.Equal(t, TestSubscriptionsOfAddressParams1, params)

	params = NewQuerySubscriptionsOfAddressParams(sdk.TestAddress2)
	require.Equal(t, TestSubscriptionsOfAddressParams2, params)
}

func Test_querySubscriptionsOfAddress(t *testing.T) {
	ctx, _, vpnkeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var subscriptions []sdk.Subscription

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", sdk.QuerierRoute, QuerySubscriptionsOfAddress),
		Data: []byte{},
	}

	res, _err := querySubscriptionsOfAddress(ctx, cdc, req, vpnkeeper)
	require.NotNil(t, _err)
	require.Len(t, res, 0)

	vpnkeeper.SetSubscription(ctx, sdk.TestSubscriptionValid)
	vpnkeeper.SetSubscriptionsCountOfAddress(ctx, sdk.TestAddress1, 1)
	vpnkeeper.SetSubscriptionIDByAddress(ctx, sdk.TestAddress1, 0, sdk.TestSubscriptionValid.ID)

	req.Data, err = cdc.MarshalJSON(NewQuerySubscriptionsOfAddressParams(sdk.TestAddressEmpty))
	require.Nil(t, err)

	res, _err = querySubscriptionsOfAddress(ctx, cdc, req, vpnkeeper)
	require.Nil(t, _err)

	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.NotEqual(t, sdk.TestSubscriptionsValid, subscriptions)

	req.Data, err = cdc.MarshalJSON(NewQuerySubscriptionsOfAddressParams(sdk.TestAddress1))
	require.Nil(t, err)

	res, _err = querySubscriptionsOfAddress(ctx, cdc, req, vpnkeeper)
	require.Nil(t, _err)

	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.Equal(t, sdk.TestSubscriptionsValid, subscriptions)

	req.Data, err = cdc.MarshalJSON(NewQuerySubscriptionsOfAddressParams(sdk.TestAddress2))
	require.Nil(t, err)

	res, _err = querySubscriptionsOfAddress(ctx, cdc, req, vpnkeeper)
	require.Nil(t, _err)

	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.NotEqual(t, sdk.TestSubscriptionsValid, subscriptions)
}

func Test_queryAllSubscriptions(t *testing.T) {
	ctx, _, vpnkeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var subscriptions []sdk.Subscription

	res, _err := queryAllSubscriptions(ctx, cdc, vpnkeeper)
	require.Nil(t, _err)

	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.NotEqual(t, sdk.TestSubscriptionsValid, subscriptions)

	vpnkeeper.SetSubscription(ctx, sdk.TestSubscriptionEmpty)
	res, _err = queryAllSubscriptions(ctx, cdc, vpnkeeper)
	require.Nil(t, _err)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.NotEqual(t, sdk.TestSubscriptionsValid, subscriptions)

	vpnkeeper.SetSubscription(ctx, sdk.TestSubscriptionValid)
	require.Nil(t, err)

	res, _err = queryAllSubscriptions(ctx, cdc, vpnkeeper)
	require.Nil(t, _err)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.Equal(t, sdk.TestSubscriptionsValid, subscriptions)

	subscription := sdk.TestSubscriptionValid
	subscription.ID = sdk.TestIDPos
	vpnkeeper.SetSubscription(ctx, subscription)
	require.Nil(t, err)

	res, _err = queryAllSubscriptions(ctx, cdc, vpnkeeper)
	require.Nil(t, _err)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &subscriptions)
	require.Nil(t, err)
	require.Equal(t, append(sdk.TestSubscriptionsValid, subscription), subscriptions)
}

func TestNewQuerySessionsCountOfSubscriptionParams(t *testing.T) {
	params := NewQuerySessionsCountOfSubscriptionParams(sdk.TestIDZero)
	require.Equal(t, TestSessionsCountOfSubscriptionParamsZero, params)

	params = NewQuerySessionsCountOfSubscriptionParams(sdk.TestIDPos)
	require.Equal(t, TestSessionsCountOfSubscriptionParamsPos, params)
}

func Test_querySessionsCountOfSubscription(t *testing.T) {
	ctx, _, vpnkeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var count uint64

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", sdk.QuerierRoute, QuerySessionsCountOfSubscription),
		Data: []byte{},
	}

	res, _err := querySessionsCountOfSubscription(ctx, cdc, req, vpnkeeper)
	require.NotNil(t, _err)
	require.Len(t, res, 0)

	vpnkeeper.SetSessionsCountOfSubscription(ctx, sdk.TestIDZero, 1)
	req.Data, err = cdc.MarshalJSON(NewQuerySessionsCountOfSubscriptionParams(sdk.TestIDZero))
	require.Nil(t, err)

	res, _err = querySessionsCountOfSubscription(ctx, cdc, req, vpnkeeper)
	require.Nil(t, _err)

	err = cdc.UnmarshalJSON(res, &count)
	require.Nil(t, err)

	require.Equal(t, uint64(1), count)

	vpnkeeper.SetSessionsCountOfSubscription(ctx, sdk.TestIDZero, 2)
	req.Data, err = cdc.MarshalJSON(NewQuerySessionsCountOfSubscriptionParams(sdk.TestIDZero))
	require.Nil(t, err)

	res, _err = querySessionsCountOfSubscription(ctx, cdc, req, vpnkeeper)
	require.Nil(t, _err)

	err = cdc.UnmarshalJSON(res, &count)
	require.Nil(t, err)
	require.Equal(t, uint64(2), count)

	vpnkeeper.SetSessionsCountOfSubscription(ctx, sdk.TestIDPos, 1)
	req.Data, err = cdc.MarshalJSON(NewQuerySessionsCountOfSubscriptionParams(sdk.TestIDPos))
	require.Nil(t, err)

	res, _err = querySessionsCountOfSubscription(ctx, cdc, req, vpnkeeper)
	require.Nil(t, _err)

	err = cdc.UnmarshalJSON(res, &count)
	require.Nil(t, err)
	require.Equal(t, uint64(1), count)

	vpnkeeper.SetSessionsCountOfSubscription(ctx, sdk.TestIDPos, 2)
	req.Data, err = cdc.MarshalJSON(NewQuerySessionsCountOfSubscriptionParams(sdk.TestIDPos))
	require.Nil(t, err)

	res, _err = querySessionsCountOfSubscription(ctx, cdc, req, vpnkeeper)
	require.Nil(t, _err)

	err = cdc.UnmarshalJSON(res, &count)
	require.Nil(t, err)
	require.Equal(t, uint64(2), count)

}
