package vpn

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

const (
	codespaceVPN = csdkTypes.CodespaceType("node")

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
	return csdkTypes.NewError(codespaceVPN, sdkTypes.ErrCodeMarshal, sdkTypes.ErrMsgMarshal)
}

func errorUnmarshal() csdkTypes.Error {
	return csdkTypes.NewError(codespaceVPN, sdkTypes.ErrCodeUnmarshal, sdkTypes.ErrMsgUnmarshal)
}

func errorUnknownMsgType(msgType string) csdkTypes.Error {
	return csdkTypes.NewError(codespaceVPN, errCodeUnknownMsgType, errMsgUnknownMsgType+msgType)
}

func errorNodeAlreadyExists() csdkTypes.Error {
	return csdkTypes.NewError(codespaceVPN, errCodeNodeAlreadyExists, errMsgNodeAlreadyExists)
}

func errorNodeNotExists() csdkTypes.Error {
	return csdkTypes.NewError(codespaceVPN, errCodeNodeNotExists, errMsgNodeNotExists)
}

func errorUnauthorized() csdkTypes.Error {
	return csdkTypes.NewError(codespaceVPN, errCodeUnauthorized, errMsgUnauthorized)
}

func errorInvalidField(field string) csdkTypes.Error {
	return csdkTypes.NewError(codespaceVPN, errCodeInvalidField, errMsgInvalidField+field)
}

func errorInvalidLockDenom() csdkTypes.Error {
	return csdkTypes.NewError(codespaceVPN, errCodeInvalidLockDenom, errMsgInvalidLockDenom)
}

func errorInvalidNodeStatus() csdkTypes.Error {
	return csdkTypes.NewError(codespaceVPN, errCodeInvalidNodeStatus, errMsgInvalidNodeStatus)
}
