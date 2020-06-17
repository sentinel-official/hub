package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	Codespace = sdk.CodespaceType("provider")
)

const (
	errorCodeUnknownMsgType = iota + 101
	errorCodeUnknownQueryType
	errorCodeInvalidField
	errorCodeDuplicateProviderAddress
	errorCodeNoProviderFound
	errorCodeUnauthorised
)

const (
	errorMsgUnknownMsgType           = "unknown message type: %s"
	errorMsgUnknownQueryType         = "unknown query type: %s"
	errorMsgInvalidField             = "invalid field: %s"
	errorMsgDuplicateProviderAddress = "duplicate provider address"
	errorMsgNoProviderFound          = "no provider found"
	errorMsgUnauthorised             = "unauthorised"
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

func ErrorDuplicateProviderAddress() sdk.Error {
	return sdk.NewError(Codespace, errorCodeDuplicateProviderAddress, errorMsgDuplicateProviderAddress)
}

func ErrorNoProviderFound() sdk.Error {
	return sdk.NewError(Codespace, errorCodeNoProviderFound, errorMsgNoProviderFound)
}

func ErrorUnauthorised() sdk.Error {
	return sdk.NewError(Codespace, errorCodeUnauthorised, errorMsgUnauthorised)
}
