package vpn

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

func TestKeeper_SetVPNDetails(t *testing.T) {
	cdc := codec.New()

	ms, _, vpnKey, sessionKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abciTypes.Header{}, false, log.NewNopLogger())
	vpnKeeper := NewKeeper(cdc, vpnKey, sessionKey)

	var err csdkTypes.Error

	vpnDetails0, err := vpnKeeper.GetVPNDetails(ctx, "vpn_id")
	require.Nil(t, err)
	require.Nil(t, vpnDetails0)

	vpnDetails1 := vpnDetails

	err = vpnKeeper.SetVPNDetails(ctx, "vpn_id", &vpnDetails1)
	require.Nil(t, err)

	vpnDetails2, err := vpnKeeper.GetVPNDetails(ctx, "vpn_id")
	require.Nil(t, err)
	require.Equal(t, &vpnDetails1, vpnDetails2)
}

func TestKeeper_SetActiveNodeIDs(t *testing.T) {
	cdc := codec.New()

	ms, _, vpnKey, sessionKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abciTypes.Header{}, false, log.NewNopLogger())
	vpnKeeper := NewKeeper(cdc, vpnKey, sessionKey)

	var err csdkTypes.Error

	activeNodeIDs, err := vpnKeeper.GetActiveNodeIDs(ctx)
	require.Nil(t, err)
	require.Equal(t, []string(nil), activeNodeIDs)

	err = vpnKeeper.SetActiveNodeIDs(ctx, []string{"node_id_0", "node_id_1"})
	require.Nil(t, err)

	activeNodeIDs, err = vpnKeeper.GetActiveNodeIDs(ctx)
	require.Nil(t, err)
	require.Equal(t, []string{"node_id_0", "node_id_1"}, activeNodeIDs)
}

func TestKeeper_SetVPNsCount(t *testing.T) {
	cdc := codec.New()

	ms, _, vpnKey, sessionKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abciTypes.Header{}, false, log.NewNopLogger())
	vpnKeeper := NewKeeper(cdc, vpnKey, sessionKey)

	var err csdkTypes.Error

	vpnsCount, err := vpnKeeper.GetVPNsCount(ctx, accAddress1)
	require.Nil(t, err)
	require.Equal(t, uint64(0), vpnsCount)

	err = vpnKeeper.SetVPNsCount(ctx, accAddress1, uint64(1))
	require.Nil(t, err)

	vpnsCount, err = vpnKeeper.GetVPNsCount(ctx, accAddress1)
	require.Nil(t, err)
	require.Equal(t, uint64(1), vpnsCount)
}

func TestKeeper_SetSessionDetails(t *testing.T) {
	cdc := codec.New()

	ms, _, vpnKey, sessionKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abciTypes.Header{}, false, log.NewNopLogger())
	vpnKeeper := NewKeeper(cdc, vpnKey, sessionKey)

	var err csdkTypes.Error

	sessionDetails0, err := vpnKeeper.GetSessionDetails(ctx, "session_id")
	require.Nil(t, err)
	require.Nil(t, sessionDetails0)

	sessionDetails1 := sessionDetails

	err = vpnKeeper.SetSessionDetails(ctx, "session_id", &sessionDetails1)
	require.Nil(t, err)

	sessionDetails2, err := vpnKeeper.GetSessionDetails(ctx, "session_id")
	require.Nil(t, err)
	require.Equal(t, &sessionDetails1, sessionDetails2)
}

func TestKeeper_SetActiveSessionIDs(t *testing.T) {
	cdc := codec.New()

	ms, _, vpnKey, sessionKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abciTypes.Header{}, false, log.NewNopLogger())
	vpnKeeper := NewKeeper(cdc, vpnKey, sessionKey)

	var err csdkTypes.Error

	activeSessionIDs, err := vpnKeeper.GetActiveSessionIDs(ctx)
	require.Nil(t, err)
	require.Equal(t, []string(nil), activeSessionIDs)

	err = vpnKeeper.SetActiveSessionIDs(ctx, []string{"session_id_0", "session_id_1"})
	require.Nil(t, err)

	activeSessionIDs, err = vpnKeeper.GetActiveSessionIDs(ctx)
	require.Nil(t, err)
	require.Equal(t, []string{"session_id_0", "session_id_1"}, activeSessionIDs)
}

func TestKeeper_SetSessionsCount(t *testing.T) {
	cdc := codec.New()

	ms, _, vpnKey, sessionKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abciTypes.Header{}, false, log.NewNopLogger())
	vpnKeeper := NewKeeper(cdc, vpnKey, sessionKey)

	var err csdkTypes.Error

	sessionsCount, err := vpnKeeper.GetSessionsCount(ctx, accAddress1)
	require.Nil(t, err)
	require.Equal(t, uint64(0), sessionsCount)

	err = vpnKeeper.SetSessionsCount(ctx, accAddress1, uint64(1))
	require.Nil(t, err)

	sessionsCount, err = vpnKeeper.GetSessionsCount(ctx, accAddress1)
	require.Nil(t, err)
	require.Equal(t, uint64(1), sessionsCount)
}

func TestKeeper_SetVPNStatus(t *testing.T) {
	cdc := codec.New()

	ms, _, vpnKey, sessionKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abciTypes.Header{}, false, log.NewNopLogger())
	vpnKeeper := NewKeeper(cdc, vpnKey, sessionKey)

	var err csdkTypes.Error

	vpnDetails0, err := vpnKeeper.GetVPNDetails(ctx, "vpn_id")
	require.Nil(t, err)
	require.Nil(t, vpnDetails0)

	vpnDetails1 := vpnDetails

	err = vpnKeeper.SetVPNDetails(ctx, "vpn_id", &vpnDetails1)
	require.Nil(t, err)

	vpnDetails2, err := vpnKeeper.GetVPNDetails(ctx, "vpn_id")
	require.Nil(t, err)
	require.Equal(t, &vpnDetails1, vpnDetails2)

	err = vpnKeeper.SetVPNStatus(ctx, "vpn_id", sdkTypes.StatusActive)
	require.Nil(t, err)

	vpnDetails1.Status = sdkTypes.StatusActive
	vpnDetails3, err := vpnKeeper.GetVPNDetails(ctx, "vpn_id")
	require.Nil(t, err)
	require.Equal(t, &vpnDetails1, vpnDetails3)
}

func TestKeeper_AddActiveNodeID(t *testing.T) {
	cdc := codec.New()

	ms, _, vpnKey, sessionKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abciTypes.Header{}, false, log.NewNopLogger())
	vpnKeeper := NewKeeper(cdc, vpnKey, sessionKey)

	var err csdkTypes.Error

	activeNodeIDs, err := vpnKeeper.GetActiveNodeIDs(ctx)
	require.Nil(t, err)
	require.Equal(t, []string(nil), activeNodeIDs)

	err = vpnKeeper.AddActiveNodeID(ctx, "node_id")
	require.Nil(t, err)

	activeNodeIDs, err = vpnKeeper.GetActiveNodeIDs(ctx)
	require.Nil(t, err)
	require.Equal(t, []string{"node_id"}, activeNodeIDs)
}

func TestKeeper_RemoveActiveNodeID(t *testing.T) {
	cdc := codec.New()

	ms, _, vpnKey, sessionKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abciTypes.Header{}, false, log.NewNopLogger())
	vpnKeeper := NewKeeper(cdc, vpnKey, sessionKey)

	var err csdkTypes.Error

	err = vpnKeeper.AddActiveNodeID(ctx, "node_id_0")
	require.Nil(t, err)

	err = vpnKeeper.AddActiveNodeID(ctx, "node_id_1")
	require.Nil(t, err)

	activeNodeIDs, err := vpnKeeper.GetActiveNodeIDs(ctx)
	require.Nil(t, err)
	require.Equal(t, []string{"node_id_0", "node_id_1"}, activeNodeIDs)

	err = vpnKeeper.RemoveActiveNodeID(ctx, "node_id_0")
	require.Nil(t, err)

	activeNodeIDs, err = vpnKeeper.GetActiveNodeIDs(ctx)
	require.Nil(t, err)
	require.Equal(t, []string{"node_id_1"}, activeNodeIDs)
}

func TestKeeper_SetSessionStatus(t *testing.T) {
	cdc := codec.New()

	ms, _, vpnKey, sessionKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abciTypes.Header{}, false, log.NewNopLogger())
	vpnKeeper := NewKeeper(cdc, vpnKey, sessionKey)

	var err csdkTypes.Error

	sessionDetails0, err := vpnKeeper.GetSessionDetails(ctx, "session_id")
	require.Nil(t, err)
	require.Nil(t, sessionDetails0)

	sessionDetails1 := sessionDetails

	err = vpnKeeper.SetSessionDetails(ctx, "session_id", &sessionDetails1)
	require.Nil(t, err)

	sessionDetails2, err := vpnKeeper.GetSessionDetails(ctx, "session_id")
	require.Nil(t, err)
	require.Equal(t, &sessionDetails1, sessionDetails2)

	err = vpnKeeper.SetSessionStatus(ctx, "session_id", sdkTypes.StatusActive)
	require.Nil(t, err)

	sessionDetails1.Status = sdkTypes.StatusActive
	sessionDetails2, err = vpnKeeper.GetSessionDetails(ctx, "session_id")
	require.Nil(t, err)
	require.Equal(t, &sessionDetails1, sessionDetails2)
}

func TestKeeper_AddActiveSessionID(t *testing.T) {
	cdc := codec.New()

	ms, _, vpnKey, sessionKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abciTypes.Header{}, false, log.NewNopLogger())
	vpnKeeper := NewKeeper(cdc, vpnKey, sessionKey)

	var err csdkTypes.Error

	activeSessionIDs, err := vpnKeeper.GetActiveSessionIDs(ctx)
	require.Nil(t, err)
	require.Equal(t, []string(nil), activeSessionIDs)

	err = vpnKeeper.AddActiveSessionID(ctx, "session_id")
	require.Nil(t, err)

	activeSessionIDs, err = vpnKeeper.GetActiveSessionIDs(ctx)
	require.Nil(t, err)
	require.Equal(t, []string{"session_id"}, activeSessionIDs)
}

func TestKeeper_RemoveActiveSessionID(t *testing.T) {
	cdc := codec.New()

	ms, _, vpnKey, sessionKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abciTypes.Header{}, false, log.NewNopLogger())
	vpnKeeper := NewKeeper(cdc, vpnKey, sessionKey)

	var err csdkTypes.Error

	err = vpnKeeper.AddActiveSessionID(ctx, "session_id_0")
	require.Nil(t, err)

	err = vpnKeeper.AddActiveSessionID(ctx, "session_id_1")
	require.Nil(t, err)

	activeSessionIDs, err := vpnKeeper.GetActiveSessionIDs(ctx)
	require.Nil(t, err)
	require.Equal(t, []string{"session_id_0", "session_id_1"}, activeSessionIDs)

	err = vpnKeeper.RemoveActiveSessionID(ctx, "session_id_0")
	require.Nil(t, err)

	activeSessionIDs, err = vpnKeeper.GetActiveSessionIDs(ctx)
	require.Nil(t, err)
	require.Equal(t, []string{"session_id_1"}, activeSessionIDs)
}

func TestKeeper_AddVPN(t *testing.T) {
	cdc := codec.New()

	ms, _, vpnKey, sessionKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abciTypes.Header{}, false, log.NewNopLogger())
	vpnKeeper := NewKeeper(cdc, vpnKey, sessionKey)

	var err csdkTypes.Error

	vpnsCount, err := vpnKeeper.GetVPNsCount(ctx, accAddress1)
	require.Nil(t, err)
	require.Equal(t, uint64(0), vpnsCount)

	vpnDetails0, err := vpnKeeper.GetVPNDetails(ctx, "vpn_id")
	require.Nil(t, err)
	require.Nil(t, vpnDetails0)

	vpnDetails1 := vpnDetails

	err = vpnKeeper.AddVPN(ctx, "vpn_id", &vpnDetails1)
	require.Nil(t, err)

	vpnsCount, err = vpnKeeper.GetVPNsCount(ctx, accAddress1)
	require.Nil(t, err)
	require.Equal(t, uint64(1), vpnsCount)

	vpnDetails2, err := vpnKeeper.GetVPNDetails(ctx, "vpn_id")
	require.Nil(t, err)
	require.Equal(t, &vpnDetails1, vpnDetails2)
}

func TestKeeper_AddSession(t *testing.T) {
	cdc := codec.New()

	ms, _, vpnKey, sessionKey := setupMultiStore()
	ctx := csdkTypes.NewContext(ms, abciTypes.Header{}, false, log.NewNopLogger())
	vpnKeeper := NewKeeper(cdc, vpnKey, sessionKey)

	var err csdkTypes.Error

	sessionsCount, err := vpnKeeper.GetSessionsCount(ctx, accAddress1)
	require.Nil(t, err)
	require.Equal(t, uint64(0), sessionsCount)

	sessionDetails0, err := vpnKeeper.GetSessionDetails(ctx, "session_id")
	require.Nil(t, err)
	require.Nil(t, sessionDetails0)

	sessionDetails1 := sessionDetails

	err = vpnKeeper.AddSession(ctx, "session_id", &sessionDetails1)
	require.Nil(t, err)

	sessionsCount, err = vpnKeeper.GetSessionsCount(ctx, accAddress1)
	require.Nil(t, err)
	require.Equal(t, uint64(1), sessionsCount)

	sessionDetails2, err := vpnKeeper.GetSessionDetails(ctx, "session_id")
	require.Nil(t, err)
	require.Equal(t, &sessionDetails1, sessionDetails2)
}
