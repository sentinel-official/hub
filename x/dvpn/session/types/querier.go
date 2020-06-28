package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	QuerySession                 = "query_session"
	QuerySessions                = "query_sessions"
	QuerySessionsForSubscription = "query_sessions_for_subscription"
	QuerySessionsForNode         = "query_sessions_for_node"
	QuerySessionsForAddress      = "query_sessions_for_address"
)

type QuerySessionParams struct {
	ID uint64 `json:"id"`
}

func NewQuerySessionParams(id uint64) QuerySessionParams {
	return QuerySessionParams{
		ID: id,
	}
}

type QuerySessionsForSubscriptionParams struct {
	ID uint64 `json:"id"`
}

func NewQuerySessionsForSubscriptionParams(id uint64) QuerySessionsForSubscriptionParams {
	return QuerySessionsForSubscriptionParams{
		ID: id,
	}
}

type QuerySessionsForNodeParams struct {
	Address hub.NodeAddress `json:"address"`
}

func NewQuerySessionsForNodeParams(address hub.NodeAddress) QuerySessionsForNodeParams {
	return QuerySessionsForNodeParams{
		Address: address,
	}
}

type QuerySessionsForAddressParams struct {
	Address sdk.AccAddress `json:"address"`
}

func NewQuerySessionsForAddressParams(address sdk.AccAddress) QuerySessionsForAddressParams {
	return QuerySessionsForAddressParams{
		Address: address,
	}
}
