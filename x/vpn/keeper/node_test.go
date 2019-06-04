package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeeper_SetNodesCount(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	keeper.SetNodesCount(ctx, uint64(1))
	count := keeper.GetNodesCount(ctx)
	require.Equal(t, uint64(1), count)

	keeper.SetNodesCount(ctx, uint64(0))
	count = keeper.GetNodesCount(ctx)
	require.Equal(t, uint64(0), count)

	keeper.SetNodesCount(ctx, uint64(2))
	count = keeper.GetNodesCount(ctx)
	require.Equal(t, uint64(2), count)
}

func TestKeeper_GetNodesCount(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	count := keeper.GetNodesCount(ctx)
	require.Equal(t, uint64(0), count)

	keeper.SetNodesCount(ctx, uint64(1))
	count = keeper.GetNodesCount(ctx)
	require.Equal(t, uint64(1), count)
}

func TestKeeper_SetNode(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	keeper.SetNode(ctx, TestNodeValid)
	result, found := keeper.GetNode(ctx, TestNodeValid.ID)
	require.Equal(t, true, found)
	require.Equal(t, TestNodeValid, result)

	keeper.SetNode(ctx, TestNodeEmpty)
	result, found = keeper.GetNode(ctx, TestNodeEmpty.ID)
	require.Equal(t, true, found)
}

func TestKeeper_GetNode(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	_, found := keeper.GetNode(ctx, TestIDPos)
	require.Equal(t, false, found)

	keeper.SetNode(ctx, TestNodeValid)
	result, found := keeper.GetNode(ctx, TestNodeValid.ID)
	require.Equal(t, true, found)
	require.Equal(t, TestNodeValid, result)
}

func TestKeeper_SetNodesCountOfAddress(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	keeper.SetNodesCountOfAddress(ctx, TestAddress1, uint64(1))
	count := keeper.GetNodesCountOfAddress(ctx, TestAddress1)
	require.Equal(t, uint64(1), count)

	keeper.SetNodesCountOfAddress(ctx, TestAddress1, uint64(10))
	count = keeper.GetNodesCountOfAddress(ctx, TestAddress1)
	require.Equal(t, uint64(10), count)
}

func TestKeeper_GetNodesCountOfAddress(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	count := keeper.GetNodesCountOfAddress(ctx, TestAddress1)
	require.Equal(t, uint64(0), count)

	keeper.SetNodesCountOfAddress(ctx, TestAddress1, uint64(10))
	count = keeper.GetNodesCountOfAddress(ctx, TestAddress1)
	require.Equal(t, uint64(10), count)
}

func TestKeeper_SetNodeIDByAddress(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	keeper.SetNodeIDByAddress(ctx, TestAddress1, uint64(0), TestIDPos)
	id, found := keeper.GetNodeIDByAddress(ctx, TestAddress1, uint64(0))
	require.Equal(t, true, found)
	require.Equal(t, TestIDPos, id)

	keeper.SetNodeIDByAddress(ctx, TestAddress2, uint64(10), TestIDZero)
	id, found = keeper.GetNodeIDByAddress(ctx, TestAddress2, uint64(10))
	require.Equal(t, true, found)
	require.Equal(t, TestIDZero, id)
}

func TestKeeper_GetNodeIDByAddress(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	_, found := keeper.GetNodeIDByAddress(ctx, TestAddress1, uint64(0))
	require.Equal(t, false, found)

	keeper.SetNodeIDByAddress(ctx, TestAddress1, uint64(1), TestIDPos)
	res, found := keeper.GetNodeIDByAddress(ctx, TestAddress1, uint64(1))
	require.Equal(t, true, found)
	require.Equal(t, TestIDPos, res)
}

func TestKeeper_SetActiveNodeIDs(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	keeper.SetActiveNodeIDs(ctx, 10, TestIDsValid)
	ids := keeper.GetActiveNodeIDs(ctx, 10)
	require.Equal(t, TestIDsValid, ids)

	keeper.SetActiveNodeIDs(ctx, 1, TestIDsEmpty)
	ids = keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, TestIDsEmpty, ids)
}

func TestKeeper_GetActiveNodeIDs(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	ids := keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, TestIDsEmpty, ids)

	keeper.SetActiveNodeIDs(ctx, 10, TestIDsValid)
	ids = keeper.GetActiveNodeIDs(ctx, 10)
	require.Equal(t, TestIDsValid, ids)
}

func TestKeeper_DeleteActiveNodeIDs(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	keeper.SetActiveNodeIDs(ctx, 1, TestIDsValid)
	ids := keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, TestIDsValid, ids)

	keeper.DeleteActiveNodeIDs(ctx, 1)
	ids = keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, TestIDsEmpty, ids)
}

func TestKeeper_GetNodesOfAddress(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	nodes := keeper.GetNodesOfAddress(ctx, TestAddress1)
	require.Equal(t, TestNodesEmpty, nodes)

	keeper.SetNode(ctx, TestNodeValid)
	keeper.SetNodeIDByAddress(ctx, TestAddress1, 0, TestIDZero)
	keeper.SetNodesCountOfAddress(ctx, TestAddress1, 1)

	nodes = keeper.GetNodesOfAddress(ctx, TestAddress1)
	require.Equal(t, TestNodesValid, nodes)
}

func TestKeeper_GetAllNodes(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	nodes := keeper.GetAllNodes(ctx)
	require.Equal(t, TestNodesNil, nodes)

	keeper.SetNode(ctx, TestNodeValid)
	nodes = keeper.GetAllNodes(ctx)
	require.Equal(t, TestNodesValid, nodes)
}

func TestKeeper_AddNodeIDToActiveList(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	ids := keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, TestIDsEmpty, ids)

	keeper.AddNodeIDToActiveList(ctx, 1, TestIDPos)
	ids = keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, TestIDsValid, ids)
	require.Equal(t, TestIDsValid.Len(), ids.Len())

	keeper.AddNodeIDToActiveList(ctx, 1, TestIDZero)
	ids = keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, TestIDsValid.Append(TestIDZero).Sort(), ids)
	require.Equal(t, TestIDsValid.Append(TestIDZero).Len(), ids.Len())
}

func TestKeeper_RemoveNodeIDFromActiveList(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	keeper.AddNodeIDToActiveList(ctx, 1, TestIDPos)
	keeper.DeleteActiveNodeIDs(ctx, 1)
	ids := keeper.GetActiveNodeIDs(ctx, 1)
	require.Equal(t, TestIDsEmpty, ids)
	require.Equal(t, TestIDsEmpty.Len(), ids.Len())
}
