package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *Swap) GetTxHash() (hash EthereumHash) {
	return BytesToHash(s.TxHash)
}

func (s *Swap) Validate() error {
	if s.TxHash == nil || len(s.TxHash) != EthereumHashLength {
		return fmt.Errorf("tx_hash length should be 32")
	}

	if _, err := sdk.AccAddressFromBech32(s.Receiver); err != nil {
		return err
	}

	if !s.Amount.IsValid() {
		return fmt.Errorf("amount should be valid")
	}

	return nil
}

type Swaps []Swap
