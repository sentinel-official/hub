package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/types"
)

func QuerySession(ctx context.CLIContext, id uint64) (*types.Session, error) {
	params := types.NewQuerySessionParams(id)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QuerySession)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no session found")
	}

	var session types.Session
	if err := ctx.Codec.UnmarshalJSON(res, &session); err != nil {
		return nil, err
	}

	return &session, nil
}

func QuerySessions(ctx context.CLIContext, page, limit int) (types.Sessions, error) {
	params := types.NewQuerySessionsParams(page, limit)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QuerySessions)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no sessions found")
	}

	var sessions types.Sessions
	if err := ctx.Codec.UnmarshalJSON(res, &sessions); err != nil {
		return nil, err
	}

	return sessions, nil
}

func QuerySessionsForSubscription(ctx context.CLIContext, id uint64, page, limit int) (types.Sessions, error) {
	params := types.NewQuerySessionsForSubscriptionParams(id, page, limit)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QuerySessionsForSubscription)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no sessions found")
	}

	var sessions types.Sessions
	if err := ctx.Codec.UnmarshalJSON(res, &sessions); err != nil {
		return nil, err
	}

	return sessions, nil
}

func QuerySessionsForNode(ctx context.CLIContext, address hub.NodeAddress, page, limit int) (types.Sessions, error) {
	params := types.NewQuerySessionsForNodeParams(address, page, limit)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QuerySessionsForNode)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no sessions found")
	}

	var sessions types.Sessions
	if err := ctx.Codec.UnmarshalJSON(res, &sessions); err != nil {
		return nil, err
	}

	return sessions, nil
}

func QuerySessionsForAddress(ctx context.CLIContext, address sdk.AccAddress, page, limit int) (types.Sessions, error) {
	params := types.NewQuerySessionsForAddressParams(address, page, limit)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QuerySessionsForAddress)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no sessions found")
	}

	var sessions types.Sessions
	if err := ctx.Codec.UnmarshalJSON(res, &sessions); err != nil {
		return nil, err
	}

	return sessions, nil
}

func QueryOngoingSession(ctx context.CLIContext, id uint64, address sdk.AccAddress) (*types.Session, error) {
	params := types.NewQueryOngoingSessionParams(id, address)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QueryOngoingSession)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no session found")
	}

	var session types.Session
	if err := ctx.Codec.UnmarshalJSON(res, &session); err != nil {
		return nil, err
	}

	return &session, nil
}
