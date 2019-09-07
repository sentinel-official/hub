package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	QueryNode           = "node"
	QueryNodesOfAddress = "nodes_of_address"
	QueryAllNodes       = "all_nodes"

	QuerySubscription                = "subscription"
	QuerySubscriptionsOfNode         = "subscriptions_of_node"
	QuerySubscriptionsOfAddress      = "subscriptions_of_address"
	QueryAllSubscriptions            = "all_subscriptions"
	QuerySessionsCountOfSubscription = "sessions_count_of_subscription"

	QuerySession                = "session"
	QuerySessionOfSubscription  = "session_of_subscription"
	QuerySessionsOfSubscription = "sessions_of_subscription"
	QueryAllSessions            = "all_sessions"
)

type QueryNodeParams struct {
	ID hub.ID
}

func NewQueryNodeParams(id hub.ID) QueryNodeParams {
	return QueryNodeParams{
		ID: id,
	}
}

type QueryNodesOfAddressPrams struct {
	Address sdk.AccAddress
}

func NewQueryNodesOfAddressParams(address sdk.AccAddress) QueryNodesOfAddressPrams {
	return QueryNodesOfAddressPrams{
		Address: address,
	}
}

type QuerySubscriptionParams struct {
	ID hub.ID
}

func NewQuerySubscriptionParams(id hub.ID) QuerySubscriptionParams {
	return QuerySubscriptionParams{
		ID: id,
	}
}

type QuerySubscriptionsOfNodePrams struct {
	ID hub.ID
}

func NewQuerySubscriptionsOfNodePrams(id hub.ID) QuerySubscriptionsOfNodePrams {
	return QuerySubscriptionsOfNodePrams{
		ID: id,
	}
}

type QuerySubscriptionsOfAddressParams struct {
	Address sdk.AccAddress
}

func NewQuerySubscriptionsOfAddressParams(address sdk.AccAddress) QuerySubscriptionsOfAddressParams {
	return QuerySubscriptionsOfAddressParams{
		Address: address,
	}
}

type QuerySessionsCountOfSubscriptionParams struct {
	ID hub.ID
}

func NewQuerySessionsCountOfSubscriptionParams(id hub.ID) QuerySessionsCountOfSubscriptionParams {
	return QuerySessionsCountOfSubscriptionParams{
		ID: id,
	}
}

type QuerySessionParams struct {
	ID hub.ID
}

func NewQuerySessionParams(id hub.ID) QuerySessionParams {
	return QuerySessionParams{
		ID: id,
	}
}

type QuerySessionOfSubscriptionPrams struct {
	ID    hub.ID
	Index uint64
}

func NewQuerySessionOfSubscriptionPrams(id hub.ID, index uint64) QuerySessionOfSubscriptionPrams {
	return QuerySessionOfSubscriptionPrams{
		ID:    id,
		Index: index,
	}
}

type QuerySessionsOfSubscriptionPrams struct {
	ID hub.ID
}

func NewQuerySessionsOfSubscriptionPrams(id hub.ID) QuerySessionsOfSubscriptionPrams {
	return QuerySessionsOfSubscriptionPrams{
		ID: id,
	}
}
