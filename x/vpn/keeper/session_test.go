package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sentinel-official/sentinel-hub/x/vpn/types"
)

func TestKeeper_SetSessionsCount(t *testing.T) {
	ctx, _, keeper, _ := TestCreateInput()

	count := keeper.GetSessionsCount(ctx)
	require.Equal(t, uint64(0), count)

	keeper.SetSessionsCount(ctx, 1)
	count = keeper.GetSessionsCount(ctx)
	require.Equal(t, uint64(1), count)

	keeper.SetSessionsCount(ctx, 0)
	count = keeper.GetSessionsCount(ctx)
	require.Equal(t, uint64(0), count)

	keeper.SetSessionsCount(ctx, 2)
	count = keeper.GetSessionsCount(ctx)
	require.Equal(t, uint64(2), count)
}

func TestKeeper_GetSessionsCount(t *testing.T) {
	TestKeeper_SetSessionsCount(t)
}

func TestKeeper_SetSession(t *testing.T) {
	ctx, _, keeper, _ := TestCreateInput()

	_, found := keeper.GetSession(ctx, types.TestIDZero)
	require.Equal(t, false, found)

	keeper.SetSession(ctx, types.TestSessionEmpty)
	result, found := keeper.GetSession(ctx, types.TestSessionEmpty.ID)
	require.Equal(t, true, found)

	keeper.SetSession(ctx, types.TestSessionValid)
	result, found = keeper.GetSession(ctx, types.TestSessionValid.ID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestSessionValid, result)
}

func TestKeeper_GetSession(t *testing.T) {
	TestKeeper_SetNode(t)
}

func TestKeeper_SetSessionsCountOfSubscription(t *testing.T) {
	ctx, _, keeper, _ := TestCreateInput()

	count := keeper.GetSessionsCountOfSubscription(ctx, types.TestIDZero)
	require.Equal(t, uint64(0), count)

	keeper.SetSessionsCountOfSubscription(ctx, types.TestIDZero, 1)
	count = keeper.GetSessionsCountOfSubscription(ctx, types.TestIDZero)
	require.Equal(t, uint64(1), count)

	keeper.SetSessionsCountOfSubscription(ctx, types.TestIDPos, 1)
	count = keeper.GetSessionsCountOfSubscription(ctx, types.TestIDPos)
	require.Equal(t, uint64(1), count)

	keeper.SetSessionsCountOfSubscription(ctx, types.TestIDZero, 2)
	count = keeper.GetSessionsCountOfSubscription(ctx, types.TestIDZero)
	require.Equal(t, uint64(2), count)
}

func TestKeeper_GetSessionsCountOfSubscription(t *testing.T) {
	TestKeeper_SetSessionsCountOfSubscription(t)
}

func TestKeeper_SetSessionIDBySubscriptionID(t *testing.T) {
	ctx, _, keeper, _ := TestCreateInput()

	_, found := keeper.GetSessionIDBySubscriptionID(ctx, types.TestIDPos, 1)
	require.Equal(t, false, found)

	keeper.SetSessionIDBySubscriptionID(ctx, types.TestIDZero, 2, types.TestIDZero)
	id, found := keeper.GetSessionIDBySubscriptionID(ctx, types.TestIDZero, 2)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDZero, id)

	keeper.SetSessionIDBySubscriptionID(ctx, types.TestIDZero, 3, types.TestIDPos)
	id, found = keeper.GetSessionIDBySubscriptionID(ctx, types.TestIDZero, 3)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDPos, id)

	keeper.SetSessionIDBySubscriptionID(ctx, types.TestIDPos, 4, types.TestIDZero)
	id, found = keeper.GetSessionIDBySubscriptionID(ctx, types.TestIDPos, 4)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDZero, id)

	keeper.SetSessionIDBySubscriptionID(ctx, types.TestIDPos, 5, types.TestIDPos)
	id, found = keeper.GetSessionIDBySubscriptionID(ctx, types.TestIDPos, 5)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDPos, id)

}

func TestKeeper_GetSessionIDBySubscriptionID(t *testing.T) {
	TestKeeper_SetSessionIDBySubscriptionID(t)
}

func TestKeeper_SetActiveSessionIDs(t *testing.T) {
	ctx, _, keeper, _ := TestCreateInput()

	ids := keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.SetActiveSessionIDs(ctx, 1, types.TestIDsEmpty)
	ids = keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.SetActiveSessionIDs(ctx, 1, types.TestIDsNil)
	ids = keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.SetActiveSessionIDs(ctx, 1, types.TestIDsValid)
	ids = keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, types.TestIDsValid, ids)

	keeper.SetActiveSessionIDs(ctx, 1, types.TestIDsValid.Append(types.TestIDZero))
	ids = keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, types.TestIDsValid.Append(types.TestIDZero).Sort(), ids)

	keeper.SetActiveSessionIDs(ctx, 1, types.TestIDsEmpty)
	ids = keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, types.TestIDsNil, ids)

	ids = keeper.GetActiveSessionIDs(ctx, 2)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.SetActiveSessionIDs(ctx, 2, types.TestIDsEmpty)
	ids = keeper.GetActiveSessionIDs(ctx, 2)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.SetActiveSessionIDs(ctx, 2, types.TestIDsNil)
	ids = keeper.GetActiveSessionIDs(ctx, 2)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.SetActiveSessionIDs(ctx, 2, types.TestIDsValid)
	ids = keeper.GetActiveSessionIDs(ctx, 2)
	require.Equal(t, types.TestIDsValid, ids)

	keeper.SetActiveSessionIDs(ctx, 2, types.TestIDsValid.Append(types.TestIDZero))
	ids = keeper.GetActiveSessionIDs(ctx, 2)
	require.Equal(t, types.TestIDsValid.Append(types.TestIDZero).Sort(), ids)

	keeper.SetActiveSessionIDs(ctx, 2, types.TestIDsEmpty)
	ids = keeper.GetActiveSessionIDs(ctx, 2)
	require.Equal(t, types.TestIDsNil, ids)

}

func TestKeeper_GetActiveSessionIDs(t *testing.T) {
	TestKeeper_SetActiveSessionIDs(t)
}

func TestKeeper_DeleteActiveSessionIDs(t *testing.T) {
	ctx, _, keeper, _ := TestCreateInput()

	ids := keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, types.TestIDsNil, ids)
	keeper.DeleteActiveSessionIDs(ctx, 1)
	ids = keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.SetActiveSessionIDs(ctx, 1, types.TestIDsValid)
	keeper.DeleteActiveSessionIDs(ctx, 1)
	ids = keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.SetActiveSessionIDs(ctx, 1, types.TestIDsValid.Append(types.TestIDPos))
	keeper.DeleteActiveSessionIDs(ctx, 2)
	ids = keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, types.TestIDsValid.Append(types.TestIDPos), ids)

	keeper.DeleteActiveSessionIDs(ctx, 1)
	ids = keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, types.TestIDsNil, ids)
}

func TestKeeper_GetSessionsOfSubscription(t *testing.T) {
	ctx, _, keeper, _ := TestCreateInput()

	sessions := keeper.GetSessionsOfSubscription(ctx, types.TestIDZero)
	require.Equal(t, types.TestSessionsEmpty, sessions)

	keeper.SetSession(ctx, types.TestSessionValid)
	keeper.SetSessionIDBySubscriptionID(ctx, types.TestIDZero, 0, types.TestIDZero)
	keeper.SetSessionsCountOfSubscription(ctx, types.TestIDZero, 1)

	sessions = keeper.GetSessionsOfSubscription(ctx, types.TestIDPos)
	require.Equal(t, types.TestSessionsEmpty, sessions)
	sessions = keeper.GetSessionsOfSubscription(ctx, types.TestIDZero)
	require.Equal(t, types.TestSessionsValid, sessions)

	session := types.TestSessionValid
	session.ID = types.TestIDPos
	keeper.SetSession(ctx, session)
	keeper.SetSessionIDBySubscriptionID(ctx, types.TestIDZero, 1, types.TestIDPos)
	keeper.SetSessionsCountOfSubscription(ctx, types.TestIDZero, 2)
	sessions = keeper.GetSessionsOfSubscription(ctx, types.TestIDZero)
	require.Equal(t, append(types.TestSessionsValid, session), sessions)
}

func TestKeeper_GetAllSessions(t *testing.T) {
	ctx, _, keeper, _ := TestCreateInput()

	sessions := keeper.GetAllSessions(ctx)
	require.Equal(t, types.TestSessionsNil, sessions)

	keeper.SetSession(ctx, types.TestSessionValid)
	sessions = keeper.GetAllSessions(ctx)
	require.Equal(t, types.TestSessionsValid, sessions)

	session := types.TestSessionValid
	session.ID = types.TestIDPos
	keeper.SetSession(ctx, session)
	sessions = keeper.GetAllSessions(ctx)
	require.Equal(t, append(types.TestSessionsValid, session), sessions)

	keeper.SetSession(ctx, types.TestSessionValid)
	sessions = keeper.GetAllSessions(ctx)
	require.Equal(t, append(types.TestSessionsValid, session), sessions)

}

func TestKeeper_AddSessionIDToActiveList(t *testing.T) {
	ctx, _, keeper, _ := TestCreateInput()

	ids := keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.AddSessionIDToActiveList(ctx, 1, types.TestIDZero)
	ids = keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, types.TestIDsValid, ids)

	keeper.AddSessionIDToActiveList(ctx, 1, types.TestIDPos)
	ids = keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, types.TestIDsValid.Append(types.TestIDPos), ids)

	keeper.AddSessionIDToActiveList(ctx, 2, types.TestIDZero)
	keeper.AddSessionIDToActiveList(ctx, 2, types.TestIDPos)
	require.Equal(t, types.TestIDsValid.Append(types.TestIDPos), ids)
}

func TestKeeper_RemoveSessionIDFromActiveList(t *testing.T) {
	ctx, _, keeper, _ := TestCreateInput()

	ids := keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.AddSessionIDToActiveList(ctx, 1, types.TestIDZero)
	keeper.RemoveSessionIDFromActiveList(ctx, 1, types.TestIDPos)
	ids = keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, types.TestIDsValid, ids)

	keeper.RemoveSessionIDFromActiveList(ctx, 1, types.TestIDZero)
	ids = keeper.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, types.TestIDsNil, ids)

	keeper.AddSessionIDToActiveList(ctx, 3, types.TestIDZero)
	keeper.AddSessionIDToActiveList(ctx, 3, types.TestIDPos)
	keeper.RemoveSessionIDFromActiveList(ctx, 3, types.TestIDPos)
	ids = keeper.GetActiveSessionIDs(ctx, 3)
	require.Equal(t, types.TestIDsValid, ids)

	keeper.RemoveSessionIDFromActiveList(ctx, 3, types.TestIDZero)
	ids = keeper.GetActiveSessionIDs(ctx, 3)
	require.Equal(t, types.TestIDsNil, ids)
}
