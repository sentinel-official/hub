package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	Codespace = sdk.CodespaceType("plan")
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
)

const (
	errorMsgUnknownMsgType       = "unknown message type: %s"
	errorMsgUnknownQueryType     = "unknown query type: %s"
	errorMsgInvalidField         = "invalid field: %s"
	errorMsgProviderDoesNotExist = "provider does not exist"
	errorMsgPlanDoesNotExist     = "plan does not exist"
	errorMsgNodeDoesNotExist     = "node does not exist"
	errorMsgUnauthorized         = "unauthorized"
	errorMsgDuplicateNode        = "duplicate node"
	errorMsgNodeWasNotAdded      = "node was not added"
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
