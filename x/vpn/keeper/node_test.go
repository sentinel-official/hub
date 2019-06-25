package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sentinel-official/sentinel-hub/x/vpn/types"
)

func TestKeeper_SetNodesCount(t *testing.T) {
	ctx, _, keeper, _ := TestCreateInput()

	count := keeper.GetNodesCount(ctx)
	require.Equal(t, uint64(0), count)

	keeper.SetNodesCount(ctx, 0)
	count = keeper.GetNodesCount(ctx)
	require.Equal(t, uint64(0), count)

	keeper.SetNodesCount(ctx, 1)
	count = keeper.GetNodesCount(ctx)
	require.Equal(t, uint64(1), count)

	keeper.SetNodesCount(ctx, 2)
	count = keeper.GetNodesCount(ctx)
	require.Equal(t, uint64(2), count)
}

func TestKeeper_GetNodesCount(t *testing.T) {
	TestKeeper_SetNodesCount(t)
}

func TestKeeper_SetNode(t *testing.T) {
	ctx, _, keeper, _ := TestCreateInput()

	_, found := keeper.GetNode(ctx, types.TestIDZero)
	require.Equal(t, false, found)

	keeper.SetNode(ctx, types.TestNodeEmpty)
	_, found = keeper.GetNode(ctx, types.TestNodeEmpty.ID)
	require.Equal(t, true, found)

	keeper.SetNode(ctx, types.TestNodeValid)
	node, found := keeper.GetNode(ctx, types.TestNodeValid.ID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestNodeValid, node)
}

func TestKeeper_GetNode(t *testing.T) {
	TestKeeper_SetNode(t)
}

func TestKeeper_SetNodesCountOfAddress(t *testing.T) {
	ctx, _, keeper, _ := TestCreateInput()

	count := keeper.GetNodesCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(0), count)

	keeper.SetNodesCountOfAddress(ctx, types.TestAddressEmpty, 1)
	count = keeper.GetNodesCountOfAddress(ctx, types.TestAddressEmpty)
	require.Equal(t, uint64(1), count)

	keeper.SetNodesCountOfAddress(ctx, types.TestAddress1, 1)
	count = keeper.GetNodesCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(1), count)

	keeper.SetNodesCountOfAddress(ctx, types.TestAddress2, 2)
	count = keeper.GetNodesCountOfAddress(ctx, types.TestAddress2)
	require.Equal(t, uint64(2), count)

	count = keeper.GetNodesCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(1), count)

	keeper.SetNodesCountOfAddress(ctx, types.TestAddress1, 0)
	count = keeper.GetNodesCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(0), count)

	count = keeper.GetNodesCountOfAddress(ctx, types.TestAddress2)
	require.Equal(t, uint64(2), count)

	keeper.SetNodesCountOfAddress(ctx, types.TestAddress2, 0)
	count = keeper.GetNodesCountOfAddress(ctx, types.TestAddress2)
	require.Equal(t, uint64(0), count)
}

func TestKeeper_GetNodesCountOfAddress(t *testing.T) {
	TestKeeper_SetNodesCountOfAddress(t)
}

func TestKeeper_SetNodeIDByAddress(t *testing.T) {
	ctx, _, keeper, _ := TestCreateInput()

	id, found := keeper.GetNodeIDByAddress(ctx, types.TestAddress1, 0)
	require.Equal(t, false, found)
	require.Equal(t, types.TestIDZero, id)

	keeper.SetNodeIDByAddress(ctx, types.TestAddressEmpty, 0, types.TestIDZero)
	id, found = keeper.GetNodeIDByAddress(ctx, types.TestAddressEmpty, 0)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDZero, id)

	keeper.SetNodeIDByAddress(ctx, types.TestAddress1, 0, types.TestIDZero)
	id, found = keeper.GetNodeIDByAddress(ctx, types.TestAddress1, 0)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDZero, id)

	keeper.SetNodeIDByAddress(ctx, types.TestAddress2, 0, types.TestIDZero)
	id, found = keeper.GetNodeIDByAddress(ctx, types.TestAddress2, 0)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDZero, id)

	keeper.SetNodeIDByAddress(ctx, types.TestAddress1, 1, types.TestIDPos)
	id, found = keeper.GetNodeIDByAddress(ctx, types.TestAddress1, 1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDPos, id)

	keeper.SetNodeIDByAddress(ctx, types.TestAddress2, 1, types.TestIDPos)
	id, found = keeper.GetNodeIDByAddress(ctx, types.TestAddress2, 1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDPos, id)
}

func TestKeeper_GetNodeIDByAddress(t *testing.T) {
	TestKeeper_SetNodeIDByAddress(t)
}

func TestKeeper_SetActiveNodeIDs(t *testing.T) {
	ctx, _, keeper, _ := TestCreateInput()

	ids := keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.SetActiveNodeIDs(ctx, 1, types.TestIDsEmpty)
	ids = keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.SetActiveNodeIDs(ctx, 1, types.TestIDsNil)
	ids = keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.SetActiveNodeIDs(ctx, 1, types.TestIDsValid)
	ids = keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, types.TestIDsValid, ids)

	keeper.SetActiveNodeIDs(ctx, 1, types.TestIDsValid.Append(types.TestIDZero))
	ids = keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, types.TestIDsValid.Append(types.TestIDZero).Sort(), ids)

	keeper.SetActiveNodeIDs(ctx, 1, types.TestIDsEmpty)
	ids = keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, types.TestIDsNil, ids)

	ids = keeper.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.SetActiveNodeIDs(ctx, 2, types.TestIDsEmpty)
	ids = keeper.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.SetActiveNodeIDs(ctx, 2, types.TestIDsNil)
	ids = keeper.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.SetActiveNodeIDs(ctx, 2, types.TestIDsValid)
	ids = keeper.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, types.TestIDsValid, ids)

	keeper.SetActiveNodeIDs(ctx, 2, types.TestIDsValid.Append(types.TestIDZero))
	ids = keeper.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, types.TestIDsValid.Append(types.TestIDZero).Sort(), ids)

	keeper.SetActiveNodeIDs(ctx, 2, types.TestIDsEmpty)
	ids = keeper.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, types.TestIDsNil, ids)
}

func TestKeeper_GetActiveNodeIDs(t *testing.T) {
	TestKeeper_SetActiveNodeIDs(t)
}

func TestKeeper_DeleteActiveNodeIDs(t *testing.T) {
	ctx, _, keeper, _ := TestCreateInput()

	ids := keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, types.TestIDsNil, ids)
	keeper.DeleteActiveNodeIDs(ctx, 1)
	ids = keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.SetActiveNodeIDs(ctx, 1, types.TestIDsValid)
	keeper.DeleteActiveNodeIDs(ctx, 1)
	ids = keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.SetActiveNodeIDs(ctx, 2, types.TestIDsValid.Append(types.TestIDPos))
	keeper.DeleteActiveNodeIDs(ctx, 3)
	ids = keeper.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, types.TestIDsValid.Append(types.TestIDPos), ids)

	keeper.DeleteActiveNodeIDs(ctx, 2)
	ids = keeper.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, types.TestIDsNil, ids)
}

func TestKeeper_GetNodesOfAddress(t *testing.T) {
	ctx, _, keeper, _ := TestCreateInput()

	nodes := keeper.GetNodesOfAddress(ctx, types.TestAddress1)
	require.Equal(t, types.TestNodesEmpty, nodes)

	keeper.SetNode(ctx, types.TestNodeValid)
	keeper.SetNodeIDByAddress(ctx, types.TestAddress1, 0, types.TestIDZero)
	keeper.SetNodesCountOfAddress(ctx, types.TestAddress1, 1)

	nodes = keeper.GetNodesOfAddress(ctx, types.TestAddressEmpty)
	require.Equal(t, types.TestNodesEmpty, nodes)
	nodes = keeper.GetNodesOfAddress(ctx, types.TestAddress2)
	require.Equal(t, types.TestNodesEmpty, nodes)
	nodes = keeper.GetNodesOfAddress(ctx, types.TestAddress1)
	require.Equal(t, types.TestNodesValid, nodes)

	keeper.SetNode(ctx, types.TestNodeValid)
	keeper.SetNodeIDByAddress(ctx, types.TestAddress2, 0, types.TestIDZero)
	keeper.SetNodesCountOfAddress(ctx, types.TestAddress2, 1)

	nodes = keeper.GetNodesOfAddress(ctx, types.TestAddress2)
	require.Equal(t, types.TestNodesValid, nodes)

	node := types.TestNodeValid
	node.ID = types.TestIDPos
	keeper.SetNode(ctx, node)
	keeper.SetNodeIDByAddress(ctx, types.TestAddress2, 1, types.TestIDPos)
	keeper.SetNodesCountOfAddress(ctx, types.TestAddress2, 2)

	nodes = keeper.GetNodesOfAddress(ctx, types.TestAddress2)
	require.Equal(t, append(types.TestNodesValid, node), nodes)
}

func TestKeeper_GetAllNodes(t *testing.T) {
	ctx, _, keeper, _ := TestCreateInput()

	nodes := keeper.GetAllNodes(ctx)
	require.Equal(t, types.TestNodesNil, nodes)

	keeper.SetNode(ctx, types.TestNodeValid)
	nodes = keeper.GetAllNodes(ctx)
	require.Equal(t, types.TestNodesValid, nodes)

	node := types.TestNodeValid
	node.ID = types.TestIDPos
	keeper.SetNode(ctx, node)
	nodes = keeper.GetAllNodes(ctx)
	require.Equal(t, append(types.TestNodesValid, node), nodes)

	keeper.SetNode(ctx, types.TestNodeValid)
	nodes = keeper.GetAllNodes(ctx)
	require.Equal(t, append(types.TestNodesValid, node), nodes)
}

func TestKeeper_AddNodeIDToActiveList(t *testing.T) {
	ctx, _, keeper, _ := TestCreateInput()

	ids := keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.AddNodeIDToActiveList(ctx, 1, types.TestIDZero)
	ids = keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, types.TestIDsValid, ids)

	keeper.AddNodeIDToActiveList(ctx, 2, types.TestIDZero)
	keeper.AddNodeIDToActiveList(ctx, 2, types.TestIDPos)
	ids = keeper.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, types.TestIDsValid.Append(types.TestIDPos).Sort(), ids)
}

func TestKeeper_RemoveNodeIDFromActiveList(t *testing.T) {
	ctx, _, keeper, _ := TestCreateInput()

	ids := keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.AddNodeIDToActiveList(ctx, 1, types.TestIDZero)
	keeper.RemoveNodeIDFromActiveList(ctx, 1, types.TestIDPos)
	ids = keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, types.TestIDsValid, ids)

	keeper.RemoveNodeIDFromActiveList(ctx, 1, types.TestIDZero)
	ids = keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.AddNodeIDToActiveList(ctx, 2, types.TestIDZero)
	keeper.AddNodeIDToActiveList(ctx, 2, types.TestIDPos)
	keeper.RemoveNodeIDFromActiveList(ctx, 2, types.TestIDPos)
	ids = keeper.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, types.TestIDsValid, ids)

	keeper.RemoveNodeIDFromActiveList(ctx, 2, types.TestIDZero)
	ids = keeper.GetActiveNodeIDs(ctx, 2)
	require.Equal(t, types.TestIDsNil, ids)
}
