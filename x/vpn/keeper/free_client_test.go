package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func TestKeeper_FreeClients(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	freeclients := k.GetFreeClientsOfNode(ctx, hub.NewNodeID(1))
	require.Equal(t, 0, len(freeclients))

	nodes := k.GetFreeNodesOfClient(ctx, types.TestAddress2)
	require.Equal(t, 0, len(nodes))

	address, found := k.GetFreeClientOfNode(ctx, hub.NewNodeID(1), types.TestAddress1)
	require.False(t, found)
	require.Nil(t, address)

	node, found := k.GetFreeNodeOfClient(ctx, types.TestAddress1, hub.NewNodeID(1))
	require.False(t, found)
	require.Nil(t, node)

	client := types.TestFreeClient
	k.SetFreeNodeOfClient(ctx, client)
	k.SetFreeClientOfNode(ctx, client)

	address, found = k.GetFreeClientOfNode(ctx, client.NodeID, client.Client)
	require.True(t, found)
	require.Equal(t, types.TestAddress1, address)

	node, found = k.GetFreeNodeOfClient(ctx, client.Client, client.NodeID)
	require.True(t, found)
	require.Equal(t, client.NodeID, node)

	client.Client = types.TestAddress2
	k.SetFreeNodeOfClient(ctx, client)
	k.SetFreeClientOfNode(ctx, client)

	address, found = k.GetFreeClientOfNode(ctx, client.NodeID, client.Client)
	require.True(t, found)
	require.Equal(t, types.TestAddress2, address)

	freeclients = k.GetFreeClientsOfNode(ctx, client.NodeID)
	require.Equal(t, 2, len(freeclients))

	k.RemoveFreeClient(ctx, client.NodeID, types.TestAddress1)
	nodes = k.GetFreeNodesOfClient(ctx, types.TestAddress1)
	require.Equal(t, 0, len(nodes))

	nodes = k.GetFreeNodesOfClient(ctx, types.TestAddress2)
	require.Equal(t, 1, len(nodes))
}
