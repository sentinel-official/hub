package types

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
	errCodeInvalidQueryType  = 208

	errMsgUnknownMsgType    = "Unknown message type: "
	errMsgNodeAlreadyExists = "Node already exists"
	errMsgNodeNotExists     = "Node not exists"
	errMsgUnauthorized      = "Transaction is unauthorized"
	errMsgInvalidField      = "Invalid field: "
	errMsgInvalidLockDenom  = "Invalid lock denom"
	errMsgInvalidNodeStatus = "Invalid node status"
	errMsgInvalidQueryType  = "Invalid query type: "
)

func ErrorMarshal() csdkTypes.Error {
	return csdkTypes.NewError(codespace, sdkTypes.ErrCodeMarshal, sdkTypes.ErrMsgMarshal)
}

func ErrorUnmarshal() csdkTypes.Error {
	return csdkTypes.NewError(codespace, sdkTypes.ErrCodeUnmarshal, sdkTypes.ErrMsgUnmarshal)
}

func ErrorUnknownMsgType(msgType string) csdkTypes.Error {
	return csdkTypes.NewError(codespace, errCodeUnknownMsgType, errMsgUnknownMsgType+msgType)
}

func ErrorNodeAlreadyExists() csdkTypes.Error {
	return csdkTypes.NewError(codespace, errCodeNodeAlreadyExists, errMsgNodeAlreadyExists)
}

func ErrorNodeNotExists() csdkTypes.Error {
	return csdkTypes.NewError(codespace, errCodeNodeNotExists, errMsgNodeNotExists)
}

func ErrorUnauthorized() csdkTypes.Error {
	return csdkTypes.NewError(codespace, errCodeUnauthorized, errMsgUnauthorized)
}

func ErrorInvalidField(field string) csdkTypes.Error {
	return csdkTypes.NewError(codespace, errCodeInvalidField, errMsgInvalidField+field)
}

func ErrorInvalidLockDenom() csdkTypes.Error {
	return csdkTypes.NewError(codespace, errCodeInvalidLockDenom, errMsgInvalidLockDenom)
}

func ErrorInvalidNodeStatus() csdkTypes.Error {
	return csdkTypes.NewError(codespace, errCodeInvalidNodeStatus, errMsgInvalidNodeStatus)
}

func ErrorInvalidQueryType(queryType string) csdkTypes.Error {
	return csdkTypes.NewError(codespace, errCodeInvalidQueryType, errMsgInvalidQueryType+queryType)
}
