package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeeper_SetSessionsCount(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	keeper.SetSessionsCount(ctx, 1)
	count := keeper.GetSessionsCount(ctx)
	require.Equal(t, uint64(1), count)

	keeper.SetSessionsCount(ctx, 0)
	count = keeper.GetSessionsCount(ctx)
	require.Equal(t, uint64(0), count)

	keeper.SetSessionsCount(ctx, 2)
	count = keeper.GetSessionsCount(ctx)
	require.Equal(t, uint64(2), count)
}

func TestKeeper_GetSessionsCount(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	count := keeper.GetSessionsCount(ctx)
	require.Equal(t, uint64(0), count)

	keeper.SetSessionsCount(ctx, 1)
	count = keeper.GetSessionsCount(ctx)
	require.Equal(t, uint64(1), count)
}

func TestKeeper_SetSession(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	_, found := keeper.GetSession(ctx, TestSessionEmpty.ID)
	require.Equal(t, false, found)

	keeper.SetSession(ctx, TestSessionValid)
	result, found := keeper.GetSession(ctx, TestSessionValid.ID)
	require.Equal(t, true, found)
	require.Equal(t, TestSessionValid, result)
}

func TestKeeper_GetSession(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	_, found := keeper.GetSession(ctx, TestIDPos)
	require.Equal(t, false, found)

	keeper.SetSession(ctx, TestSessionValid)
	result, found := keeper.GetSession(ctx, TestSessionValid.ID)
	require.Equal(t, true, found)
	require.Equal(t, TestSessionValid, result)
}

func TestKeeper_SetSessionsCountOfSubscription(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	keeper.SetSessionsCountOfSubscription(ctx, TestIDPos, 10)
	count := keeper.GetSessionsCountOfSubscription(ctx, TestIDPos)
	require.Equal(t, uint64(10), count)

	keeper.SetSessionsCountOfSubscription(ctx, TestIDPos, 1)
	count = keeper.GetSessionsCountOfSubscription(ctx, TestIDPos)
	require.Equal(t, uint64(1), count)
}

func TestKeeper_GetSessionsCountOfSubscription(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	count := keeper.GetSessionsCountOfSubscription(ctx, TestIDPos)
	require.Equal(t, uint64(0), count)

	keeper.SetSessionsCountOfSubscription(ctx, TestIDPos, 1)
	count = keeper.GetSessionsCountOfSubscription(ctx, TestIDPos)
	require.Equal(t, uint64(1), count)
}

func TestKeeper_SetSessionIDBySubscriptionID(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	keeper.SetSessionIDBySubscriptionID(ctx, TestIDPos, 1, TestIDZero)
	id, found := keeper.GetSessionIDBySubscriptionID(ctx, TestIDPos, 1)
	require.Equal(t, true, found)
	require.Equal(t, TestIDZero, id)

	keeper.SetSessionIDBySubscriptionID(ctx, TestIDPos, 1, TestIDPos)
	id, found = keeper.GetSessionIDBySubscriptionID(ctx, TestIDPos, 1)
	require.Equal(t, true, found)
	require.Equal(t, TestIDPos, id)

}

func TestKeeper_GetSessionIDBySubscriptionID(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	_, found := keeper.GetSessionIDBySubscriptionID(ctx, TestIDPos, 1)
	require.Equal(t, false, found)

	keeper.SetSessionIDBySubscriptionID(ctx, TestIDPos, 1, TestIDPos)
	id, found := keeper.GetSessionIDBySubscriptionID(ctx, TestIDPos, 1)
	require.Equal(t, true, found)
	require.Equal(t, TestIDPos, id)
}

func TestKeeper_SetActiveSessionIDs(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	keeper.SetActiveSessionIDs(ctx, 10, TestIDsValid)
	ids := keeper.GetActiveSessionIDs(ctx, 10)
	require.Equal(t, TestIDsValid, ids)

	keeper.SetActiveSessionIDs(ctx, 1, TestIDsEmpty)
	ids = keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, TestIDsEmpty, ids)
}

func TestKeeper_GetActiveSessionIDs(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	ids := keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, TestIDsEmpty, ids)

	keeper.SetActiveSessionIDs(ctx, 1, TestIDsEmpty)
	ids = keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, TestIDsEmpty, ids)
}

func TestKeeper_DeleteActiveSessionIDs(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	keeper.SetActiveSessionIDs(ctx, 10, TestIDsValid)
	ids := keeper.GetActiveSessionIDs(ctx, 10)
	require.Equal(t, TestIDsValid, ids)

	keeper.DeleteActiveSessionIDs(ctx, 10)
	ids = keeper.GetActiveSessionIDs(ctx, 10)
	require.Equal(t, TestIDsEmpty, ids)
}

func TestKeeper_GetSessionsOfSubscription(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	sessions := keeper.GetSessionsOfSubscription(ctx, TestIDPos)
	require.Equal(t, TestSessionsEmpty, sessions)

	keeper.SetSession(ctx, TestSessionValid)
	keeper.SetSessionIDBySubscriptionID(ctx, TestIDPos, 0, TestIDPos)
	keeper.SetSessionsCountOfSubscription(ctx, TestIDPos, 1)
	sessions = keeper.GetSessionsOfSubscription(ctx, TestIDPos)
	require.Equal(t, TestSessionsValid, sessions)
}

func TestKeeper_GetAllSessions(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	sessions := keeper.GetAllSessions(ctx)
	require.Equal(t, TestSessionsNil, sessions)

	keeper.SetSession(ctx, TestSessionValid)
	sessions = keeper.GetAllSessions(ctx)
	require.Equal(t, TestSessionsValid, sessions)
}

func TestKeeper_AddSessionIDToActiveList(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	ids := keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, TestIDsEmpty, ids)

	keeper.AddSessionIDToActiveList(ctx, 1, TestIDPos)
	ids = keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, TestIDsValid, ids)
}

func TestKeeper_RemoveSessionIDFromActiveList(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	ids := keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, TestIDsEmpty, ids)

	keeper.AddSessionIDToActiveList(ctx, 1, TestIDPos)
	ids = keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, TestIDsValid, ids)

	keeper.RemoveSessionIDFromActiveList(ctx, 1, TestIDPos)
	ids = keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, TestIDsEmpty, ids)
}
