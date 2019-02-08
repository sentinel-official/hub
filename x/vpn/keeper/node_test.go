package keeper

import (
	"testing"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func TestKeeper_SetNodeDetails(t *testing.T) {
	var err csdkTypes.Error
	ctx, keeper, _ := TestCreateInput()

	err = keeper.SetNodeDetails(ctx, &TestNodeEmpty)
	require.Nil(t, err)

	err = keeper.SetNodeDetails(ctx, &TestNodeValid)
	require.Nil(t, err)
	result1, err := keeper.GetNodeDetails(ctx, TestNodeValid.ID)
	require.Nil(t, err)
	require.Equal(t, &TestNodeValid, result1)
}

func TestKeeper_GetNodeDetails(t *testing.T) {
	var err csdkTypes.Error
	ctx, keeper, _ := TestCreateInput()

	result1, err := keeper.GetNodeDetails(ctx, TestNodeValid.ID)
	require.Nil(t, err)
	require.Nil(t, result1)

	err = keeper.SetNodeDetails(ctx, &TestNodeValid)
	require.Nil(t, err)
	result2, err := keeper.GetNodeDetails(ctx, TestNodeValid.ID)
	require.Nil(t, err)
	require.Equal(t, &TestNodeValid, result2)
}

func TestKeeper_SetActiveNodeIDsAtHeight(t *testing.T) {
	var err csdkTypes.Error
	ctx, keeper, _ := TestCreateInput()

	err = keeper.SetActiveNodeIDsAtHeight(ctx, 0, TestNodeIDsEmpty)
	require.Nil(t, err)
	result1, err := keeper.GetActiveNodeIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Equal(t, TestNodeIDsEmpty, result1)

	err = keeper.SetActiveNodeIDsAtHeight(ctx, 0, TestNodeIDsValid)
	require.Nil(t, err)
	result2, err := keeper.GetActiveNodeIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Equal(t, TestNodeIDsValid, result2)
}

func TestKeeper_GetActiveNodeIDsAtHeight(t *testing.T) {
	var err csdkTypes.Error
	ctx, keeper, _ := TestCreateInput()

	result1, err := keeper.GetActiveNodeIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Nil(t, result1)

	err = keeper.SetActiveNodeIDsAtHeight(ctx, 0, TestNodeIDsValid)
	require.Nil(t, err)
	result2, err := keeper.GetActiveNodeIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Equal(t, TestNodeIDsValid, result2)
}

func TestKeeper_SetNodesCount(t *testing.T) {
	var err csdkTypes.Error
	ctx, keeper, _ := TestCreateInput()

	err = keeper.SetNodesCount(ctx, types.TestAddress, 0)
	require.Nil(t, err)
	result1, err := keeper.GetNodesCount(ctx, types.TestAddress)
	require.Nil(t, err)
	require.Equal(t, uint64(0), result1)

	err = keeper.SetNodesCount(ctx, types.TestAddress, 1)
	require.Nil(t, err)
	result2, err := keeper.GetNodesCount(ctx, types.TestAddress)
	require.Nil(t, err)
	require.Equal(t, uint64(1), result2)
}

func TestKeeper_GetNodesCount(t *testing.T) {
	var err csdkTypes.Error
	ctx, keeper, _ := TestCreateInput()

	result1, err := keeper.GetNodesCount(ctx, types.TestAddress)
	require.Nil(t, err)
	require.Equal(t, uint64(0), result1)

	err = keeper.SetNodesCount(ctx, types.TestAddress, 1)
	require.Nil(t, err)
	result2, err := keeper.GetNodesCount(ctx, types.TestAddress)
	require.Nil(t, err)
	require.Equal(t, uint64(1), result2)
}

func TestKeeper_GetNodesOfOwner(t *testing.T) {
	var err csdkTypes.Error
	ctx, keeper, _ := TestCreateInput()

	result1, err := keeper.GetNodesOfOwner(ctx, types.TestAddress)
	require.Nil(t, err)
	require.Equal(t, TestNodesEmpty, result1)

	err = keeper.SetNodeDetails(ctx, &TestNodeValid)
	require.Nil(t, err)
	err = keeper.SetNodesCount(ctx, types.TestAddress, 1)
	require.Nil(t, err)
	result2, err := keeper.GetNodesOfOwner(ctx, types.TestAddress)
	require.Nil(t, err)
	require.Equal(t, []*types.NodeDetails{&TestNodeValid}, result2)
}

func TestKeeper_GetNodes(t *testing.T) {
	var err csdkTypes.Error
	ctx, keeper, _ := TestCreateInput()

	result1, err := keeper.GetNodes(ctx)
	require.Nil(t, err)
	require.Equal(t, TestNodesEmpty, result1)

	err = keeper.SetNodeDetails(ctx, &TestNodeValid)
	require.Nil(t, err)
	result2, err := keeper.GetNodes(ctx)
	require.Nil(t, err)
	require.Equal(t, []*types.NodeDetails{&TestNodeValid}, result2)
}

func TestKeeper_AddNode(t *testing.T) {
	var err csdkTypes.Error
	ctx, keeper, _ := TestCreateInput()

	tags, err := keeper.AddNode(ctx, &TestNodeValid)
	require.Nil(t, err)
	require.Equal(t, TestNodeTagsValid, tags)
	result1, err := keeper.GetNodesCount(ctx, types.TestAddress)
	require.Nil(t, err)
	require.Equal(t, uint64(1), result1)
	result2, err := keeper.GetNodeDetails(ctx, TestNodeValid.ID)
	require.Nil(t, err)
	require.Equal(t, &TestNodeValid, result2)
}

func TestKeeper_AddActiveNodeIDAtHeight(t *testing.T) {
	var err csdkTypes.Error
	ctx, keeper, _ := TestCreateInput()

	err = keeper.AddActiveNodeIDAtHeight(ctx, 0, types.TestNodeIDValid)
	require.Nil(t, err)
	result1, err := keeper.GetActiveNodeIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Equal(t, types.NodeIDs{types.TestNodeIDValid}, result1)
	err = keeper.AddActiveNodeIDAtHeight(ctx, 0, types.TestNodeIDValid)
	require.Nil(t, err)
	result2, err := keeper.GetActiveNodeIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Equal(t, types.NodeIDs{types.TestNodeIDValid}, result2)
}

func TestKeeper_RemoveActiveNodeIDAtHeight(t *testing.T) {
	var err csdkTypes.Error
	ctx, keeper, _ := TestCreateInput()

	err = keeper.RemoveActiveNodeIDAtHeight(ctx, 0, types.TestNodeIDValid)
	require.Nil(t, err)

	err = keeper.AddActiveNodeIDAtHeight(ctx, 0, types.TestNodeIDValid)
	require.Nil(t, err)
	result1, err := keeper.GetActiveNodeIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Equal(t, types.NodeIDs{types.TestNodeIDValid}, result1)

	err = keeper.RemoveActiveNodeIDAtHeight(ctx, 0, types.TestNodeIDValid)
	require.Nil(t, err)
	require.Equal(t, types.NodeIDs{types.TestNodeIDValid}, result1)
	result2, err := keeper.GetActiveNodeIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Nil(t, result2)
}
