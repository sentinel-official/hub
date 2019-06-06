package types

import (
	"fmt"

	csdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Codespace                       = csdk.CodespaceType("deposit")
	errCodeInsufficientDepositFunds = 101
	errMsgInsufficientDepositFunds  = "insufficient deposit funds: %s < %s"
)

func ErrorInsufficientDepositFunds(x, y csdk.Coins) csdk.Error {
	return csdk.NewError(Codespace, errCodeInsufficientDepositFunds,
		fmt.Sprintf(errMsgInsufficientDepositFunds, x, y))
}
