package querier

import (
	"fmt"
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	"testing"
)

func Test_queryResolver(t *testing.T) {
	ctx, k, _, _ := keeper.CreateTestInput(t, false)
	cdc := keeper.MakeTestCodec()

	var err error
	var resolvers types.Resolvers

	var resolver = types.TestResolver
	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryResolvers),
		Data: []byte{},
	}

	res, err := queryResolvers(ctx, req, k)
	require.Nil(t, err)
	require.Equal(t, []byte("null"), res)

	k.SetResolver(ctx, types.TestResolver)

	res, err = queryResolvers(ctx, req, k)
	require.Nil(t, err)
	require.Nil(t, cdc.UnmarshalJSON(res, &resolvers))
	require.Equal(t, len(resolvers), 1)

	resolver.Owner = types.TestAddress2
	k.SetResolver(ctx, resolver)

	res, err = queryResolvers(ctx, req, k)
	require.Nil(t, err)
	require.Nil(t, cdc.UnmarshalJSON(res, &resolvers))
	require.Equal(t, len(resolvers), 2)

	req.Data, err = cdc.MarshalJSON(types.TestAddress2)
	require.Nil(t, err)
	res, err = queryResolvers(ctx, req, k)
	require.Nil(t, err)
	require.Nil(t, cdc.UnmarshalJSON(res, &resolvers))
	require.Equal(t, len(resolvers), 1)

	req.Data = nil
	res, err = queryResolvers(ctx, req, k)
	require.Nil(t, err)
	require.Nil(t, cdc.UnmarshalJSON(res, &resolvers))
	require.Equal(t, len(resolvers), 2)

}
