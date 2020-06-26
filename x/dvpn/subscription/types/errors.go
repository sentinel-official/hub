package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	Codespace = sdk.CodespaceType("subscription")
)

const (
	errorCodeUnknownMsgType = iota + 101
	errorCodeUnknownQueryType
	errorCodeInvalidField
	errorCodeNoProviderFound
	errorCodeNoPlanFound
	errorCodeNoNodeFound
	errorCodeUnauthorized
	errorCodeDuplicateNode
	errorCodeNoNodeAdded
	errorCodeInvalidPlanStatus
	errorCodeNoPriceFound
	errorCodeInvalidNodeStatus
)

const (
	errorMsgUnknownMsgType    = "unknown message type: %s"
	errorMsgUnknownQueryType  = "unknown query type: %s"
	errorMsgInvalidField      = "invalid field: %s"
	errorMsgNoProviderFound   = "no provider found"
	errorMsgNoPlanFound       = "no plan found"
	errorMsgNoNodeFound       = "no node found"
	errorMsgUnauthorized      = "unauthorized"
	errorMsgDuplicateNode     = "duplicate node"
	errorMsgNoNodeAdded       = "no node added"
	errorMsgInvalidPlanStatus = "invalid plan status"
	errorMsgNoPriceFound      = "no price found"
	errorMsgInvalidNodeStatus = "invalid node status"
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

func ErrorNoProviderFound() sdk.Error {
	return sdk.NewError(Codespace, errorCodeNoProviderFound, errorMsgNoProviderFound)
}

func ErrorNoPlanFound() sdk.Error {
	return sdk.NewError(Codespace, errorCodeNoPlanFound, errorMsgNoPlanFound)
}

func ErrorNoNodeFound() sdk.Error {
	return sdk.NewError(Codespace, errorCodeNoNodeFound, errorMsgNoNodeFound)
}

func ErrorUnauthorized() sdk.Error {
	return sdk.NewError(Codespace, errorCodeUnauthorized, errorMsgUnauthorized)
}

func ErrorDuplicateNode() sdk.Error {
	return sdk.NewError(Codespace, errorCodeDuplicateNode, errorMsgDuplicateNode)
}

func ErrorNoNodeAdded() sdk.Error {
	return sdk.NewError(Codespace, errorCodeNoNodeAdded, errorMsgNoNodeAdded)
}

func ErrorInvalidPlanStatus() sdk.Error {
	return sdk.NewError(Codespace, errorCodeInvalidPlanStatus, errorMsgInvalidPlanStatus)
}

func ErrorNoPriceFound() sdk.Error {
	return sdk.NewError(Codespace, errorCodeNoPriceFound, errorMsgNoPriceFound)
}

func ErrorInvalidNodeStatus() sdk.Error {
	return sdk.NewError(Codespace, errorCodeInvalidNodeStatus, errorMsgInvalidNodeStatus)
}
