package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	QuerySession                 = "Session"
	QuerySessions                = "Sessions"
	QuerySessionsForSubscription = "SessionsForSubscription"
	QuerySessionsForNode         = "SessionsForNode"
	QuerySessionsForAddress      = "SessionsForAddress"

	QueryActiveSession = "ActiveSession"
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
	Status  hub.Status     `json:"status"`
	Skip    int            `json:"skip"`
	Limit   int            `json:"limit"`
}

func NewQuerySessionsForAddressParams(address sdk.AccAddress, status hub.Status, skip, limit int) QuerySessionsForAddressParams {
	return QuerySessionsForAddressParams{
		Address: address,
		Status:  status,
		Skip:    skip,
		Limit:   limit,
	}
}

type QueryActiveSessionParams struct {
	Address      sdk.AccAddress  `json:"address"`
	Subscription uint64          `json:"subscription"`
	Node         hub.NodeAddress `json:"node"`
}

func NewQueryActiveSessionParams(address sdk.AccAddress, subscription uint64, node hub.NodeAddress) QueryActiveSessionParams {
	return QueryActiveSessionParams{
		Address:      address,
		Subscription: subscription,
		Node:         node,
	}
}
