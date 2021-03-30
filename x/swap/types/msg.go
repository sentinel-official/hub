package types

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgSwap struct {
	From     sdk.AccAddress `json:"from"`
	TxHash   EthereumHash   `json:"tx_hash"`
	Receiver sdk.AccAddress `json:"receiver"`
	Amount   sdk.Int        `json:"amount"`
}

func NewMsgSwap(from sdk.AccAddress, txHash EthereumHash, receiver sdk.AccAddress, amount sdk.Int) MsgSwap {
	return MsgSwap{
		From:     from,
		TxHash:   txHash,
		Receiver: receiver,
		Amount:   amount,
	}
}

func (m MsgSwap) Route() string {
	return RouterKey
}

func (m MsgSwap) Type() string {
	return fmt.Sprintf("%s:swap", ModuleName)
}

func (m MsgSwap) ValidateBasic() error {
	if m.From == nil || m.From.Empty() {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	if m.Receiver == nil || m.Receiver.Empty() {
		return errors.Wrapf(ErrorInvalidField, "%s", "receiver")
	}

	if m.Amount.LT(PrecisionLoss) {
		return errors.Wrapf(ErrorInvalidField, "%s", "amount")
	}

	return nil
}

func (m MsgSwap) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgSwap) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}
