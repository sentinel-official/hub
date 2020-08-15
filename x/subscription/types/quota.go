package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

type Quota struct {
	Address   sdk.AccAddress `json:"address"`
	Consumed  hub.Bandwidth  `json:"consumed"`
	Allocated hub.Bandwidth  `json:"allocated"`
}

func (q Quota) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
Address:   %s
Consumed:  %s
Allocated: %s
`), q.Address, q.Consumed, q.Allocated)
}

func (q Quota) Validate() error {
	if q.Address == nil || q.Address.Empty() {
		return fmt.Errorf("address should not be nil or empty")
	}
	if !q.Consumed.IsValid() {
		return fmt.Errorf("consumed should be valid")
	}
	if !q.Allocated.IsValid() {
		return fmt.Errorf("allocated should be valid")
	}
	if q.Consumed.IsAnyGT(q.Allocated) {
		return fmt.Errorf("consumed should not be greater than allocated")
	}

	return nil
}

type Quotas []Quota
