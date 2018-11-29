package vpn

import (
	"sort"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	"fmt"
)

func TestKeeper_SetVPNDetails(t *testing.T) {
	cdc := codec.New()

	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	vpnsCount, err := keeper.GetVPNsCount(ctx)
	require.Nil(t, err)
	require.Equal(t, vpnsCount, uint64(0))

	keeper.SetVPNsCount(ctx, count)

	vpnsCount, err = keeper.GetVPNsCount(ctx)
	require.Nil(t, err)
	require.NotNil(t, vpnsCount)
	require.Equal(t, vpnsCount, count)

	vpnID1 := addr1.String() + "/" + strconv.Itoa(int(vpnsCount))

	vpnDetails, err := keeper.GetVPNDetails(ctx, vpnID1)
	require.Nil(t, err)
	require.Nil(t, vpnDetails)

	testVpnDetails := TestGetVPNDetails()
	keeper.SetVPNDetails(ctx, vpnID1, testVpnDetails)

	vpnDetails, err = keeper.GetVPNDetails(ctx, vpnID1)
	require.Nil(t, err)
	require.NotNil(t, vpnDetails)
	require.Equal(t, testVpnDetails, vpnDetails)

}

func TestKeeper_SetActiveNodeIDs(t *testing.T) {
	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	vpnsCount, err := keeper.GetVPNsCount(ctx)
	require.Nil(t, err)
	require.Equal(t, vpnsCount, uint64(0))

	keeper.SetVPNsCount(ctx, count)

	vpnsCount, err = keeper.GetVPNsCount(ctx)
	require.Nil(t, err)
	require.NotNil(t, vpnsCount)
	require.Equal(t, vpnsCount, count)

	vpnID1 := addr1.String() + "/" + strconv.Itoa(int(vpnsCount))
	vpnID2 := addr2.String() + "/" + strconv.Itoa(int(vpnsCount))

	nodes, err := keeper.GetActiveNodeIDs(ctx)
	require.Nil(t, err)
	require.Nil(t, nodes)

	newNodeIDs := []string{vpnID1, vpnID2}

	keeper.SetActiveNodeIDs(ctx, newNodeIDs)

	nodeIDs, err := keeper.GetActiveNodeIDs(ctx)
	require.Nil(t, err)
	require.NotNil(t, nodeIDs)
	require.Equal(t, newNodeIDs, nodeIDs)

}

func TestKeeper_SetVPNsCount(t *testing.T) {
	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	vpnsCount, err := keeper.GetVPNsCount(ctx)
	require.Nil(t, err)
	require.Equal(t, vpnsCount, uint64(0))

	keeper.SetVPNsCount(ctx, count)

	vpnsCount, err = keeper.GetVPNsCount(ctx)
	require.Nil(t, err)
	require.NotNil(t, vpnsCount)
	require.Equal(t, vpnsCount, count)

}

func TestKeeper_SetSessionDetails(t *testing.T) {

	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	sessionsCount, err := keeper.GetSessionsCount(ctx)
	require.Nil(t, err)
	require.Equal(t, sessionsCount, uint64(0))

	keeper.SetSessionsCount(ctx, count)

	sessionsCount, err = keeper.GetSessionsCount(ctx)
	require.Nil(t, err)
	require.NotNil(t, sessionsCount)
	require.Equal(t, sessionsCount, count)

	sessionID := addr1.String() + "/" + strconv.Itoa(int(sessionsCount))

	sessionDetails, err := keeper.GetSessionDetails(ctx, sessionID)
	require.Nil(t, err)
	require.Nil(t, sessionDetails)

	testSessionDetails := TestGetSessionDetails()

	keeper.SetSessionDetails(ctx, sessionID, testSessionDetails)

	sessionDetails, err = keeper.GetSessionDetails(ctx, sessionID)
	require.Nil(t, err)
	require.NotNil(t, sessionDetails)
	//	require.Equal(t, testSessionDetails, sessionDetails)

}

func TestKeeper_SetActiveSessionIDs(t *testing.T) {
	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	sessionsCount, err := keeper.GetSessionsCount(ctx)
	require.Nil(t, err)
	require.Equal(t, sessionsCount, uint64(0))

	keeper.SetSessionsCount(ctx, count)

	sessionsCount, err = keeper.GetSessionsCount(ctx)
	require.Nil(t, err)
	require.NotNil(t, sessionsCount)
	require.Equal(t, sessionsCount, count)

	sessionID1 := addr1.String() + "/" + strconv.Itoa(int(sessionsCount))
	sessionID2 := addr2.String() + "/" + strconv.Itoa(int(sessionsCount))

	session, err := keeper.GetActiveSessionIDs(ctx)
	require.Nil(t, err)
	require.Nil(t, session)

	newSessionIDs := []string{sessionID1, sessionID2}

	keeper.SetActiveSessionIDs(ctx, newSessionIDs)

	sessionIDs, err := keeper.GetActiveSessionIDs(ctx)
	require.Nil(t, err)
	require.NotNil(t, sessionIDs)
	require.Equal(t, newSessionIDs, sessionIDs)

}

func TestKeeper_SetSessionsCount(t *testing.T) {
	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	sessionsCount, err := keeper.GetSessionsCount(ctx)
	require.Nil(t, err)
	require.Equal(t, sessionsCount, uint64(0))

	keeper.SetSessionsCount(ctx, count)

	newSessionsCount, err := keeper.GetSessionsCount(ctx)
	require.Nil(t, err)
	require.NotNil(t, newSessionsCount)
	require.Equal(t, count, newSessionsCount)

}

func TestKeeper_SetVPNStatus(t *testing.T) {
	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	vpnsCount, err := keeper.GetVPNsCount(ctx)
	require.Nil(t, err)
	require.Equal(t, vpnsCount, uint64(0))

	keeper.SetVPNsCount(ctx, count)

	vpnsCount, err = keeper.GetVPNsCount(ctx)
	require.Nil(t, err)
	require.NotNil(t, vpnsCount)
	require.Equal(t, vpnsCount, count)

	vpnID1 := addr1.String() + "/" + strconv.Itoa(int(vpnsCount))

	vpnDetails, err := keeper.GetVPNDetails(ctx, vpnID1)
	require.Nil(t, err)
	require.Nil(t, vpnDetails)

	testVpnDetails := TestGetVPNDetails()
	keeper.SetVPNDetails(ctx, vpnID1, testVpnDetails)

	vpnDetails, err = keeper.GetVPNDetails(ctx, vpnID1)
	require.Nil(t, err)
	require.NotNil(t, vpnDetails)
	require.Equal(t, testVpnDetails, vpnDetails)

	keeper.SetVPNStatus(ctx, vpnID1, status)

	vpnDetails, err = keeper.GetVPNDetails(ctx, vpnID1)
	require.Nil(t, err)
	require.NotNil(t, vpnDetails)
	require.Equal(t, status, vpnDetails.Status)

}

func TestKeeper_AddActiveNodeID(t *testing.T) {
	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	vpnsCount, err := keeper.GetVPNsCount(ctx)
	require.Nil(t, err)
	require.Equal(t, vpnsCount, uint64(0))

	keeper.SetVPNsCount(ctx, count)

	vpnsCount, err = keeper.GetVPNsCount(ctx)
	require.Nil(t, err)
	require.NotNil(t, vpnsCount)
	require.Equal(t, vpnsCount, count)

	vpnID1 := addr1.String() + "/" + strconv.Itoa(int(vpnsCount))
	vpnID2 := addr2.String() + "/" + strconv.Itoa(int(vpnsCount))

	nodes, err := keeper.GetActiveNodeIDs(ctx)
	require.Nil(t, err)
	require.Nil(t, nodes)

	nodeIDs := []string{vpnID1}

	keeper.SetActiveNodeIDs(ctx, nodeIDs)

	nodeIDs, err = keeper.GetActiveNodeIDs(ctx)
	require.Nil(t, err)
	require.NotNil(t, nodeIDs)
	require.Equal(t, nodeIDs[0], vpnID1)

	keeper.AddActiveNodeID(ctx, vpnID2)

	nodeIDs, err = keeper.GetActiveNodeIDs(ctx)
	fmt.Println(nodeIDs)
	nodes = []string{vpnID1, vpnID2}
	sort.Strings(nodes)
	require.Nil(t, err)
	require.NotNil(t, nodeIDs)
	require.Equal(t, nodeIDs, nodes)

}

func TestKeeper_RemoveActiveNodeID(t *testing.T) {
	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	vpnsCount, err := keeper.GetVPNsCount(ctx)
	require.Nil(t, err)
	require.Equal(t, vpnsCount, uint64(0))

	keeper.SetVPNsCount(ctx, count)

	vpnsCount, err = keeper.GetVPNsCount(ctx)
	require.Nil(t, err)
	require.NotNil(t, vpnsCount)
	require.Equal(t, vpnsCount, count)

	vpnID1 := addr1.String() + "/" + strconv.Itoa(int(vpnsCount))
	vpnID2 := addr2.String() + "/" + strconv.Itoa(int(vpnsCount))

	nodes, err := keeper.GetActiveNodeIDs(ctx)
	require.Nil(t, err)
	require.Nil(t, nodes)

	nodeIDs := []string{vpnID1}

	keeper.SetActiveNodeIDs(ctx, nodeIDs)

	newNodeIDs, err := keeper.GetActiveNodeIDs(ctx)
	require.Nil(t, err)
	require.NotNil(t, newNodeIDs)
	require.Equal(t, newNodeIDs[0], vpnID1)

	keeper.AddActiveNodeID(ctx, vpnID2)

	nodes = append(nodeIDs, vpnID2)
	sort.Strings(nodes)

	newNodeIDs, err = keeper.GetActiveNodeIDs(ctx)
	require.Nil(t, err)
	require.Equal(t, newNodeIDs, nodes)

	keeper.RemoveActiveNodeID(ctx, vpnID2)
	newNodeIDs, err = keeper.GetActiveNodeIDs(ctx)
	require.Nil(t, err)

	require.Equal(t, newNodeIDs, nodeIDs)

}

func TestKeeper_SetSessionStatus(t *testing.T) {
	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	sessionsCount, err := keeper.GetSessionsCount(ctx)
	require.Nil(t, err)
	require.Equal(t, sessionsCount, uint64(0))

	keeper.SetSessionsCount(ctx, count)

	sessionsCount, err = keeper.GetSessionsCount(ctx)
	require.Nil(t, err)
	require.NotNil(t, sessionsCount)
	require.Equal(t, sessionsCount, count)

	sessionID1 := addr1.String() + "/" + strconv.Itoa(int(sessionsCount))

	session, err := keeper.GetActiveSessionIDs(ctx)
	require.Nil(t, err)
	require.Nil(t, session)

	newSessionIDs := []string{sessionID1}

	keeper.SetActiveSessionIDs(ctx, newSessionIDs)

	sessionIDs, err := keeper.GetActiveSessionIDs(ctx)
	require.Nil(t, err)
	require.NotNil(t, sessionIDs)
	require.Equal(t, newSessionIDs, sessionIDs)

	sessionDetails, err := keeper.GetSessionDetails(ctx, sessionID1)
	require.Nil(t, err)
	require.Nil(t, sessionDetails)

	testSessionDetails := TestGetSessionDetails()

	keeper.SetSessionDetails(ctx, sessionID1, testSessionDetails)

	sessionDetails, err = keeper.GetSessionDetails(ctx, sessionID1)
	require.Nil(t, err)
	require.NotNil(t, sessionDetails)

	keeper.SetSessionStatus(ctx, sessionID1, status)

	sessionDetails, err = keeper.GetSessionDetails(ctx, sessionID1)
	require.Nil(t, err)
	require.NotNil(t, sessionDetails)
	require.Equal(t, sessionDetails.Status, status)

}

func TestKeeper_AddActiveSessionID(t *testing.T) {
	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	sessionsCount, err := keeper.GetSessionsCount(ctx)
	require.Nil(t, err)
	require.Equal(t, sessionsCount, uint64(0))

	keeper.SetSessionsCount(ctx, count)

	sessionsCount, err = keeper.GetSessionsCount(ctx)
	require.NotNil(t, sessionsCount)
	require.Equal(t, sessionsCount, count)

	sessionID1 := addr1.String() + "/" + strconv.Itoa(int(sessionsCount))
	sessionID2 := addr2.String() + "/" + strconv.Itoa(int(sessionsCount))

	session, err := keeper.GetActiveSessionIDs(ctx)
	require.Nil(t, err)
	require.Nil(t, session)

	newSessionIDs := []string{sessionID1}

	keeper.SetActiveSessionIDs(ctx, newSessionIDs)

	sessionIDs, err := keeper.GetActiveSessionIDs(ctx)
	require.Nil(t, err)
	require.NotNil(t, sessionIDs)
	require.Equal(t, newSessionIDs, sessionIDs)

	keeper.AddActiveSessionID(ctx, sessionID2)

	sessionIDs, err = keeper.GetActiveSessionIDs(ctx)
	require.NotNil(t, sessionIDs)
	sessions := []string{sessionID1, sessionID2}
	sort.Strings(sessions)
	require.Equal(t, sessionIDs, sessions)

}

func TestKeeper_RemoveActiveSessionID(t *testing.T) {
	cdc := codec.New()
	ms, _, vpnStoreKey, sessionStoreKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, vpnStoreKey, sessionStoreKey)

	sessionsCount, err := keeper.GetSessionsCount(ctx)
	require.Nil(t, err)
	require.Equal(t, sessionsCount, uint64(0))

	keeper.SetSessionsCount(ctx, count)

	sessionsCount, err = keeper.GetSessionsCount(ctx)
	require.Nil(t, err)
	require.NotNil(t, sessionsCount)
	require.Equal(t, sessionsCount, count)

	sessionID1 := addr1.String() + "/" + strconv.Itoa(int(sessionsCount))
	sessionID2 := addr2.String() + "/" + strconv.Itoa(int(sessionsCount))

	session, err := keeper.GetActiveSessionIDs(ctx)
	require.Nil(t, err)
	require.Nil(t, session)

	newSessionIDs := []string{sessionID1}

	keeper.SetActiveSessionIDs(ctx, newSessionIDs)

	sessionIDs, err := keeper.GetActiveSessionIDs(ctx)
	require.Nil(t, err)
	require.NotNil(t, sessionIDs)
	require.Equal(t, newSessionIDs, sessionIDs)

	keeper.AddActiveSessionID(ctx, sessionID2)

	sessionIDs, err = keeper.GetActiveSessionIDs(ctx)
	require.Nil(t, err)
	require.NotNil(t, sessionIDs)
	sessions := []string{sessionID1, sessionID2}
	sort.Strings(sessions)
	require.Equal(t, sessionIDs, sessions)

	keeper.RemoveActiveSessionID(ctx, sessionID2)

	sessionIDs, err = keeper.GetActiveSessionIDs(ctx)
	require.Nil(t, err)
	require.Equal(t, sessionIDs, newSessionIDs)
}
