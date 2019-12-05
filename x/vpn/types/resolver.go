package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	hub "github.com/sentinel-official/hub/types"
)

type Resolver struct {
	ID         hub.ResolverID `json:"id"`
	Owner      sdk.AccAddress `json:"owner"`
	Commission sdk.Dec        `json:"commission"`
	Status     string         `json:"status"`
}

func (resolver Resolver) String() string {
	return fmt.Sprintf(`
ID : %s
Owner : %s
Commission : %s
Status : %s
`, resolver.ID, resolver.Owner, resolver.Commission, resolver.Status)
}
