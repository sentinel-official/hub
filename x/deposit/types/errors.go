package types

import (
	"fmt"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	
	hub "github.com/sentinel-official/hub/types"
)

const (
	Codespace = sdk.CodespaceType("deposit")
	
	errCodeUnknownQueryType         = 101
	errCodeInsufficientDepositFunds = 102
	
	errMsgUnknownQueryType         = "Invalid query type: "
	errMsgInsufficientDepositFunds = "insufficient deposit funds: %s < %s"
)

func ErrorMarshal() sdk.Error {
	return sdk.NewError(Codespace, hub.ErrCodeMarshal, hub.ErrMsgMarshal)
}

func ErrorUnmarshal() sdk.Error {
	return sdk.NewError(Codespace, hub.ErrCodeUnmarshal, hub.ErrMsgUnmarshal)
}

func ErrorInvalidQueryType(queryType string) sdk.Error {
	return sdk.NewError(Codespace, errCodeUnknownQueryType, errMsgUnknownQueryType+queryType)
}

func ErrorInsufficientDepositFunds(x, y sdk.Coins) sdk.Error {
	return sdk.NewError(Codespace, errCodeInsufficientDepositFunds, fmt.Sprintf(errMsgInsufficientDepositFunds, x, y))
}
