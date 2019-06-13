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
	ctx, _, vpnkeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var session sdk.Session

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", sdk.QuerierRoute, QuerySession),
		Data: []byte{},
	}

	res, _err := querySession(ctx, cdc, req, vpnkeeper)
	require.NotNil(t, _err)
	require.Len(t, res, 0)

	vpnkeeper.SetSession(ctx, sdk.TestSessionValid)
	req.Data, err = cdc.MarshalJSON(NewQuerySessionParams(sdk.TestIDZero))
	require.Nil(t, err)

	res, _err = querySession(ctx, cdc, req, vpnkeeper)
	require.Nil(t, _err)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &session)
	require.Nil(t, err)
	require.Equal(t, sdk.TestSessionValid, session)

	req.Data, err = cdc.MarshalJSON(NewQuerySessionParams(sdk.TestIDPos))
	require.Nil(t, err)

	res, _err = querySession(ctx, cdc, req, vpnkeeper)
	require.Nil(t, res)
	require.Len(t, res, 0)
}

func TestNewQuerySessionOfSubscriptionPrams(t *testing.T) {
	params := NewQuerySessionOfSubscriptionPrams(sdk.TestIDZero, 0)
	require.Equal(t, TestSessionOfSubscriptionPramsZero, params)

	params = NewQuerySessionOfSubscriptionPrams(sdk.TestIDPos, 0)
	require.Equal(t, TestSessionOfSubscriptionPramsPos, params)
}

func Test_querySessionOfSubscription(t *testing.T) {
	ctx, _, vpnkeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var session sdk.Session

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", sdk.QuerierRoute, QuerySessionOfSubscription),
		Data: []byte{},
	}

	res, _err := querySessionOfSubscription(ctx, cdc, req, vpnkeeper)
	require.NotNil(t, _err)
	require.Len(t, res, 0)

	vpnkeeper.SetSubscription(ctx, sdk.TestSubscriptionValid)
	vpnkeeper.SetSessionIDBySubscriptionID(ctx, sdk.TestSessionValid.ID, 0, sdk.TestSubscriptionValid.ID)
	vpnkeeper.SetSessionsCountOfSubscription(ctx, sdk.TestSubscriptionValid.ID, 1)

	req.Data, err = cdc.MarshalJSON(NewQuerySessionsOfSubscriptionPrams(sdk.TestIDZero))
	require.Nil(t, err)

	res, _err = querySessionOfSubscription(ctx, cdc, req, vpnkeeper)
	require.Nil(t, err)
	require.Nil(t, res)

	err = cdc.UnmarshalJSON(res, &session)
	require.NotEqual(t, sdk.TestSessionValid, session)

	vpnkeeper.SetSession(ctx, sdk.TestSessionValid)
	req.Data, err = cdc.MarshalJSON(NewQuerySessionsOfSubscriptionPrams(sdk.TestIDZero))
	require.Nil(t, err)

	res, _err = querySessionOfSubscription(ctx, cdc, req, vpnkeeper)
	require.Nil(t, err)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &session)
	require.Nil(t, err)
	require.Equal(t, sdk.TestSessionValid, session)

	req.Data, err = cdc.MarshalJSON(NewQuerySessionsOfSubscriptionPrams(sdk.TestIDPos))
	require.Nil(t, err)

	res, _err = querySessionOfSubscription(ctx, cdc, req, vpnkeeper)
	require.Nil(t, err)
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
	ctx, _, vpnkeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var sessions []sdk.Session

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", sdk.QuerierRoute, QuerySessionsOfSubscription),
		Data: []byte{},
	}

	res, _err := querySessionsOfSubscription(ctx, cdc, req, vpnkeeper)
	require.NotNil(t, _err)
	require.Len(t, res, 0)

	vpnkeeper.SetSubscription(ctx, sdk.TestSubscriptionValid)
	vpnkeeper.SetSession(ctx, sdk.TestSessionValid)
	vpnkeeper.SetSessionIDBySubscriptionID(ctx, sdk.TestSubscriptionValid.ID, 0, sdk.TestSessionValid.ID)
	vpnkeeper.SetSessionsCountOfSubscription(ctx, sdk.TestSubscriptionValid.ID, 1)

	req.Data, err = cdc.MarshalJSON(NewQuerySessionsOfSubscriptionPrams(sdk.TestIDZero))
	require.Nil(t, err)

	res, _err = querySessionsOfSubscription(ctx, cdc, req, vpnkeeper)
	require.Nil(t, _err)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.Equal(t, sdk.TestSessionsValid, sessions)

	session := sdk.TestSessionValid
	session.ID = sdk.TestIDPos
	vpnkeeper.SetSession(ctx, session)
	vpnkeeper.SetSessionIDBySubscriptionID(ctx, sdk.TestSubscriptionValid.ID, 0, session.ID)
	vpnkeeper.SetSessionsCountOfSubscription(ctx, sdk.TestSubscriptionValid.ID, 2)

	res, _err = querySessionsOfSubscription(ctx, cdc, req, vpnkeeper)
	require.Nil(t, _err)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.Len(t, sessions, 2)
}

func Test_queryAllSessions(t *testing.T) {
	ctx, _, vpnkeeper, _ := keeper.TestCreateInput()
	cdc := keeper.TestMakeCodec()
	var err error
	var sessions []sdk.Session

	res, _err := queryAllSessions(ctx, cdc, vpnkeeper)
	require.Nil(t, _err)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.NotEqual(t, sdk.TestSessionsValid, sessions)

	vpnkeeper.SetSession(ctx, sdk.TestSessionEmpty)
	res, _err = queryAllSessions(ctx, cdc, vpnkeeper)
	require.Nil(t, _err)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.NotEqual(t, sdk.TestSessionsValid, sessions)

	vpnkeeper.SetSession(ctx, sdk.TestSessionValid)
	require.Nil(t, err)

	res, _err = queryAllSessions(ctx, cdc, vpnkeeper)
	require.Nil(t, _err)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.Equal(t, sdk.TestSessionsValid, sessions)

	session := sdk.TestSessionValid
	session.ID = sdk.TestIDPos
	vpnkeeper.SetSession(ctx, session)
	require.Nil(t, err)

	res, _err = queryAllSessions(ctx, cdc, vpnkeeper)
	require.Nil(t, _err)
	require.NotNil(t, res)

	err = cdc.UnmarshalJSON(res, &sessions)
	require.Nil(t, err)
	require.Equal(t, append(sdk.TestSessionsValid, session), sessions)
}
