package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (q Quota) GetAddress() sdk.AccAddress {
	if q.Address == "" {
		return nil
	}

	address, err := sdk.AccAddressFromBech32(q.Address)
	if err != nil {
		panic(err)
	}

	return address
}

func (q Quota) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
Address:   %s
Consumed:  %s
Allocated: %s
`), q.Address, q.Consumed, q.Allocated)
}

func (q Quota) Validate() error {
	if _, err := sdk.AccAddressFromBech32(q.Address); err != nil {
		return fmt.Errorf("address should not be nil or empty")
	}
	if q.Consumed.IsNegative() {
		return fmt.Errorf("consumed should not be negative")
	}
	if q.Allocated.IsNegative() {
		return fmt.Errorf("allocated should not be negative")
	}
	if q.Consumed.GT(q.Allocated) {
		return fmt.Errorf("consumed should not be greater than allocated")
	}

	return nil
}

type Quotas []Quota
