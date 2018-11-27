package vpn

import (
	"testing"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stretchr/testify/require"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	"strconv"
	"sort"
)

func TestKeeper_SetVPNDetails(t *testing.T) {
	cdc := codec.New()

	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	vpnsCount := keeper.GetVPNsCount(ctx)
	require.Equal(t, vpnsCount, uint64(0))

	keeper.SetVPNsCount(ctx, count)

	vpnsCount = keeper.GetVPNsCount(ctx)
	require.NotNil(t, vpnsCount)
	require.Equal(t, vpnsCount, count)

	vpnID1 := addr1.String() + "/" + strconv.Itoa(int(vpnsCount))

	vpnDetails := keeper.GetVPNDetails(ctx, vpnID1)
	require.Nil(t, vpnDetails)

	testVpnDetails := TestGetVPNDetails(vpnID1)
	keeper.SetVPNDetails(ctx, vpnID1, testVpnDetails)

	vpnDetails = keeper.GetVPNDetails(ctx, vpnID1)
	require.NotNil(t, vpnDetails)
	require.Equal(t, testVpnDetails, vpnDetails)

}

func TestKeeper_SetActiveNodeIDs(t *testing.T) {
	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	vpnsCount := keeper.GetVPNsCount(ctx)
	require.Equal(t, vpnsCount, uint64(0))

	keeper.SetVPNsCount(ctx, count)

	vpnsCount = keeper.GetVPNsCount(ctx)
	require.NotNil(t, vpnsCount)
	require.Equal(t, vpnsCount, count)

	vpnID1 := addr1.String() + "/" + strconv.Itoa(int(vpnsCount))
	vpnID2 := addr1.String() + "/" + strconv.Itoa(int(vpnsCount))

	nodes := keeper.GetActiveNodeIDs(ctx)
	require.Nil(t, nodes)

	newNodeIDs := []string{vpnID1, vpnID2}

	keeper.SetActiveNodeIDs(ctx, newNodeIDs)

	nodeIDs := keeper.GetActiveNodeIDs(ctx)
	require.NotNil(t, nodeIDs)
	require.Equal(t, newNodeIDs, nodeIDs)

}

func TestKeeper_SetVPNsCount(t *testing.T) {
	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	vpnsCount := keeper.GetVPNsCount(ctx)
	require.Equal(t, vpnsCount, uint64(0))

	keeper.SetVPNsCount(ctx, count)

	vpnsCount = keeper.GetVPNsCount(ctx)
	require.NotNil(t, vpnsCount)
	require.Equal(t, vpnsCount, count)

}

func TestKeeper_SetSessionDetails(t *testing.T) {

	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	sessionsCount := keeper.GetSessionsCount(ctx)
	require.Equal(t, sessionsCount, uint64(0))

	keeper.SetSessionsCount(ctx, count)

	sessionsCount = keeper.GetSessionsCount(ctx)
	require.NotNil(t, sessionsCount)
	require.Equal(t, sessionsCount, count)

	sessionID := addr1.String() + "/" + strconv.Itoa(int(sessionsCount))

	sessionDetails := keeper.GetSessionDetails(ctx, sessionID)
	require.Nil(t, sessionDetails)

	testSessionDetails := TestGetSessionDetails(sessionID)

	keeper.SetSessionDetails(ctx, sessionID, testSessionDetails)

	sessionDetails = keeper.GetSessionDetails(ctx, sessionID)
	require.NotNil(t, sessionDetails)
	//	require.Equal(t, testSessionDetails, sessionDetails)

}

func TestKeeper_SetActiveSessionIDs(t *testing.T) {
	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	sessionsCount := keeper.GetSessionsCount(ctx)
	require.Equal(t, sessionsCount, uint64(0))

	keeper.SetSessionsCount(ctx, count)

	sessionsCount = keeper.GetSessionsCount(ctx)
	require.NotNil(t, sessionsCount)
	require.Equal(t, sessionsCount, count)

	sessionID1 := addr1.String() + "/" + strconv.Itoa(int(sessionsCount))
	sessionID2 := addr1.String() + "/" + strconv.Itoa(int(sessionsCount))

	session := keeper.GetActiveSessionIDs(ctx)
	require.Nil(t, session)

	newSessionIDs := []string{sessionID1, sessionID2}

	keeper.SetActiveSessionIDs(ctx, newSessionIDs)

	sessionIDs := keeper.GetActiveSessionIDs(ctx)
	require.NotNil(t, sessionIDs)
	require.Equal(t, newSessionIDs, sessionIDs)

}

func TestKeeper_SetSessionsCount(t *testing.T) {
	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	sessionsCount := keeper.GetSessionsCount(ctx)
	require.Equal(t, sessionsCount, uint64(0))

	keeper.SetSessionsCount(ctx, count)

	newSessionsCount := keeper.GetSessionsCount(ctx)
	require.NotNil(t, newSessionsCount)
	require.Equal(t, count, newSessionsCount)

}

func TestKeeper_SetVPNStatus(t *testing.T) {
	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	vpnsCount := keeper.GetVPNsCount(ctx)
	require.Equal(t, vpnsCount, uint64(0))

	keeper.SetVPNsCount(ctx, count)

	vpnsCount = keeper.GetVPNsCount(ctx)
	require.NotNil(t, vpnsCount)
	require.Equal(t, vpnsCount, count)

	vpnID1 := addr1.String() + "/" + strconv.Itoa(int(vpnsCount))

	vpnDetails := keeper.GetVPNDetails(ctx, vpnID1)
	require.Nil(t, vpnDetails)

	testVpnDetails := TestGetVPNDetails(vpnID1)
	keeper.SetVPNDetails(ctx, vpnID1, testVpnDetails)

	vpnDetails = keeper.GetVPNDetails(ctx, vpnID1)
	require.NotNil(t, vpnDetails)
	require.Equal(t, testVpnDetails, vpnDetails)

	keeper.SetVPNStatus(ctx, vpnID1, status)

	vpnDetails = keeper.GetVPNDetails(ctx, vpnID1)
	require.NotNil(t, vpnDetails)
	require.Equal(t, status, vpnDetails.Status)

}

func TestKeeper_AddActiveNodeID(t *testing.T) {
	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	vpnsCount := keeper.GetVPNsCount(ctx)
	require.Equal(t, vpnsCount, uint64(0))

	keeper.SetVPNsCount(ctx, count)

	vpnsCount = keeper.GetVPNsCount(ctx)
	require.NotNil(t, vpnsCount)
	require.Equal(t, vpnsCount, count)

	vpnID1 := addr1.String() + "/" + strconv.Itoa(int(vpnsCount))
	vpnID2 := addr1.String() + "/" + strconv.Itoa(int(vpnsCount))

	nodes := keeper.GetActiveNodeIDs(ctx)
	require.Nil(t, nodes)

	nodeIDs := []string{vpnID1}

	keeper.SetActiveNodeIDs(ctx, nodeIDs)

	nodeIDs = keeper.GetActiveNodeIDs(ctx)
	require.NotNil(t, nodeIDs)
	require.Equal(t, nodeIDs[0], vpnID1)

	keeper.AddActiveNodeID(ctx, vpnID2)

	nodeIDs = keeper.GetActiveNodeIDs(ctx)
	require.NotNil(t, nodeIDs)
	require.Equal(t, nodeIDs[0], vpnID1)
	require.Equal(t, nodeIDs[1], vpnID2)

}

func TestKeeper_RemoveActiveNodeID(t *testing.T) {
	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	vpnsCount := keeper.GetVPNsCount(ctx)
	require.Equal(t, vpnsCount, uint64(0))

	keeper.SetVPNsCount(ctx, count)

	vpnsCount = keeper.GetVPNsCount(ctx)
	require.NotNil(t, vpnsCount)
	require.Equal(t, vpnsCount, count)

	vpnID1 := addr1.String() + "/" + strconv.Itoa(int(vpnsCount))
	vpnID2 := addr2.String() + "/" + strconv.Itoa(int(vpnsCount))

	nodes := keeper.GetActiveNodeIDs(ctx)
	require.Nil(t, nodes)

	nodeIDs := []string{vpnID1}

	keeper.SetActiveNodeIDs(ctx, nodeIDs)

	newNodeIDs := keeper.GetActiveNodeIDs(ctx)
	require.NotNil(t, newNodeIDs)
	require.Equal(t, newNodeIDs[0], vpnID1)

	keeper.AddActiveNodeID(ctx, vpnID2)

	nodes = append(nodeIDs, vpnID2)
	sort.Strings(nodes)

	newNodeIDs = keeper.GetActiveNodeIDs(ctx)
	require.Equal(t, newNodeIDs, nodes)

	keeper.RemoveActiveNodeID(ctx, vpnID2)
	newNodeIDs = keeper.GetActiveNodeIDs(ctx)

	require.Equal(t, newNodeIDs, nodeIDs)

}

func TestKeeper_SetSessionStatus(t *testing.T) {
	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	sessionsCount := keeper.GetSessionsCount(ctx)
	require.Equal(t, sessionsCount, uint64(0))

	keeper.SetSessionsCount(ctx, count)

	sessionsCount = keeper.GetSessionsCount(ctx)
	require.NotNil(t, sessionsCount)
	require.Equal(t, sessionsCount, count)

	sessionID1 := addr1.String() + "/" + strconv.Itoa(int(sessionsCount))

	session := keeper.GetActiveSessionIDs(ctx)
	require.Nil(t, session)

	newSessionIDs := []string{sessionID1}

	keeper.SetActiveSessionIDs(ctx, newSessionIDs)

	sessionIDs := keeper.GetActiveSessionIDs(ctx)
	require.NotNil(t, sessionIDs)
	require.Equal(t, newSessionIDs, sessionIDs)

	keeper.SetSessionStatus(ctx, sessionID1, status)

	sessionDetails := keeper.GetSessionDetails(ctx, sessionID1)
	require.NotNil(t, sessionDetails)
	require.Equal(t, sessionDetails.Status, status)

}

func TestKeeper_AddActiveSessionID(t *testing.T) {
	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	sessionsCount := keeper.GetSessionsCount(ctx)
	require.Equal(t, sessionsCount, uint64(0))

	keeper.SetSessionsCount(ctx, count)

	sessionsCount = keeper.GetSessionsCount(ctx)
	require.NotNil(t, sessionsCount)
	require.Equal(t, sessionsCount, count)

	sessionID1 := addr1.String() + "/" + strconv.Itoa(int(sessionsCount))
	sessionID2 := addr2.String() + "/" + strconv.Itoa(int(sessionsCount))

	session := keeper.GetActiveSessionIDs(ctx)
	require.Nil(t, session)

	newSessionIDs := []string{sessionID1}

	keeper.SetActiveSessionIDs(ctx, newSessionIDs)

	sessionIDs := keeper.GetActiveSessionIDs(ctx)
	require.NotNil(t, sessionIDs)
	require.Equal(t, newSessionIDs, sessionIDs)

	keeper.AddActiveSessionID(ctx, sessionID2)

	sessionIDs = keeper.GetActiveSessionIDs(ctx)
	require.NotNil(t, sessionIDs)
	require.Equal(t, sessionIDs[0], sessionID1)
	require.Equal(t, sessionIDs[1], sessionID2)

}

func TestKeeper_RemoveActiveSessionID(t *testing.T) {
	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	sessionsCount := keeper.GetSessionsCount(ctx)
	require.Equal(t, sessionsCount, uint64(0))

	keeper.SetSessionsCount(ctx, count)

	sessionsCount = keeper.GetSessionsCount(ctx)
	require.NotNil(t, sessionsCount)
	require.Equal(t, sessionsCount, count)

	sessionID1 := addr1.String() + "/" + strconv.Itoa(int(sessionsCount))
	sessionID2 := addr2.String() + "/" + strconv.Itoa(int(sessionsCount))

	session := keeper.GetActiveSessionIDs(ctx)
	require.Nil(t, session)

	newSessionIDs := []string{sessionID1}

	keeper.SetActiveSessionIDs(ctx, newSessionIDs)

	sessionIDs := keeper.GetActiveSessionIDs(ctx)
	require.NotNil(t, sessionIDs)
	require.Equal(t, newSessionIDs, sessionIDs)

	keeper.AddActiveSessionID(ctx, sessionID2)

	sessionIDs = keeper.GetActiveSessionIDs(ctx)
	require.NotNil(t, sessionIDs)
	require.Equal(t, sessionIDs[0], sessionID1)
	require.Equal(t, sessionIDs[1], sessionID2)

	keeper.RemoveActiveSessionID(ctx, sessionID2)

	sessionIDs = keeper.GetActiveSessionIDs(ctx)
	require.Equal(t, sessionIDs, newSessionIDs)
}
