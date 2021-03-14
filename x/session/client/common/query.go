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

func QuerySessions(ctx context.CLIContext, skip, limit int) (types.Sessions, error) {
	params := types.NewQuerySessionsParams(skip, limit)
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

func QuerySessionsForSubscription(ctx context.CLIContext, id uint64, skip, limit int) (types.Sessions, error) {
	params := types.NewQuerySessionsForSubscriptionParams(id, skip, limit)
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

func QuerySessionsForNode(ctx context.CLIContext, address hub.NodeAddress, skip, limit int) (types.Sessions, error) {
	params := types.NewQuerySessionsForNodeParams(address, skip, limit)
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

func QuerySessionsForAddress(ctx context.CLIContext, address sdk.AccAddress, status hub.Status, skip, limit int) (types.Sessions, error) {
	params := types.NewQuerySessionsForAddressParams(address, status, skip, limit)
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

func QueryActiveSession(ctx context.CLIContext, address sdk.AccAddress, id uint64, node hub.NodeAddress) (*types.Session, error) {
	params := types.NewQueryActiveSessionParams(address, id, node)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QueryActiveSession)
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
