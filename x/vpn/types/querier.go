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
	ID hub.NodeID
}

func NewQueryNodeParams(id hub.NodeID) QueryNodeParams {
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
	ID hub.SubscriptionID
}

func NewQuerySubscriptionParams(id hub.SubscriptionID) QuerySubscriptionParams {
	return QuerySubscriptionParams{
		ID: id,
	}
}

type QuerySubscriptionsOfNodePrams struct {
	ID hub.NodeID
}

func NewQuerySubscriptionsOfNodePrams(id hub.NodeID) QuerySubscriptionsOfNodePrams {
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
	ID hub.SubscriptionID
}

func NewQuerySessionsCountOfSubscriptionParams(id hub.SubscriptionID) QuerySessionsCountOfSubscriptionParams {
	return QuerySessionsCountOfSubscriptionParams{
		ID: id,
	}
}

type QuerySessionParams struct {
	ID hub.SessionID
}

func NewQuerySessionParams(id hub.SessionID) QuerySessionParams {
	return QuerySessionParams{
		ID: id,
	}
}

type QuerySessionOfSubscriptionPrams struct {
	ID    hub.SubscriptionID
	Index uint64
}

func NewQuerySessionOfSubscriptionPrams(id hub.SubscriptionID, index uint64) QuerySessionOfSubscriptionPrams {
	return QuerySessionOfSubscriptionPrams{
		ID:    id,
		Index: index,
	}
}

type QuerySessionsOfSubscriptionPrams struct {
	ID hub.SubscriptionID
}

func NewQuerySessionsOfSubscriptionPrams(id hub.SubscriptionID) QuerySessionsOfSubscriptionPrams {
	return QuerySessionsOfSubscriptionPrams{
		ID: id,
	}
}
