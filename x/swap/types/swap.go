package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (m *Swap) GetTxHash() (hash EthereumHash) {
	return BytesToHash(m.TxHash)
}

func (m *Swap) Validate() error {
	if m.TxHash == nil {
		return fmt.Errorf("tx_hash cannot be nil")
	}
	if len(m.TxHash) == 0 {
		return fmt.Errorf("tx_hash cannot be empty")
	}
	if len(m.TxHash) < EthereumHashLength {
		return fmt.Errorf("tx_hash length cannot be less than %d", EthereumHashLength)
	}
	if len(m.TxHash) > EthereumHashLength {
		return fmt.Errorf("tx_hash length cannot be greater than %d", EthereumHashLength)
	}
	if m.Receiver == "" {
		return fmt.Errorf("receiver cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Receiver); err != nil {
		return errors.Wrapf(err, "invalid receiver %s", m.Receiver)
	}
	if m.Amount.IsNegative() {
		return fmt.Errorf("amount cannot be negative")
	}
	if m.Amount.IsZero() {
		return fmt.Errorf("amount cannot be zero")
	}
	if m.Amount.Amount.LT(PrecisionLoss) {
		return fmt.Errorf("amount cannot be less than %s", PrecisionLoss)
	}
	if !m.Amount.IsValid() {
		return fmt.Errorf("amount must be valid")
	}

	return nil
}

type (
	Swaps []Swap
)
