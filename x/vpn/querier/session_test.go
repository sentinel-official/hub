package querier

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"
	sdk "github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func TestNewQuerySessionParams(t *testing.T) {
	params := NewQuerySessionParams(sdk.TestIDZero)
	require.Equal(t, TestSessionParamsZero, params)

	params = NewQuerySessionParams(sdk.TestIDPos)
	require.Equal(t, TestSessionParamsPos, params)
}

func Test_querySession(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var session sdk.Session

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", sdk.QuerierRoute, QuerySession),
		Data: []byte{},
	}

	res, _err := querySession(ctx, cdc, req, vpnKeeper)
	require.NotNil(t, _err)
	require.Equal(t,[]byte(nil),res)
	require.Len(t, res, 0)

	vpnKeeper.SetSession(ctx, sdk.TestSessionValid)
	req.Data, err = cdc.MarshalJSON(NewQuerySessionParams(sdk.TestIDZero))
	require.Nil(t, err)

	res, _err = querySession(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t,[]byte(nil),res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &session)
	require.Nil(t, err)
	require.Equal(t, sdk.TestSessionValid, session)

	req.Data, err = cdc.MarshalJSON(NewQuerySessionParams(sdk.TestIDPos))
	require.Nil(t, err)

	res, _err = querySession(ctx, cdc, req, vpnKeeper)
	require.Nil(t, res)
	require.Equal(t,[]byte(nil),res)
	require.Len(t, res, 0)
}

func TestNewQuerySessionOfSubscriptionPrams(t *testing.T) {
	params := NewQuerySessionOfSubscriptionPrams(sdk.TestIDZero, 0)
	require.Equal(t, TestSessionOfSubscriptionPramsZero, params)

	params = NewQuerySessionOfSubscriptionPrams(sdk.TestIDPos, 0)
	require.Equal(t, TestSessionOfSubscriptionPramsPos, params)
}

func Test_querySessionOfSubscription(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var session sdk.Session

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", sdk.QuerierRoute, QuerySessionOfSubscription),
		Data: []byte{},
	}

	res, _err := querySessionOfSubscription(ctx, cdc, req, vpnKeeper)
	require.NotNil(t, _err)
	require.Equal(t,[]byte(nil),res)
	require.Len(t, res, 0)

	vpnKeeper.SetSubscription(ctx, sdk.TestSubscriptionValid)
	vpnKeeper.SetSessionIDBySubscriptionID(ctx, sdk.TestSessionValid.ID, 0, sdk.TestSubscriptionValid.ID)
	vpnKeeper.SetSessionsCountOfSubscription(ctx, sdk.TestSubscriptionValid.ID, 1)

	req.Data, err = cdc.MarshalJSON(NewQuerySessionsOfSubscriptionPrams(sdk.TestIDZero))
	require.Nil(t, err)

	res, _err = querySessionOfSubscription(ctx, cdc, req, vpnKeeper)
	require.Nil(t, err)
	require.Equal(t,[]byte(nil),res)
	require.Nil(t, res)

	err = cdc.UnmarshalJSON(res, &session)
	require.NotEqual(t, sdk.TestSessionValid, session)

	vpnKeeper.SetSession(ctx, sdk.TestSessionValid)
	req.Data, err = cdc.MarshalJSON(NewQuerySessionsOfSubscriptionPrams(sdk.TestIDZero))
	require.Nil(t, err)

	res, _err = querySessionOfSubscription(ctx, cdc, req, vpnKeeper)
	require.Nil(t, err)
	require.NotEqual(t,[]byte(nil),res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &session)
	require.Nil(t, err)
	require.Equal(t, sdk.TestSessionValid, session)

	req.Data, err = cdc.MarshalJSON(NewQuerySessionsOfSubscriptionPrams(sdk.TestIDPos))
	require.Nil(t, err)

	res, _err = querySessionOfSubscription(ctx, cdc, req, vpnKeeper)
	require.Nil(t, err)
	require.Equal(t,[]byte(nil),res)
	require.Nil(t, res)

	err = cdc.UnmarshalJSON(res, &session)
	require.NotNil(t, err)
}

func TestNewQuerySessionsOfSubscriptionPrams(t *testing.T) {
	params := NewQuerySessionsOfSubscriptionPrams(sdk.TestIDZero)
	require.Equal(t, TestSessionsOfSubscriptionPramsZero, params)

	params = NewQuerySessionsOfSubscriptionPrams(sdk.TestIDPos)
	require.Equal(t, TestSessionsOfSubscriptionPramsPos, params)
}

func Test_querySessionsOfSubscription(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var sessions []sdk.Session

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", sdk.QuerierRoute, QuerySessionsOfSubscription),
		Data: []byte{},
	}

	res, _err := querySessionsOfSubscription(ctx, cdc, req, vpnKeeper)
	require.NotNil(t, _err)
	require.Equal(t,[]byte(nil),res)
	require.Len(t, res, 0)

	vpnKeeper.SetSubscription(ctx, sdk.TestSubscriptionValid)
	vpnKeeper.SetSession(ctx, sdk.TestSessionValid)
	vpnKeeper.SetSessionIDBySubscriptionID(ctx, sdk.TestSubscriptionValid.ID, 0, sdk.TestSessionValid.ID)
	vpnKeeper.SetSessionsCountOfSubscription(ctx, sdk.TestSubscriptionValid.ID, 1)

	req.Data, err = cdc.MarshalJSON(NewQuerySessionsOfSubscriptionPrams(sdk.TestIDZero))
	require.Nil(t, err)

	res, _err = querySessionsOfSubscription(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t,[]byte(nil),res)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.Equal(t, sdk.TestSessionsValid, sessions)

	session := sdk.TestSessionValid
	session.ID = sdk.TestIDPos
	vpnKeeper.SetSession(ctx, session)
	vpnKeeper.SetSessionIDBySubscriptionID(ctx, sdk.TestSubscriptionValid.ID, 0, session.ID)
	vpnKeeper.SetSessionsCountOfSubscription(ctx, sdk.TestSubscriptionValid.ID, 2)

	res, _err = querySessionsOfSubscription(ctx, cdc, req, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t,[]byte(nil),res)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.Len(t, sessions, 2)
}

func Test_queryAllSessions(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var sessions []sdk.Session

	res, _err := queryAllSessions(ctx, cdc, vpnKeeper)
	require.Nil(t, _err)
	require.Equal(t,[]byte("null"),res)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.NotEqual(t, sdk.TestSessionsValid, sessions)

	vpnKeeper.SetSession(ctx, sdk.TestSessionEmpty)
	res, _err = queryAllSessions(ctx, cdc, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t,[]byte(nil),res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.NotEqual(t, sdk.TestSessionsValid, sessions)

	vpnKeeper.SetSession(ctx, sdk.TestSessionValid)
	require.Nil(t, err)

	res, _err = queryAllSessions(ctx, cdc, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t,[]byte(nil),res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.Equal(t, sdk.TestSessionsValid, sessions)

	session := sdk.TestSessionValid
	session.ID = sdk.TestIDPos
	vpnKeeper.SetSession(ctx, session)
	require.Nil(t, err)

	res, _err = queryAllSessions(ctx, cdc, vpnKeeper)
	require.Nil(t, _err)
	require.NotEqual(t,[]byte(nil),res)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.Equal(t, append(sdk.TestSessionsValid, session), sessions)
}
