package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	Codespace = sdk.CodespaceType("dvpn")
)

const (
	errCodeUnknownMsgType = iota + 101
	errCodeUnknownQueryType
)

const (
	errMsgUnknownMsgType   = "unknown message type: %s"
	errMsgUnknownQueryType = "unknown query type: %s"
)

func ErrorMarshal() sdk.Error {
	return sdk.NewError(Codespace, hub.ErrCodeMarshal, hub.ErrMsgMarshal)
}

func ErrorUnmarshal() sdk.Error {
	return sdk.NewError(Codespace, hub.ErrCodeUnmarshal, hub.ErrMsgUnmarshal)
}

func ErrorUnknownMsgType(v string) sdk.Error {
	return sdk.NewError(Codespace, errCodeUnknownMsgType, fmt.Sprintf(errMsgUnknownMsgType, v))
}

func ErrorUnknownQueryType(v string) sdk.Error {
	return sdk.NewError(Codespace, errCodeUnknownQueryType, fmt.Sprintf(errMsgUnknownQueryType, v))
}
