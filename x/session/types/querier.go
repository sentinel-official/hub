package types

import (
	"github.com/cosmos/cosmos-sdk/types/query"
)

const (
	QuerySession                 = "Session"
	QuerySessions                = "Sessions"
	QuerySessionsForSubscription = "SessionsForSubscription"
	QuerySessionsForNode         = "SessionsForNode"
	QuerySessionsForAddress      = "SessionsForAddress"

	QueryOngoingSession = "OngoingSession"
)

func NewQuerySessionRequest(id uint64) QuerySessionRequest {
	return QuerySessionRequest{
		Id: id,
	}
}

func NewQuerySessionsRequest(pagination *query.PageRequest) QuerySessionsRequest {
	return QuerySessionsRequest{
		Pagination: pagination,
	}
}

func NewQuerySessionsForSubscriptionRequest(id uint64, pagination *query.PageRequest) QuerySessionsForSubscriptionRequest {
	return QuerySessionsForSubscriptionRequest{
		Id:         id,
		Pagination: pagination,
	}
}

func NewQuerySessionsForNodeRequest(address string, pagination *query.PageRequest) QuerySessionsForNodeRequest {
	return QuerySessionsForNodeRequest{
		Address:    address,
		Pagination: pagination,
	}
}

func NewQuerySessionsForAddressRequest(address string, pagination *query.PageRequest) QuerySessionsForAddressRequest {
	return QuerySessionsForAddressRequest{
		Address:    address,
		Pagination: pagination,
	}
}

func NewQueryOngoingSessionRequest(id uint64, address string) QueryOngoingSessionRequest {
	return QueryOngoingSessionRequest{
		Id:      id,
		Address: address,
	}
}
