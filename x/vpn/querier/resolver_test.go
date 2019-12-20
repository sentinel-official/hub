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

func Test_queryResolver(t *testing.T) {
	ctx, k, _, _ := keeper.CreateTestInput(t, false)
	cdc := keeper.MakeTestCodec()
	
	var err error
	var resolvers types.Resolvers
	
	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryResolvers),
		Data: []byte{},
	}
	
	res, err := queryResolvers(ctx, req, k)
	require.Nil(t, err)
	require.Equal(t, []byte("null"), res)
	
	k.SetResolver(ctx, types.TestResolver)
	
	resolverID, err := hub.NewResolverIDFromString("reso1")
	require.Nil(t, err)
	
	req.Data = cdc.MustMarshalJSON(resolverID)
	res, err = queryResolvers(ctx, req, k)
	require.NotNil(t, err)
	require.Equal(t, types.ErrorResolverDoesNotExist(), err)
	
	req.Data = cdc.MustMarshalJSON(types.TestResolver.ID)
	res, err = queryResolvers(ctx, req, k)
	require.Nil(t, err)
	require.Nil(t, cdc.UnmarshalJSON(res, &resolvers))
	require.Equal(t, len(resolvers), 1)
	
	req.Data, err = cdc.MarshalJSON(hub.NewResolverID(2))
	require.Nil(t, err)
	res, err = queryResolvers(ctx, req, k)
	require.NotNil(t, err)
	require.Equal(t, types.ErrorResolverDoesNotExist(), err)
	
	resolver := types.TestResolver
	resolver.ID = hub.NewResolverID(1)
	k.SetResolver(ctx, resolver)
	req.Data = nil
	res, err = queryResolvers(ctx, req, k)
	require.Nil(t, err)
	require.Nil(t, cdc.UnmarshalJSON(res, &resolvers))
	require.Equal(t, len(resolvers), 2)
	
	req.Path = fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryNodesOfResolver)
	req.Data = []byte{}
	
	k.SetNodeOfResolver(ctx, hub.NewResolverID(1), hub.NewNodeID(1))
	k.SetResolverOfNode(ctx, hub.NewNodeID(1), hub.NewResolverID(1))
	
	var nodes []hub.NodeID
	req.Data = types.TestAddress2
	res, err = queryNodesOfResolver(ctx, req, k)
	require.NotNil(t, err)
	
	req.Data = cdc.MustMarshalJSON(types.NewQueryNodesOfResolverPrams(hub.NewResolverID(1)))
	res, err = queryNodesOfResolver(ctx, req, k)
	require.Nil(t, err)
	
	_ = cdc.UnmarshalJSON(res, &nodes)
	require.Equal(t, hub.NewNodeID(1), nodes[0])
	
	req.Path = fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryResolversOfNode)
	req.Data = types.TestAddress1
	res, err = queryResolversOfNode(ctx, req, k)
	require.NotNil(t, err)
	
	var resolverIDs []hub.ResolverID
	req.Data = cdc.MustMarshalJSON(types.NewQueryResolversOfNodeParams(hub.NewNodeID(1)))
	res, err = queryResolversOfNode(ctx, req, k)
	require.Nil(t, err)
	
	_ = cdc.UnmarshalJSON(res, &resolverIDs)
	require.Equal(t, hub.NewResolverID(1), resolverIDs[0])
	
}
