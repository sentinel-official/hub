package types

import (
	"fmt"

	csdk "github.com/cosmos/cosmos-sdk/types"

	sdk "github.com/ironman0x7b2/sentinel-sdk/types"
)

const (
	Codespace = csdk.CodespaceType("deposit")

	errCodeUnknownQueryType         = 101
	errCodeInsufficientDepositFunds = 102

	errMsgUnknownQueryType         = "Invalid query type: "
	errMsgInsufficientDepositFunds = "insufficient deposit funds: %s < %s"
)

func ErrorMarshal() csdk.Error {
	return csdk.NewError(Codespace, sdk.ErrCodeMarshal, sdk.ErrMsgMarshal)
}

func ErrorUnmarshal() csdk.Error {
	return csdk.NewError(Codespace, sdk.ErrCodeUnmarshal, sdk.ErrMsgUnmarshal)
}

func ErrorInvalidQueryType(queryType string) csdk.Error {
	return csdk.NewError(Codespace, errCodeUnknownQueryType, errMsgUnknownQueryType+queryType)
}

func ErrorInsufficientDepositFunds(x, y csdk.Coins) csdk.Error {
	return csdk.NewError(Codespace, errCodeInsufficientDepositFunds, fmt.Sprintf(errMsgInsufficientDepositFunds, x, y))
}
