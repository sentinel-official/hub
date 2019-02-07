package keeper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func TestKeeper_NodeDetails(t *testing.T) {
	keeper, ctx := CreateTestInput()

	for _, vpnDetails := range ParamsOfNodeDetails() {
		err := keeper.SetNodeDetails(ctx, &vpnDetails)
		require.Nil(t, err)
	}

	res, err := keeper.GetNodeDetails(ctx, ParamsOfNodeDetails()[0].ID)
	require.Nil(t, err)
	require.Equal(t, sdkTypes.NewNodeID("new-node-id-1"), res.ID)

	nilRes, err := keeper.GetNodeDetails(ctx, "data/0")
	require.Nil(t, err)
	require.Nil(t, nilRes)
}

func TestKeeper_NodesCount(t *testing.T) {
	keeper, ctx := CreateTestInput()

	err := keeper.SetNodesCount(ctx, sdkTypes.NodeAddress1, uint64(10))
	require.Nil(t, err)

	res, err := keeper.GetNodesCount(ctx, sdkTypes.NodeAddress2)
	require.Equal(t, uint64(0), res)
	require.Nil(t, err)

	res, err = keeper.GetNodesCount(ctx, sdkTypes.NodeAddress1)
	assert.NotNil(t, res)
	assert.Equal(t, uint64(10), res)
}

func TestKeeper_GetNodes(t *testing.T) {
	keeper, ctx := CreateTestInput()

	for _, vpnDetails := range ParamsOfNodeDetails() {
		err := keeper.SetNodeDetails(ctx, &vpnDetails)
		require.Nil(t, err)
	}

	res, err := keeper.GetNodes(ctx)
	require.Nil(t, err)
	require.Equal(t, sdkTypes.NewNodeID("new-node-id-1"), res[1].ID)
	require.NotEqual(t, []*sdkTypes.NodeDetails{}, res)
}

func TestKeeper_GetNodesOfOwner(t *testing.T) {
	keeper, ctx := CreateTestInput()

	for _, vpnDetails := range ParamsOfNodeDetails() {
		_, err := keeper.AddNode(ctx, &vpnDetails)
		require.Nil(t, err)
	}

	res2, err := keeper.GetNodesOfOwner(ctx, sdkTypes.NodeAddress2)
	require.Nil(t, err)
	require.Equal(t, sdkTypes.NodeAddress2, res2[0].Owner)
}

func TestKeeper_NodeIDAtHeight(t *testing.T) {
	keeper, ctx := CreateTestInput()

	for _, vpnDetails := range ParamsOfNodeDetails() {
		err := keeper.SetNodeDetails(ctx, &vpnDetails)
		require.Nil(t, err)

		err = keeper.AddActiveNodeIDAtHeight(ctx, 6, vpnDetails.ID)
		require.Nil(t, err)
	}

	res, err := keeper.GetActiveNodeIDsAtHeight(ctx, 6)

	err = keeper.AddActiveNodeIDAtHeight(ctx, 6, sdkTypes.NewNodeID("str1"))
	require.Equal(t, nil, err)

	res2, err := keeper.GetNodeDetails(ctx, res[1])
	require.Nil(t, err)
	require.Equal(t, int64(8), res2.StatusAtHeight)

	err = keeper.AddActiveNodeIDAtHeight(ctx, 6, sdkTypes.NewNodeID("str1"))
	require.Equal(t, nil, err)

	err = keeper.RemoveActiveNodeIDAtHeight(ctx, 6, sdkTypes.NewNodeID("str1"))
	require.Nil(t, err)

	res, err = keeper.GetActiveNodeIDsAtHeight(ctx, 6)
	require.Equal(t, 3, res.Len())
}
