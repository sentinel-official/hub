package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	QuerySession                 = "session"
	QuerySessions                = "sessions"
	QuerySessionsForSubscription = "sessions_for_subscription"
	QuerySessionsForNode         = "sessions_for_node"
	QuerySessionsForAddress      = "sessions_for_address"

	QueryOngoingSession = "ongoing_session"
)

type QuerySessionParams struct {
	ID uint64 `json:"id"`
}

func NewQuerySessionParams(id uint64) QuerySessionParams {
	return QuerySessionParams{
		ID: id,
	}
}

type QuerySessionsParams struct {
	Skip  int `json:"skip"`
	Limit int `json:"limit"`
}

func NewQuerySessionsParams(skip, limit int) QuerySessionsParams {
	return QuerySessionsParams{
		Skip:  skip,
		Limit: limit,
	}
}

type QuerySessionsForSubscriptionParams struct {
	ID    uint64 `json:"id"`
	Skip  int    `json:"skip"`
	Limit int    `json:"limit"`
}

func NewQuerySessionsForSubscriptionParams(id uint64, skip, limit int) QuerySessionsForSubscriptionParams {
	return QuerySessionsForSubscriptionParams{
		ID:    id,
		Skip:  skip,
		Limit: limit,
	}
}

type QuerySessionsForNodeParams struct {
	Address hub.NodeAddress `json:"address"`
	Skip    int             `json:"skip"`
	Limit   int             `json:"limit"`
}

func NewQuerySessionsForNodeParams(address hub.NodeAddress, skip, limit int) QuerySessionsForNodeParams {
	return QuerySessionsForNodeParams{
		Address: address,
		Skip:    skip,
		Limit:   limit,
	}
}

type QuerySessionsForAddressParams struct {
	Address sdk.AccAddress `json:"address"`
	Skip    int            `json:"skip"`
	Limit   int            `json:"limit"`
}

func NewQuerySessionsForAddressParams(address sdk.AccAddress, skip, limit int) QuerySessionsForAddressParams {
	return QuerySessionsForAddressParams{
		Address: address,
		Skip:    skip,
		Limit:   limit,
	}
}

type QueryOngoingSessionParams struct {
	ID      uint64         `json:"id"`
	Address sdk.AccAddress `json:"address"`
}

func NewQueryOngoingSessionParams(id uint64, address sdk.AccAddress) QueryOngoingSessionParams {
	return QueryOngoingSessionParams{
		ID:      id,
		Address: address,
	}
}
