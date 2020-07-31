package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	Codespace = sdk.CodespaceType(ModuleName)
)

const (
	errorCodeUnknownMsgType sdk.CodeType = iota + 101
	errorCodeUnknownQueryType
	errorCodeInvalidField
	errorCodePlanDoesNotExist
	errorCodeNodeDoesNotExist
	errorCodeUnauthorized
	errorCodeInvalidPlanStatus
	errorCodePriceDoesNotExist
	errorCodeInvalidNodeStatus
	errorCodeSubscriptionDoesNotExist
	errorCodeInvalidSubscriptionStatus
	errorCodeDuplicateAddress
	errorCodeAddressWasNotAdded
	errorCodeCanNotSubscribe
)

const (
	errorMsgUnknownMsgType            = "unknown message type: %s"
	errorMsgUnknownQueryType          = "unknown query type: %s"
	errorMsgInvalidField              = "invalid field: %s"
	errorMsgPlanDoesNotExist          = "plan does not exist"
	errorMsgNodeDoesNotExist          = "node does not exist"
	errorMsgUnauthorized              = "unauthorized"
	errorMsgInvalidPlanStatus         = "invalid plan status"
	errorMsgPriceDoesNotExist         = "price does not exist"
	errorMsgInvalidNodeStatus         = "invalid node status"
	errorMsgSubscriptionDoesNotExist  = "subscription does not exist"
	errorMsgInvalidSubscriptionStatus = "invalid subscription status"
	errorMsgDuplicateAddress          = "duplicate address"
	errorMsgAddressWasNotAdded        = "address was not added"
	errorMsgCanNotSubscribe           = "can not subscribe"
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

func ErrorPlanDoesNotExist() sdk.Error {
	return sdk.NewError(Codespace, errorCodePlanDoesNotExist, errorMsgPlanDoesNotExist)
}

func ErrorNodeDoesNotExist() sdk.Error {
	return sdk.NewError(Codespace, errorCodeNodeDoesNotExist, errorMsgNodeDoesNotExist)
}

func ErrorUnauthorized() sdk.Error {
	return sdk.NewError(Codespace, errorCodeUnauthorized, errorMsgUnauthorized)
}

func ErrorInvalidPlanStatus() sdk.Error {
	return sdk.NewError(Codespace, errorCodeInvalidPlanStatus, errorMsgInvalidPlanStatus)
}

func ErrorPriceDoesNotExist() sdk.Error {
	return sdk.NewError(Codespace, errorCodePriceDoesNotExist, errorMsgPriceDoesNotExist)
}

func ErrorInvalidNodeStatus() sdk.Error {
	return sdk.NewError(Codespace, errorCodeInvalidNodeStatus, errorMsgInvalidNodeStatus)
}

func ErrorSubscriptionDoesNotExist() sdk.Error {
	return sdk.NewError(Codespace, errorCodeSubscriptionDoesNotExist, errorMsgSubscriptionDoesNotExist)
}

func ErrorInvalidSubscriptionStatus() sdk.Error {
	return sdk.NewError(Codespace, errorCodeInvalidSubscriptionStatus, errorMsgInvalidSubscriptionStatus)
}

func ErrorDuplicateAddress() sdk.Error {
	return sdk.NewError(Codespace, errorCodeDuplicateAddress, errorMsgDuplicateAddress)
}

func ErrorAddressWasNotAdded() sdk.Error {
	return sdk.NewError(Codespace, errorCodeAddressWasNotAdded, errorMsgAddressWasNotAdded)
}

func ErrorCanNotSubscribe() sdk.Error {
	return sdk.NewError(Codespace, errorCodeCanNotSubscribe, errorMsgCanNotSubscribe)
}
