package vpn

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

const (
	codespace = csdkTypes.CodespaceType("vpn")

	errCodeUnknownMsgType    = 201
	errCodeNodeAlreadyExists = 202
	errCodeNodeNotExists     = 203
	errCodeUnauthorized      = 204
	errCodeInvalidField      = 205
	errCodeInvalidLockDenom  = 206
	errCodeInvalidNodeStatus = 207

	errMsgUnknownMsgType    = "Unknown message type: "
	errMsgNodeAlreadyExists = "Node already exists"
	errMsgNodeNotExists     = "Node not exists"
	errMsgUnauthorized      = "Transaction is unauthorized"
	errMsgInvalidField      = "Invalid field: "
	errMsgInvalidLockDenom  = "Invalid lock denom"
	errMsgInvalidNodeStatus = "Invalid node status"
)

func errorMarshal() csdkTypes.Error {
	return csdkTypes.NewError(codespace, sdkTypes.ErrCodeMarshal, sdkTypes.ErrMsgMarshal)
}

func errorUnmarshal() csdkTypes.Error {
	return csdkTypes.NewError(codespace, sdkTypes.ErrCodeUnmarshal, sdkTypes.ErrMsgUnmarshal)
}

func errorUnknownMsgType(msgType string) csdkTypes.Error {
	return csdkTypes.NewError(codespace, errCodeUnknownMsgType, errMsgUnknownMsgType+msgType)
}

func errorNodeAlreadyExists() csdkTypes.Error {
	return csdkTypes.NewError(codespace, errCodeNodeAlreadyExists, errMsgNodeAlreadyExists)
}

func errorNodeNotExists() csdkTypes.Error {
	return csdkTypes.NewError(codespace, errCodeNodeNotExists, errMsgNodeNotExists)
}

func errorUnauthorized() csdkTypes.Error {
	return csdkTypes.NewError(codespace, errCodeUnauthorized, errMsgUnauthorized)
}

func errorInvalidField(field string) csdkTypes.Error {
	return csdkTypes.NewError(codespace, errCodeInvalidField, errMsgInvalidField+field)
}

func errorInvalidLockDenom() csdkTypes.Error {
	return csdkTypes.NewError(codespace, errCodeInvalidLockDenom, errMsgInvalidLockDenom)
}

func errorInvalidNodeStatus() csdkTypes.Error {
	return csdkTypes.NewError(codespace, errCodeInvalidNodeStatus, errMsgInvalidNodeStatus)
}
