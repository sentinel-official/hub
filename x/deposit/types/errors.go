package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	Codespace = sdk.CodespaceType("deposit")

	errCodeInvalidQueryType         = 101
	errCodeInsufficientDepositFunds = 102
	errCodeDepositDoesNotExist      = 103

	errMsgInvalidQueryType         = "invalid query type: %s"
	errMsgInsufficientDepositFunds = "insufficient deposit funds: %s < %s"
	errMsgDepositDoesNotExist      = "deposit does not exist"
)

func ErrorMarshal() sdk.Error {
	return sdk.NewError(Codespace, hub.ErrCodeMarshal, hub.ErrMsgMarshal)
}

func ErrorUnmarshal() sdk.Error {
	return sdk.NewError(Codespace, hub.ErrCodeUnmarshal, hub.ErrMsgUnmarshal)
}

func ErrorInvalidQueryType(queryType string) sdk.Error {
	return sdk.NewError(Codespace, errCodeInvalidQueryType, fmt.Sprintf(errMsgInvalidQueryType, queryType))
}

func ErrorInsufficientDepositFunds(x, y sdk.Coins) sdk.Error {
	return sdk.NewError(Codespace, errCodeInsufficientDepositFunds, fmt.Sprintf(errMsgInsufficientDepositFunds, x, y))
}

func ErrorDepositDoesNotExist() sdk.Error {
	return sdk.NewError(Codespace, errCodeDepositDoesNotExist, errMsgDepositDoesNotExist)
}
