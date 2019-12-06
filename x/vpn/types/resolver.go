package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Resolver struct {
	Owner      sdk.AccAddress `json:"owner"`
	Commission sdk.Dec        `json:"commission"`
	Status     string         `json:"status"`
}

func (resolver Resolver) String() string {
	return fmt.Sprintf(`
Owner : %s
Commission : %s
Status : %s
`, resolver.Owner, resolver.Commission, resolver.Status)
}

func (resolver Resolver) UpdateInfo(_resolver Resolver) Resolver {
	if _resolver.Commission.GTE(sdk.ZeroDec()) && _resolver.Commission.LTE(sdk.OneDec()) { // commission rate between 0 to 1
		resolver.Commission = _resolver.Commission
	}

	return resolver
}

func (resolver Resolver) GetCommission(pay sdk.Coin) sdk.Coin {
	commission := resolver.Commission.MulInt(pay.Amount).Quo(sdk.NewDec(100))
	pay.Amount = commission.RoundInt()

	return pay
}
