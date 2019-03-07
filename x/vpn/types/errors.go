package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

const (
	Codespace = csdkTypes.CodespaceType("vpn")

	errCodeUnknownMsgType       = 101
	errCodeUnknownQueryType     = 102
	errCodeInvalidField         = 103
	errCodeNodeNotExists        = 104
	errCodeUnauthorized         = 105
	errCodeInvalidLockDenom     = 106
	errCodeInvalidNodeStatus    = 107
	errCodeInvalidPriceDenom    = 108
	errCodeSessionNotExists     = 109
	errCodeBandwidthUpdate      = 110
	errCodeInvalidSessionStatus = 111

	errMsgUnknownMsgType        = "Unknown message type: "
	errMsgUnknownQueryType      = "Invalid query type: "
	errMsgInvalidField          = "Invalid field: "
	errMsgNodeNotExists         = "Node not exists"
	errMsgUnauthorized          = "Transaction is unauthorized"
	errMsgInvalidLockDenom      = "Invalid lock denom"
	errMsgInvalidNodeStatus     = "Invalid node status"
	errMsgInvalidPriceDenom     = "Invalid price denom"
	errMsgSessionNotExists      = "Session not exists"
	errMsgInvalidBandwidth      = "Invalid bandwidth"
	errMsgInvalidBandwidthSigns = "Invalid client sign or node owner sign"
	errMsgInvalidSessionStatus  = "Invalid session status"
)

func ErrorMarshal() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, sdkTypes.ErrCodeMarshal, sdkTypes.ErrMsgMarshal)
}

func ErrorUnmarshal() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, sdkTypes.ErrCodeUnmarshal, sdkTypes.ErrMsgUnmarshal)
}

func ErrorUnknownMsgType(msgType string) csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeUnknownMsgType, errMsgUnknownMsgType+msgType)
}

func ErrorInvalidQueryType(queryType string) csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeUnknownQueryType, errMsgUnknownQueryType+queryType)
}

func ErrorInvalidField(field string) csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeInvalidField, errMsgInvalidField+field)
}

func ErrorNodeNotExists() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeNodeNotExists, errMsgNodeNotExists)
}

func ErrorUnauthorized() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeUnauthorized, errMsgUnauthorized)
}

func ErrorInvalidLockDenom() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeInvalidLockDenom, errMsgInvalidLockDenom)
}

func ErrorInvalidNodeStatus() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeInvalidNodeStatus, errMsgInvalidNodeStatus)
}

func ErrorInvalidPriceDenom() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeInvalidPriceDenom, errMsgInvalidPriceDenom)
}

func ErrorSessionNotExists() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeSessionNotExists, errMsgSessionNotExists)
}

func ErrorBandwidthUpdate(msg string) csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeBandwidthUpdate, msg)
}

func ErrorInvalidSessionStatus() csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeInvalidSessionStatus, errMsgInvalidSessionStatus)
}
