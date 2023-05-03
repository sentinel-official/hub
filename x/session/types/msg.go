package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgStartRequest)(nil)
	_ sdk.Msg = (*MsgUpdateDetailsRequest)(nil)
	_ sdk.Msg = (*MsgEndRequest)(nil)
)

func NewMsgStartRequest(from sdk.AccAddress, id uint64, addr hubtypes.NodeAddress) *MsgStartRequest {
	return &MsgStartRequest{
		From:    from.String(),
		ID:      id,
		Address: addr.String(),
	}
}

func (m *MsgStartRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return errors.Wrap(ErrorInvalidMessage, "id cannot be zero")
	}
	if m.Address == "" {
		return errors.Wrap(ErrorInvalidMessage, "address cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.Address); err != nil {
		return errors.Wrap(ErrorInvalidMessage, err.Error())
	}

	return nil
}

func (m *MsgStartRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgUpdateDetailsRequest(from hubtypes.NodeAddress, proof Proof, signature []byte) *MsgUpdateDetailsRequest {
	return &MsgUpdateDetailsRequest{
		From:      from.String(),
		Proof:     proof,
		Signature: signature,
	}
}

func (m *MsgUpdateDetailsRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.From); err != nil {
		return errors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if m.Proof.ID == 0 {
		return errors.Wrap(ErrorInvalidMessage, "proof.id cannot be zero")
	}
	if m.Proof.Bandwidth.IsAnyNil() {
		return errors.Wrap(ErrorInvalidMessage, "proof.bandwidth cannot be empty")
	}
	if m.Proof.Bandwidth.IsAnyNegative() {
		return errors.Wrap(ErrorInvalidMessage, "proof.bandwidth cannot be negative")
	}
	if m.Proof.Duration < 0 {
		return errors.Wrap(ErrorInvalidMessage, "proof.duration cannot be negative")
	}
	if m.Signature != nil {
		if len(m.Signature) < 64 {
			return errors.Wrapf(ErrorInvalidMessage, "signature length cannot be less than %d", 64)
		}
		if len(m.Signature) > 64 {
			return errors.Wrapf(ErrorInvalidMessage, "signature length cannot be greater than %d", 64)
		}
	}

	return nil
}

func (m *MsgUpdateDetailsRequest) GetSigners() []sdk.AccAddress {
	from, err := hubtypes.NodeAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgEndRequest(from sdk.AccAddress, id uint64, rating uint64) *MsgEndRequest {
	return &MsgEndRequest{
		From:   from.String(),
		ID:     id,
		Rating: rating,
	}
}

func (m *MsgEndRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return errors.Wrap(ErrorInvalidMessage, "id cannot be zero")
	}
	if m.Rating > 10 {
		return errors.Wrapf(ErrorInvalidMessage, "rating cannot be greater than %d", 10)
	}

	return nil
}

func (m *MsgEndRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}
