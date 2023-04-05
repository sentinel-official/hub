package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = (*MsgSwapRequest)(nil)
)

func NewMsgSwapRequest(from sdk.AccAddress, txHash EthereumHash, receiver sdk.AccAddress, amount sdk.Int) *MsgSwapRequest {
	return &MsgSwapRequest{
		From:     from.String(),
		TxHash:   txHash.Bytes(),
		Receiver: receiver.String(),
		Amount:   amount,
	}
}

func (m *MsgSwapRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if m.Receiver == "" {
		return errors.Wrap(ErrorInvalidReceiver, "receiver cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Receiver); err != nil {
		return errors.Wrapf(ErrorInvalidReceiver, "%s", err)
	}
	if m.TxHash == nil {
		return errors.Wrap(ErrorInvalidTxHash, "tx_hash cannot be nil")
	}
	if len(m.TxHash) == 0 {
		return errors.Wrap(ErrorInvalidTxHash, "tx_hash cannot be empty")
	}
	if len(m.TxHash) < EthereumHashLength {
		return errors.Wrapf(ErrorInvalidTxHash, "tx_hash length cannot be less than %d", EthereumHashLength)
	}
	if len(m.TxHash) > EthereumHashLength {
		return errors.Wrapf(ErrorInvalidTxHash, "tx_hash length cannot be greater than %d", EthereumHashLength)
	}
	if m.Amount.IsNegative() {
		return errors.Wrap(ErrorInvalidAmount, "amount cannot be negative")
	}
	if m.Amount.IsZero() {
		return errors.Wrap(ErrorInvalidAmount, "amount cannot be zero")
	}
	if m.Amount.LT(PrecisionLoss) {
		return errors.Wrapf(ErrorInvalidAmount, "amount cannot be less than %s", PrecisionLoss)
	}

	return nil
}

func (m *MsgSwapRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}
