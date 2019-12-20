package keeper

import (
	"testing"
	
	"github.com/stretchr/testify/require"
	
	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func TestKeeper_SetNodesCount(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)
	
	count := k.GetNodesCount(ctx)
	require.Equal(t, uint64(0), count)
	
	k.SetNodesCount(ctx, 0)
	count = k.GetNodesCount(ctx)
	require.Equal(t, uint64(0), count)
	
	k.SetNodesCount(ctx, 1)
	count = k.GetNodesCount(ctx)
	require.Equal(t, uint64(1), count)
	
	k.SetNodesCount(ctx, 2)
	count = k.GetNodesCount(ctx)
	require.Equal(t, uint64(2), count)
}

func TestKeeper_GetNodesCount(t *testing.T) {
	TestKeeper_SetNodesCount(t)
}

func TestKeeper_SetNode(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)
	
	_, found := k.GetNode(ctx, hub.NewNodeID(0))
	require.Equal(t, false, found)
	
	k.SetNode(ctx, types.TestNode)
	node, found := k.GetNode(ctx, types.TestNode.ID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestNode, node)
}

func TestKeeper_GetNode(t *testing.T) {
	TestKeeper_SetNode(t)
}

func TestKeeper_SetNodesCountOfAddress(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)
	
	count := k.GetNodesCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(0), count)
	
	k.SetNodesCountOfAddress(ctx, []byte(""), 1)
	count = k.GetNodesCountOfAddress(ctx, []byte(""))
	require.Equal(t, uint64(1), count)
	
	k.SetNodesCountOfAddress(ctx, types.TestAddress1, 1)
	count = k.GetNodesCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(1), count)
	
	k.SetNodesCountOfAddress(ctx, types.TestAddress2, 2)
	count = k.GetNodesCountOfAddress(ctx, types.TestAddress2)
	require.Equal(t, uint64(2), count)
	
	count = k.GetNodesCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(1), count)
	
	k.SetNodesCountOfAddress(ctx, types.TestAddress1, 0)
	count = k.GetNodesCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(0), count)
	
	count = k.GetNodesCountOfAddress(ctx, types.TestAddress2)
	require.Equal(t, uint64(2), count)
	
	k.SetNodesCountOfAddress(ctx, types.TestAddress2, 0)
	count = k.GetNodesCountOfAddress(ctx, types.TestAddress2)
	require.Equal(t, uint64(0), count)
}

func TestKeeper_GetNodesCountOfAddress(t *testing.T) {
	TestKeeper_SetNodesCountOfAddress(t)
}

func TestKeeper_SetNodeIDByAddress(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)
	
	id, found := k.GetNodeIDByAddress(ctx, types.TestAddress1, 0)
	require.Equal(t, false, found)
	require.Equal(t, hub.NewNodeID(0), id)
	
	k.SetNodeIDByAddress(ctx, []byte(""), 0, hub.NewNodeID(0))
	id, found = k.GetNodeIDByAddress(ctx, []byte(""), 0)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewNodeID(0), id)
	
	k.SetNodeIDByAddress(ctx, types.TestAddress1, 0, hub.NewNodeID(0))
	id, found = k.GetNodeIDByAddress(ctx, types.TestAddress1, 0)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewNodeID(0), id)
	
	k.SetNodeIDByAddress(ctx, types.TestAddress2, 0, hub.NewNodeID(0))
	id, found = k.GetNodeIDByAddress(ctx, types.TestAddress2, 0)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewNodeID(0), id)
	
	k.SetNodeIDByAddress(ctx, types.TestAddress1, 1, hub.NewNodeID(1))
	id, found = k.GetNodeIDByAddress(ctx, types.TestAddress1, 1)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewNodeID(1), id)
	
	k.SetNodeIDByAddress(ctx, types.TestAddress2, 1, hub.NewNodeID(1))
	id, found = k.GetNodeIDByAddress(ctx, types.TestAddress2, 1)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewNodeID(1), id)
}

func TestKeeper_GetNodeIDByAddress(t *testing.T) {
	TestKeeper_SetNodeIDByAddress(t)
}

func TestKeeper_SetActiveNodeIDs(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)
	
	ids := k.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, hub.IDs(nil), ids)
	
	k.SetActiveNodeIDs(ctx, 1, hub.IDs{})
	ids = k.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, hub.IDs(nil), ids)
	
	k.SetActiveNodeIDs(ctx, 1, hub.IDs(nil))
	ids = k.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, hub.IDs(nil), ids)
	
	k.SetActiveNodeIDs(ctx, 1, hub.IDs{hub.NewNodeID(0)})
	ids = k.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, hub.IDs{hub.NewNodeID(0)}, ids)
	
	k.SetActiveNodeIDs(ctx, 1, hub.IDs{hub.NewNodeID(0)}.Append(hub.NewNodeID(0)))
	ids = k.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, hub.IDs{hub.NewNodeID(0)}.Append(hub.NewNodeID(0)).Sort(), ids)
	
	k.SetActiveNodeIDs(ctx, 1, hub.IDs{})
	ids = k.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, hub.IDs(nil), ids)
	
	ids = k.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, hub.IDs(nil), ids)
	
	k.SetActiveNodeIDs(ctx, 2, hub.IDs{})
	ids = k.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, hub.IDs(nil), ids)
	
	k.SetActiveNodeIDs(ctx, 2, hub.IDs(nil))
	ids = k.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, hub.IDs(nil), ids)
	
	k.SetActiveNodeIDs(ctx, 2, hub.IDs{hub.NewNodeID(0)})
	ids = k.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, hub.IDs{hub.NewNodeID(0)}, ids)
	
	k.SetActiveNodeIDs(ctx, 2, hub.IDs{hub.NewNodeID(0)}.Append(hub.NewNodeID(0)))
	ids = k.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, hub.IDs{hub.NewNodeID(0)}.Append(hub.NewNodeID(0)).Sort(), ids)
	
	k.SetActiveNodeIDs(ctx, 2, hub.IDs{})
	ids = k.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, hub.IDs(nil), ids)
	
}

func TestKeeper_GetActiveNodeIDs(t *testing.T) {
	TestKeeper_SetActiveNodeIDs(t)
}

func TestKeeper_DeleteActiveNodeIDs(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)
	
	ids := k.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, hub.IDs(nil), ids)
	
	k.DeleteActiveNodeIDs(ctx, 1)
	ids = k.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, hub.IDs(nil), ids)
	
	k.SetActiveNodeIDs(ctx, 1, hub.IDs{hub.NewNodeID(0)})
	k.DeleteActiveNodeIDs(ctx, 1)
	ids = k.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, hub.IDs(nil), ids)
	
	k.SetActiveNodeIDs(ctx, 2, hub.IDs{hub.NewNodeID(0), hub.NewNodeID(1)})
	k.DeleteActiveNodeIDs(ctx, 3)
	ids = k.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, hub.IDs{hub.NewNodeID(0), hub.NewNodeID(1)}, ids)
	
	k.DeleteActiveNodeIDs(ctx, 2)
	ids = k.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, hub.IDs(nil), ids)
}

func TestKeeper_GetNodesOfAddress(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)
	
	nodes := k.GetNodesOfAddress(ctx, types.TestAddress1)
	require.Equal(t, []types.Node{}, nodes)
	
	k.SetNode(ctx, types.TestNode)
	k.SetNodeIDByAddress(ctx, types.TestAddress1, 0, hub.NewNodeID(0))
	k.SetNodesCountOfAddress(ctx, types.TestAddress1, 1)
	
	nodes = k.GetNodesOfAddress(ctx, []byte(""))
	require.Equal(t, []types.Node{}, nodes)
	nodes = k.GetNodesOfAddress(ctx, types.TestAddress2)
	require.Equal(t, []types.Node{}, nodes)
	nodes = k.GetNodesOfAddress(ctx, types.TestAddress1)
	require.Equal(t, []types.Node{types.TestNode}, nodes)
	
	k.SetNode(ctx, types.TestNode)
	k.SetNodeIDByAddress(ctx, types.TestAddress2, 0, hub.NewNodeID(0))
	k.SetNodesCountOfAddress(ctx, types.TestAddress2, 1)
	
	nodes = k.GetNodesOfAddress(ctx, types.TestAddress2)
	require.Equal(t, []types.Node{types.TestNode}, nodes)
	
	node := types.TestNode
	node.ID = hub.NewNodeID(1)
	k.SetNode(ctx, node)
	k.SetNodeIDByAddress(ctx, types.TestAddress2, 1, hub.NewNodeID(1))
	k.SetNodesCountOfAddress(ctx, types.TestAddress2, 2)
	
	nodes = k.GetNodesOfAddress(ctx, types.TestAddress2)
	require.Equal(t, append([]types.Node{types.TestNode}, node), nodes)
}

func TestKeeper_GetAllNodes(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)
	
	nodes := k.GetAllNodes(ctx)
	require.Equal(t, []types.Node(nil), nodes)
	
	k.SetNode(ctx, types.TestNode)
	nodes = k.GetAllNodes(ctx)
	require.Equal(t, []types.Node{types.TestNode}, nodes)
	
	node := types.TestNode
	node.ID = hub.NewNodeID(1)
	k.SetNode(ctx, node)
	nodes = k.GetAllNodes(ctx)
	require.Equal(t, append([]types.Node{types.TestNode}, node), nodes)
	
	k.SetNode(ctx, types.TestNode)
	nodes = k.GetAllNodes(ctx)
	require.Equal(t, append([]types.Node{types.TestNode}, node), nodes)
}

func TestKeeper_AddNodeIDToActiveList(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)
	
	ids := k.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, hub.IDs(nil), ids)
	
	k.AddNodeIDToActiveList(ctx, 1, hub.NewNodeID(0))
	ids = k.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, hub.IDs{hub.NewNodeID(0)}, ids)
	
	k.AddNodeIDToActiveList(ctx, 2, hub.NewNodeID(0))
	k.AddNodeIDToActiveList(ctx, 2, hub.NewNodeID(1))
	ids = k.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, hub.IDs{hub.NewNodeID(0)}.Append(hub.NewNodeID(1)).Sort(), ids)
}

func TestKeeper_RemoveNodeIDFromActiveList(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)
	
	ids := k.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, hub.IDs(nil), ids)
	
	k.AddNodeIDToActiveList(ctx, 1, hub.NewNodeID(0))
	k.RemoveNodeIDFromActiveList(ctx, 1, hub.NewNodeID(1))
	ids = k.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, hub.IDs{hub.NewNodeID(0)}, ids)
	
	k.RemoveNodeIDFromActiveList(ctx, 1, hub.NewNodeID(0))
	ids = k.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, hub.IDs(nil), ids)
	
	k.AddNodeIDToActiveList(ctx, 2, hub.NewNodeID(0))
	k.AddNodeIDToActiveList(ctx, 2, hub.NewNodeID(1))
	k.RemoveNodeIDFromActiveList(ctx, 2, hub.NewNodeID(1))
	ids = k.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, hub.IDs{hub.NewNodeID(0)}, ids)
	
	k.RemoveNodeIDFromActiveList(ctx, 2, hub.NewNodeID(0))
	ids = k.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, hub.IDs(nil), ids)
}
