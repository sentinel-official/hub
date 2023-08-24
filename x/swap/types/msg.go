package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgSwapRequest)(nil)
)

func NewMsgSwapRequest(from sdk.AccAddress, txHash EthereumHash, receiver sdk.AccAddress, amount sdkmath.Int) *MsgSwapRequest {
	return &MsgSwapRequest{
		From:     from.String(),
		TxHash:   txHash.Bytes(),
		Receiver: receiver.String(),
		Amount:   amount,
	}
}

func (m *MsgSwapRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if m.Receiver == "" {
		return sdkerrors.Wrap(ErrorInvalidReceiver, "receiver cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Receiver); err != nil {
		return sdkerrors.Wrapf(ErrorInvalidReceiver, "%s", err)
	}
	if m.TxHash == nil {
		return sdkerrors.Wrap(ErrorInvalidTxHash, "tx_hash cannot be nil")
	}
	if len(m.TxHash) == 0 {
		return sdkerrors.Wrap(ErrorInvalidTxHash, "tx_hash cannot be empty")
	}
	if len(m.TxHash) < EthereumHashLength {
		return sdkerrors.Wrapf(ErrorInvalidTxHash, "tx_hash length cannot be less than %d", EthereumHashLength)
	}
	if len(m.TxHash) > EthereumHashLength {
		return sdkerrors.Wrapf(ErrorInvalidTxHash, "tx_hash length cannot be greater than %d", EthereumHashLength)
	}
	if m.Amount.IsNegative() {
		return sdkerrors.Wrap(ErrorInvalidAmount, "amount cannot be negative")
	}
	if m.Amount.IsZero() {
		return sdkerrors.Wrap(ErrorInvalidAmount, "amount cannot be zero")
	}
	if m.Amount.LT(PrecisionLoss) {
		return sdkerrors.Wrapf(ErrorInvalidAmount, "amount cannot be less than %s", PrecisionLoss)
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
