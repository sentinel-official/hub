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
	errorCodeProviderDoesNotExist
	errorCodePlanDoesNotExist
	errorCodeNodeDoesNotExist
	errorCodeUnauthorized
	errorCodeDuplicateNode
	errorCodeNodeWasNotAdded
	errorCodeInvalidPlanStatus
	errorCodePriceDoesNotExist
	errorCodeInvalidNodeStatus
	errorCodePriceDoesNotExit
	errorCodeSubscriptionDoesNotExist
	errorCodeInvalidSubscriptionStatus
	errorCodeDuplicateAddress
)

const (
	errorMsgUnknownMsgType            = "unknown message type: %s"
	errorMsgUnknownQueryType          = "unknown query type: %s"
	errorMsgInvalidField              = "invalid field: %s"
	errorMsgProviderDoesNotExist      = "provider does not exist"
	errorMsgPlanDoesNotExist          = "plan does not exist"
	errorMsgNodeDoesNotExist          = "node does not exist"
	errorMsgUnauthorized              = "unauthorized"
	errorMsgDuplicateNode             = "duplicate node"
	errorMsgNodeWasNotAdded           = "node was not added"
	errorMsgInvalidPlanStatus         = "invalid plan status"
	errorMsgPriceDoesNotExist         = "price does not exist"
	errorMsgInvalidNodeStatus         = "invalid node status"
	errorMsgPriceDoesNotExit          = "price does not exist"
	errorMsgSubscriptionDoesNotExist  = "subscription does not exist"
	errorMsgInvalidSubscriptionStatus = "invalid subscription status"
	errorMsgDuplicateAddress          = "duplicate address"
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

func ErrorProviderDoesNotExist() sdk.Error {
	return sdk.NewError(Codespace, errorCodeProviderDoesNotExist, errorMsgProviderDoesNotExist)
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

func ErrorDuplicateNode() sdk.Error {
	return sdk.NewError(Codespace, errorCodeDuplicateNode, errorMsgDuplicateNode)
}

func ErrorNodeWasNotAdded() sdk.Error {
	return sdk.NewError(Codespace, errorCodeNodeWasNotAdded, errorMsgNodeWasNotAdded)
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

func ErrorPriceDoesNotExit() sdk.Error {
	return sdk.NewError(Codespace, errorCodePriceDoesNotExit, errorMsgPriceDoesNotExit)
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
