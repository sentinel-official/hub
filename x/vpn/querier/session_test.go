package querier

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func Test_querySession(t *testing.T) {
	ctx, k, _, _ := keeper.CreateTestInput(t, false)
	cdc := keeper.MakeTestCodec()
	var err error
	var session types.Session

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySession),
		Data: []byte{},
	}

	res, _err := querySession(ctx, req, k)
	require.NotNil(t, _err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)

	k.SetSession(ctx, types.TestSession)
	req.Data, err = cdc.MarshalJSON(types.NewQuerySessionParams(hub.NewSessionID(0)))
	require.Nil(t, err)

	res, _err = querySession(ctx, req, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &session)
	require.Nil(t, err)
	require.Equal(t, types.TestSession, session)

	req.Data, err = cdc.MarshalJSON(types.NewQuerySessionParams(hub.NewSessionID(1)))
	require.Nil(t, err)

	res, _err = querySession(ctx, req, k)
	require.Nil(t, res)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)
}

func Test_querySessionOfSubscription(t *testing.T) {
	ctx, k, _, _ := keeper.CreateTestInput(t, false)
	cdc := keeper.MakeTestCodec()
	var err error
	var session types.Session

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySessionOfSubscription),
		Data: []byte{},
	}

	res, _err := querySessionOfSubscription(ctx, req, k)
	require.NotNil(t, _err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)

	k.SetSubscription(ctx, types.TestSubscription)
	k.SetSessionIDBySubscriptionID(ctx, types.TestSubscription.ID, 0, types.TestSession.ID)
	k.SetSessionsCountOfSubscription(ctx, types.TestSubscription.ID, 1)

	req.Data, err = cdc.MarshalJSON(types.NewQuerySessionsOfSubscriptionPrams(hub.NewSubscriptionID(0)))
	require.Nil(t, err)

	res, _err = querySessionOfSubscription(ctx, req, k)
	require.Nil(t, err)
	require.Equal(t, []byte(nil), res)
	require.Nil(t, res)

	err = cdc.UnmarshalJSON(res, &session)
	require.NotEqual(t, types.TestSession, session)

	k.SetSession(ctx, types.TestSession)
	req.Data, err = cdc.MarshalJSON(types.NewQuerySessionsOfSubscriptionPrams(hub.NewSubscriptionID(0)))
	require.Nil(t, err)

	res, _err = querySessionOfSubscription(ctx, req, k)
	require.Nil(t, err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &session)
	require.Nil(t, err)
	require.Equal(t, types.TestSession, session)

	req.Data, err = cdc.MarshalJSON(types.NewQuerySessionsOfSubscriptionPrams(hub.NewSubscriptionID(1)))
	require.Nil(t, err)

	res, _err = querySessionOfSubscription(ctx, req, k)
	require.Nil(t, err)
	require.Equal(t, []byte(nil), res)
	require.Nil(t, res)

	err = cdc.UnmarshalJSON(res, &session)
	require.NotNil(t, err)
}

func Test_querySessionsOfSubscription(t *testing.T) {
	ctx, k, _, _ := keeper.CreateTestInput(t, false)
	cdc := keeper.MakeTestCodec()
	var err error
	var sessions []types.Session

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySessionsOfSubscription),
		Data: []byte{},
	}

	res, _err := querySessionsOfSubscription(ctx, req, k)
	require.NotNil(t, _err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)

	k.SetSubscription(ctx, types.TestSubscription)
	k.SetSession(ctx, types.TestSession)
	k.SetSessionIDBySubscriptionID(ctx, types.TestSubscription.ID, 0, types.TestSession.ID)
	k.SetSessionsCountOfSubscription(ctx, types.TestSubscription.ID, 1)

	req.Data, err = cdc.MarshalJSON(types.NewQuerySessionsOfSubscriptionPrams(hub.NewSubscriptionID(0)))
	require.Nil(t, err)

	res, _err = querySessionsOfSubscription(ctx, req, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.Equal(t, []types.Session{types.TestSession}, sessions)

	session := types.TestSession
	session.ID = hub.NewSessionID(1)
	k.SetSession(ctx, session)
	k.SetSessionIDBySubscriptionID(ctx, types.TestSubscription.ID, 0, session.ID)
	k.SetSessionsCountOfSubscription(ctx, types.TestSubscription.ID, 2)

	res, _err = querySessionsOfSubscription(ctx, req, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.Len(t, sessions, 2)
}

func Test_queryAllSessions(t *testing.T) {
	ctx, k, _, _ := keeper.CreateTestInput(t, false)
	cdc := keeper.MakeTestCodec()
	var err error
	var sessions []types.Session

	res, _err := queryAllSessions(ctx, k)
	require.Nil(t, _err)
	require.Equal(t, []byte("null"), res)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.NotEqual(t, []types.Session{types.TestSession}, sessions)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.NotEqual(t, []types.Session{types.TestSession}, sessions)

	k.SetSession(ctx, types.TestSession)
	require.Nil(t, err)

	res, _err = queryAllSessions(ctx, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.Equal(t, []types.Session{types.TestSession}, sessions)

	session := types.TestSession
	session.ID = hub.NewSessionID(1)
	k.SetSession(ctx, session)
	require.Nil(t, err)

	res, _err = queryAllSessions(ctx, k)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.Equal(t, append([]types.Session{types.TestSession}, session), sessions)
}
