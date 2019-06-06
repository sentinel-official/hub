package types

import (
	csdk "github.com/cosmos/cosmos-sdk/types"

	sdk "github.com/ironman0x7b2/sentinel-sdk/types"
)

const (
	Codespace = csdk.CodespaceType("vpn")

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
	errCodeInvalidBandwidthSignature = 112
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
	errMsgInvalidBandwidthSignature = "Invalid bandwidth signature"
	errMsgSessionAlreadyExists      = "Session is active"
	errMsgInvalidSessionStatus      = "Invalid session status"
)

func ErrorMarshal() csdk.Error {
	return csdk.NewError(Codespace, sdk.ErrCodeMarshal, sdk.ErrMsgMarshal)
}

func ErrorUnmarshal() csdk.Error {
	return csdk.NewError(Codespace, sdk.ErrCodeUnmarshal, sdk.ErrMsgUnmarshal)
}

func ErrorUnknownMsgType(msgType string) csdk.Error {
	return csdk.NewError(Codespace, errCodeUnknownMsgType, errMsgUnknownMsgType+msgType)
}

func ErrorInvalidQueryType(queryType string) csdk.Error {
	return csdk.NewError(Codespace, errCodeUnknownQueryType, errMsgUnknownQueryType+queryType)
}

func ErrorInvalidField(field string) csdk.Error {
	return csdk.NewError(Codespace, errCodeInvalidField, errMsgInvalidField+field)
}

func ErrorUnauthorized() csdk.Error {
	return csdk.NewError(Codespace, errCodeUnauthorized, errMsgUnauthorized)
}

func ErrorNodeDoesNotExist() csdk.Error {
	return csdk.NewError(Codespace, errCodeNodeDoesNotExist, errMsgNodeDoesNotExist)
}

func ErrorInvalidNodeStatus() csdk.Error {
	return csdk.NewError(Codespace, errCodeInvalidNodeStatus, errMsgInvalidNodeStatus)
}

func ErrorInvalidDeposit() csdk.Error {
	return csdk.NewError(Codespace, errCodeInvalidDeposit, errMsgInvalidDeposit)
}

func ErrorSubscriptionDoesNotExist() csdk.Error {
	return csdk.NewError(Codespace, errCodeSubscriptionDoesNotExist, errMsgSubscriptionDoesNotExist)
}

func ErrorSubscriptionAlreadyExists() csdk.Error {
	return csdk.NewError(Codespace, errCodeSubscriptionAlreadyExists, errMsgSubscriptionAlreadyExists)
}

func ErrorInvalidSubscriptionStatus() csdk.Error {
	return csdk.NewError(Codespace, errCodeInvalidSubscriptionStatus, errMsgInvalidSubscriptionStatus)
}

func ErrorInvalidBandwidth() csdk.Error {
	return csdk.NewError(Codespace, errCodeInvalidBandwidth, errMsgInvalidBandwidth)
}

func ErrorInvalidBandwidthSignature() csdk.Error {
	return csdk.NewError(Codespace, errCodeInvalidBandwidthSignature, errMsgInvalidBandwidthSignature)
}

func ErrorSessionAlreadyExists() csdk.Error {
	return csdk.NewError(Codespace, errCodeSessionAlreadyExists, errMsgSessionAlreadyExists)
}

func ErrorInvalidSessionStatus() csdk.Error {
	return csdk.NewError(Codespace, errCodeInvalidSessionStatus, errMsgInvalidSessionStatus)
}
