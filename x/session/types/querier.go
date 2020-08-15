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
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func NewQuerySessionsParams(page, limit int) QuerySessionsParams {
	return QuerySessionsParams{
		Page:  page,
		Limit: limit,
	}
}

type QuerySessionsForSubscriptionParams struct {
	ID    uint64 `json:"id"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

func NewQuerySessionsForSubscriptionParams(id uint64, page, limit int) QuerySessionsForSubscriptionParams {
	return QuerySessionsForSubscriptionParams{
		ID:    id,
		Page:  page,
		Limit: limit,
	}
}

type QuerySessionsForNodeParams struct {
	Address hub.NodeAddress `json:"address"`
	Page    int             `json:"page"`
	Limit   int             `json:"limit"`
}

func NewQuerySessionsForNodeParams(address hub.NodeAddress, page, limit int) QuerySessionsForNodeParams {
	return QuerySessionsForNodeParams{
		Address: address,
		Page:    page,
		Limit:   limit,
	}
}

type QuerySessionsForAddressParams struct {
	Address sdk.AccAddress `json:"address"`
	Page    int            `json:"page"`
	Limit   int            `json:"limit"`
}

func NewQuerySessionsForAddressParams(address sdk.AccAddress, page, limit int) QuerySessionsForAddressParams {
	return QuerySessionsForAddressParams{
		Address: address,
		Page:    page,
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
