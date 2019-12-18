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

func Test_queryFreeClients(t *testing.T) {
	ctx, k, _, _ := keeper.CreateTestInput(t, false)
	cdc := keeper.MakeTestCodec()

	var err error
	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryFreeClientsOfNode),
		Data: []byte{},
	}

	clients, err := queryFreeClientsOfNode(ctx, req, k)
	require.Nil(t, clients)
	require.Equal(t, types.ErrorUnmarshal(), err)

	req.Path = fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryFreeNodesOfClient)
	res, err := queryFreeNodesOfClient(ctx, req, k)
	require.NotNil(t, err)
	require.Equal(t, types.ErrorUnmarshal(), err)

	k.SetFreeClientOfNode(ctx, hub.NewNodeID(0), types.TestAddress2)
	k.SetFreeNodeOfClient(ctx, types.TestAddress2, hub.NewNodeID(0))

	req.Path = fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryFreeClientsOfNode)
	req.Data = cdc.MustMarshalJSON(types.NewQueryFreeClientsOfNodeParams(hub.NewNodeID(0)))
	res, err = queryFreeClientsOfNode(ctx, req, k)
	require.Nil(t, err)

	var freeClient []sdk.AccAddress
	_ = cdc.UnmarshalJSON(res, &freeClient)
	require.Equal(t, types.TestAddress2, freeClient[0])

	req.Path = fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryFreeNodesOfClient)
	res, err = queryFreeNodesOfClient(ctx, req, k)
	require.Nil(t, err)

	var nodes []hub.NodeID
	_ = cdc.UnmarshalJSON(res, &nodes)
	require.Equal(t, hub.NewNodeID(0), nodes[0])

}
