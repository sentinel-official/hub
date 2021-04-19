package types

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func NewMsgSwapRequest(from sdk.AccAddress, txHash EthereumHash, receiver sdk.AccAddress, amount sdk.Int) *MsgSwapRequest {
	return &MsgSwapRequest{
		From:     from.String(),
		TxHash:   txHash.Bytes(),
		Receiver: receiver.String(),
		Amount:   amount,
	}
}

func (m *MsgSwapRequest) Route() string {
	return RouterKey
}

func (m *MsgSwapRequest) Type() string {
	return fmt.Sprintf("%s:swap", ModuleName)
}

func (m *MsgSwapRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	if _, err := sdk.AccAddressFromBech32(m.Receiver); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "receiver")
	}

	if m.Amount.LT(PrecisionLoss) {
		return errors.Wrapf(ErrorInvalidField, "%s", "amount")
	}

	return nil
}

func (m *MsgSwapRequest) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m *MsgSwapRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}
