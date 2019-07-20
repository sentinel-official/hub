package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	Codespace = sdk.CodespaceType("vpn")

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

func ErrorMarshal() sdk.Error {
	return sdk.NewError(Codespace, hub.ErrCodeMarshal, hub.ErrMsgMarshal)
}

func ErrorUnmarshal() sdk.Error {
	return sdk.NewError(Codespace, hub.ErrCodeUnmarshal, hub.ErrMsgUnmarshal)
}

func ErrorUnknownMsgType(msgType string) sdk.Error {
	return sdk.NewError(Codespace, errCodeUnknownMsgType, errMsgUnknownMsgType+msgType)
}

func ErrorInvalidQueryType(queryType string) sdk.Error {
	return sdk.NewError(Codespace, errCodeUnknownQueryType, errMsgUnknownQueryType+queryType)
}

func ErrorInvalidField(field string) sdk.Error {
	return sdk.NewError(Codespace, errCodeInvalidField, errMsgInvalidField+field)
}

func ErrorUnauthorized() sdk.Error {
	return sdk.NewError(Codespace, errCodeUnauthorized, errMsgUnauthorized)
}

func ErrorNodeDoesNotExist() sdk.Error {
	return sdk.NewError(Codespace, errCodeNodeDoesNotExist, errMsgNodeDoesNotExist)
}

func ErrorInvalidNodeStatus() sdk.Error {
	return sdk.NewError(Codespace, errCodeInvalidNodeStatus, errMsgInvalidNodeStatus)
}

func ErrorInvalidDeposit() sdk.Error {
	return sdk.NewError(Codespace, errCodeInvalidDeposit, errMsgInvalidDeposit)
}

func ErrorSubscriptionDoesNotExist() sdk.Error {
	return sdk.NewError(Codespace, errCodeSubscriptionDoesNotExist, errMsgSubscriptionDoesNotExist)
}

func ErrorSubscriptionAlreadyExists() sdk.Error {
	return sdk.NewError(Codespace, errCodeSubscriptionAlreadyExists, errMsgSubscriptionAlreadyExists)
}

func ErrorInvalidSubscriptionStatus() sdk.Error {
	return sdk.NewError(Codespace, errCodeInvalidSubscriptionStatus, errMsgInvalidSubscriptionStatus)
}

func ErrorInvalidBandwidth() sdk.Error {
	return sdk.NewError(Codespace, errCodeInvalidBandwidth, errMsgInvalidBandwidth)
}

func ErrorInvalidBandwidthSignature() sdk.Error {
	return sdk.NewError(Codespace, errCodeInvalidBandwidthSignature, errMsgInvalidBandwidthSignature)
}

func ErrorSessionAlreadyExists() sdk.Error {
	return sdk.NewError(Codespace, errCodeSessionAlreadyExists, errMsgSessionAlreadyExists)
}

func ErrorInvalidSessionStatus() sdk.Error {
	return sdk.NewError(Codespace, errCodeInvalidSessionStatus, errMsgInvalidSessionStatus)
}
