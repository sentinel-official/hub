package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Swap struct {
	TxHash   EthereumHash   `json:"tx_hash"`
	Receiver sdk.AccAddress `json:"receiver"`
	Amount   sdk.Coin       `json:"amount"`
}

func (s Swap) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
Tx hash:  %s
Receiver: %s
Amount:   %s
`), s.TxHash, s.Receiver, s.Amount)
}

func (s Swap) Validate() error {
	if s.Receiver == nil || s.Receiver.Empty() {
		return fmt.Errorf("receiver should not be nil or empty")
	}
	if !s.Amount.IsValid() {
		return fmt.Errorf("amount should be valid")
	}

	return nil
}

type Swaps []Swap
