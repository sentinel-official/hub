package querier

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"
	sdk "github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func TestNewQueryNodeParams(t *testing.T) {
	params := NewQueryNodeParams(sdk.TestIDZero)
	require.Equal(t, TestNodeParamsZero, params)

	params = NewQueryNodeParams(sdk.TestIDPos)
	require.Equal(t, TestNodeParamsPos, params)
}

func Test_queryNode(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var node sdk.Node

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", sdk.QuerierRoute, QueryNode),
		Data: []byte{},
	}

	res, _err := queryNode(ctx, cdc, req, vpnKeeper)
	require.NotNil(t, _err)
	require.Equal(t,[]byte(nil),res)
	require.Len(t, res, 0)

	vpnKeeper.SetNode(ctx, sdk.TestNodeValid)
	req.Data, err = cdc.MarshalJSON(NewQueryNodeParams(sdk.TestIDZero))
	require.Nil(t, err)

	res, _err = queryNode(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t,[]byte(nil),res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &node)
	require.Nil(t, err)
	require.Equal(t, sdk.TestNodeValid, node)

	req.Data, err = cdc.MarshalJSON(NewQueryNodeParams(sdk.TestIDPos))
	require.Nil(t, err)

	res, _err = queryNode(ctx, cdc, req, vpnKeeper)
	require.Nil(t, res)
	require.Equal(t,[]byte(nil),res)
	require.Len(t, res, 0)
}

func TestNewQueryNodesOfAddressParams(t *testing.T) {
	params := NewQueryNodesOfAddressParams(sdk.TestAddressEmpty)
	require.Equal(t, TestNodeOfAddressParamsEmpty, params)

	params = NewQueryNodesOfAddressParams(sdk.TestAddress1)
	require.Equal(t, TestNodeOfAddressParams1, params)

	params = NewQueryNodesOfAddressParams(sdk.TestAddress2)
	require.Equal(t, TestNodeOfAddressParams2, params)
}

func Test_queryNodesOfAddress(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var nodes []sdk.Node

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", sdk.QuerierRoute, QueryNodesOfAddress),
		Data: []byte{},
	}

	res, _err := queryNodesOfAddress(ctx, cdc, req, vpnKeeper)
	require.NotNil(t, _err)
	require.Equal(t,[]byte(nil),res)
	require.Len(t, res, 0)

	vpnKeeper.SetNode(ctx, sdk.TestNodeValid)
	vpnKeeper.SetNodesCountOfAddress(ctx, sdk.TestAddress1, 1)
	vpnKeeper.SetNodeIDByAddress(ctx, sdk.TestAddress1, 0, sdk.TestIDZero)

	req.Data, err = cdc.MarshalJSON(NewQueryNodesOfAddressParams(sdk.TestAddressEmpty))
	require.Nil(t, err)

	res, _err = queryNodesOfAddress(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t,[]byte(nil),res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err)
	require.NotEqual(t, sdk.TestNodesValid, nodes)

	vpnKeeper.SetNode(ctx, sdk.TestNodeValid)
	req.Data, err = cdc.MarshalJSON(NewQueryNodesOfAddressParams(sdk.TestAddress1))
	require.Nil(t, err)

	res, _err = queryNodesOfAddress(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t,[]byte(nil),res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err)
	require.Equal(t, sdk.TestNodesValid, nodes)

	req.Data, err = cdc.MarshalJSON(NewQueryNodesOfAddressParams(sdk.TestAddress2))
	require.Nil(t, err)

	res, _err = queryNodesOfAddress(ctx, cdc, req, vpnKeeper)
	require.NotNil(t, res)
	require.NotEqual(t,[]byte(nil),res)

	err = cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err)
	require.NotEqual(t, sdk.TestNodesValid, nodes)
}

func Test_queryAllNodes(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var nodes []sdk.Node

	res, _err := queryAllNodes(ctx, cdc, vpnKeeper)
	require.Nil(t, _err)
	require.Equal(t,[]byte("null"),res)

	err = cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err)
	require.NotEqual(t, sdk.TestNodesValid, nodes)

	vpnKeeper.SetNode(ctx, sdk.TestNodeEmpty)
	res, _err = queryAllNodes(ctx, cdc, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t,[]byte(nil),res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err)
	require.NotEqual(t, sdk.TestNodesValid, nodes)

	vpnKeeper.SetNode(ctx, sdk.TestNodeValid)
	require.Nil(t, err)

	res, _err = queryAllNodes(ctx, cdc, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t,[]byte(nil),res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err)
	require.Equal(t, sdk.TestNodesValid, nodes)

	node := sdk.TestNodeValid
	node.ID = sdk.TestIDPos
	vpnKeeper.SetNode(ctx, node)
	require.Nil(t, err)

	res, _err = queryAllNodes(ctx, cdc, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t,[]byte(nil),res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err)
	require.Equal(t, append(sdk.TestNodesValid, node), nodes)
}
