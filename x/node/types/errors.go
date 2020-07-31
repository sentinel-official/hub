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
	errorCodeProviderDoesNotExist
	errorCodeDuplicateNode
	errorCodeNodeDoesNotExist
	errorCodeCanNotUpdate
)

const (
	errorMsgUnknownMsgType       = "unknown message type: %s"
	errorMsgUnknownQueryType     = "unknown query type: %s"
	errorMsgInvalidField         = "invalid field: %s"
	errorMsgProviderDoesNotExist = "provider does not exist"
	errorMsgDuplicateNode        = "duplicate node"
	errorMsgNodeDoesNotExist     = "node does not exist"
	errorMsgCanNotUpdate         = "can not update"
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

func ErrorDuplicateNode() sdk.Error {
	return sdk.NewError(Codespace, errorCodeDuplicateNode, errorMsgDuplicateNode)
}

func ErrorNodeDoesNotExist() sdk.Error {
	return sdk.NewError(Codespace, errorCodeNodeDoesNotExist, errorMsgNodeDoesNotExist)
}

func ErrorCanNotUpdate() sdk.Error {
	return sdk.NewError(Codespace, errorCodeCanNotUpdate, errorMsgCanNotUpdate)
}
