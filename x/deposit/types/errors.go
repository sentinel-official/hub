package types

import (
	"fmt"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

const (
	Codespace                       = csdkTypes.CodespaceType("deposit")
	errCodeInsufficientDepositFunds = 101
	errMsgInsufficientDepositFunds  = "insufficient deposit funds: %s < %s"
)

func ErrorInsufficientDepositFunds(x, y csdkTypes.Coins) csdkTypes.Error {
	return csdkTypes.NewError(Codespace, errCodeInsufficientDepositFunds,
		fmt.Sprintf(errMsgInsufficientDepositFunds, x, y))
}
