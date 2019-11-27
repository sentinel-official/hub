package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func TestKeeper_SetSessionsCount(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	count := k.GetSessionsCount(ctx)
	require.Equal(t, uint64(0), count)

	k.SetSessionsCount(ctx, 1)
	count = k.GetSessionsCount(ctx)
	require.Equal(t, uint64(1), count)

	k.SetSessionsCount(ctx, 0)
	count = k.GetSessionsCount(ctx)
	require.Equal(t, uint64(0), count)

	k.SetSessionsCount(ctx, 2)
	count = k.GetSessionsCount(ctx)
	require.Equal(t, uint64(2), count)
}

func TestKeeper_GetSessionsCount(t *testing.T) {
	TestKeeper_SetSessionsCount(t)
}

func TestKeeper_SetSession(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	_, found := k.GetSession(ctx, hub.NewSessionID(0))
	require.Equal(t, false, found)

	k.SetSession(ctx, types.TestSession)
	result, found := k.GetSession(ctx, types.TestSession.ID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestSession, result)
}

func TestKeeper_GetSession(t *testing.T) {
	TestKeeper_SetNode(t)
}

func TestKeeper_SetSessionsCountOfSubscription(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	count := k.GetSessionsCountOfSubscription(ctx, hub.NewSubscriptionID(0))
	require.Equal(t, uint64(0), count)

	k.SetSessionsCountOfSubscription(ctx, hub.NewSubscriptionID(0), 1)
	count = k.GetSessionsCountOfSubscription(ctx, hub.NewSubscriptionID(0))
	require.Equal(t, uint64(1), count)

	k.SetSessionsCountOfSubscription(ctx, hub.NewSubscriptionID(1), 1)
	count = k.GetSessionsCountOfSubscription(ctx, hub.NewSubscriptionID(1))
	require.Equal(t, uint64(1), count)

	k.SetSessionsCountOfSubscription(ctx, hub.NewSubscriptionID(0), 2)
	count = k.GetSessionsCountOfSubscription(ctx, hub.NewSubscriptionID(0))
	require.Equal(t, uint64(2), count)
}

func TestKeeper_GetSessionsCountOfSubscription(t *testing.T) {
	TestKeeper_SetSessionsCountOfSubscription(t)
}

func TestKeeper_SetSessionIDBySubscriptionID(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	_, found := k.GetSessionIDBySubscriptionID(ctx, hub.NewSubscriptionID(1), 1)
	require.Equal(t, false, found)

	k.SetSessionIDBySubscriptionID(ctx, hub.NewSubscriptionID(0), 2, hub.NewSessionID(0))
	id, found := k.GetSessionIDBySubscriptionID(ctx, hub.NewSubscriptionID(0), 2)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewSessionID(0), id)

	k.SetSessionIDBySubscriptionID(ctx, hub.NewSubscriptionID(0), 3, hub.NewSessionID(1))
	id, found = k.GetSessionIDBySubscriptionID(ctx, hub.NewSubscriptionID(0), 3)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewSessionID(1), id)

	k.SetSessionIDBySubscriptionID(ctx, hub.NewSubscriptionID(1), 4, hub.NewSessionID(0))
	id, found = k.GetSessionIDBySubscriptionID(ctx, hub.NewSubscriptionID(1), 4)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewSessionID(0), id)

	k.SetSessionIDBySubscriptionID(ctx, hub.NewSubscriptionID(1), 5, hub.NewSessionID(1))
	id, found = k.GetSessionIDBySubscriptionID(ctx, hub.NewSubscriptionID(1), 5)
	require.Equal(t, true, found)
	require.Equal(t, hub.NewSessionID(1), id)

}

func TestKeeper_GetSessionIDBySubscriptionID(t *testing.T) {
	TestKeeper_SetSessionIDBySubscriptionID(t)
}

func TestKeeper_SetActiveSessionIDs(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	ids := k.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, hub.IDs(nil), ids)

	k.SetActiveSessionIDs(ctx, 1, hub.IDs{})
	ids = k.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, hub.IDs(nil), ids)

	k.SetActiveSessionIDs(ctx, 1, hub.IDs(nil))
	ids = k.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, hub.IDs(nil), ids)

	k.SetActiveSessionIDs(ctx, 1, hub.IDs{hub.NewSessionID(0)})
	ids = k.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, hub.IDs{hub.NewSessionID(0)}, ids)

	k.SetActiveSessionIDs(ctx, 1, hub.IDs{hub.NewSessionID(0)}.Append(hub.NewSessionID(0)))
	ids = k.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, hub.IDs{hub.NewSessionID(0)}.Append(hub.NewSessionID(0)).Sort(), ids)

	k.SetActiveSessionIDs(ctx, 1, hub.IDs{})
	ids = k.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, hub.IDs(nil), ids)

	ids = k.GetActiveSessionIDs(ctx, 2)
	require.Equal(t, hub.IDs(nil), ids)

	k.SetActiveSessionIDs(ctx, 2, hub.IDs{})
	ids = k.GetActiveSessionIDs(ctx, 2)
	require.Equal(t, hub.IDs(nil), ids)

	k.SetActiveSessionIDs(ctx, 2, hub.IDs(nil))
	ids = k.GetActiveSessionIDs(ctx, 2)
	require.Equal(t, hub.IDs(nil), ids)

	k.SetActiveSessionIDs(ctx, 2, hub.IDs{hub.NewSessionID(0)})
	ids = k.GetActiveSessionIDs(ctx, 2)
	require.Equal(t, hub.IDs{hub.NewSessionID(0)}, ids)

	k.SetActiveSessionIDs(ctx, 2, hub.IDs{hub.NewSessionID(0)}.Append(hub.NewSessionID(0)))
	ids = k.GetActiveSessionIDs(ctx, 2)
	require.Equal(t, hub.IDs{hub.NewSessionID(0)}.Append(hub.NewSessionID(0)).Sort(), ids)

	k.SetActiveSessionIDs(ctx, 2, hub.IDs{})
	ids = k.GetActiveSessionIDs(ctx, 2)
	require.Equal(t, hub.IDs(nil), ids)

}

func TestKeeper_GetActiveSessionIDs(t *testing.T) {
	TestKeeper_SetActiveSessionIDs(t)
}

func TestKeeper_DeleteActiveSessionIDs(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	ids := k.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, hub.IDs(nil), ids)
	k.DeleteActiveSessionIDs(ctx, 1)
	ids = k.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, hub.IDs(nil), ids)

	k.SetActiveSessionIDs(ctx, 1, hub.IDs{hub.NewSessionID(0)})
	k.DeleteActiveSessionIDs(ctx, 1)
	ids = k.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, hub.IDs(nil), ids)

	k.SetActiveSessionIDs(ctx, 1, hub.IDs{hub.NewSessionID(0)}.Append(hub.NewSessionID(1)))
	k.DeleteActiveSessionIDs(ctx, 2)
	ids = k.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, hub.IDs{hub.NewSessionID(0)}.Append(hub.NewSessionID(1)), ids)

	k.DeleteActiveSessionIDs(ctx, 1)
	ids = k.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, hub.IDs(nil), ids)
}

func TestKeeper_GetSessionsOfSubscription(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	sessions := k.GetSessionsOfSubscription(ctx, hub.NewSubscriptionID(0))
	require.Equal(t, []types.Session{}, sessions)

	k.SetSession(ctx, types.TestSession)
	k.SetSessionIDBySubscriptionID(ctx, hub.NewSubscriptionID(0), 0, hub.NewSessionID(0))
	k.SetSessionsCountOfSubscription(ctx, hub.NewSubscriptionID(0), 1)

	sessions = k.GetSessionsOfSubscription(ctx, hub.NewSubscriptionID(1))
	require.Equal(t, []types.Session{}, sessions)
	sessions = k.GetSessionsOfSubscription(ctx, hub.NewSubscriptionID(0))
	require.Equal(t, []types.Session{types.TestSession}, sessions)

	session := types.TestSession
	session.ID = hub.NewSessionID(1)
	k.SetSession(ctx, session)
	k.SetSessionIDBySubscriptionID(ctx, hub.NewSubscriptionID(0), 1, hub.NewSessionID(1))
	k.SetSessionsCountOfSubscription(ctx, hub.NewSubscriptionID(0), 2)
	sessions = k.GetSessionsOfSubscription(ctx, hub.NewSubscriptionID(0))
	require.Equal(t, append([]types.Session{types.TestSession}, session), sessions)
}

func TestKeeper_GetAllSessions(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	sessions := k.GetAllSessions(ctx)
	require.Equal(t, []types.Session(nil), sessions)

	k.SetSession(ctx, types.TestSession)
	sessions = k.GetAllSessions(ctx)
	require.Equal(t, []types.Session{types.TestSession}, sessions)

	session := types.TestSession
	session.ID = hub.NewSessionID(1)
	k.SetSession(ctx, session)
	sessions = k.GetAllSessions(ctx)
	require.Equal(t, append([]types.Session{types.TestSession}, session), sessions)

	k.SetSession(ctx, types.TestSession)
	sessions = k.GetAllSessions(ctx)
	require.Equal(t, append([]types.Session{types.TestSession}, session), sessions)
}

func TestKeeper_AddSessionIDToActiveList(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	ids := k.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, hub.IDs(nil), ids)

	k.AddSessionIDToActiveList(ctx, 1, hub.NewSessionID(0))
	ids = k.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, hub.IDs{hub.NewSessionID(0)}, ids)

	k.AddSessionIDToActiveList(ctx, 1, hub.NewSessionID(1))
	ids = k.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, hub.IDs{hub.NewSessionID(0)}.Append(hub.NewSessionID(1)), ids)

	k.AddSessionIDToActiveList(ctx, 2, hub.NewSessionID(0))
	k.AddSessionIDToActiveList(ctx, 2, hub.NewSessionID(1))
	require.Equal(t, hub.IDs{hub.NewSessionID(0)}.Append(hub.NewSessionID(1)), ids)
}

func TestKeeper_RemoveSessionIDFromActiveList(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	ids := k.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, hub.IDs(nil), ids)

	k.AddSessionIDToActiveList(ctx, 1, hub.NewSessionID(0))
	k.RemoveSessionIDFromActiveList(ctx, 1, hub.NewSessionID(1))
	ids = k.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, hub.IDs{hub.NewSessionID(0)}, ids)

	k.RemoveSessionIDFromActiveList(ctx, 1, hub.NewSessionID(0))
	ids = k.GetActiveSessionIDs(ctx, 1)
	require.Equal(t, hub.IDs(nil), ids)

	k.AddSessionIDToActiveList(ctx, 3, hub.NewSessionID(0))
	k.AddSessionIDToActiveList(ctx, 3, hub.NewSessionID(1))
	k.RemoveSessionIDFromActiveList(ctx, 3, hub.NewSessionID(1))
	ids = k.GetActiveSessionIDs(ctx, 3)
	require.Equal(t, hub.IDs{hub.NewSessionID(0)}, ids)

	k.RemoveSessionIDFromActiveList(ctx, 3, hub.NewSessionID(0))
	ids = k.GetActiveSessionIDs(ctx, 3)
	require.Equal(t, hub.IDs(nil), ids)
}
