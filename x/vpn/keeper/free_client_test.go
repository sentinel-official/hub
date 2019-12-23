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
	
	address, found := k.GetFreeClientOfNode(ctx, hub.NewNodeID(0), types.TestAddress2)
	require.False(t, found)
	require.Nil(t, address)
	
	node, found := k.GetFreeNodeOfClient(ctx, types.TestAddress2, hub.NewNodeID(0))
	require.False(t, found)
	require.Nil(t, node)
	
	k.SetFreeNodeOfClient(ctx, types.TestAddress2, hub.NewNodeID(0))
	k.SetFreeClientOfNode(ctx, hub.NewNodeID(0), types.TestAddress2)
	
	address, found = k.GetFreeClientOfNode(ctx, hub.NewNodeID(0), types.TestAddress2)
	require.True(t, found)
	require.Equal(t, types.TestAddress2, address)
	
	node, found = k.GetFreeNodeOfClient(ctx, types.TestAddress2, hub.NewNodeID(0))
	require.True(t, found)
	require.Equal(t, hub.NewNodeID(0), node)
	
	k.SetFreeNodeOfClient(ctx, types.TestAddress3, hub.NewNodeID(0))
	k.SetFreeClientOfNode(ctx, hub.NewNodeID(0), types.TestAddress3)
	
	address, found = k.GetFreeClientOfNode(ctx, hub.NewNodeID(0), types.TestAddress3)
	require.True(t, found)
	require.Equal(t, types.TestAddress3, address)
	
	freeclients = k.GetFreeClientsOfNode(ctx, hub.NewNodeID(0))
	require.Equal(t, 2, len(freeclients))
	
	k.RemoveFreeClientOfNode(ctx, hub.NewNodeID(0), types.TestAddress2)
	nodes = k.GetFreeNodesOfClient(ctx, types.TestAddress2)
	require.Equal(t, 0, len(nodes))
	
	nodes = k.GetFreeNodesOfClient(ctx, types.TestAddress3)
	require.Equal(t, 1, len(nodes))
}
