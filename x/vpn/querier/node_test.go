package querier

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func TestNewQueryNodeParams(t *testing.T) {
	cdc := keeper.TestMakeCodec()
	ctx, _, vpnKeeper, _, _ := keeper.TestCreateInput()
	querier := NewQuerier(vpnKeeper, cdc)

	data := NewQueryNodeParams(types.TestNodeIDValid)
	require.NotNil(t, data)

	query := abciTypes.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", types.QuerierRoute, QueryNode),
		Data: cdc.MustMarshalJSON(data),
	}
	res1, err := querier(ctx, []string{QueryNode}, query)
	require.Nil(t, err)
	require.Nil(t, res1)

	tags, err := vpnKeeper.AddNode(ctx, &keeper.TestNodeValid)
	require.Nil(t, err)
	require.NotNil(t, tags)

	res2, err := querier(ctx, []string{QueryNode}, query)
	require.Nil(t, err)
	require.NotNil(t, res2)

	var node types.NodeDetails
	err1 := cdc.UnmarshalJSON(res2, &node)
	require.Nil(t, err1)
	require.Equal(t, keeper.TestNodeValid, node)

	query.Data = cdc.MustMarshalJSON(NewQueryNodeParams(types.TestNodeIDEmpty))
	res3, err := querier(ctx, []string{QueryNode}, query)
	require.Nil(t, err)
	require.Nil(t, res3)

	query.Data = cdc.MustMarshalJSON("")
	res4, err := querier(ctx, []string{QueryNode}, query)
	require.Equal(t, err, types.ErrorUnmarshal())
	require.Nil(t, res4)
}

func TestNewQueryNodesOfOwnerParams(t *testing.T) {
	cdc := keeper.TestMakeCodec()
	ctx, _, vpnKeeper, _, _ := keeper.TestCreateInput()
	querier := NewQuerier(vpnKeeper, cdc)

	data := NewQueryNodesOfOwnerParams(types.TestAddress1)
	require.NotNil(t, data)

	query := abciTypes.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", types.QuerierRoute, QueryNodesOfOwner),
		Data: cdc.MustMarshalJSON(data),
	}
	res1, err := querier(ctx, []string{QueryNodesOfOwner}, query)
	require.Nil(t, err)
	require.Equal(t, []byte("null"), res1)

	tags, err := vpnKeeper.AddNode(ctx, &keeper.TestNodeValid)
	require.Nil(t, err)
	require.NotNil(t, tags)

	res2, err := querier(ctx, []string{QueryNodesOfOwner}, query)
	require.Nil(t, err)
	require.NotNil(t, res2)

	var nodes []types.NodeDetails
	err1 := cdc.UnmarshalJSON(res2, &nodes)
	require.Nil(t, err1)
	require.Equal(t, []types.NodeDetails{keeper.TestNodeValid}, nodes)

	query.Data = cdc.MustMarshalJSON(NewQueryNodeParams(types.TestNodeIDEmpty))
	res3, err := querier(ctx, []string{QueryNodesOfOwner}, query)
	require.Nil(t, err)
	require.Equal(t, []byte("null"), res3)

	query.Data = cdc.MustMarshalJSON("")
	res4, err := querier(ctx, []string{QueryNodesOfOwner}, query)
	require.Equal(t, err, types.ErrorUnmarshal())
	require.Nil(t, res4)
}

func TestNewQuerier(t *testing.T) {
	cdc := keeper.TestMakeCodec()
	ctx, _, vpnKeeper, _, _ := keeper.TestCreateInput()
	querier := NewQuerier(vpnKeeper, cdc)

	data := NewQueryNodeParams(types.TestNodeIDValid)
	require.NotNil(t, data)

	query := abciTypes.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", types.QuerierRoute, "query_type"),
		Data: cdc.MustMarshalJSON(data),
	}
	res1, err := querier(ctx, []string{"query_type"}, query)
	require.Equal(t, err, types.ErrorInvalidQueryType("query_type"))
	require.Nil(t, res1)
}
