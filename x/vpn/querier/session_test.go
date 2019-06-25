package querier

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/sentinel-hub/x/vpn/keeper"
	hub "github.com/sentinel-official/sentinel-hub/x/vpn/types"
)

func TestNewQuerySessionParams(t *testing.T) {
	params := NewQuerySessionParams(hub.TestIDZero)
	require.Equal(t, TestSessionParamsZero, params)

	params = NewQuerySessionParams(hub.TestIDPos)
	require.Equal(t, TestSessionParamsPos, params)
}

func Test_querySession(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var session hub.Session

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", hub.QuerierRoute, QuerySession),
		Data: []byte{},
	}

	res, _err := querySession(ctx, cdc, req, vpnKeeper)
	require.NotNil(t, _err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)

	vpnKeeper.SetSession(ctx, hub.TestSessionValid)
	req.Data, err = cdc.MarshalJSON(NewQuerySessionParams(hub.TestIDZero))
	require.Nil(t, err)

	res, _err = querySession(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &session)
	require.Nil(t, err)
	require.Equal(t, hub.TestSessionValid, session)

	req.Data, err = cdc.MarshalJSON(NewQuerySessionParams(hub.TestIDPos))
	require.Nil(t, err)

	res, _err = querySession(ctx, cdc, req, vpnKeeper)
	require.Nil(t, res)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)
}

func TestNewQuerySessionOfSubscriptionPrams(t *testing.T) {
	params := NewQuerySessionOfSubscriptionPrams(hub.TestIDZero, 0)
	require.Equal(t, TestSessionOfSubscriptionPramsZero, params)

	params = NewQuerySessionOfSubscriptionPrams(hub.TestIDPos, 0)
	require.Equal(t, TestSessionOfSubscriptionPramsPos, params)
}

func Test_querySessionOfSubscription(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var session hub.Session

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", hub.QuerierRoute, QuerySessionOfSubscription),
		Data: []byte{},
	}

	res, _err := querySessionOfSubscription(ctx, cdc, req, vpnKeeper)
	require.NotNil(t, _err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)

	vpnKeeper.SetSubscription(ctx, hub.TestSubscriptionValid)
	vpnKeeper.SetSessionIDBySubscriptionID(ctx, hub.TestSessionValid.ID, 0, hub.TestSubscriptionValid.ID)
	vpnKeeper.SetSessionsCountOfSubscription(ctx, hub.TestSubscriptionValid.ID, 1)

	req.Data, err = cdc.MarshalJSON(NewQuerySessionsOfSubscriptionPrams(hub.TestIDZero))
	require.Nil(t, err)

	res, _err = querySessionOfSubscription(ctx, cdc, req, vpnKeeper)
	require.Nil(t, err)
	require.Equal(t, []byte(nil), res)
	require.Nil(t, res)

	err = cdc.UnmarshalJSON(res, &session)
	require.NotEqual(t, hub.TestSessionValid, session)

	vpnKeeper.SetSession(ctx, hub.TestSessionValid)
	req.Data, err = cdc.MarshalJSON(NewQuerySessionsOfSubscriptionPrams(hub.TestIDZero))
	require.Nil(t, err)

	res, _err = querySessionOfSubscription(ctx, cdc, req, vpnKeeper)
	require.Nil(t, err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &session)
	require.Nil(t, err)
	require.Equal(t, hub.TestSessionValid, session)

	req.Data, err = cdc.MarshalJSON(NewQuerySessionsOfSubscriptionPrams(hub.TestIDPos))
	require.Nil(t, err)

	res, _err = querySessionOfSubscription(ctx, cdc, req, vpnKeeper)
	require.Nil(t, err)
	require.Equal(t, []byte(nil), res)
	require.Nil(t, res)

	err = cdc.UnmarshalJSON(res, &session)
	require.NotNil(t, err)
}

func TestNewQuerySessionsOfSubscriptionPrams(t *testing.T) {
	params := NewQuerySessionsOfSubscriptionPrams(hub.TestIDZero)
	require.Equal(t, TestSessionsOfSubscriptionPramsZero, params)

	params = NewQuerySessionsOfSubscriptionPrams(hub.TestIDPos)
	require.Equal(t, TestSessionsOfSubscriptionPramsPos, params)
}

func Test_querySessionsOfSubscription(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var sessions []hub.Session

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", hub.QuerierRoute, QuerySessionsOfSubscription),
		Data: []byte{},
	}

	res, _err := querySessionsOfSubscription(ctx, cdc, req, vpnKeeper)
	require.NotNil(t, _err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)

	vpnKeeper.SetSubscription(ctx, hub.TestSubscriptionValid)
	vpnKeeper.SetSession(ctx, hub.TestSessionValid)
	vpnKeeper.SetSessionIDBySubscriptionID(ctx, hub.TestSubscriptionValid.ID, 0, hub.TestSessionValid.ID)
	vpnKeeper.SetSessionsCountOfSubscription(ctx, hub.TestSubscriptionValid.ID, 1)

	req.Data, err = cdc.MarshalJSON(NewQuerySessionsOfSubscriptionPrams(hub.TestIDZero))
	require.Nil(t, err)

	res, _err = querySessionsOfSubscription(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.Equal(t, hub.TestSessionsValid, sessions)

	session := hub.TestSessionValid
	session.ID = hub.TestIDPos
	vpnKeeper.SetSession(ctx, session)
	vpnKeeper.SetSessionIDBySubscriptionID(ctx, hub.TestSubscriptionValid.ID, 0, session.ID)
	vpnKeeper.SetSessionsCountOfSubscription(ctx, hub.TestSubscriptionValid.ID, 2)

	res, _err = querySessionsOfSubscription(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.Len(t, sessions, 2)
}

func Test_queryAllSessions(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var sessions []hub.Session

	res, _err := queryAllSessions(ctx, cdc, vpnKeeper)
	require.Nil(t, _err)
	require.Equal(t, []byte("null"), res)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.NotEqual(t, hub.TestSessionsValid, sessions)

	vpnKeeper.SetSession(ctx, hub.TestSessionEmpty)
	res, _err = queryAllSessions(ctx, cdc, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.NotEqual(t, hub.TestSessionsValid, sessions)

	vpnKeeper.SetSession(ctx, hub.TestSessionValid)
	require.Nil(t, err)

	res, _err = queryAllSessions(ctx, cdc, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.Equal(t, hub.TestSessionsValid, sessions)

	session := hub.TestSessionValid
	session.ID = hub.TestIDPos
	vpnKeeper.SetSession(ctx, session)
	require.Nil(t, err)

	res, _err = queryAllSessions(ctx, cdc, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.Equal(t, append(hub.TestSessionsValid, session), sessions)
}
