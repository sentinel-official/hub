package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

const (
	Codespace = csdkTypes.CodespaceType("vpn")

	errCodeUnknownMsgType            = 101
	errCodeUnknownQueryType          = 102
	errCodeInvalidField              = 103
	errCodeUnauthorized              = 104
	errCodeNodeDoesNotExist          = 105
	errCodeInvalidNodeStatus         = 106
	errCodeInvalidDeposit            = 107
	errCodeSubscriptionDoesNotExist  = 108
	errCodeSubscriptionAlreadyExists = 109
	errCodeInvalidSubscriptionStatus = 110
	errCodeInvalidBandwidth          = 111
	errCodeInvalidBandwidthSign      = 112
	errCodeSessionAlreadyExists      = 113
	errCodeInvalidSessionStatus      = 114

	errMsgUnknownMsgType            = "Unknown message type: "
	errMsgUnknownQueryType          = "Invalid query type: "
	errMsgInvalidField              = "Invalid field: "
	errMsgUnauthorized              = "Unauthorized"
	errMsgNodeDoesNotExist          = "Node does not exist"
	errMsgInvalidNodeStatus         = "Invalid node status"
	errMsgInvalidDeposit            = "Invalid deposit"
	errMsgSubscriptionDoesNotExist  = "Subscription does not exist"
	errMsgSubscriptionAlreadyExists = "Subscription already exists"
	errMsgInvalidSubscriptionStatus = "Invalid subscription status"
	errMsgInvalidBandwidth          = "Invalid bandwidth"
	errMsgInvalidBandwidthSign      = "Invalid bandwidth sign"
	errMsgSessionAlreadyExists      = "Session is active"
	errMsgInvalidSessionStatus      = "Invalid session status"
)

func ErrorMarshal() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, sdkTypes.ErrCodeMarshal, sdkTypes.ErrMsgMarshal)
}

func ErrorUnmarshal() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, sdkTypes.ErrCodeUnmarshal, sdkTypes.ErrMsgUnmarshal)
}

func ErrorUnknownMsgType(msgType string) csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeUnknownMsgType, errMsgUnknownMsgType+msgType)
}

func ErrorInvalidQueryType(queryType string) csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeUnknownQueryType, errMsgUnknownQueryType+queryType)
}

func ErrorInvalidField(field string) csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeInvalidField, errMsgInvalidField+field)
}

func ErrorUnauthorized() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeUnauthorized, errMsgUnauthorized)
}

func ErrorNodeDoesNotExist() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeNodeDoesNotExist, errMsgNodeDoesNotExist)
}

func ErrorInvalidNodeStatus() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeInvalidNodeStatus, errMsgInvalidNodeStatus)
}

func ErrorInvalidDeposit() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeInvalidDeposit, errMsgInvalidDeposit)
}

func ErrorSubscriptionDoesNotExist() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeSubscriptionDoesNotExist, errMsgSubscriptionDoesNotExist)
}

func ErrorSubscriptionAlreadyExists() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeSubscriptionAlreadyExists, errMsgSubscriptionAlreadyExists)
}

func ErrorInvalidSubscriptionStatus() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeInvalidSubscriptionStatus, errMsgInvalidSubscriptionStatus)
}

func ErrorInvalidBandwidth() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeInvalidBandwidth, errMsgInvalidBandwidth)
}

func ErrorInvalidBandwidthSign() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeInvalidBandwidthSign, errMsgInvalidBandwidthSign)
}

func ErrorSessionAlreadyExists() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeSessionAlreadyExists, errMsgSessionAlreadyExists)
}

func ErrorInvalidSessionStatus() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeInvalidSessionStatus, errMsgInvalidSessionStatus)
}
