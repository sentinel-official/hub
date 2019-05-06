package keeper

import (
	"testing"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func TestKeeper_SetSession(t *testing.T) {
	var err csdkTypes.Error
	ctx, _, keeper, _, _ := TestCreateInput()

	err = keeper.SetSession(ctx, &TestSessionEmpty)
	require.Nil(t, err)

	err = keeper.SetSession(ctx, &TestSessionValid)
	require.Nil(t, err)
	result1, err := keeper.GetSession(ctx, TestSessionValid.ID)
	require.Nil(t, err)
	require.Equal(t, &TestSessionValid, result1)
}

func TestKeeper_GetSession(t *testing.T) {
	var err csdkTypes.Error
	ctx, _, keeper, _, _ := TestCreateInput()

	result1, err := keeper.GetSession(ctx, TestSessionValid.ID)
	require.Nil(t, err)
	require.Nil(t, result1)

	err = keeper.SetSession(ctx, &TestSessionValid)
	require.Nil(t, err)
	result2, err := keeper.GetSession(ctx, TestSessionValid.ID)
	require.Nil(t, err)
	require.Equal(t, &TestSessionValid, result2)
}

func TestKeeper_SetActiveSessionIDsAtHeight(t *testing.T) {
	var err csdkTypes.Error
	ctx, _, keeper, _, _ := TestCreateInput()

	err = keeper.SetActiveSessionIDsAtHeight(ctx, 0, TestSessionIDsEmpty)
	require.Nil(t, err)
	result1, err := keeper.GetActiveSessionIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Equal(t, TestSessionIDsEmpty, result1)

	err = keeper.SetActiveSessionIDsAtHeight(ctx, 0, TestSessionIDsValid)
	require.Nil(t, err)
	result2, err := keeper.GetActiveSessionIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Equal(t, TestSessionIDsValid, result2)
}

func TestKeeper_GetActiveSessionIDsAtHeight(t *testing.T) {
	var err csdkTypes.Error
	ctx, _, keeper, _, _ := TestCreateInput()

	result1, err := keeper.GetActiveSessionIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Nil(t, result1)

	err = keeper.SetActiveSessionIDsAtHeight(ctx, 0, TestSessionIDsValid)
	require.Nil(t, err)
	result2, err := keeper.GetActiveSessionIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Equal(t, TestSessionIDsValid, result2)
}

func TestKeeper_SetSessionsCount(t *testing.T) {
	var err csdkTypes.Error
	ctx, _, keeper, _, _ := TestCreateInput()

	err = keeper.SetSessionsCount(ctx, types.TestAddress2, 0)
	require.Nil(t, err)
	result1, err := keeper.GetSessionsCount(ctx, types.TestAddress2)
	require.Nil(t, err)
	require.Equal(t, uint64(0), result1)

	err = keeper.SetSessionsCount(ctx, types.TestAddress2, 1)
	require.Nil(t, err)
	result2, err := keeper.GetSessionsCount(ctx, types.TestAddress2)
	require.Nil(t, err)
	require.Equal(t, uint64(1), result2)
}

func TestKeeper_GetSessionsCount(t *testing.T) {
	var err csdkTypes.Error
	ctx, _, keeper, _, _ := TestCreateInput()

	result1, err := keeper.GetSessionsCount(ctx, types.TestAddress2)
	require.Nil(t, err)
	require.Equal(t, uint64(0), result1)

	err = keeper.SetSessionsCount(ctx, types.TestAddress2, 1)
	require.Nil(t, err)
	result2, err := keeper.GetSessionsCount(ctx, types.TestAddress2)
	require.Nil(t, err)
	require.Equal(t, uint64(1), result2)
}

func TestKeeper_AddSession(t *testing.T) {
	var err csdkTypes.Error
	ctx, _, keeper, _, _ := TestCreateInput()

	tags, err := keeper.AddSession(ctx, &TestSessionValid)
	require.Nil(t, err)
	require.Equal(t, TestSessionTagsValid, tags)
	result1, err := keeper.GetSessionsCount(ctx, types.TestAddress2)
	require.Nil(t, err)
	require.Equal(t, uint64(1), result1)
	result2, err := keeper.GetSession(ctx, TestSessionValid.ID)
	require.Nil(t, err)
	require.Equal(t, &TestSessionValid, result2)
}

func TestKeeper_AddActiveSessionIDsAtHeight(t *testing.T) {
	var err csdkTypes.Error
	ctx, _, keeper, _, _ := TestCreateInput()

	err = keeper.AddActiveSessionIDsAtHeight(ctx, 0, types.TestSessionIDValid)
	require.Nil(t, err)
	result1, err := keeper.GetActiveSessionIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Equal(t, sdkTypes.IDs{types.TestSessionIDValid}, result1)
	err = keeper.AddActiveSessionIDsAtHeight(ctx, 0, types.TestSessionIDValid)
	require.Nil(t, err)
	result2, err := keeper.GetActiveSessionIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Equal(t, sdkTypes.IDs{types.TestSessionIDValid}, result2)
}

func TestKeeper_RemoveActiveSessionIDsAtHeight(t *testing.T) {
	var err csdkTypes.Error
	ctx, _, keeper, _, _ := TestCreateInput()

	err = keeper.RemoveActiveSessionIDsAtHeight(ctx, 0, types.TestSessionIDValid)
	require.Nil(t, err)

	err = keeper.AddActiveSessionIDsAtHeight(ctx, 0, types.TestSessionIDValid)
	require.Nil(t, err)
	result1, err := keeper.GetActiveSessionIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Equal(t, sdkTypes.IDs{types.TestSessionIDValid}, result1)

	err = keeper.RemoveActiveSessionIDsAtHeight(ctx, 0, types.TestSessionIDValid)
	require.Nil(t, err)
	require.Equal(t, sdkTypes.IDs{types.TestSessionIDValid}, result1)
	result2, err := keeper.GetActiveSessionIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Nil(t, result2)
}
