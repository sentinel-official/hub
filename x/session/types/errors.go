package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	Codespace = sdk.CodespaceType("session")
)

const (
	errorCodeUnknownMsgType sdk.CodeType = iota + 101
	errorCodeUnknownQueryType
	errorCodeInvalidField
	errorCodeSubscriptionDoesNotExit
	errorCodeInvalidSubscriptionStatus
	errorCodeUnauthorized
	errorCodeQuotaDoesNotExist
	errorCodeInvalidBandwidth
)

const (
	errorMsgUnknownMsgType            = "unknown message type: %s"
	errorMsgUnknownQueryType          = "unknown query type: %s"
	errorMsgInvalidField              = "invalid field: %s"
	errorMsgSubscriptionDoesNotExit   = "subscription does not exist"
	errorMsgInvalidSubscriptionStatus = "invalid subscription status"
	errorMsgUnauthorized              = "unauthorized"
	errorMsgQuotaDoesNotExist         = "quota does not exist"
	errorMsgInvalidBandwidth          = "invalid bandwidth"
)

func ErrorMarshal() sdk.Error {
	return sdk.NewError(Codespace, hub.ErrorCodeMarshal, hub.ErrorMsgMarshal)
}

func ErrorUnmarshal() sdk.Error {
	return sdk.NewError(Codespace, hub.ErrorCodeUnmarshal, hub.ErrorMsgUnmarshal)
}

func ErrorUnknownMsgType(v string) sdk.Error {
	return sdk.NewError(Codespace, errorCodeUnknownMsgType, fmt.Sprintf(errorMsgUnknownMsgType, v))
}

func ErrorUnknownQueryType(v string) sdk.Error {
	return sdk.NewError(Codespace, errorCodeUnknownQueryType, fmt.Sprintf(errorMsgUnknownQueryType, v))
}

func ErrorInvalidField(v string) sdk.Error {
	return sdk.NewError(Codespace, errorCodeInvalidField, fmt.Sprintf(errorMsgInvalidField, v))
}

func ErrorSubscriptionDoesNotExit() sdk.Error {
	return sdk.NewError(Codespace, errorCodeSubscriptionDoesNotExit, errorMsgSubscriptionDoesNotExit)
}

func ErrorInvalidSubscriptionStatus() sdk.Error {
	return sdk.NewError(Codespace, errorCodeInvalidSubscriptionStatus, errorMsgInvalidSubscriptionStatus)
}

func ErrorUnauthorized() sdk.Error {
	return sdk.NewError(Codespace, errorCodeUnauthorized, errorMsgUnauthorized)
}

func ErrorQuotaDoesNotExist() sdk.Error {
	return sdk.NewError(Codespace, errorCodeQuotaDoesNotExist, errorMsgQuotaDoesNotExist)
}

func ErrorInvalidBandwidth() sdk.Error {
	return sdk.NewError(Codespace, errorCodeInvalidBandwidth, errorMsgInvalidBandwidth)
}
