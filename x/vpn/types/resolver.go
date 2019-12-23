package types

import (
	"fmt"
	"strings"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	
	hub "github.com/sentinel-official/hub/types"
)

type Resolver struct {
	ID               hub.ResolverID `json:"id"`
	Owner            sdk.AccAddress `json:"owner"`
	Commission       sdk.Dec        `json:"commission"`
	Status           string         `json:"status"`
	StatusModifiedAt int64          `json:"status_modified_at"`
}

func (resolver Resolver) String() string {
	return fmt.Sprintf(`
  ID : %s
  Owner : %s
  Commission : %s
  Status : %s
  StatusModifiedAt : %d
`, resolver.ID.String(), resolver.Owner, resolver.Commission,
		resolver.Status, resolver.StatusModifiedAt)
}

func (resolver Resolver) UpdateInfo(_resolver Resolver) Resolver {
	if _resolver.Commission.GTE(sdk.ZeroDec()) && _resolver.Commission.LTE(sdk.OneDec()) {
		// commission rate between 0 to 1
		resolver.Commission = _resolver.Commission
	}
	
	return resolver
}

func (resolver Resolver) GetCommission(pay sdk.Coin) sdk.Coin {
	commission := resolver.Commission.Mul(sdk.NewDec(100)).MulInt(pay.Amount).Quo(sdk.NewDec(100))
	pay.Amount = commission.RoundInt()
	
	return pay
}

type Resolvers []Resolver

func (resolvers Resolvers) String() string {
	var out string
	for _, resolver := range resolvers {
		out += resolver.String() + "\n"
	}
	
	return strings.TrimSpace(out)
}
