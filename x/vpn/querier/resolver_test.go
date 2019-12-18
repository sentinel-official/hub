package querier

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

	req.Data = cdc.MustMarshalJSON(sdk.AccAddress("0x1234"))
	res, err = queryResolvers(ctx, req, k)
	require.NotNil(t, err)
	require.Equal(t, types.ErrorUnmarshal(), err)

	req.Data = cdc.MustMarshalJSON(types.TestResolver.Owner)
	res, err = queryResolvers(ctx, req, k)
	require.Nil(t, err)
	require.Nil(t, cdc.UnmarshalJSON(res, &resolvers))
	require.Equal(t, len(resolvers), 1)

	res, err = queryResolvers(ctx, req, k)
	require.Nil(t, err)
	require.Nil(t, cdc.UnmarshalJSON(res, &resolvers))
	require.Equal(t, len(resolvers), 1)

	req.Data, err = cdc.MarshalJSON(types.TestAddress2)
	require.Nil(t, err)
	res, err = queryResolvers(ctx, req, k)
	require.NotNil(t, err)
	require.Equal(t, types.ErrorResolverDoesNotExist(), err)

	resolver := types.TestResolver
	resolver.Owner = types.TestAddress2
	k.SetResolver(ctx, resolver)
	req.Data = nil
	res, err = queryResolvers(ctx, req, k)
	require.Nil(t, err)
	require.Nil(t, cdc.UnmarshalJSON(res, &resolvers))
	require.Equal(t, len(resolvers), 2)

	req.Path = fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryNodesOfResolver)
	req.Data = []byte{}

	k.SetNodeOfResolver(ctx, types.TestAddress2, hub.NewNodeID(1))
	k.SetResolverOfNode(ctx, hub.NewNodeID(1), types.TestAddress2)

	var nodes []hub.NodeID
	req.Data = types.TestAddress2
	res, err = queryNodesOfResolver(ctx, req, k)
	require.NotNil(t, err)

	req.Data = cdc.MustMarshalJSON(types.NewQueryNodesOfResolverPrams(types.TestAddress2))
	res, err = queryNodesOfResolver(ctx, req, k)
	require.Nil(t, err)

	_ = cdc.UnmarshalJSON(res, &nodes)
	require.Equal(t, hub.NewNodeID(1), nodes[0])

	req.Path = fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryResolversOfNode)
	req.Data = types.TestAddress1
	res, err = queryResolversOfNode(ctx, req, k)
	require.NotNil(t, err)

	var addresses []sdk.AccAddress
	req.Data = cdc.MustMarshalJSON(types.NewQueryResolversOfNodeParams(hub.NewNodeID(1)))
	res, err = queryResolversOfNode(ctx, req, k)
	require.Nil(t, err)

	_ = cdc.UnmarshalJSON(res, &addresses)
	require.Equal(t, types.TestAddress2, addresses[0])

}
