package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

type Provider struct {
	ID          hub.ProviderID `json:"id"`
	Address     sdk.AccAddress `json:"address"`
	Name        string         `json:"name"`
	Website     string         `json:"website"`
	Description string         `json:"description"`
}

func (p Provider) String() string {
	return strings.TrimSpace(fmt.Sprintf(`
ID: %s
Address: %s
Name: %s
Website: %s
Description: %s
`, p.ID, p.Address, p.Name, p.Website, p.Description))
}

type Providers []Provider
