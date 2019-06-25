package querier

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/sentinel-hub/x/vpn/keeper"
	hub "github.com/sentinel-official/sentinel-hub/x/vpn/types"
)

func TestNewQueryNodeParams(t *testing.T) {
	params := NewQueryNodeParams(hub.TestIDZero)
	require.Equal(t, TestNodeParamsZero, params)

	params = NewQueryNodeParams(hub.TestIDPos)
	require.Equal(t, TestNodeParamsPos, params)
}

func Test_queryNode(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var node hub.Node

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", hub.QuerierRoute, QueryNode),
		Data: []byte{},
	}

	res, _err := queryNode(ctx, cdc, req, vpnKeeper)
	require.NotNil(t, _err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)

	vpnKeeper.SetNode(ctx, hub.TestNodeValid)
	req.Data, err = cdc.MarshalJSON(NewQueryNodeParams(hub.TestIDZero))
	require.Nil(t, err)

	res, _err = queryNode(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &node)
	require.Nil(t, err)
	require.Equal(t, hub.TestNodeValid, node)

	req.Data, err = cdc.MarshalJSON(NewQueryNodeParams(hub.TestIDPos))
	require.Nil(t, err)

	res, _err = queryNode(ctx, cdc, req, vpnKeeper)
	require.Nil(t, res)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)
}

func TestNewQueryNodesOfAddressParams(t *testing.T) {
	params := NewQueryNodesOfAddressParams(hub.TestAddressEmpty)
	require.Equal(t, TestNodeOfAddressParamsEmpty, params)

	params = NewQueryNodesOfAddressParams(hub.TestAddress1)
	require.Equal(t, TestNodeOfAddressParams1, params)

	params = NewQueryNodesOfAddressParams(hub.TestAddress2)
	require.Equal(t, TestNodeOfAddressParams2, params)
}

func Test_queryNodesOfAddress(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var nodes []hub.Node

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", hub.QuerierRoute, QueryNodesOfAddress),
		Data: []byte{},
	}

	res, _err := queryNodesOfAddress(ctx, cdc, req, vpnKeeper)
	require.NotNil(t, _err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)

	vpnKeeper.SetNode(ctx, hub.TestNodeValid)
	vpnKeeper.SetNodesCountOfAddress(ctx, hub.TestAddress1, 1)
	vpnKeeper.SetNodeIDByAddress(ctx, hub.TestAddress1, 0, hub.TestIDZero)

	req.Data, err = cdc.MarshalJSON(NewQueryNodesOfAddressParams(hub.TestAddressEmpty))
	require.Nil(t, err)

	res, _err = queryNodesOfAddress(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err)
	require.NotEqual(t, hub.TestNodesValid, nodes)

	vpnKeeper.SetNode(ctx, hub.TestNodeValid)
	req.Data, err = cdc.MarshalJSON(NewQueryNodesOfAddressParams(hub.TestAddress1))
	require.Nil(t, err)

	res, _err = queryNodesOfAddress(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err)
	require.Equal(t, hub.TestNodesValid, nodes)

	req.Data, err = cdc.MarshalJSON(NewQueryNodesOfAddressParams(hub.TestAddress2))
	require.Nil(t, err)

	res, _err = queryNodesOfAddress(ctx, cdc, req, vpnKeeper)
	require.NotNil(t, res)
	require.NotEqual(t, []byte(nil), res)

	err = cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err)
	require.NotEqual(t, hub.TestNodesValid, nodes)
}

func Test_queryAllNodes(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var nodes []hub.Node

	res, _err := queryAllNodes(ctx, cdc, vpnKeeper)
	require.Nil(t, _err)
	require.Equal(t, []byte("null"), res)

	err = cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err)
	require.NotEqual(t, hub.TestNodesValid, nodes)

	vpnKeeper.SetNode(ctx, hub.TestNodeEmpty)
	res, _err = queryAllNodes(ctx, cdc, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err)
	require.NotEqual(t, hub.TestNodesValid, nodes)

	vpnKeeper.SetNode(ctx, hub.TestNodeValid)
	require.Nil(t, err)

	res, _err = queryAllNodes(ctx, cdc, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err)
	require.Equal(t, hub.TestNodesValid, nodes)

	node := hub.TestNodeValid
	node.ID = hub.TestIDPos
	vpnKeeper.SetNode(ctx, node)
	require.Nil(t, err)

	res, _err = queryAllNodes(ctx, cdc, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err)
	require.Equal(t, append(hub.TestNodesValid, node), nodes)
}
