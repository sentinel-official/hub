package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

type Quota struct {
	Address sdk.AccAddress `json:"address"`
	Current hub.Bandwidth  `json:"current"`
	Maximum hub.Bandwidth  `json:"maximum"`
}

func (q Quota) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
Address: %s
Current: %s
Maximum: %s
`), q.Address, q.Current, q.Maximum)
}

func (q Quota) Validate() error {
	if q.Address == nil || q.Address.Empty() {
		return fmt.Errorf("address should not be nil and empty")
	}
	if !q.Current.IsValid() {
		return fmt.Errorf("current should be valid")
	}
	if !q.Maximum.IsValid() {
		return fmt.Errorf("maximum should be valid")
	}

	return nil
}

type Quotas []Quota
