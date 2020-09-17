package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Quota struct {
	Address   sdk.AccAddress `json:"address"`
	Consumed  sdk.Int        `json:"consumed"`
	Allocated sdk.Int        `json:"allocated"`
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
		return fmt.Errorf("address is nil or empty")
	}
	if q.Consumed.IsNegative() {
		return fmt.Errorf("consumed is negative")
	}
	if q.Allocated.IsNegative() {
		return fmt.Errorf("allocated is netgative")
	}
	if q.Consumed.GT(q.Allocated) {
		return fmt.Errorf("consumed is greater than allocated")
	}

	return nil
}

type Quotas []Quota
