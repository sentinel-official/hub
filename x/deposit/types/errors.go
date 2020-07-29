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
	errorCodeUnknownQueryType = iota + 101
	errorCodeInsufficientDepositFunds
	errorCodeDepositDoesNotExist
)

const (
	errorMsgUnknownQueryType         = "unknown query type: %s"
	errorMsgInsufficientDepositFunds = "insufficient deposit funds"
	errorMsgDepositDoesNotExist      = "deposit does not exist"
)

func ErrorMarshal() sdk.Error {
	return sdk.NewError(Codespace, hub.ErrorCodeMarshal, hub.ErrorMsgMarshal)
}

func ErrorUnmarshal() sdk.Error {
	return sdk.NewError(Codespace, hub.ErrorCodeUnmarshal, hub.ErrorMsgUnmarshal)
}

func ErrorUnknownQueryType(v string) sdk.Error {
	return sdk.NewError(Codespace, errorCodeUnknownQueryType, fmt.Sprintf(errorMsgUnknownQueryType, v))
}

func ErrorInsufficientDepositFunds() sdk.Error {
	return sdk.NewError(Codespace, errorCodeInsufficientDepositFunds, errorMsgInsufficientDepositFunds)
}

func ErrorDepositDoesNotExist() sdk.Error {
	return sdk.NewError(Codespace, errorCodeDepositDoesNotExist, errorMsgDepositDoesNotExist)
}
