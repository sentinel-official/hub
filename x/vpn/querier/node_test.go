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

func Test_queryNode(t *testing.T) {
	ctx, k, _, _ := keeper.CreateTestInput(t, false)
	cdc := keeper.MakeTestCodec()
	
	var err error
	var node types.Node
	
	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryNode),
		Data: []byte{},
	}
	
	res, _err := queryNode(ctx, req, k)
	require.NotNil(t, _err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)
	
	k.SetNode(ctx, types.TestNode)
	req.Data, err = cdc.MarshalJSON(types.NewQueryNodeParams(hub.NewNodeID(0)))
	require.Nil(t, err)
	
	res, _err = queryNode(ctx, req, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)
	
	err = cdc.UnmarshalJSON(res, &node)
	require.Nil(t, err)
	require.Equal(t, types.TestNode, node)
	
	a := hub.NewNodeID(1)
	req.Data, err = cdc.MarshalJSON(types.NewQueryNodeParams(a))
	require.Nil(t, err)
	
	res, _err = queryNode(ctx, req, k)
	require.Nil(t, res)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)
}

func Test_queryNodesOfAddress(t *testing.T) {
	ctx, k, _, _ := keeper.CreateTestInput(t, false)
	cdc := keeper.MakeTestCodec()
	var err error
	var nodes []types.Node
	
	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryNodesOfAddress),
		Data: []byte{},
	}
	
	res, _err := queryNodesOfAddress(ctx, req, k)
	require.NotNil(t, _err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)
	
	k.SetNode(ctx, types.TestNode)
	k.SetNodesCountOfAddress(ctx, types.TestAddress1, 1)
	k.SetNodeIDByAddress(ctx, types.TestAddress1, 0, hub.NewNodeID(0))
	
	req.Data, err = cdc.MarshalJSON(types.NewQueryNodesOfAddressParams([]byte("")))
	require.Nil(t, err)
	
	res, _err = queryNodesOfAddress(ctx, req, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)
	
	err = cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err)
	require.NotEqual(t, []types.Node{types.TestNode}, nodes)
	
	k.SetNode(ctx, types.TestNode)
	req.Data, err = cdc.MarshalJSON(types.NewQueryNodesOfAddressParams(types.TestAddress1))
	require.Nil(t, err)
	
	res, _err = queryNodesOfAddress(ctx, req, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)
	
	err = cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err)
	require.Equal(t, []types.Node{types.TestNode}, nodes)
	
	req.Data, err = cdc.MarshalJSON(types.NewQueryNodesOfAddressParams(types.TestAddress2))
	require.Nil(t, err)
	
	res, _err = queryNodesOfAddress(ctx, req, k)
	require.NotNil(t, res)
	require.NotEqual(t, []byte(nil), res)
	
	err = cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err)
	require.NotEqual(t, []types.Node{types.TestNode}, nodes)
}

func Test_queryAllNodes(t *testing.T) {
	ctx, k, _, _ := keeper.CreateTestInput(t, false)
	cdc := keeper.MakeTestCodec()
	var err error
	var nodes []types.Node
	
	res, _err := queryAllNodes(ctx, k)
	require.Nil(t, _err)
	require.Equal(t, []byte("null"), res)
	
	err = cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err)
	require.NotEqual(t, []types.Node{types.TestNode}, nodes)
	
	err = cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err)
	require.NotEqual(t, []types.Node{types.TestNode}, nodes)
	
	k.SetNode(ctx, types.TestNode)
	require.Nil(t, err)
	
	res, _err = queryAllNodes(ctx, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)
	
	err = cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err)
	require.Equal(t, []types.Node{types.TestNode}, nodes)
	
	node := types.TestNode
	node.ID = hub.NewNodeID(1)
	k.SetNode(ctx, node)
	require.Nil(t, err)
	
	res, _err = queryAllNodes(ctx, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)
	
	err = cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err)
	require.Equal(t, append([]types.Node{types.TestNode}, node), nodes)
}
